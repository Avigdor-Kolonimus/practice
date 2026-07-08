package summercommon2026

import (
	"bufio"
	"os"
)

const (
	MAXNSUDOKUGRID = 16
)

// ---------- DLX ----------
type Node struct {
	left, right, up, down *Node
	head                  *ColumnNode
	rowID                 int
}

type ColumnNode struct {
	node *Node
	size int
	id   int
}

type DLX struct {
	root        *Node
	headers     []*ColumnNode
	solution    []int
	solCount    int
	maxRowID    int
	gridAnswers []int
	N           int
}

func NewDLX(numCols int, maxRows int) *DLX {
	root := &Node{}
	root.left = root
	root.right = root
	root.up = root
	root.down = root

	headers := make([]*ColumnNode, numCols)
	last := root
	for i := 0; i < numCols; i++ {
		h := &ColumnNode{node: &Node{}, id: i}
		h.node.head = h
		h.node.up = h.node
		h.node.down = h.node

		h.node.left = last
		h.node.right = root
		last.right = h.node
		root.left = h.node

		headers[i] = h
		last = h.node
	}

	return &DLX{
		root:        root,
		headers:     headers,
		solution:    make([]int, maxRows),
		gridAnswers: make([]int, maxRows),
	}
}

func (dlx *DLX) AddRow(rowID int, colIndices []int) {
	var first *Node
	var last *Node

	for _, colIdx := range colIndices {
		col := dlx.headers[colIdx]
		col.size++

		node := &Node{
			rowID: rowID,
			head:  col,
		}

		node.up = col.node.up
		node.down = col.node
		col.node.up.down = node
		col.node.up = node

		if first == nil {
			first = node
			last = node
			node.left = node
			node.right = node
		} else {
			node.left = last
			node.right = first
			last.right = node
			first.left = node
			last = node
		}
	}
}

func (dlx *DLX) cover(col *ColumnNode) {
	col.node.right.left = col.node.left
	col.node.left.right = col.node.right

	for row := col.node.down; row != col.node; row = row.down {
		for n := row.right; n != row; n = n.right {
			n.down.up = n.up
			n.up.down = n.down
			n.head.size--
		}
	}
}

func (dlx *DLX) uncover(col *ColumnNode) {
	for row := col.node.up; row != col.node; row = row.up {
		for n := row.left; n != row; n = n.left {
			n.head.size++
			n.down.up = n
			n.up.down = n
		}
	}

	col.node.left.right = col.node
	col.node.right.left = col.node
}

func (dlx *DLX) search(k int) bool {
	if dlx.root.right == dlx.root {
		dlx.solCount = k
		for i := 0; i < k; i++ {
			dlx.gridAnswers[i] = dlx.solution[i]
		}
		return true
	}

	var bestCol *ColumnNode
	minSize := 999999

	for n := dlx.root.right; n != dlx.root; n = n.right {
		col := n.head
		if col.size < minSize {
			minSize = col.size
			bestCol = col

			if minSize == 0 {
				return false
			}
		}
	}

	dlx.cover(bestCol)

	for row := bestCol.node.down; row != bestCol.node; row = row.down {
		dlx.solution[k] = row.rowID

		for n := row.right; n != row; n = n.right {
			dlx.cover(n.head)
		}

		if dlx.search(k + 1) {
			return true
		}

		for n := row.left; n != row; n = n.left {
			dlx.uncover(n.head)
		}
	}

	dlx.uncover(bestCol)

	return false
}

// https://coderun.yandex.ru/seasons/2026-summer/tracks/common/problem/beach-sudoku
// BeachSudoku - problem 10
func BeachSudoku() {
	fs := NewFastScanner()
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	tCount := fs.NextInt()

	var outBuf [20]byte
	grid := make([][]int, MAXNSUDOKUGRID)
	for i := range MAXNSUDOKUGRID {
		grid[i] = make([]int, MAXNSUDOKUGRID)
	}

	for t := 0; t < tCount; t++ {
		n := fs.NextInt()
		N := n * n

		numCols := 4 * N * N
		maxRows := N * N * N

		dlx := NewDLX(numCols, maxRows)
		dlx.N = N

		initialValid := true
		rowsMask := [MAXNSUDOKUGRID]int{}
		colsMask := [MAXNSUDOKUGRID]int{}
		boxMask := [MAXNSUDOKUGRID]int{}

		for r := range N {
			for c := range N {
				val := fs.NextInt()
				grid[r][c] = val
				if val != 0 {
					bit := 1 << (val - 1)
					b := (r/n)*n + (c / n)

					if (rowsMask[r]&bit) != 0 || (colsMask[c]&bit) != 0 || (boxMask[b]&bit) != 0 {
						initialValid = false
					}

					rowsMask[r] |= bit
					colsMask[c] |= bit
					boxMask[b] |= bit
				}
			}
		}

		if !initialValid {
			writer.WriteString("NO\n")

			continue
		}

		for r := range N {
			for c := range N {
				b := (r/n)*n + (c / n)
				fixedVal := grid[r][c]

				for v := 1; v <= N; v++ {
					if fixedVal != 0 && fixedVal != v {
						continue
					}

					rowID := (r * N * N) + (c * N) + (v - 1)

					c1 := r*N + c
					c2 := N*N + r*N + (v - 1)
					c3 := 2*N*N + c*N + (v - 1)
					c4 := 3*N*N + b*N + (v - 1)

					dlx.AddRow(rowID, []int{c1, c2, c3, c4})
				}
			}
		}

		if dlx.search(0) {
			writer.WriteString("YES\n")
			for i := 0; i < dlx.solCount; i++ {
				id := dlx.gridAnswers[i]
				v := (id % N) + 1
				id /= N
				c := id % N
				r := id / N
				grid[r][c] = v
			}

			for r := range N {
				for c := range N {
					x := grid[r][c]
					pos := len(outBuf)
					for x > 0 {
						pos--
						outBuf[pos] = byte(x%10) + '0'
						x /= 10
					}
					writer.Write(outBuf[pos:])
					if c < N-1 {
						writer.WriteByte(' ')
					}
				}
				writer.WriteByte('\n')
			}
		} else {
			writer.WriteString("NO\n")
		}
	}
}
