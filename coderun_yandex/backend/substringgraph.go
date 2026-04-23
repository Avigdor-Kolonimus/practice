package backend

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	From string
	To   string
}

func validateSubstringGraphTInput(n int) bool {
	return n >= 1 && n <= 40_000
}

// https://coderun.yandex.ru/selections/backend/problems/substring-graph
// SubstringGraph - problem 54
func SubstringGraph() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// T input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	t, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateSubstringGraphTInput(t) {
		panic("number T out of range")
	}

	graph := make(map[Edge]int)
	uniqueNodes := make(map[string]struct{})
	for range t {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		word := strings.TrimSpace(line)

		for idx := range len(word) - 3 {
			edge := Edge{
				From: word[idx : idx+3],
				To:   word[idx+1 : idx+4],
			}

			graph[edge] += 1
			uniqueNodes[edge.From] = struct{}{}
			uniqueNodes[edge.To] = struct{}{}
		}
	}

	writer.WriteString(strconv.Itoa(len(uniqueNodes)))
	writer.WriteByte('\n')

	writer.WriteString(strconv.Itoa(len(graph)))
	writer.WriteByte('\n')

	for edge, weight := range graph {
		writer.WriteString(edge.From)
		writer.WriteByte(' ')
		writer.WriteString(edge.To)
		writer.WriteByte(' ')
		writer.WriteString(strconv.Itoa(weight))
		writer.WriteByte('\n')
	}
}
