package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type NodeSecondMaximum struct {
	val   int
	left  *NodeSecondMaximum
	right *NodeSecondMaximum
}

func (n *NodeSecondMaximum) InsertSecondMaximum(val int) *NodeSecondMaximum {
	if n == nil {
		return &NodeSecondMaximum{val: val}
	}

	if val < n.val {
		n.left = n.left.InsertSecondMaximum(val)
	} else if val > n.val {
		n.right = n.right.InsertSecondMaximum(val)
	}

	return n
}

// https://coderun.yandex.ru/problem/second-maximum
// SecondMaximum - problem 240
func SecondMaximum() {
	var root *NodeSecondMaximum

	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	readLine := func() string {
		line, err := reader.ReadString('\n')
		if err != nil && !(err == io.EOF && len(line) > 0) {
			panic(err)
		}
		return strings.TrimRight(line, "\r\n")
	}

	// input
	line := readLine()
	nums := strings.Fields(line)

	for _, num := range nums {
		input, err := strconv.Atoi(num)
		if err != nil || input == 0 {
			break
		}

		root = root.InsertSecondMaximum(input)
	}

	var parent *NodeSecondMaximum
	cur := root
	for cur.right != nil {
		parent = cur
		cur = cur.right
	}

	if cur.left != nil {
		t := cur.left

		for t.right != nil {
			t = t.right
		}

		writer.WriteString(strconv.Itoa(t.val))
		writer.WriteByte('\n')
	} else {
		writer.WriteString(strconv.Itoa(parent.val))
		writer.WriteByte('\n')
	}
}
