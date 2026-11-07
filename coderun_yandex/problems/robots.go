package problems

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Edge struct {
	u, v int
}

type Event struct {
	vertex int
	robot  int
}

// https://coderun.yandex.ru/problem/robots
// Robots - problem 46
func Robots() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	graph := make([][]int, n)
	edges := make([]Edge, k)

	for i := 0; i < k; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)

		u--
		v--

		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)

		edges[i] = Edge{u, v}
	}

	var m int
	fmt.Fscan(reader, &m)

	robots := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &robots[i])
		robots[i]--
	}

	dist := make([][]int, m)
	for r := 0; r < m; r++ {
		dist[r] = make([]int, 2*n)

		for i := range dist[r] {
			dist[r][i] = -1
		}

		start := robots[r] * 2

		queue := make([]int, 0, 2*n)
		queue = append(queue, start)
		dist[r][start] = 0
		for head := 0; head < len(queue); head++ {
			state := queue[head]

			v := state / 2
			parity := state % 2

			for _, to := range graph[v] {
				next := to*2 + (parity ^ 1)

				if dist[r][next] != -1 {
					continue
				}

				dist[r][next] = dist[r][state] + 1
				queue = append(queue, next)
			}
		}
	}

	events := make([][]Event, 2*n)
	for r := 0; r < m; r++ {
		for v := 0; v < n; v++ {
			for p := 0; p < 2; p++ {
				d := dist[r][v*2+p]

				if d == -1 {
					continue
				}

				events[d] = append(events[d], Event{
					vertex: v,
					robot:  r,
				})
			}
		}
	}

	words := (m + 63) / 64
	masks := make([][][]uint64, 2)
	for p := 0; p < 2; p++ {
		masks[p] = make([][]uint64, n)

		for v := 0; v < n; v++ {
			masks[p][v] = make([]uint64, words)
		}
	}

	full := make([]uint64, words)
	for i := range full {
		full[i] = ^uint64(0)
	}

	if m%64 != 0 {
		full[words-1] = (uint64(1) << (m % 64)) - 1
	}

	for t := 0; t < 2*n; t++ {
		p := t % 2

		for _, e := range events[t] {
			masks[p][e.vertex][e.robot/64] |=
				uint64(1) << (e.robot % 64)
		}

		for v := 0; v < n; v++ {
			ok := true

			for w := 0; w < words; w++ {
				if masks[p][v][w] != full[w] {
					ok = false
					break
				}
			}

			if ok {
				writer.WriteString(strconv.Itoa(t))
				writer.WriteByte('\n')
				return
			}
		}

		for _, e := range edges {
			ok := true

			for w := 0; w < words; w++ {
				combined := masks[p][e.u][w] |
					masks[p][e.v][w]

				if combined != full[w] {
					ok = false
					break
				}
			}

			if ok {
				writer.WriteString(strconv.Itoa(t))
				writer.WriteString(".5\n")
				return
			}
		}
	}

	writer.WriteString("-1\n")
}
