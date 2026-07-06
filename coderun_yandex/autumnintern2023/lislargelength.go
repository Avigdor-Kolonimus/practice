package autumnintern2023

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type SegTree struct {
	size int
	data []int
}

func NewSegTree(n int) *SegTree {
	size := 1
	for size < n {
		size <<= 1
	}
	return &SegTree{
		size: size,
		data: make([]int, 2*size),
	}
}

func (st *SegTree) Set(pos, val int) {
	i := pos + st.size
	st.data[i] = val
	for i > 1 {
		i >>= 1
		st.data[i] = max(st.data[i<<1], st.data[i<<1|1])
	}
}

func (st *SegTree) Query(l, r int) int {
	res := 0
	l += st.size
	r += st.size

	for l < r {
		if l&1 == 1 {
			res = max(res, st.data[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, st.data[r])
		}
		l >>= 1
		r >>= 1
	}

	return res
}

// https://coderun.yandex.ru/selections/autumn-intern-2023/problems/lis-large-length
// LisLargeLength - problem 4
func LisLargeLength() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// A and B input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}
	a, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// sequence input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNum = strings.Fields(line)
	if len(strNum) != a {
		panic("numbers count does not match a")
	}

	sequence := make([]int, a)
	for i := 0; i < a; i++ {
		sequence[i], err = strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}
	}

	st := NewSegTree(a + 1)
	ans := 0
	for i := 0; i < a; i++ {
		x := sequence[i]

		l := x - b
		if l < 0 {
			l = 0
		}

		cur := st.Query(l, x+1) + 1

		st.Set(x, cur)

		if cur > ans {
			ans = cur
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
