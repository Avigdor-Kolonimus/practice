package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/topological-sorting
// TopologicalSorting - problem 10
func TopologicalSorting() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	tokens := strings.Fields(line)
	if len(tokens) != 2 {
		panic("invalid input")
	}

	n, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}

	// graph input
	graph := make([][]int, n+1)
	indegree := make([]int, n+1)
	for i := 0; i < m; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		tokens := strings.Fields(line)
		if len(tokens) != 2 {
			panic("invalid input")
		}

		u, err := strconv.Atoi(tokens[0])
		if err != nil {
			panic(err)
		}
		v, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}

		graph[u] = append(graph[u], v)
		indegree[v]++
	}

	queue := make([]int, 0)
	for v := 1; v <= n; v++ {
		if indegree[v] == 0 {
			queue = append(queue, v)
		}
	}

	order := make([]int, 0, n)
	head := 0
	for head < len(queue) {
		v := queue[head]
		head++

		order = append(order, v)

		for _, to := range graph[v] {
			indegree[to]--

			if indegree[to] == 0 {
				queue = append(queue, to)
			}
		}
	}

	if len(order) != n {
		writer.WriteString(strconv.Itoa(-1))
		writer.WriteByte('\n')

		return
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			writer.WriteByte(' ')
		}
		writer.WriteString(strconv.Itoa(order[i]))
	}

	writer.WriteByte('\n')
}
