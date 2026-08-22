package backend

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type SegTree struct {
	n    int
	lazy []int
}

func NewSegTree(a []int) *SegTree {
	n := len(a)

	st := &SegTree{
		n:    n,
		lazy: make([]int, 4*n),
	}

	st.build(1, 0, n-1, a)

	return st
}

// lazy[v] != 0 means that the whole segment
// has the same server number.
//
// lazy[v] == 0 means that the segment is not uniform.
func (st *SegTree) build(v, l, r int, a []int) {
	if l == r {
		st.lazy[v] = a[l]
		return
	}

	mid := (l + r) / 2

	st.build(v*2, l, mid, a)
	st.build(v*2+1, mid+1, r, a)

	if st.lazy[v*2] == st.lazy[v*2+1] {
		st.lazy[v] = st.lazy[v*2]
	}
}

func (st *SegTree) push(v int) {
	if st.lazy[v] == 0 {
		return
	}

	st.lazy[v*2] = st.lazy[v]
	st.lazy[v*2+1] = st.lazy[v]
	st.lazy[v] = 0
}

// check returns true if all chunks in [ql, qr]
// are located on server expected.
func (st *SegTree) check(v, l, r, ql, qr, expected int) bool {
	if qr < l || r < ql {
		return true
	}

	if ql <= l && r <= qr && st.lazy[v] != 0 {
		return st.lazy[v] == expected
	}

	if l == r {
		return st.lazy[v] == expected
	}

	st.push(v)

	mid := (l + r) / 2

	return st.check(v*2, l, mid, ql, qr, expected) &&
		st.check(v*2+1, mid+1, r, ql, qr, expected)
}

// assign moves all chunks in [ql, qr]
// to the specified server.
func (st *SegTree) assign(v, l, r, ql, qr, server int) {
	if qr < l || r < ql {
		return
	}

	if ql <= l && r <= qr {
		st.lazy[v] = server
		return
	}

	st.push(v)

	mid := (l + r) / 2

	st.assign(v*2, l, mid, ql, qr, server)
	st.assign(v*2+1, mid+1, r, ql, qr, server)

	if st.lazy[v*2] != 0 &&
		st.lazy[v*2] == st.lazy[v*2+1] {
		st.lazy[v] = st.lazy[v*2]
	} else {
		st.lazy[v] = 0
	}
}

func (st *SegTree) Check(l, r, server int) bool {
	return st.check(1, 0, st.n-1, l, r, server)
}

func (st *SegTree) Assign(l, r, server int) {
	st.assign(1, 0, st.n-1, l, r, server)
}

type FastReader struct {
	reader *bufio.Reader
}

func NewFastReader() *FastReader {
	return &FastReader{
		reader: bufio.NewReaderSize(os.Stdin, 1<<20),
	}
}

func (fr *FastReader) NextInt() int {
	for {
		line, err := fr.reader.ReadString('\n')

		if err != nil && len(line) == 0 {
			return 0
		}

		fields := strings.Fields(line)

		for _, field := range fields {
			value, err := strconv.Atoi(field)

			if err == nil {
				return value
			}
		}
	}
}

// https://coderun.yandex.ru/selections/backend/problems/lollipops
// Lollipops - problem 5
func Lollipops() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// We need to read arbitrary numbers of integers from lines,
	// so keep a small token buffer.
	var tokens []string

	nextInt := func() int {
		for len(tokens) == 0 {
			line, _ := reader.ReadString('\n')
			tokens = strings.Fields(line)
		}

		value, _ := strconv.Atoi(tokens[0])
		tokens = tokens[1:]

		return value
	}

	n := nextInt()
	m := nextInt()
	q := nextInt()

	_ = m

	servers := make([]int, n)
	for i := range n {
		servers[i] = nextInt()
	}

	tree := NewSegTree(servers)
	for i := 0; i < q; i++ {
		a := nextInt()
		b := nextInt()
		l := nextInt()
		r := nextInt()

		// Convert chunk numbers to zero-based indices.
		l--
		r--

		if tree.Check(l, r, a) {
			writer.WriteString("1\n")
			tree.Assign(l, r, b)
		} else {
			writer.WriteString("0\n")
		}
	}
}
