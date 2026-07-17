package backend

import (
	"bufio"
	"fmt"
	"os"
)

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n+1)
	size := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &DSU{
		parent: parent,
		size:   size,
	}
}

func (d *DSU) Find(v int) int {
	if d.parent[v] != v {
		d.parent[v] = d.Find(d.parent[v])
	}
	return d.parent[v]
}

func (d *DSU) Union(a, b int) bool {
	a = d.Find(a)
	b = d.Find(b)

	if a == b {
		return false
	}

	if d.size[a] < d.size[b] {
		a, b = b, a
	}

	d.parent[b] = a
	d.size[a] += d.size[b]
	return true
}

type EdgeDsuYaIntern struct {
	u, v int
}

// https://coderun.yandex.ru/selections/backend/problems/dsu-ya-intern
// DsuYaIntern - problem 33
func DsuYaIntern() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	edges := make([]EdgeDsuYaIntern, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscan(reader, &edges[i].u, &edges[i].v)
	}

	var q int
	fmt.Fscan(reader, &q)

	queries := make([]int, q)
	removed := make([]bool, m+1)

	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &queries[i])
		removed[queries[i]] = true
	}

	dsu := NewDSU(n)
	components := n
	for i := 1; i <= m; i++ {
		if removed[i] {
			continue
		}
		if dsu.Union(edges[i].u, edges[i].v) {
			components--
		}
	}

	ans := make([]int, q)

	for i := q - 1; i >= 0; i-- {
		ans[i] = components

		e := edges[queries[i]]
		if dsu.Union(e.u, e.v) {
			components--
		}
	}

	for _, x := range ans {
		fmt.Fprint(writer, x, " ")
	}
}
