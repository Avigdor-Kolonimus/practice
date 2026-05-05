package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type NodeBSL struct {
	val   int
	left  *NodeBSL
	right *NodeBSL
}

func insert(root *NodeBSL, x int) *NodeBSL {
	if root == nil {
		return &NodeBSL{val: x}
	}

	if x < root.val {
		root.left = insert(root.left, x)
	} else if x > root.val {
		root.right = insert(root.right, x)
	}

	return root
}

// Left → Root → Right
func inorder(writer *bufio.Writer, node *NodeBSL) {
	if node == nil {
		return
	}

	inorder(writer, node.left)

	if node.left != nil && node.right != nil {
		writer.WriteString(strconv.Itoa(node.val))
		writer.WriteByte('\n')
	}

	inorder(writer, node.right)
}

// Root → Left → Right
func preorder(writer *bufio.Writer, node *NodeBSL) {
	if node == nil {
		return
	}

	writer.WriteString(strconv.Itoa(node.val))
	writer.WriteByte('\n')

	preorder(writer, node.left)
	preorder(writer, node.right)
}

// Left → Right → Root
func postorder(writer *bufio.Writer, node *NodeBSL) {
	if node == nil {
		return
	}

	postorder(writer, node.left)
	postorder(writer, node.right)

	writer.WriteString(strconv.Itoa(node.val))
	writer.WriteByte('\n')
}

// https://coderun.yandex.ru/problem/fork-conclusion
// ForkConclusion - problem 87
func ForkConclusion() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// tree line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	rootValue, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	strNum = strNum[1:]

	// root
	tree := &NodeBSL{
		rootValue,
		nil,
		nil,
	}

	for _, v := range strNum {
		inp, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		if inp == 0 {
			break
		}

		tree = insert(tree, inp)
	}

	inorder(writer, tree)
}
