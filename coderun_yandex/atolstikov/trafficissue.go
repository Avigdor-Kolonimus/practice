package atolstikov

import (
	"bufio"
	"fmt"
	"os"
)

type EdgeTrafficIssue struct {
	to   int
	next int
}

type CondEdge struct {
	to   int
	w    int64
	next int
}

// https://coderun.yandex.ru/selections/atolstikov/problems/traffic-issue
// TrafficIssue - problem 12
func TrafficIssue() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	tax := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &tax[i])
	}

	head := make([]int, n+1)
	headR := make([]int, n+1)
	for i := range head {
		head[i] = -1
		headR[i] = -1
	}

	edges := make([]EdgeTrafficIssue, m)
	reverseEdges := make([]EdgeTrafficIssue, m)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)

		edges[i] = EdgeTrafficIssue{
			to:   v,
			next: head[u],
		}
		head[u] = i

		reverseEdges[i] = EdgeTrafficIssue{
			to:   u,
			next: headR[v],
		}
		headR[v] = i
	}

	visited := make([]bool, n+1)
	iter := make([]int, n+1)
	order := make([]int, 0, n)
	stack := make([]int, 0, n)
	for start := 1; start <= n; start++ {
		if visited[start] {
			continue
		}

		visited[start] = true
		iter[start] = head[start]
		stack = append(stack, start)

		for len(stack) > 0 {
			v := stack[len(stack)-1]
			e := iter[v]

			if e == -1 {
				order = append(order, v)
				stack = stack[:len(stack)-1]
				continue
			}

			iter[v] = edges[e].next
			u := edges[e].to

			if !visited[u] {
				visited[u] = true
				iter[u] = head[u]
				stack = append(stack, u)
			}
		}
	}

	comp := make([]int, n+1)
	compCount := 0
	compSize := make([]int, n+1)
	stack = stack[:0]
	for i := n - 1; i >= 0; i-- {
		start := order[i]

		if comp[start] != 0 {
			continue
		}

		compCount++
		comp[start] = compCount
		stack = append(stack, start)

		for len(stack) > 0 {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			compSize[compCount]++

			for e := headR[v]; e != -1; e = reverseEdges[e].next {
				u := reverseEdges[e].to

				if comp[u] == 0 {
					comp[u] = compCount
					stack = append(stack, u)
				}
			}
		}
	}

	headC := make([]int, compCount+1)
	indeg := make([]int, compCount+1)

	for i := range headC {
		headC[i] = -1
	}

	condEdges := make([]CondEdge, 0, m)
	for u := 1; u <= n; u++ {
		cu := comp[u]

		for e := head[u]; e != -1; e = edges[e].next {
			v := edges[e].to
			cv := comp[v]

			if cu == cv {
				continue
			}

			condEdges = append(condEdges, CondEdge{
				to:   cv,
				w:    tax[v],
				next: headC[cu],
			})

			headC[cu] = len(condEdges) - 1
			indeg[cv]++
		}
	}

	const INF int64 = 1 << 62
	dist := make([]int64, compCount+1)
	for i := range dist {
		dist[i] = INF
	}

	startComp := comp[1]
	endComp := comp[n]

	if compSize[startComp] == 1 {
		dist[startComp] = 0
	}

	for e := head[1]; e != -1; e = edges[e].next {
		v := edges[e].to
		cv := comp[v]

		if cv == startComp {
			if tax[v] < dist[startComp] {
				dist[startComp] = tax[v]
			}
		} else {
			if tax[v] < dist[cv] {
				dist[cv] = tax[v]
			}
		}
	}

	queue := make([]int, 0, compCount)
	for c := 1; c <= compCount; c++ {
		if indeg[c] == 0 {
			queue = append(queue, c)
		}
	}

	for q := 0; q < len(queue); q++ {
		u := queue[q]

		for e := headC[u]; e != -1; e = condEdges[e].next {
			v := condEdges[e].to
			w := condEdges[e].w

			if dist[u] != INF {
				nd := dist[u] + w

				if nd < dist[v] {
					dist[v] = nd
				}
			}

			indeg[v]--

			if indeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	fmt.Fprintln(out, dist[endComp])
}
