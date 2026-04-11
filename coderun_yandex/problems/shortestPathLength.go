package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateSPLInput(n int) bool {
	return n >= 1 && n <= 100
}

func bfs(matrix [][]int, n, start, end int) int {
	startIndex := start - 1
	endIndex := end - 1

	dist := make([]int, n)
	for i := range dist {
		dist[i] = -1
	}

	queue := []int{startIndex}
	dist[startIndex] = 0

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		if v == endIndex {
			break
		}

		for i := range n {
			if matrix[v][i] == 1 && dist[i] == -1 {
				dist[i] = dist[v] + 1
				queue = append(queue, i)
			}
		}
	}

	return dist[endIndex]
}

// https://coderun.yandex.ru/problem/shortest-path-length
// ShortestPathLength - problem 3
func ShortestPathLength() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first line
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	n, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}
	if !validateSPLInput(n) {
		panic("n out of range")
	}

	// map input
	matrix := make([][]int, n)
	for i := range n {
		matrix[i] = make([]int, n)

		// line input
		line, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		matrixLine := strings.Fields(line)
		mlCount := len(matrixLine)
		if mlCount != n {
			panic("matrix line count does not match n")
		}
		for j, v := range matrixLine {
			conn, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			if conn != 0 && conn != 1 {
				panic("illegal parameter")
			}

			matrix[i][j] = conn
		}
	}

	// last line
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	start, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}

	end, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	spl := bfs(matrix, n, start, end)

	writer.WriteString(strconv.Itoa(spl))
	writer.WriteByte('\n')
}
