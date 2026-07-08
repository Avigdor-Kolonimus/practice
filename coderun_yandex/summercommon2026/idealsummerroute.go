package summercommon2026

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

// API structures

type ScoreWeights struct {
	Rating            float64 `json:"rating"`
	Diversity         float64 `json:"diversity"`
	TravelPenalty     float64 `json:"travel_penalty"`
	BadWeatherPenalty float64 `json:"bad_weather_penalty"`
}

type DayResponse struct {
	DayID            string       `json:"day_id"`
	TimeLimitMinutes int          `json:"time_limit_minutes"`
	Budget           int          `json:"budget"`
	ScoreWeights     ScoreWeights `json:"score_weights"`
}

type Place struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

type PlacesResponse struct {
	Places []Place `json:"places"`
}

type PlaceDetail struct {
	ID               int     `json:"id"`
	Type             string  `json:"type"`
	VisitTimeMinutes int     `json:"visit_time_minutes"`
	Cost             int     `json:"cost"`
	Rating           float64 `json:"rating"`
	WeatherSensitive bool    `json:"weather_sensitive"`
	OpenFrom         string  `json:"open_from"`
	OpenTo           string  `json:"open_to"`
}

type BatchPlacesResponse struct {
	Places []PlaceDetail `json:"places"`
}

type MatrixResponse struct {
	IDs    []string `json:"ids"`
	Matrix [][]int  `json:"matrix"`
}

type WeatherResponse struct {
	Weather string `json:"weather"`
}

type PlaceNode struct {
	ID           int
	Index        int
	Type         string
	VisitTime    int
	Cost         int
	Rating       float64
	IsBadWeather bool
}

var (
	bestScore float64 = -1e18
	bestPath  []int
)

