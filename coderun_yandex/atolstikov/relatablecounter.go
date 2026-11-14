package atolstikov

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type MinHeap []int64

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int64))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *MinHeap) push(x int64) {
	heap.Push(h, x)
}

func (h *MinHeap) pop() int64 {
	return heap.Pop(h).(int64)
}

type Solver struct {
	small    MinHeap
	large    MinHeap
	sumSmall int64
	k        int
}

func NewSolver(arr []int64, k int) *Solver {
	s := &Solver{
		k:     k,
		small: make(MinHeap, 0, len(arr)),
		large: make(MinHeap, 0, len(arr)),
	}

	for i, v := range arr {
		if i < k {
			s.small = append(s.small, v)
			s.sumSmall += v
		} else {
			s.large = append(s.large, v)
		}
	}

	heap.Init(&s.small)
	heap.Init(&s.large)

	return s
}

func (s *Solver) run(r int) []int64 {
	for i := 0; i < r; i++ {
		newValue := s.sumSmall

		minValue := s.small.pop()
		s.sumSmall -= minValue
		if s.k == s.small.Len()+1 && s.large.Len() == 0 {
			s.small.push(newValue)
			s.sumSmall += newValue
		} else {
			s.large.push(newValue)

			nextMin := s.large.pop()
			s.small.push(nextMin)
			s.sumSmall += nextMin
		}
	}

	return s.result()
}

func (s *Solver) result() []int64 {
	res := make([]int64, 0, s.small.Len()+s.large.Len())
	for s.small.Len() > 0 {
		res = append(res, s.small.pop())
	}

	for s.large.Len() > 0 {
		res = append(res, s.large.pop())
	}

	return res
}

// https://coderun.yandex.ru/selections/atolstikov/problems/relatable-counter
// RelatableCounter - problem 8
func RelatableCounter() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var n, k, r int
	fmt.Fscan(reader, &n, &k, &r)

	arr := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	solver := NewSolver(arr, k)
	ans := solver.run(r)

	for i, v := range ans {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
