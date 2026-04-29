package yandexinterview

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type PointInterestingJourney struct {
	x int
	y int
}

func manhattanGeometryDist(a, b PointInterestingJourney) int {
	distX := a.x - b.x
	if distX < 0 {
		distX = -distX
	}

	distY := a.y - b.y
	if distY < 0 {
		distY = -distY
	}

	return distX + distY
}

func bfsInterestingJourney(cities []PointInterestingJourney, startPoint, endPoint, n, k int) int {
	queue := make([]int, 0, n)
	queue = append(queue, startPoint)

	distances := make([]int, n)
	for i := range distances {
		distances[i] = -1
	}
	distances[startPoint] = 0

	head := 0
	for head < len(queue) {
		u := queue[head]
		head++

		if u == endPoint {
			return distances[u]
		}

		for v := 0; v < n; v++ {
			if distances[v] == -1 && manhattanGeometryDist(cities[u], cities[v]) <= k {
				distances[v] = distances[u] + 1
				queue = append(queue, v)
			}
		}
	}

	return -1
}

func validateInterestingJourneyNInput(n int) bool {
	return n >= 2 && n <= 1_000
}

func validateInterestingJourneyKInput(k int) bool {
	return k <= 2_000_000_000
}

// https://coderun.yandex.ru/selections/yandex-interview/problems/interesting-journey
// InterestingJourney - problem 5
func InterestingJourney() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateInterestingJourneyNInput(n) {
		panic("number N out of range")
	}

	// coordinates input
	cityCoordinates := make([]PointInterestingJourney, n)
	for i := 0; i < n; i++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		strNum := strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		x, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		if !validateInterestingJourneyKInput(x) {
			panic("X out of range")
		}

		y, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}
		if !validateInterestingJourneyKInput(y) {
			panic("Y out of range")
		}

		cityCoordinates[i].x = x
		cityCoordinates[i].y = y
	}

	// K input
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	k, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateInterestingJourneyKInput(k) {
		panic("number K out of range")
	}

	// start and end input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	startPoint, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}

	endPoint, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	startPoint--
	endPoint--

	result := bfsInterestingJourney(cityCoordinates, startPoint, endPoint, n, k)

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
