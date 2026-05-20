package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"mcp_weather/config"
	geo "mcp_weather/geo"
)

func main() {
	// Creating the MCP Server
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "mcp-weather-sse",
		Version: "1.0.0",
	}, nil)

	cfgPath := config.DefaultWeatherServerConfigPath
	if p := os.Getenv("MCP_WEATHER_CONFIG"); p != "" {
		cfgPath = p
	}
	cfg, err := config.Load(cfgPath)
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	weather := geo.NewWeatherService(cfg)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "ask_weather",
		Description: "Returns city-based live weather forecast",
	}, weather.AskWeather)

	// Configuring the Transport
	addr := ":8080"
	fmt.Printf("MCP SSE server listening on %s\n", addr)

	handler := mcp.NewSSEHandler(func(request *http.Request) *mcp.Server {
		return server
	}, nil)

	if err := http.ListenAndServe(addr, handler); err != nil {
		panic(err)
	}
}
