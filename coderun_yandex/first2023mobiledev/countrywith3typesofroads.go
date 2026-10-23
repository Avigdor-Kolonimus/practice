package first2023mobiledev

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/country-with-3-types-of-roads
// CountryWith3TypesOfRoads - problem 3
func CountryWith3TypesOfRoads() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	readInt := func() int {
		line, _ := reader.ReadString('\n')
		n, _ := strconv.Atoi(strings.TrimSpace(line))
		return n
	}

	readEdge := func() (int, int, byte) {
		line, _ := reader.ReadString('\n')
		fields := strings.Fields(line)

		src, _ := strconv.Atoi(fields[0])
		dst, _ := strconv.Atoi(fields[1])

		return src, dst, fields[2][0]
	}

	line, _ := reader.ReadString('\n')
	fields := strings.Fields(line)
	n, _ := strconv.Atoi(fields[0])
	m, _ := strconv.Atoi(fields[1])

	graph := make([][3][]int, n+1)

	for i := 0; i < m; i++ {
		src, dst, color := readEdge()

		switch color {
		case 'g':
			graph[src][0] = append(graph[src][0], dst)
		case 'y':
			graph[src][1] = append(graph[src][1], dst)
		case 'r':
			graph[src][2] = append(graph[src][2], dst)
		}
	}

	getResult := func(src int) []int64 {
		result := make([][]int64, 4)
		for i := range result {
			result[i] = make([]int64, n+1)
		}
		result[0][src] = 1

		queue := []int{src}

		for step := 0; step < 3; step++ {
			nextQueue := make([]int, 0)

			for _, city := range queue {
				for _, nextCity := range graph[city][step] {
					if result[step+1][nextCity] == 0 {
						nextQueue = append(nextQueue, nextCity)
					}
					result[step+1][nextCity] += result[step][city]
				}
			}

			queue = nextQueue
		}

		return result[3]
	}

	q := readInt()
	cache := make(map[int][]int64)
	var output strings.Builder

	for range q {
		line, _ := reader.ReadString('\n')
		fields := strings.Fields(line)

		src, _ := strconv.Atoi(fields[0])
		dst, _ := strconv.Atoi(fields[1])

		if _, ok := cache[src]; !ok {
			cache[src] = getResult(src)
		}

		output.WriteString(strconv.FormatInt(cache[src][dst], 10))
		output.WriteByte('\n')
	}

	writer.WriteString(output.String())
}
