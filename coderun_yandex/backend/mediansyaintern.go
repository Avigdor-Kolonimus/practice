package backend

import (
	"bufio"
	"container/heap"
	"io"
	"os"
	"strconv"
	"strings"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// https://coderun.yandex.ru/selections/backend/problems/medians-ya-intern
// MediansYaIntern - problem 36
func MediansYaIntern() {
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

	// array input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	left := &MaxHeap{}
	right := &MinHeap{}
	heap.Init(left)
	heap.Init(right)
	var ans int
	for i := 0; i < n; i++ {
		x, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		if left.Len() == 0 || x <= (*left)[0] {
			heap.Push(left, x)
		} else {
			heap.Push(right, x)
		}

		if left.Len() > right.Len()+1 {
			heap.Push(right, heap.Pop(left))
		}

		if right.Len() > left.Len() {
			heap.Push(left, heap.Pop(right))
		}

		ans += (*left)[0]
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
