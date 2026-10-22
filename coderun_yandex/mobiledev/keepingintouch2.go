package mobiledev

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INF = int(1e9)
	V   = 86400
)

type SegmentTree struct {
	size int
	tree []int
}

func NewSegmentTree(n int) *SegmentTree {
	size := 1
	for size < n {
		size <<= 1
	}

	tree := make([]int, size*2)
	for i := range tree {
		tree[i] = INF
	}

	return &SegmentTree{size: size, tree: tree}
}

func (st *SegmentTree) Update(pos, value int) {
	i := pos + st.size
	st.tree[i] = value

	for i >>= 1; i > 0; i >>= 1 {
		st.tree[i] = min(st.tree[i<<1], st.tree[i<<1|1])
	}
}

func (st *SegmentTree) Query(l, r int) int {
	if l > r {
		return INF
	}

	l += st.size
	r += st.size
	res := INF

	for l <= r {
		if l&1 == 1 {
			res = min(res, st.tree[l])
			l++
		}
		if r&1 == 0 {
			res = min(res, st.tree[r])
			r--
		}
		l >>= 1
		r >>= 1
	}

	return res
}

func parseTime(s string) int {
	h := int(s[0]-'0')*10 + int(s[1]-'0')
	m := int(s[3]-'0')*10 + int(s[4]-'0')
	sec := int(s[6]-'0')*10 + int(s[7]-'0')

	return h*3600 + m*60 + sec
}

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/keeping-in-touch-2
// KeepingInTouch2 - problem 29
func KeepingInTouch2() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20)
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	t := make([]int, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		t[i] = parseTime(s)
	}

	segments := make([][]int, 0)
	start := 0
	for i := 1; i < n; i++ {
		if t[i] <= t[i-1] {
			segments = append(segments, t[start:i])
			start = i
		}
	}
	segments = append(segments, t[start:n])

	if len(segments) == 1 {
		fmt.Fprintln(out, n-1)
		return
	}

	st := NewSegmentTree(V)

	prev := segments[0]
	dp := make([]int, len(prev))
	for i, x := range prev {
		dp[i] = 1
		st.Update(x, 1)
	}

	active := append([]int(nil), prev...)
	for day := 1; day < len(segments); day++ {
		cur := segments[day]
		next := make([]int, len(cur))

		prefix := INF

		for i, x := range cur {
			best := st.Query(x, V-1)
			if best < INF {
				best++
			}

			if prefix < INF {
				best = min(best, prefix+1)
			}

			next[i] = best
			prefix = min(prefix, best)
		}

		for _, x := range active {
			st.Update(x, INF)
		}

		active = active[:0]

		for i, x := range cur {
			if next[i] < INF {
				st.Update(x, next[i])
				active = append(active, x)
			}
		}

		prev = cur
		dp = next
	}

	minKept := INF
	for _, x := range dp {
		minKept = min(minKept, x)
	}

	fmt.Fprintln(out, n-minKept)
}
