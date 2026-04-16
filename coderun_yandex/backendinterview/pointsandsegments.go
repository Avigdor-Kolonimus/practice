package backendinterview

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Event struct {
	x   int
	val int
}

type Point struct {
	x   int
	idx int
}

func validatePointsAndSegmentsInput(p int) bool {
	return p >= 1 && p <= 100_000
}

// https://coderun.yandex.ru/selections/backend-interview/problems/points-and-segments
// PointsAndSegments - assignment 5
func PointsAndSegments() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first line
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	// N
	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	if !validatePointsAndSegmentsInput(n) {
		panic("N out of range")
	}

	// M
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	if !validatePointsAndSegmentsInput(m) {
		panic("M out of range")
	}

	events := make([]Event, 0, 2*n)
	// input segments
	for range n {
		line, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		strNum = strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		// A
		a, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}

		// B
		b, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		if a > b {
			a, b = b, a
		}

		events = append(events, Event{a, +1})
		events = append(events, Event{b + 1, -1})
	}

	// points
	points := make([]Point, m)

	// last line
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	strNum = strings.Fields(line)
	if len(strNum) != m {
		panic("numbers count does not match M")
	}

	// process points
	for i := range m {
		x, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		points[i].x = x
		points[i].idx = i
	}

	// sort events
	sort.Slice(events, func(i, j int) bool {
		return events[i].x < events[j].x
	})

	// sort points
	sort.Slice(points, func(i, j int) bool {
		return points[i].x < points[j].x
	})

	result := make([]int, m)

	cur := 0
	j := 0

	// sweep line
	for _, p := range points {
		for j < len(events) && events[j].x <= p.x {
			cur += events[j].val
			j++
		}
		result[p.idx] = cur
	}

	// output
	for i := 0; i < m; i++ {
		writer.WriteString(strconv.Itoa(result[i]) + " ")
	}
	writer.WriteByte('\n')
}
