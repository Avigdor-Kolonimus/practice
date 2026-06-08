package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type EdgeConnectivityComponents struct {
	v1 int
	v2 int
}

func dfs(start int, graph [][]int, visited []bool) []int {
	stack := []int{start}
	visited[start] = true
	component := []int{}
	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		component = append(component, v)

		for _, to := range graph[v] {
			if !visited[to] {
				visited[to] = true
				stack = append(stack, to)
			}
		}
	}

	return component
}

// https://coderun.yandex.ru/problem/connectivity-components
// ConnectivityComponents - problem 8
func ConnectivityComponents() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// edges
	edges := make([]EdgeConnectivityComponents, m)
	for i := 0; i < m; i++ {
		// edge input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		x, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		edges[i] = EdgeConnectivityComponents{
			v1: x,
			v2: y,
		}
	}

	graph := make([][]int, n+1)
	for _, e := range edges {
		graph[e.v1] = append(graph[e.v1], e.v2)
		graph[e.v2] = append(graph[e.v2], e.v1)
	}

	var components [][]int
	visited := make([]bool, n+1)
	for v := 1; v <= n; v++ {
		if !visited[v] {
			component := dfs(v, graph, visited)
			components = append(components, component)
		}
	}

	writer.WriteString(strconv.Itoa(len(components)))
	writer.WriteByte('\n')

	for _, comp := range components {
		writer.WriteString(strconv.Itoa(len(comp)))
		writer.WriteByte('\n')

		for _, v := range comp {
			writer.WriteString(strconv.Itoa(v))
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')
	}
}
