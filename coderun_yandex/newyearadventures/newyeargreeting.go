package newyearadventures

import (
	"bufio"
	"container/heap"
	"io"
	"os"
	"strconv"
	"strings"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

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

// https://coderun.yandex.ru/selections/new-year-adventures/problems/new_year_greeting
// NewYearGreeting - problem 7
func NewYearGreeting() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and T input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}

	tTotal, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// friends input
	x := make([]int, n)
	t := make([]int, n)
	for i := range n {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		xi, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		ti, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		x[i] = xi
		t[i] = ti
	}

	h := &MaxHeap{}
	heap.Init(h)

	cur, ans := 0, 0
	for i := 0; i < n; i++ {
		if x[i] > tTotal {
			break
		}

		cur += t[i]
		heap.Push(h, t[i])

		for h.Len() > 0 && cur > tTotal-x[i] {
			cur -= heap.Pop(h).(int)
		}

		if h.Len() > ans {
			ans = h.Len()
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
