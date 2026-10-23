package first2023dataanalytics

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Segment struct {
	l, r int64
	c    int
}

const INF int64 = 1 << 62

type SegTree struct {
	n    int
	tree []int64
}

func NewSegTree(n int) *SegTree {
	size := 1
	for size < n {
		size <<= 1
	}

	tree := make([]int64, 2*size)
	for i := range tree {
		tree[i] = INF
	}

	return &SegTree{n: size, tree: tree}
}

func (st *SegTree) Update(pos int, val int64) {
	pos += st.n
	st.tree[pos] = val

	for pos >>= 1; pos > 0; pos >>= 1 {
		a, b := st.tree[pos*2], st.tree[pos*2+1]
		if a < b {
			st.tree[pos] = a
		} else {
			st.tree[pos] = b
		}
	}
}

func (st *SegTree) Query(l, r int) int64 {
	res := INF
	l += st.n
	r += st.n

	for l < r {
		if l&1 == 1 {
			if st.tree[l] < res {
				res = st.tree[l]
			}
			l++
		}
		if r&1 == 1 {
			r--
			if st.tree[r] < res {
				res = st.tree[r]
			}
		}
		l >>= 1
		r >>= 1
	}

	return res
}

// https://coderun.yandex.ru/selections/first-2023-data-analytics/problems/cuts
// Cuts - problem 38
func Cuts() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	a := make([]Segment, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i].l, &a[i].r, &a[i].c)
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].l < a[j].l
	})

	st := NewSegTree(n + 1)

	for _, cur := range a {
		leftMin := st.Query(0, cur.c)
		rightMin := st.Query(cur.c+1, n+1)

		minOther := leftMin
		if rightMin < minOther {
			minOther = rightMin
		}

		if minOther < cur.r {
			writer.WriteString("NO\n")
			return
		}

		old := st.Query(cur.c, cur.c+1)
		if cur.r < old {
			st.Update(cur.c, cur.r)
		}
	}

	writer.WriteString("YES\n")
}
