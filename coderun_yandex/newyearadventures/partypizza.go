package newyearadventures

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	val int // candidate
	cnt int // balance
}

func merge(a, b Node) Node {
	if a.val == b.val {
		return Node{a.val, a.cnt + b.cnt}
	}

	if a.cnt > b.cnt {
		return Node{a.val, a.cnt - b.cnt}
	}

	return Node{b.val, b.cnt - a.cnt}
}

var (
	n, m int
	a    []int
	tree []Node
	pos  [][]int
)

func build(v, tl, tr int) {
	if tl == tr {
		tree[v] = Node{a[tl], 1}
		return
	}

	tm := (tl + tr) / 2

	build(v*2, tl, tm)
	build(v*2+1, tm+1, tr)

	tree[v] = merge(tree[v*2], tree[v*2+1])
}

func query(v, tl, tr, l, r int) Node {
	if l > r {
		return Node{0, 0}
	}

	if l == tl && r == tr {
		return tree[v]
	}

	tm := (tl + tr) / 2

	left := query(v*2, tl, tm, l, min(r, tm))
	right := query(v*2+1, tm+1, tr, max(l, tm+1), r)

	return merge(left, right)
}

// https://coderun.yandex.ru/selections/new-year-adventures/problems/party-pizza
// PartyPizza - problem 10
func PartyPizza() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// friends input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum = strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	a = make([]int, n+1)
	pos = make([][]int, n+1)
	for i := 1; i <= n; i++ {
		ai, err := strconv.Atoi(strNum[i-1])
		if err != nil {
			panic(err)
		}

		a[i] = ai
		pos[ai] = append(pos[ai], i)
	}

	// requests input
	tree = make([]Node, 4*n+5)
	build(1, 1, n)
	for i := 0; i < m; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		l, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		cand := query(1, 1, n, l, r).val

		positions := pos[cand]

		left := sort.SearchInts(positions, l)
		right := sort.SearchInts(positions, r+1)

		cnt := right - left
		length := r - l + 1

		if cnt > length/2 {
			writer.WriteString(strconv.Itoa(cand))
		} else {
			writer.WriteByte('0')
		}
		writer.WriteByte('\n')
	}
}
