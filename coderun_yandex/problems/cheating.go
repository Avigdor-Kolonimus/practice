package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/cheating
// Cheating - problem 9
func Cheating() {
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
	graph := make([][]int, n+1)
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

		u, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		v, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// 0  -> no colored
	// 1  -> 1 group
	// -1 -> 2 group
	color := make([]int, n+1)
	for start := 1; start <= n; start++ {
		if color[start] != 0 {
			continue
		}

		stack := []int{start}
		color[start] = 1
		for len(stack) > 0 {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			for _, to := range graph[v] {
				if color[to] == 0 {
					color[to] = -color[v]
					stack = append(stack, to)
				} else if color[to] == color[v] {
					writer.WriteString("NO")
					writer.WriteByte('\n')

					return
				}
			}
		}
	}

	writer.WriteString("YES")
	writer.WriteByte('\n')
}
