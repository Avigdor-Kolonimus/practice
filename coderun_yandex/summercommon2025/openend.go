package summercommon2025

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

const INF = 1_000_000_000

type Query struct {
	l, r, k int
	id      int
}

type Pair struct {
	g, r int
}

type Event struct {
	g, l, r int
}

type QueryKey struct {
	l, r, k int
}

var answerMap map[QueryKey]int64

// -------------------- GCD --------------------

func gcdOpenEnd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// -------------------- Segment Tree --------------------

type SegmentTree struct {
	n   int
	mx  []int
	sum []int64
}

func NewSegmentTree(n int) *SegmentTree {
	st := &SegmentTree{
		n:   n,
		mx:  make([]int, 4*n+5),
		sum: make([]int64, 4*n+5),
	}

	if n > 0 {
		st.build(1, 0, n-1)
	}

	return st
}

func (st *SegmentTree) build(v, tl, tr int) {
	if tl == tr {
		st.mx[v] = tl - 1
		st.sum[v] = int64(tl - 1)
		return
	}

	tm := (tl + tr) / 2

	st.build(v*2, tl, tm)
	st.build(v*2+1, tm+1, tr)

	st.pull(v)
}

func (st *SegmentTree) pull(v int) {
	st.mx[v] = max(st.mx[v*2], st.mx[v*2+1])
	st.sum[v] = st.sum[v*2] + st.sum[v*2+1]
}

func (st *SegmentTree) update(v, tl, tr, pos, val int) {
	if tl == tr {
		st.mx[v] = val
		st.sum[v] = int64(val)
		return
	}

	tm := (tl + tr) / 2

	if pos <= tm {
		st.update(v*2, tl, tm, pos, val)
	} else {
		st.update(v*2+1, tm+1, tr, pos, val)
	}

	st.pull(v)
}

func (st *SegmentTree) Update(pos, val int) {
	st.update(1, 0, st.n-1, pos, val)
}

// Find the first position whose value is >= val.
func (st *SegmentTree) goFirst(v, tl, tr, val int) int {
	if st.mx[v] < val {
		return -1
	}

	if tl == tr {
		return tl
	}

	tm := (tl + tr) / 2

	res := st.goFirst(v*2, tl, tm, val)

	if res != -1 {
		return res
	}

	return st.goFirst(v*2+1, tm+1, tr, val)
}

func (st *SegmentTree) GoFirst(val int) int {
	return st.goFirst(1, 0, st.n-1, val)
}

func (st *SegmentTree) get(v, tl, tr, l, r int) int64 {
	if l > r {
		return 0
	}

	if l == tl && r == tr {
		return st.sum[v]
	}

	tm := (tl + tr) / 2

	left := st.get(
		v*2,
		tl,
		tm,
		l,
		min(r, tm),
	)

	right := st.get(
		v*2+1,
		tm+1,
		tr,
		max(l, tm+1),
		r,
	)

	return left + right
}

func (st *SegmentTree) Get(l, r int) int64 {
	if l > r {
		return 0
	}

	return st.get(1, 0, st.n-1, l, r)
}

// -------------------- Math --------------------

func gg(r int) int64 {
	return int64(r) * int64(r+1) / 2
}

func ff(l, r int) int64 {
	return gg(r) - gg(l-1)
}

// -------------------- Preprocessing --------------------

func prepareAnswers(n int, a []int, queries []Query) {
	answerMap = make(map[QueryKey]int64, len(queries))

	m := n - 1

	if m <= 0 {
		return
	}

	// Build difference array.
	b := make([]int, m)

	for i := 0; i < m; i++ {
		b[i] = a[i] - a[i+1]

		if b[i] < 0 {
			b[i] = -b[i]
		}
	}

	// Sort queries by descending k.
	sortedQueries := make([]Query, len(queries))
	copy(sortedQueries, queries)

	sort.Slice(sortedQueries, func(i, j int) bool {
		return sortedQueries[i].k > sortedQueries[j].k
	})

	// Generate GCD events.
	events := make([]Event, 0, m*16)

	cur := make([]Pair, 0, 32)

	for i := m - 1; i >= 0; i-- {
		pr := Pair{
			g: b[i],
			r: i,
		}

		tmp := make([]Pair, 0, len(cur)+1)

		for _, p := range cur {
			val := gcdOpenEnd(p.g, pr.g)

			if pr.g != val {
				tmp = append(tmp, pr)

				pr.g = val
				pr.r = p.r
			} else {
				pr.g = val
				pr.r = p.r
			}
		}

		tmp = append(tmp, pr)
		cur = tmp

		for _, p := range cur {
			g := p.g

			if g == 0 {
				g = INF
			}

			events = append(events, Event{
				g: g,
				l: i,
				r: p.r,
			})
		}
	}

	// Sort events by descending GCD.
	sort.Slice(events, func(i, j int) bool {
		return events[i].g > events[j].g
	})

	st := NewSegmentTree(m)

	ptr := 0

	// Process queries.
	for _, q := range sortedQueries {
		l := q.l - 1
		r := q.r - 2

		// Empty difference interval.
		if l > r {
			answerMap[QueryKey{q.l, q.r, q.k}] = 0
			continue
		}

		// Activate all events with GCD >= k.
		for ptr < len(events) && events[ptr].g >= q.k {
			e := events[ptr]

			st.Update(e.l, e.r)

			ptr++
		}

		pos := st.GoFirst(r)

		if pos == -1 || pos > r {
			pos = r + 1
		} else if pos < l {
			pos = l
		}

		// Calculate answer.
		result := st.Get(l, pos-1)

		result += int64(r-pos+1) * int64(r)
		result -= ff(l, r)
		result += int64(r - l + 1)

		answerMap[QueryKey{q.l, q.r, q.k}] = result
	}
}

// -------------------- Required function --------------------

func calculateAnswerOpenEnd(_ int, _ []int, l int, r int, k int) int64 {
	return answerMap[QueryKey{
		l: l,
		r: r,
		k: k,
	}]
}

// -------------------- Input / Output --------------------
// https://coderun.yandex.ru/selections/2025-summer-common/problems/open-end
// OpenEnd - problem 25
func OpenEndSolve() {
	input := NewFastScanner(os.Stdin)

	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	n := input.readInt()
	q := input.readInt()

	a := input.readIntArray(n)

	queries := make([]Query, q)

	for i := 0; i < q; i++ {
		queries[i] = Query{
			l:  input.readInt(),
			r:  input.readInt(),
			k:  input.readInt(),
			id: i,
		}
	}

	// Precompute all answers.
	prepareAnswers(n, a, queries)

	// Output in original order.
	for i := 0; i < q; i++ {
		query := queries[i]

		answer := calculateAnswerOpenEnd(
			n,
			a,
			query.l,
			query.r,
			query.k,
		)

		output.WriteString(strconv.FormatInt(answer, 10))
		output.WriteByte('\n')
	}
}
