package problems

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func validateSearchInDepthNInput(n int) bool {
	return n >= 1 && n <= 1_000
}

func validateSearchInDepthMInput(m int) bool {
	return m >= 0 && m <= 500_000
}

func intsToString(arr []int) string {
	var sb strings.Builder

	for i, v := range arr {
		if i > 0 {
			sb.WriteByte(' ')
		}
		sb.WriteString(strconv.Itoa(v))
	}

	return sb.String()
}

func dfs(component map[int]bool, grid [][]int, v int) {
	component[v] = true

	for _, to := range grid[v] {
		if !component[to] {
			dfs(component, grid, to)
		}
	}
}

// https://coderun.yandex.ru/selections/hr-tech-interview/problems/search-in-depth
// SearchInDepth - problem 9
func SearchInDepth() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first line
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	nm := strings.Fields(line)
	if len(nm) != 2 {
		panic("nm count does not match 2")
	}

	n, err := strconv.Atoi(strings.TrimSpace(nm[0]))
	if err != nil {
		panic(err)
	}
	if !validateSearchInDepthNInput(n) {
		panic("n out of range")
	}

	m, err := strconv.Atoi(strings.TrimSpace(nm[1]))
	if err != nil {
		panic(err)
	}
	if !validateSearchInDepthMInput(m) {
		panic("m out of range")
	}

	// grid input
	grid := make([][]int, n+1)
	for range m {
		line, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		uv := strings.Fields(line)
		if len(uv) != 2 {
			panic("nm count does not match 2")
		}

		u, err := strconv.Atoi(strings.TrimSpace(uv[0]))
		if err != nil {
			panic(err)
		}

		v, err := strconv.Atoi(strings.TrimSpace(uv[1]))
		if err != nil {
			panic(err)
		}

		grid[u] = append(grid[u], v)
		grid[v] = append(grid[v], u)
	}

	component := make(map[int]bool)
	dfs(component, grid, 1)

	nodesInComponent := make([]int, 0, len(component))
	for v := range component {
		nodesInComponent = append(nodesInComponent, v)
	}

	sort.Ints(nodesInComponent)
	result := intsToString(nodesInComponent)

	writer.WriteString(strconv.Itoa(len(component)))
	writer.WriteByte('\n')

	writer.WriteString(result)
	writer.WriteByte('\n')
}
