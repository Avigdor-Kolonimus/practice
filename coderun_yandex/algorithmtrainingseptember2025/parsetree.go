package algorithmtrainingseptember2025

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type NodeParseTree struct {
	val   byte
	left  *NodeParseTree
	right *NodeParseTree
}

var s string
var pos int

func expr() *NodeParseTree {
	cur := term()

	for pos < len(s) && (s[pos] == '+' || s[pos] == '-') {
		op := s[pos]
		pos++

		right := term()
		cur = &NodeParseTree{
			val:   op,
			left:  cur,
			right: right,
		}
	}

	return cur
}

func term() *NodeParseTree {
	cur := power()

	for pos < len(s) && (s[pos] == '*' || s[pos] == '/') {
		op := s[pos]
		pos++

		right := power()
		cur = &NodeParseTree{
			val:   op,
			left:  cur,
			right: right,
		}
	}

	return cur
}

func power() *NodeParseTree {
	left := atom()

	if pos < len(s) && s[pos] == '^' {
		pos++

		right := power() // правоассоциативность

		return &NodeParseTree{
			val:   '^',
			left:  left,
			right: right,
		}
	}

	return left
}

func atom() *NodeParseTree {
	if s[pos] == '(' {
		pos++

		res := expr()

		pos++ // ')'
		return res
	}

	v := &NodeParseTree{val: s[pos]}
	pos++

	return v
}

type Picture struct {
	lines []string
	h     int
	w     int
	root  int
}

func build(v *NodeParseTree) Picture {
	if v.left == nil && v.right == nil {
		return Picture{
			lines: []string{string(v.val)},
			h:     1,
			w:     1,
			root:  0,
		}
	}

	L := build(v.left)
	R := build(v.right)

	H := max(L.h, R.h) + 2
	W := L.w + R.w + 5

	canvas := make([][]byte, H)
	for i := range canvas {
		canvas[i] = make([]byte, W)
		for j := range canvas[i] {
			canvas[i][j] = ' '
		}
	}

	opPos := L.w + 1

	canvas[0][opPos] = '['
	canvas[0][opPos+1] = v.val
	canvas[0][opPos+2] = ']'

	leftRoot := L.root
	rightRoot := L.w + 5 + R.root

	canvas[0][leftRoot] = '.'
	for c := leftRoot + 1; c < opPos; c++ {
		canvas[0][c] = '-'
	}

	for c := opPos + 3; c < rightRoot; c++ {
		canvas[0][c] = '-'
	}
	canvas[0][rightRoot] = '.'

	canvas[1][leftRoot] = '|'
	canvas[1][rightRoot] = '|'

	for i := 0; i < L.h; i++ {
		for j := 0; j < L.w; j++ {
			canvas[i+2][j] = L.lines[i][j]
		}
	}

	offset := L.w + 5
	for i := 0; i < R.h; i++ {
		for j := 0; j < R.w; j++ {
			canvas[i+2][offset+j] = R.lines[i][j]
		}
	}

	lines := make([]string, H)
	for i := range canvas {
		lines[i] = string(canvas[i])
	}

	return Picture{
		lines: lines,
		h:     H,
		w:     W,
		root:  L.w + 2,
	}
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/parse-tree
// ParseTree - assignment 29
func ParseTree() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// line input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	s = strings.TrimRight(line, "\r\n")

	root := expr()
	ans := build(root)

	for _, row := range ans.lines {
		last := len(row) - 1
		for last >= 0 && row[last] == ' ' {
			last--
		}

		if last >= 0 {
			writer.WriteString(row[:last+1])
		}
		writer.WriteByte('\n')
	}
}
