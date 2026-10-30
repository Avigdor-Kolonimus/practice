package backend

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type EdgeGraphSplit struct {
	u, v, w int
}

type DSUGraphSplit struct {
	parent []int
	size   []int
	parity []int
}

func NewDSUGraphSplit(n int) *DSUGraphSplit {
	d := &DSUGraphSplit{
		parent: make([]int, n+1),
		size:   make([]int, n+1),
		parity: make([]int, n+1),
	}

	for i := 1; i <= n; i++ {
		d.parent[i] = i
		d.size[i] = 1
	}

	return d
}

func (d *DSUGraphSplit) find(v int) (int, int) {
	if d.parent[v] == v {
		return v, 0
	}

	p := d.parent[v]
	root, parity := d.find(p)

	d.parity[v] ^= parity
	d.parent[v] = root

	return root, d.parity[v]
}

func (d *DSUGraphSplit) add(u, v int) bool {
	ru, pu := d.find(u)
	rv, pv := d.find(v)

	if ru == rv {
		return pu^pv == 1
	}

	if d.size[ru] < d.size[rv] {
		ru, rv = rv, ru
		pu, pv = pv, pu
	}

	d.parent[rv] = ru
	d.parity[rv] = pu ^ pv ^ 1
	d.size[ru] += d.size[rv]

	return true
}

// https://coderun.yandex.ru/selections/first-2023-backend/problems/graph-split
// GraphSplit - problem 4
func GraphSplit() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	edges := make([]EdgeGraphSplit, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &edges[i].u, &edges[i].v, &edges[i].w)
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	dsu := NewDSUGraphSplit(n)

	for _, e := range edges {
		if !dsu.add(e.u, e.v) {
			fmt.Fprintln(writer, e.w)
			return
		}
	}
}