// https://coderun.yandex.ru/seasons/2026-summer/tracks/common/problem/ideal-summer-route
// IdealSummerRoute - problem 7
func IdealSummerRoute() {
	in := bufio.NewReader(os.Stdin)
	baseURL, _ := in.ReadString('\n')
	dayID, _ := in.ReadString('\n')
	baseURL = strings.TrimSpace(baseURL)
	dayID = strings.TrimSpace(dayID)

	if baseURL == "" || dayID == "" {
		fmt.Println(0)
		return
	}

	// 1. GET /day/{day_id}
	respDay, err := http.Get(fmt.Sprintf("%s/day/%s", baseURL, dayID))
	if err != nil {
		fmt.Println(0)

		return
	}
	defer respDay.Body.Close()

	var dayData DayResponse
	json.NewDecoder(respDay.Body).Decode(&dayData)

	// 2. GET /places?day_id={day_id}
	respPlaces, err := http.Get(fmt.Sprintf("%s/places?day_id=%s", baseURL, dayID))
	if err != nil {
		fmt.Println(0)

		return
	}
	defer respPlaces.Body.Close()

	var placesData PlacesResponse
	json.NewDecoder(respPlaces.Body).Decode(&placesData)

	if len(placesData.Places) == 0 {
		fmt.Println(0)

		return
	}

	var ids []int
	for _, p := range placesData.Places {
		ids = append(ids, p.ID)
	}

	// 3. POST /places/batch
	reqBody, _ := json.Marshal(map[string][]int{"ids": ids})
	respBatch, err := http.Post(fmt.Sprintf("%s/places/batch", baseURL), "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(0)
		return
	}
	defer respBatch.Body.Close()

	var batchData BatchPlacesResponse
	json.NewDecoder(respBatch.Body).Decode(&batchData)

	// 4. POST /travel-time/matrix
	matrixIDs := []string{"start"}
	for _, id := range ids {
		matrixIDs = append(matrixIDs, strconv.Itoa(id))
	}
	reqBodyMatrix, _ := json.Marshal(map[string][]string{"ids": matrixIDs})
	respMatrix, err := http.Post(fmt.Sprintf("%s/travel-time/matrix", baseURL), "application/json", bytes.NewBuffer(reqBodyMatrix))
	if err != nil {
		fmt.Println(0)

		return
	}
	defer respMatrix.Body.Close()

	var matrixData MatrixResponse
	json.NewDecoder(respMatrix.Body).Decode(&matrixData)

	idToIndex := make(map[string]int)
	for i, idStr := range matrixData.IDs {
		idToIndex[idStr] = i
	}

	// 5. GET /weather for all place
	n := len(batchData.Places)
	nodes := make([]PlaceNode, n)
	var wg sync.WaitGroup
	for i := range n {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			p := batchData.Places[idx]
			url := fmt.Sprintf("%s/weather?day_id=%s&place_id=%d", baseURL, dayID, p.ID)
			respW, err := http.Get(url)
			var isBad bool
			if err == nil {
				var wData WeatherResponse
				json.NewDecoder(respW.Body).Decode(&wData)
				respW.Body.Close()
				if p.WeatherSensitive && (wData.Weather == "rainy" || wData.Weather == "windy") {
					isBad = true
				}
			}
			nodes[idx] = PlaceNode{
				ID:           p.ID,
				Index:        idToIndex[strconv.Itoa(p.ID)],
				Type:         p.Type,
				VisitTime:    p.VisitTimeMinutes,
				Cost:         p.Cost,
				Rating:       p.Rating,
				IsBadWeather: isBad,
			}
		}(i)
	}
	wg.Wait()

	visited := make([]bool, n)
	currentPath := make([]PlaceNode, 0, n)
	typeCounts := make(map[string]int)

	startIndex := idToIndex["start"]

	var dfs func(currIdx int, remTime int, remBudget int, currentRatingSum float64, badWeatherCount int)
	dfs = func(currIdx int, remTime int, remBudget int, currentRatingSum float64, badWeatherCount int) {
		diversity := len(typeCounts)
		travelSum := dayData.TimeLimitMinutes - remTime

		var totalVisitTime int
		for _, node := range currentPath {
			totalVisitTime += node.VisitTime
		}
		pureTravelTime := travelSum - totalVisitTime

		score := currentRatingSum*dayData.ScoreWeights.Rating +
			float64(diversity)*dayData.ScoreWeights.Diversity -
			float64(pureTravelTime)*dayData.ScoreWeights.TravelPenalty -
			float64(badWeatherCount)*dayData.ScoreWeights.BadWeatherPenalty

		if len(currentPath) > 0 && score > bestScore {
			bestScore = score
			bestPath = make([]int, len(currentPath))
			for i, node := range currentPath {
				bestPath[i] = node.ID
			}
		}

		for i := range n {
			if visited[i] {
				continue
			}

			node := nodes[i]
			travelTimeToNext := matrixData.Matrix[currIdx][node.Index]
			totalNeededTime := travelTimeToNext + node.VisitTime

			if remTime >= totalNeededTime && remBudget >= node.Cost {
				visited[i] = true
				currentPath = append(currentPath, node)
				typeCounts[node.Type]++

				nextBadWeather := badWeatherCount
				if node.IsBadWeather {
					nextBadWeather++
				}

				dfs(node.Index, remTime-totalNeededTime, remBudget-node.Cost, currentRatingSum+node.Rating, nextBadWeather)

				typeCounts[node.Type]--
				if typeCounts[node.Type] == 0 {
					delete(typeCounts, node.Type)
				}
				currentPath = currentPath[:len(currentPath)-1]
				visited[i] = false
			}
		}
	}

	dfs(startIndex, dayData.TimeLimitMinutes, dayData.Budget, 0.0, 0)

	// 6. Result
	if len(bestPath) == 0 || bestScore <= -1e17 {
		fmt.Println(0)
	} else {
		fmt.Println(len(bestPath))

		strIDs := make([]string, len(bestPath))
		for i, id := range bestPath {
			strIDs[i] = strconv.Itoa(id)
		}

		fmt.Println(strings.Join(strIDs, " "))
	}
}
