package backend

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/backend/problems/two-cliques
// TwoCliques - problem 27
func TwoCliques() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	line, _ := reader.ReadString('\n')
	first := strings.Fields(line)
	n, _ := strconv.Atoi(first[0])
	m, _ := strconv.Atoi(first[1])

	// Edges input
	adj := make([]map[int]bool, n+1)
	for i := 1; i <= n; i++ {
		adj[i] = make(map[int]bool)
	}
	for range m {
		line, _ = reader.ReadString('\n')
		fields := strings.Fields(line)
		u, _ := strconv.Atoi(fields[0])
		v, _ := strconv.Atoi(fields[1])
		adj[u][v] = true
		adj[v][u] = true
	}

	color := make([]int, n+1)
	for i := range color {
		color[i] = -1
	}

	queue := make([]int, 0, n)

	for start := 1; start <= n; start++ {
		if color[start] != -1 {
			continue
		}

		color[start] = 0
		queue = queue[:0]
		queue = append(queue, start)

		for head := 0; head < len(queue); head++ {
			v := queue[head]

			for u := 1; u <= n; u++ {
				if u == v || adj[v][u] {
					continue
				}

				switch color[u] {
				case -1:
					color[u] = color[v] ^ 1
					queue = append(queue, u)
				case color[v]:
					writer.WriteString("-1")
					writer.WriteByte('\n')

					return
				}
			}
		}
	}

	part1 := make([]int, 0)
	part2 := make([]int, 0)

	for i := 1; i <= n; i++ {
		if color[i] == 0 {
			part1 = append(part1, i)
		} else {
			part2 = append(part2, i)
		}
	}

	if len(part1) == 0 || len(part2) == 0 {
		writer.WriteString("1\n")
		writer.WriteString("1\n")
		for i := 2; i <= n; i++ {
			if i > 2 {
				writer.WriteByte(' ')
			}
			writer.WriteString(strconv.Itoa(i))
		}
		writer.WriteByte('\n')
		return
	}

	writer.WriteString(strconv.Itoa(len(part1)))
	writer.WriteByte('\n')

	for i, v := range part1 {
		if i > 0 {
			writer.WriteByte(' ')
		}
		writer.WriteString(strconv.Itoa(v))
	}
	writer.WriteByte('\n')

	for i, v := range part2 {
		if i > 0 {
			writer.WriteByte(' ')
		}
		writer.WriteString(strconv.Itoa(v))
	}
	writer.WriteByte('\n')
}
