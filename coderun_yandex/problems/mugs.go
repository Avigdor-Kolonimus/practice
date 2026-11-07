package problems

import (
	"bufio"
	"container/heap"
	"fmt"
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

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)

	x := old[n-1]
	*h = old[:n-1]

	return x
}

// https://coderun.yandex.ru/problem/mugs
// Mugs - problem 42
func Mugs() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	graph := make([][]int, n+1)
	inDegree := make([]int, n+1)
	for u := 1; u <= n; u++ {
		var k int
		fmt.Fscan(reader, &k)

		for j := 0; j < k; j++ {
			var v int
			fmt.Fscan(reader, &v)

			graph[u] = append(graph[u], v)
			inDegree[v]++
		}
	}

	h := &MaxHeap{}
	heap.Init(h)

	for v := 1; v <= n; v++ {
		if inDegree[v] == 0 {
			heap.Push(h, v)
		}
	}

	answer := make([]int, 0, n)
	for h.Len() > 0 {
		u := heap.Pop(h).(int)

		answer = append(answer, u)
		for _, v := range graph[u] {
			inDegree[v]--

			if inDegree[v] == 0 {
				heap.Push(h, v)
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		if i < n-1 {
			writer.WriteByte(' ')
		}

		writer.WriteString(strconv.Itoa(answer[i]))
	}

	writer.WriteByte('\n')
}
