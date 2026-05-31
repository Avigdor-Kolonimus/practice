package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/metro-2
// Metro2 - problem 16
func Metro2() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// M input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	m, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// lines input
	lines := make([][]int, m)
	// stationToLines[s] contains all metro lines passing through station s
	stationToLines := make([][]int, n+1)
	for i := 0; i < m; i++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum := strings.Fields(line)

		p, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}

		lines[i] = make([]int, p)

		for j := 0; j < p; j++ {
			station, err := strconv.Atoi(strNum[j+1])
			if err != nil {
				panic(err)
			}

			lines[i][j] = station
			stationToLines[station] = append(stationToLines[station], i)
		}
	}

	// A and B input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	a, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	if a == b {
		writer.WriteByte('0')
		writer.WriteByte('\n')

		return
	}

	visitedStations := make([]bool, n+1)
	visitedLines := make([]bool, m)

	queue := []int{a}
	visitedStations[a] = true

	transfers := 0
	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			station := queue[0]
			queue = queue[1:]

			if station == b {
				writer.WriteString(strconv.Itoa(transfers))
				writer.WriteByte('\n')

				return
			}

			for _, line := range stationToLines[station] {
				if visitedLines[line] {
					continue
				}

				visitedLines[line] = true

				for _, nextStation := range lines[line] {
					if visitedStations[nextStation] {
						continue
					}

					if nextStation == b {
						writer.WriteString(strconv.Itoa(transfers))
						writer.WriteByte('\n')

						return
					}

					visitedStations[nextStation] = true
					queue = append(queue, nextStation)
				}
			}
		}

		transfers++
	}

	writer.WriteString("-1")
	writer.WriteByte('\n')
}
