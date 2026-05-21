package algorithmtrainingseptember2025

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/from-dead-end-to-dead-end
// FromDeadEndToDeadEnd - assignment 22
func FromDeadEndToDeadEnd() {
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

	// edge input
	g := make([][]int, n+1)
	deg := make([]int, n+1)
	for range n - 1 {
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

		g[u] = append(g[u], v)
		g[v] = append(g[v], u)

		deg[u]++
		deg[v]++
	}

	dist := make([]int, n+1)
	owner := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = -1
	}

	q := make([]int, 0)

	for i := 1; i <= n; i++ {
		if deg[i] == 1 {
			dist[i] = 0
			owner[i] = i
			q = append(q, i)
		}
	}

	ans := int(1e9)
	head := 0
	for head < len(q) {
		v := q[head]
		head++

		for _, to := range g[v] {
			if dist[to] == -1 {
				dist[to] = dist[v] + 1
				owner[to] = owner[v]
				q = append(q, to)
			} else if owner[to] != owner[v] {
				cand := dist[to] + dist[v] + 1
				ans = min(ans, cand)
			}
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
