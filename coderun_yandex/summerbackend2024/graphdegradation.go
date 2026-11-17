package summerbackend2024

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	MOD1                 int64 = 1_000_000_007
	MOD2                 int64 = 1_000_000_009
	INF_GraphDegradation int64 = 1 << 60
)

type Edge struct {
	from int
	to   int
}

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/graph-degradation
// GraphDegradation - problem 33
func GraphDegradation() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	w := make([][]int64, n)
	for i := 0; i < n; i++ {
		w[i] = make([]int64, n)

		for j := 0; j < n; j++ {
			fmt.Fscan(in, &w[i][j])
		}
	}

	dist := make([]int64, n)
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		dist[i] = INF_GraphDegradation
	}

	dist[0] = 0
	for step := 0; step < n; step++ {
		u := -1
		best := INF_GraphDegradation

		for i := 0; i < n; i++ {
			if !used[i] && dist[i] < best {
				best = dist[i]
				u = i
			}
		}

		if u == -1 {
			break
		}

		used[u] = true

		for v := 0; v < n; v++ {
			if w[u][v] > 0 {
				nd := dist[u] + w[u][v]

				if nd < dist[v] {
					dist[v] = nd
				}
			}
		}
	}

	adj := make([][]int, n)
	revAdj := make([][]int, n)
	edges := make([]Edge, 0)
	for u := 0; u < n; u++ {
		if dist[u] == INF_GraphDegradation {
			continue
		}

		for v := 0; v < n; v++ {
			if u == v || w[u][v] <= 0 {
				continue
			}

			if dist[u]+w[u][v] == dist[v] {
				adj[u] = append(adj[u], v)
				revAdj[v] = append(revAdj[v], u)

				edges = append(edges, Edge{
					from: u,
					to:   v,
				})
			}
		}
	}

	reachable := make([]int, 0, n)
	for v := 0; v < n; v++ {
		if dist[v] != INF_GraphDegradation {
			reachable = append(reachable, v)
		}
	}

	sort.Slice(reachable, func(i, j int) bool {
		return dist[reachable[i]] < dist[reachable[j]]
	})

	cnt1 := make([]int64, n)
	cnt2 := make([]int64, n)
	cnt1[0] = 1
	cnt2[0] = 1
	for _, v := range reachable {
		if v == 0 {
			continue
		}

		for _, u := range revAdj[v] {
			cnt1[v] = (cnt1[v] + cnt1[u]) % MOD1
			cnt2[v] = (cnt2[v] + cnt2[u]) % MOD2
		}
	}

	cntTo1 := make([][]int64, n)
	cntTo2 := make([][]int64, n)
	for i := 0; i < n; i++ {
		cntTo1[i] = make([]int64, n)
		cntTo2[i] = make([]int64, n)

		cntTo1[i][i] = 1
		cntTo2[i][i] = 1
	}

	for i := len(reachable) - 1; i >= 0; i-- {
		v := reachable[i]

		for _, u := range adj[v] {
			for x := 0; x < n; x++ {
				cntTo1[v][x] =
					(cntTo1[v][x] + cntTo1[u][x]) % MOD1

				cntTo2[v][x] =
					(cntTo2[v][x] + cntTo2[u][x]) % MOD2
			}
		}
	}

	answer := 0
	for _, e := range edges {
		u := e.from
		v := e.to

		current := 0
		for _, x := range reachable {
			if x == 0 {
				continue
			}

			ways1 := cnt1[u] * cntTo1[v][x] % MOD1
			ways2 := cnt2[u] * cntTo2[v][x] % MOD2

			if ways1 == cnt1[x] &&
				ways2 == cnt2[x] {
				current++
			}
		}

		if current > answer {
			answer = current
		}
	}

	fmt.Fprintln(out, answer)
}
