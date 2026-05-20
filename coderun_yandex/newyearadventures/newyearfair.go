package newyearadventures

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func dfsNewYearFair(v int, g [][]int, used []bool) {
	used[v] = true

	for _, to := range g[v] {
		if !used[to] {
			dfsNewYearFair(to, g, used)
		}
	}
}

// https://coderun.yandex.ru/selections/new-year-adventures/problems/new-year-fair
// NewYearFair - problem 5
func NewYearFair() {
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

	// N
	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}

	// M
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// trips input
	g := make([][]int, n)
	for i := 0; i < m; i++ {
		// coordinate input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
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

		a--
		b--

		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	components := 0
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		if !used[i] {
			components++
			dfsNewYearFair(i, g, used)
		}
	}

	answer := m - (n - components)

	writer.WriteString(strconv.Itoa(answer))
	writer.WriteByte('\n')
}
