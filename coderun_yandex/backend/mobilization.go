package backend

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
)

type Item struct {
	value int64
	idx   int
	ver   int
}

type Heap struct {
	items []Item
	max   bool
}

func (h Heap) Len() int {
	return len(h.items)
}

func (h Heap) Less(i, j int) bool {
	if h.max {
		return h.items[i].value > h.items[j].value
	}
	return h.items[i].value < h.items[j].value
}

func (h Heap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *Heap) Push(x any) {
	h.items = append(h.items, x.(Item))
}

func (h *Heap) Pop() any {
	old := h.items
	n := len(old)
	item := old[n-1]
	h.items = old[:n-1]
	return item
}

// https://coderun.yandex.ru/selections/first-2023-backend/problems/mobilization
// Mobilization - problem 21
func Mobilization() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	readInt := func() int64 {
		var x int64
		c, err := reader.ReadByte()

		for err == nil && c <= ' ' {
			c, err = reader.ReadByte()
		}

		for err == nil && c >= '0' && c <= '9' {
			x = x*10 + int64(c-'0')
			c, err = reader.ReadByte()
		}

		return x
	}

	n := int(readInt())
	k := n / 2

	a := make([]int64, n)
	b := make([]int64, n)
	diff := make([]int64, n)
	ver := make([]int, n)
	side := make([]int, n)

	for i := range n {
		a[i] = readInt()
	}

	for i := range n {
		b[i] = readInt()
	}

	var sumB int64

	for i := range n {
		sumB += b[i]
		diff[i] = a[i] - b[i]
	}

	selected := &Heap{}
	rest := &Heap{max: true}

	heap.Init(selected)
	heap.Init(rest)

	for i := range n {
		heap.Push(rest, Item{
			value: diff[i],
			idx:   i,
			ver:   0,
		})
	}

	selectedCount := 0
	var selectedSum int64

	clean := func(h *Heap, expectedSide int) {
		for h.Len() > 0 {
			top := h.items[0]

			if top.ver == ver[top.idx] &&
				side[top.idx] == expectedSide {
				break
			}

			heap.Pop(h)
		}
	}

	rebalance := func() {
		clean(selected, 1)
		clean(rest, 0)

		for selectedCount < k {
			clean(rest, 0)
			item := heap.Pop(rest).(Item)

			side[item.idx] = 1
			selectedCount++
			selectedSum += diff[item.idx]

			heap.Push(selected, Item{
				value: diff[item.idx],
				idx:   item.idx,
				ver:   ver[item.idx],
			})
		}

		for selectedCount > k {
			clean(selected, 1)
			item := heap.Pop(selected).(Item)

			side[item.idx] = 0
			selectedCount--
			selectedSum -= diff[item.idx]

			heap.Push(rest, Item{
				value: diff[item.idx],
				idx:   item.idx,
				ver:   ver[item.idx],
			})
		}

		for {
			clean(selected, 1)
			clean(rest, 0)

			if selected.Len() == 0 || rest.Len() == 0 {
				break
			}

			if selected.items[0].value >= rest.items[0].value {
				break
			}

			s := heap.Pop(selected).(Item)
			r := heap.Pop(rest).(Item)

			side[s.idx] = 0
			side[r.idx] = 1

			selectedSum += diff[r.idx] - diff[s.idx]

			heap.Push(rest, Item{
				value: diff[s.idx],
				idx:   s.idx,
				ver:   ver[s.idx],
			})

			heap.Push(selected, Item{
				value: diff[r.idx],
				idx:   r.idx,
				ver:   ver[r.idx],
			})
		}
	}

	rebalance()

	m := int(readInt())

	for range m {
		num := int(readInt()) - 1
		typ := int(readInt())
		d := readInt()

		if side[num] == 1 {
			selectedSum -= diff[num]
			selectedCount--
			side[num] = 0
		}

		if typ == 1 {
			a[num] += d
			diff[num] += d
		} else {
			b[num] += d
			sumB += d
			diff[num] -= d
		}

		ver[num]++

		heap.Push(rest, Item{
			value: diff[num],
			idx:   num,
			ver:   ver[num],
		})

		rebalance()

		answer := sumB + selectedSum
		writer.WriteString(strconv.FormatInt(answer, 10))
		writer.WriteByte('\n')
	}
}
