package problems

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/party
// Party - problem 408
func Party() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	data, _ := reader.ReadString(0)
	fields := strings.Fields(data)

	n, _ := strconv.Atoi(fields[0])

	a := make([]int, n)
	graph := make([][]int, n)
	indegree := make([]int, n)
	byCount := make([]int, 0, n)

	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(fields[i+1])

		if a[i] >= 0 {
			byCount = append(byCount, i)
		} else {
			favorite := -a[i] - 1
			graph[favorite] = append(graph[favorite], i)
			indegree[i]++
		}
	}

	sort.Slice(byCount, func(i, j int) bool {
		return a[byCount[i]] < a[byCount[j]]
	})

	queue := make([]int, 0, n)

	for i := 0; i < n; i++ {
		if a[i] < 0 && indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	answer := make([]int, 0, n)
	used := make([]bool, n)

	count, pos, ptr := 0, 0, 0

	for count < n {
		for pos < len(byCount) && a[byCount[pos]] <= count {
			queue = append(queue, byCount[pos])
			pos++
		}

		if ptr >= len(queue) {
			break
		}

		v := queue[ptr]
		ptr++

		if used[v] {
			continue
		}

		used[v] = true
		answer = append(answer, v+1)
		count++

		for _, u := range graph[v] {
			indegree[u]--

			if indegree[u] == 0 {
				queue = append(queue, u)
			}
		}
	}

	if len(answer) != n {
		writer.WriteString("0\n")
		return
	}

	var sb strings.Builder
	for i, v := range answer {
		if i > 0 {
			sb.WriteByte(' ')
		}
		sb.WriteString(strconv.Itoa(v))
	}
	sb.WriteByte('\n')

	writer.WriteString(sb.String())
}
