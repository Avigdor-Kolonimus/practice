package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type BST struct {
	root *Node
}

type Node struct {
	val   int
	left  *Node
	right *Node
}

// AddVal adds a value to the BST and returns the depth of the node
func (bst *BST) AddVal(val int) int {
	if bst.root == nil {
		bst.root = &Node{val: val}
		return 1
	}

	cur := bst.root
	depth := 1
	for {
		if val == cur.val {
			return -1 // already exists
		}
		depth++
		if val < cur.val {
			if cur.left == nil {
				cur.left = &Node{val: val}
				return depth
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = &Node{val: val}
				return depth
			}
			cur = cur.right
		}
	}
}

// https://coderun.yandex.ru/problem/depth-added-elements
// DepthAddedElements - problem 86
func DepthAddedElements() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNums := strings.Fields(line)

	bst := &BST{}
	result := make([]int, 0)
	nums := make([]int, len(strNums)-1)
	for i, v := range strNums {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		if n == 0 {
			break
		}

		nums[i] = n

		if d := bst.AddVal(n); d != -1 {
			result = append(result, d)
		}
	}

	for i := 0; i < len(result); i++ {
		writer.WriteString(strconv.Itoa(result[i]) + " ")
	}
	writer.WriteByte('\n')
}
