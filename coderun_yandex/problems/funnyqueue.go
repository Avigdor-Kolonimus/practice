package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const MAX_BITS = 19

type NodePoolFunnyQueue struct {
	l      []int
	r      []int
	maxVal []int
	lazy   []int

	count int
}

func NewNodePool(capacity int) NodePoolFunnyQueue {
	return NodePoolFunnyQueue{
		l:      make([]int, capacity),
		r:      make([]int, capacity),
		maxVal: make([]int, capacity),
		lazy:   make([]int, capacity),
		count:  1,
	}
}

func (p *NodePoolFunnyQueue) NewNode() int {
	id := p.count
	p.count++

	return id
}

func (p *NodePoolFunnyQueue) Pull(u int) {
	ui := u
	m := 0

	left := p.l[ui]
	if left != 0 && p.maxVal[left] > m {
		m = p.maxVal[left]
	}

	right := p.r[ui]
	if right != 0 && p.maxVal[right] > m {
		m = p.maxVal[right]
	}

	p.maxVal[ui] = m
}

type Trie struct {
	pool NodePoolFunnyQueue
	root int
}

func NewTrie() Trie {
	pool := NewNodePool(10_000_000)

	return Trie{
		pool: pool,
		root: pool.NewNode(),
	}
}

func (t *Trie) pushLazy(u, bit int) {
	ui := int(u)
	mask := t.pool.lazy[ui]

	if mask == 0 {
		return
	}

	t.pool.lazy[ui] = 0

	if ((mask >> bit) & 1) == 1 {
		t.pool.l[ui], t.pool.r[ui] =
			t.pool.r[ui], t.pool.l[ui]
	}

	left := t.pool.l[ui]
	right := t.pool.r[ui]

	if left != 0 {
		t.pool.lazy[left] ^= mask
	}

	if right != 0 {
		t.pool.lazy[right] ^= mask
	}
}

func (t *Trie) Insert(index, val int) {
	t.root = t.insertRec(t.root, index, val, 0)
}

func (t *Trie) insertRec(u, key, val, bit int) int {
	node := u

	if node == 0 {
		node = t.pool.NewNode()
	}

	if bit == MAX_BITS {
		t.pool.maxVal[node] = val
		return node
	}

	t.pushLazy(node, bit)

	ni := int(node)

	if ((key >> bit) & 1) == 0 {
		newLeft := t.insertRec(
			t.pool.l[ni],
			key,
			val,
			bit+1,
		)
		t.pool.l[ni] = newLeft
	} else {
		newRight := t.insertRec(
			t.pool.r[ni],
			key,
			val,
			bit+1,
		)
		t.pool.r[ni] = newRight
	}

	t.pool.Pull(node)

	return node
}

func (t *Trie) Delete(index int) int {
	ans := 0

	t.root = t.deleteRec(
		t.root,
		index,
		0,
		&ans,
	)

	return ans
}

func (t *Trie) deleteRec(u int, key int, bit int, out *int) int {
	if u == 0 {
		return 0
	}

	if bit == MAX_BITS {
		*out = t.pool.maxVal[u]
		return 0
	}

	t.pushLazy(u, bit)

	ui := int(u)

	if ((key >> bit) & 1) == 0 {
		newLeft := t.deleteRec(
			t.pool.l[ui],
			key,
			bit+1,
			out,
		)
		t.pool.l[ui] = newLeft
	} else {
		newRight := t.deleteRec(
			t.pool.r[ui],
			key,
			bit+1,
			out,
		)
		t.pool.r[ui] = newRight
	}

	t.pool.Pull(u)

	if t.pool.l[ui] == 0 && t.pool.r[ui] == 0 {
		return 0
	}

	return u
}

func (t *Trie) ShiftPlusOne() {
	if t.root != 0 {
		t.addOneRec(t.root, 0)
	}
}

func (t *Trie) addOneRec(u, bit int) {
	if bit >= MAX_BITS {
		return
	}

	t.pushLazy(u, bit)

	ui := int(u)

	t.pool.l[ui], t.pool.r[ui] =
		t.pool.r[ui], t.pool.l[ui]

	left := t.pool.l[ui]
	if left != 0 {
		t.addOneRec(left, bit+1)
	}

	t.pool.Pull(u)
}

func (t *Trie) ShiftMinusOne() {
	if t.root != 0 {
		t.subOneRec(t.root, 0)
	}
}

func (t *Trie) subOneRec(u, bit int) {
	if bit >= MAX_BITS {
		return
	}

	t.pushLazy(u, bit)

	ui := int(u)

	left := t.pool.l[ui]
	if left != 0 {
		t.subOneRec(left, bit+1)
	}

	t.pool.l[ui], t.pool.r[ui] =
		t.pool.r[ui], t.pool.l[ui]

	t.pool.Pull(u)
}

func (t *Trie) ApplyXor(x int) {
	if t.root != 0 {
		t.pool.lazy[int(t.root)] ^= x
	}
}

func (t *Trie) GetMax() int {
	if t.root == 0 {
		return 0
	}
	return int(t.pool.maxVal[int(t.root)])
}

// https://coderun.yandex.ru/problem/funny-queue
// FunnyQueue - problem 573
func FunnyQueue() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	trie := NewTrie()
	size := 0
	for i := 0; i < n; i++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		parts := strings.Fields(line)
		t, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		switch t {

		// add to end
		case 1:
			x, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}

			trie.Insert(size, x)
			size++

		// add to start
		case 2:
			x, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}

			trie.ShiftPlusOne()
			trie.Insert(0, x)
			size++

		// remove from end
		case 3:
			v := trie.Delete(size - 1)
			size--
			writer.WriteString(strconv.Itoa(v))
			writer.WriteByte('\n')

		// remove from start
		case 4:
			v := trie.Delete(0)
			trie.ShiftMinusOne()
			size--
			writer.WriteString(strconv.Itoa(v))
			writer.WriteByte('\n')

		// XOR permutation
		case 5:
			x, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			trie.ApplyXor(x)

		// get maximum element
		case 6:
			writer.WriteString(strconv.Itoa(trie.GetMax()))
			writer.WriteByte('\n')
		}
	}
}
