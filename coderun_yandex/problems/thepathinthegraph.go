package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/the-path-in-the-graph
// ThePathInTheGraph - problem 13
func ThePathInTheGraph() {
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

	// graph input
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		tokens := strings.Fields(line)
		if len(tokens) != n {
			panic("invalid input")
		}

		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			x, err := strconv.Atoi(tokens[j])
			if err != nil {
				panic(err)
			}

			graph[i][j] = x
		}
	}

	// start and end point
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	tokens := strings.Fields(line)
	if len(tokens) != 2 {
		panic("invalid input")
	}
	start, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}

	start--
	end--

	// BFS
	dist := make([]int, n)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1
		parent[i] = -1
	}

	queue := []int{start}
	dist[start] = 0
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		for to := 0; to < n; to++ {
			if graph[v][to] == 1 && dist[to] == -1 {
				dist[to] = dist[v] + 1
				parent[to] = v
				queue = append(queue, to)
			}
		}
	}

	if dist[end] == -1 {
		writer.WriteString(strconv.Itoa(-1))
		writer.WriteByte('\n')

		return
	}

	writer.WriteString(strconv.Itoa(dist[end]))
	writer.WriteByte('\n')

	if dist[end] == 0 {
		return
	}

	// restore path
	path := []int{}
	for v := end; v != -1; v = parent[v] {
		path = append(path, v)
	}

	// reverse
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	for _, v := range path {
		writer.WriteString(strconv.Itoa(v+1) + " ")
	}
	writer.WriteByte('\n')
}
