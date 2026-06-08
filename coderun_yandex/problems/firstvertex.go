package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/first-vertex
// FirstVertex - problem 39
func FirstVertex() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M line
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

	// inverse graph
	rev := make([][]int, n+1)
	for i := 0; i < m; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum := strings.Fields(line)
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

		rev[v] = append(rev[v], u) // inverse
	}

	// BFS from 1
	visited := make([]bool, n+1)
	visited[1] = true
	queue := []int{1}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		for _, to := range rev[v] {
			if !visited[to] {
				visited[to] = true
				queue = append(queue, to)
			}
		}
	}

	// output
	for i := 1; i <= n; i++ {
		if visited[i] {
			writer.WriteString(strconv.Itoa(i))
			writer.WriteByte(' ')
		}
	}
	writer.WriteByte('\n')
}
