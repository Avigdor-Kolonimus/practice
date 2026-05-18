package geo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"mcp_weather/config"
)

const (
	urlGeocodingDefault = "https://geocoding-api.open-meteo.com/v1/search"
	urlForecastDefault  = "https://api.open-meteo.com/v1/forecast"

	defaultRequestTimeout = 8 * time.Second
	defaultToolTimeout    = 10 * time.Second
	defaultCacheTTL       = time.Hour
)

// Config holds optional WeatherService settings. Zero values use defaults.
type Config struct {
	HTTPClient  *http.Client
	CacheTTL    time.Duration
	ToolTimeout time.Duration
}

// WeatherService fetches weather via Open-Meteo with geocoding cache.
type WeatherService struct {
	http        *http.Client
	cacheTTL    time.Duration
	toolTimeout time.Duration

	cacheMu sync.RWMutex
	cache   map[string]cacheEntry
}

type cacheEntry struct {
	lat, lon float64
	label    string
	expires  time.Time
}

// New creates a WeatherService. Pass Config{} for defaults.
func NewWeatherService(cfg *config.WeatherServerConfig) *WeatherService {
	var httpClient *http.Client
	if cfg.RequestTimeout == 0 {
		httpClient = &http.Client{Timeout: defaultRequestTimeout}
	} else {
		httpClient = &http.Client{Timeout: time.Duration(cfg.RequestTimeout) * time.Second}
	}

	cacheTTL := time.Duration(cfg.CacheTTL) * time.Second
	if cacheTTL == 0 {
		cacheTTL = defaultCacheTTL
	}

	toolTimeout := time.Duration(cfg.ToolTimeout) * time.Second
	if toolTimeout == 0 {
		toolTimeout = defaultToolTimeout
	}

	return &WeatherService{
		http:        httpClient,
		cacheTTL:    cacheTTL,
		toolTimeout: toolTimeout,
		cache:       make(map[string]cacheEntry),
	}
}

func cacheKey(city string) string {
	return strings.ToLower(strings.TrimSpace(city))
}

func (ws *WeatherService) cacheGet(key string) (float64, float64, string, bool) {
	ws.cacheMu.RLock()
	entry, ok := ws.cache[key]
	ws.cacheMu.RUnlock()

	if !ok || time.Now().After(entry.expires) {
		return 0, 0, "", false
	}

	return entry.lat, entry.lon, entry.label, true
}

func (ws *WeatherService) cachePut(key string, lat, lon float64, label string) {
	ws.cacheMu.Lock()
	defer ws.cacheMu.Unlock()

	ws.cache[key] = cacheEntry{
		lat: lat, lon: lon, label: label,
		expires: time.Now().Add(ws.cacheTTL),
	}
}

func (ws *WeatherService) getJSON(ctx context.Context, u string, dest any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return err
	}

	res, err := ws.http.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(io.LimitReader(res.Body, 1<<14))
		return fmt.Errorf("status %d: %s", res.StatusCode, string(b))
	}

	return json.NewDecoder(res.Body).Decode(dest)
}

type Geo struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Country   string  `json:"country"`
	Admin1    string  `json:"admin1"`
}

type geoResp struct {
	Results []Geo `json:"results"`
}

func (ws *WeatherService) geocode(ctx context.Context, city string) (float64, float64, string, error) {
	key := cacheKey(city)
	if lat, lon, label, ok := ws.cacheGet(key); ok {
		return lat, lon, label, nil
	}

	v := url.Values{}
	v.Set("name", strings.TrimSpace(city))
	v.Set("count", "1")
	v.Set("language", "tr")
	v.Set("format", "json")
	u := urlGeocodingDefault + "?" + v.Encode()

	var gr geoResp
	if err := ws.getJSON(ctx, u, &gr); err != nil {
		return 0, 0, "", err
	}
	if len(gr.Results) == 0 {
		return 0, 0, "", fmt.Errorf("city not found")
	}
	r := gr.Results[0]

	label := r.Name
	if r.Admin1 != "" {
		label += ", " + r.Admin1
	}
	if r.Country != "" {
		label += ", " + r.Country
	}

	ws.cachePut(key, r.Latitude, r.Longitude, label)

	return r.Latitude, r.Longitude, label, nil
}

type TempAndWind struct {
	Temperature float64 `json:"temperature_2m"`
	Wind        float64 `json:"wind_speed_10m"`
}

type wxResp struct {
	Current TempAndWind `json:"current"`
}

func (ws *WeatherService) fetchWeather(ctx context.Context, lat, lon float64) (float64, float64, error) {
	v := url.Values{}
	v.Set("latitude", fmt.Sprintf("%.4f", lat))
	v.Set("longitude", fmt.Sprintf("%.4f", lon))
	v.Set("current", "temperature_2m,wind_speed_10m")
	v.Set("timezone", "Europe/Istanbul")
	u := urlForecastDefault + "?" + v.Encode()

	var wr wxResp
	if err := ws.getJSON(ctx, u, &wr); err != nil {
		return 0, 0, err
	}

	return wr.Current.Temperature, wr.Current.Wind, nil
}

type WeatherParams struct {
	City string `json:"city" jsonschema:"name of city, example: Ankara"`
}

func (ws *WeatherService) AskWeather(ctx context.Context, _ *mcp.CallToolRequest, params WeatherParams) (*mcp.CallToolResult, any, error) {
	ctx, cancel := context.WithTimeout(ctx, ws.toolTimeout)
	defer cancel()

	lat, lon, label, err := ws.geocode(ctx, params.City)
	if err != nil {
		return nil, nil, fmt.Errorf("location not found: %w", err)
	}

	temp, wind, err := ws.fetchWeather(ctx, lat, lon)
	if err != nil {
		return nil, nil, fmt.Errorf("weather could not fetch: %w", err)
	}

	msg := fmt.Sprintf("%s %.1f°C, wind %.1f ", label, temp, wind)

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: msg}},
	}, nil, nil
}
