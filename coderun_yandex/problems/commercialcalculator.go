package problems

import (
	"bufio"
	"container/heap"
	"io"
	"os"
	"strconv"
	"strings"
)

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

// https://coderun.yandex.ru/problem/commercial-calculator
// CommercialCalculator- problem 25
func CommercialCalculator() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N  input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// price inputs
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	prices := strings.Fields(line)

	h := &MinHeap{}
	heap.Init(h)
	for i := 0; i < n; i++ {
		x, err := strconv.Atoi(prices[i])
		if err != nil {
			panic(err)
		}

		heap.Push(h, x)
	}

	totalCost := 0.0
	for h.Len() > 1 {
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)

		sum := a + b

		cost := float64(sum) * 0.05
		totalCost += cost

		heap.Push(h, sum)
	}

	writer.WriteString(strconv.FormatFloat(totalCost, 'f', 2, 64))
	writer.WriteByte('\n')
}
