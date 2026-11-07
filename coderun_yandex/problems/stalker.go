package problems

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

type EdgeStalker struct {
	to     int
	weight int
}

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	d := &DSU{
		parent: make([]int, n),
		size:   make([]int, n),
	}

	for i := 0; i < n; i++ {
		d.parent[i] = i
		d.size[i] = 1
	}

	return d
}

func (d *DSU) Find(v int) int {
	if d.parent[v] != v {
		d.parent[v] = d.Find(d.parent[v])
	}
	return d.parent[v]
}

func (d *DSU) Union(a, b int) {
	a = d.Find(a)
	b = d.Find(b)

	if a == b {
		return
	}

	if d.size[a] < d.size[b] {
		a, b = b, a
	}

	d.parent[b] = a
	d.size[a] += d.size[b]
}

func ZeroOneBFS(graph [][]EdgeStalker, start, target int) int {
	const INF = int(1e9)

	n := len(graph)

	dist := make([]int, n)

	for i := range dist {
		dist[i] = INF
	}

	dist[start] = 0

	deque := list.New()
	deque.PushFront(start)

	for deque.Len() > 0 {
		front := deque.Front()
		v := front.Value.(int)
		deque.Remove(front)

		for _, e := range graph[v] {
			newDist := dist[v] + e.weight

			if newDist >= dist[e.to] {
				continue
			}

			dist[e.to] = newDist

			if e.weight == 0 {
				deque.PushFront(e.to)
			} else {
				deque.PushBack(e.to)
			}
		}
	}

	if dist[target] == INF {
		return -1
	}

	return dist[target]
}

// https://coderun.yandex.ru/problem/stalker
// Stalker - problem 47
func Stalker() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	graph := make([][]EdgeStalker, n)
	for i := 0; i < k; i++ {
		var r int
		fmt.Fscan(reader, &r)

		dsu := NewDSU(n)

		used := make([]bool, n)
		touched := make([]int, 0)
		for j := 0; j < r; j++ {
			var a, b int
			fmt.Fscan(reader, &a, &b)

			a--
			b--

			dsu.Union(a, b)

			if !used[a] {
				used[a] = true
				touched = append(touched, a)
			}

			if !used[b] {
				used[b] = true
				touched = append(touched, b)
			}
		}

		compID := make([]int, n)
		for j := range compID {
			compID[j] = -1
		}

		for _, v := range touched {
			root := dsu.Find(v)

			if compID[root] == -1 {
				compID[root] = len(graph)
				graph = append(graph, nil)
			}

			component := compID[root]

			graph[v] = append(graph[v], EdgeStalker{
				to:     component,
				weight: 1,
			})

			graph[component] = append(
				graph[component],
				EdgeStalker{
					to:     v,
					weight: 0,
				},
			)
		}
	}

	answer := ZeroOneBFS(graph, 0, n-1)

	writer.WriteString(strconv.Itoa(answer))
	writer.WriteByte('\n')
}
