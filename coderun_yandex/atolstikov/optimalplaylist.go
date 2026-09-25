package atolstikov

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Edge struct {
	to int
	w  int
}

type Frame struct {
	v   int
	idx int
}

var (
	g     [][]Edge
	rg    [][]int
	used  []bool
	comp  []int
	order []int
)

// https://coderun.yandex.ru/selections/atolstikov/problems/optimal-playlist
// OptimalPlaylist - problem 5
func OptimalPlaylist() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20)
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	g = make([][]Edge, n)
	rg = make([][]int, n)

	weights := make([]int, 0, m)

	for i := 0; i < m; i++ {
		var a, b, w int
		fmt.Fscan(in, &a, &b, &w)
		a--
		b--

		g[a] = append(g[a], Edge{to: b, w: w})
		weights = append(weights, w)
	}

	if n == 1 {
		fmt.Fprintln(out, 0)
		return
	}

	slices.Sort(weights)
	weights = slices.Compact(weights)

	if !check(weights[len(weights)-1], n) {
		fmt.Fprintln(out, -1)
		return
	}

	left, right := 0, len(weights)-1

	for left < right {
		mid := left + (right-left)/2

		if check(weights[mid], n) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	fmt.Fprintln(out, weights[left])
}

func check(limit, n int) bool {
	used = make([]bool, n)
	comp = make([]int, n)
	order = order[:0]

	for start := 0; start < n; start++ {
		if used[start] {
			continue
		}

		used[start] = true
		stack := []Frame{{v: start}}

		for len(stack) > 0 {
			last := &stack[len(stack)-1]

			if last.idx < len(g[last.v]) {
				e := g[last.v][last.idx]
				last.idx++

				if e.w <= limit && !used[e.to] {
					used[e.to] = true
					stack = append(stack, Frame{v: e.to})
				}
			} else {
				order = append(order, last.v)
				stack = stack[:len(stack)-1]
			}
		}
	}

	for i := range rg {
		rg[i] = rg[i][:0]
	}

	for u := 0; u < n; u++ {
		for _, e := range g[u] {
			if e.w <= limit {
				rg[e.to] = append(rg[e.to], u)
			}
		}
	}

	clear(used)

	componentCount := 0

	for i := len(order) - 1; i >= 0; i-- {
		start := order[i]

		if used[start] {
			continue
		}

		used[start] = true
		comp[start] = componentCount

		stack := []int{start}

		for len(stack) > 0 {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			for _, to := range rg[v] {
				if !used[to] {
					used[to] = true
					comp[to] = componentCount
					stack = append(stack, to)
				}
			}
		}

		componentCount++
	}

	if componentCount == 1 {
		return true
	}

	dag := make([][]int, componentCount)
	indeg := make([]int, componentCount)

	for u := 0; u < n; u++ {
		cu := comp[u]

		for _, e := range g[u] {
			if e.w > limit {
				continue
			}

			cv := comp[e.to]

			if cu != cv {
				dag[cu] = append(dag[cu], cv)
				indeg[cv]++
			}
		}
	}

	queue := make([]int, 0, componentCount)

	for v := 0; v < componentCount; v++ {
		if indeg[v] == 0 {
			queue = append(queue, v)
		}
	}

	topo := make([]int, 0, componentCount)

	for head := 0; head < len(queue); head++ {
		v := queue[head]
		topo = append(topo, v)

		for _, to := range dag[v] {
			indeg[to]--
			if indeg[to] == 0 {
				queue = append(queue, to)
			}
		}
	}

	for i := 0; i+1 < len(topo); i++ {
		u := topo[i]
		v := topo[i+1]

		found := false

		for _, to := range dag[u] {
			if to == v {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}
