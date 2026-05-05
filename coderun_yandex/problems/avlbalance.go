package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type NodeAVL struct {
	value int
	Left  *NodeAVL
	Right *NodeAVL
}

type TreeAVL struct {
	root *NodeAVL
}

func buildTree(tree TreeAVL, num int) {
	curNode := tree.root
	for num != curNode.value {
		if num < curNode.value {
			if curNode.Left == nil {
				curNode.Left = addNode(num)
				return
			} else {
				curNode = curNode.Left
			}
		}

		if num > curNode.value {
			if curNode.Right == nil {
				curNode.Right = addNode(num)
				return
			} else {
				curNode = curNode.Right
			}
		}
	}
}

func addNode(num int) *NodeAVL {
	return &NodeAVL{
		num,
		nil,
		nil,
	}
}

func checkDepth(root *NodeAVL) int {
	if root == nil {
		return 0
	}

	leftDepth := checkDepth(root.Left)
	if leftDepth == -2 {
		return -2
	}

	rightDepth := checkDepth(root.Right)
	if rightDepth == -2 {
		return -2
	}

	diff := leftDepth - rightDepth
	if diff < -1 || diff > 1 {
		return -2
	}

	return max(leftDepth, rightDepth) + 1
}

// https://coderun.yandex.ru/problem/avl-balance
// AvlBalance - problem 89
func AvlBalance() {
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
	tree := TreeAVL{
		root: &NodeAVL{
			rootValue,
			nil,
			nil,
		},
	}
	for _, v := range strNum {
		inp, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		if inp == 0 {
			break
		}

		buildTree(tree, inp)
	}

	if checkDepth(tree.root) >= 1 {
		writer.WriteString("YES")
	} else {
		writer.WriteString("NO")
	}
	writer.WriteByte('\n')
}
