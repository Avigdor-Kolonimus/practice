package summercommon2026

import (
	"bufio"
	"os"
	"sort"
)

type Cell struct {
	val int
	r   int
	c   int
}

type Query struct {
	r1, c1 int
	r2, c2 int
	x      int
	id     int
}

type BIT2D struct {
	n, m int
	t    [][]int
}

func NewBIT2D(n, m int) *BIT2D {
	t := make([][]int, n+1)
	for i := range t {
		t[i] = make([]int, m+1)
	}
	return &BIT2D{
		n: n,
		m: m,
		t: t,
	}
}

func (b *BIT2D) Add(x, y int) {
	for i := x; i <= b.n; i += i & -i {
		for j := y; j <= b.m; j += j & -j {
			b.t[i][j]++
		}
	}
}

func (b *BIT2D) Sum(x, y int) int {
	res := 0
	for i := x; i > 0; i -= i & -i {
		for j := y; j > 0; j -= j & -j {
			res += b.t[i][j]
		}
	}
	return res
}

func (b *BIT2D) Query(r1, c1, r2, c2 int) int {
	return b.Sum(r2, c2) -
		b.Sum(r1-1, c2) -
		b.Sum(r2, c1-1) +
		b.Sum(r1-1, c1-1)
}

// https://coderun.yandex.ru/seasons/2026-summer/tracks/common/problem/the-map-of-hot-spots
// TheMapOfHotSpots - problem 5
func TheMapOfHotSpots() {
	sc := NewFastScanner()
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	T := sc.NextInt()

	for ; T > 0; T-- {
		n := sc.NextInt()
		m := sc.NextInt()
		q := sc.NextInt()

		cells := make([]Cell, 0, n*m)
		for i := 1; i <= n; i++ {
			for j := 1; j <= m; j++ {
				cells = append(cells, Cell{
					val: sc.NextInt(),
					r:   i,
					c:   j,
				})
			}
		}

		queries := make([]Query, q)
		for i := 0; i < q; i++ {
			queries[i] = Query{
				r1: sc.NextInt(),
				c1: sc.NextInt(),
				r2: sc.NextInt(),
				c2: sc.NextInt(),
				x:  sc.NextInt(),
				id: i,
			}
		}

		sort.Slice(cells, func(i, j int) bool {
			return cells[i].val > cells[j].val
		})

		sort.Slice(queries, func(i, j int) bool {
			return queries[i].x > queries[j].x
		})

		bit := NewBIT2D(n, m)
		ans := make([]int, q)

		ptr := 0

		for _, query := range queries {
			for ptr < len(cells) && cells[ptr].val >= query.x {
				bit.Add(cells[ptr].r, cells[ptr].c)
				ptr++
			}

			ans[query.id] = bit.Query(query.r1, query.c1, query.r2, query.c2)
		}

		for _, x := range ans {
			writeInt(out, x)
		}
	}
}
