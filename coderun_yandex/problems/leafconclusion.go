package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// Left → Root → Right
func inorderEmptyChild(writer *bufio.Writer, node *NodeBSL) {
	if node == nil {
		return
	}

	inorderEmptyChild(writer, node.left)

	if node.left == nil && node.right == nil {
		writer.WriteString(strconv.Itoa(node.val))
		writer.WriteByte('\n')
	}

	inorderEmptyChild(writer, node.right)
}

// https://coderun.yandex.ru/problem/leaf-conclusion
// LeafConclusion - problem 242
func LeafConclusion() {
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

	inorderEmptyChild(writer, tree)
}
