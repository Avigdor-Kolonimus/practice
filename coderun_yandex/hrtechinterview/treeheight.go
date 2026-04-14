package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func (n *Node) Insert(val int) *Node {
	if n == nil {
		return &Node{val: val}
	}

	if val < n.val {
		n.left = n.left.Insert(val)
	} else if val > n.val {
		n.right = n.right.Insert(val)
	}
	return n
}

func (n *Node) Height() int {
	if n == nil {
		return 0
	}

	leftH := n.left.Height()
	rightH := n.right.Height()

	if leftH > rightH {
		return leftH + 1
	}

	return rightH + 1
}

// https://coderun.yandex.ru/selections/hr-tech-interview/problems/tree-height
// TreeHeight - problem 10
func TreeHeight() {
	var root *Node

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

		root = root.Insert(input)
	}

	writer.WriteString(strconv.Itoa(root.Height()))
	writer.WriteByte('\n')
}
