package devgointerview

import (
	"bufio"
	"os"
	"strconv"
)

type MinHeap struct {
	data []int64
}

func (h *MinHeap) Push(x int64) {
	h.data = append(h.data, x)

	i := len(h.data) - 1

	for i > 0 {
		parent := (i - 1) / 2

		if h.data[parent] <= h.data[i] {
			break
		}

		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
	}
}

func (h *MinHeap) Pop() int64 {
	result := h.data[0]

	last := len(h.data) - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]

	i := 0

	for {
		left := i*2 + 1
		right := i*2 + 2
		smallest := i

		if left < len(h.data) && h.data[left] < h.data[smallest] {
			smallest = left
		}

		if right < len(h.data) && h.data[right] < h.data[smallest] {
			smallest = right
		}

		if smallest == i {
			break
		}

		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}

	return result
}

func (h *MinHeap) Top() int64 {
	return h.data[0]
}

func (h *MinHeap) Len() int {
	return len(h.data)
}

type FastScanner struct {
	reader *bufio.Reader
}

func NewFastScanner() *FastScanner {
	return &FastScanner{
		reader: bufio.NewReaderSize(os.Stdin, 1<<20),
	}
}

func (s *FastScanner) NextInt64() int64 {
	var result int64

	c, err := s.reader.ReadByte()

	for err == nil && (c < '0' || c > '9') {
		c, err = s.reader.ReadByte()
	}

	for err == nil && c >= '0' && c <= '9' {
		result = result*10 + int64(c-'0')
		c, err = s.reader.ReadByte()
	}

	return result
}

// https://coderun.yandex.ru/selections/dev-go-interview/problems/couriers-on-line
// CouriersOnLine - problem 11
func CouriersOnLine() {
	reader := NewFastScanner()
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and K input
	n := int(reader.NextInt64())
	k := int(reader.NextInt64())

	heap := MinHeap{
		data: make([]int64, 0, k),
	}

	// start and end input
	var answer int64
	var lastAccounted int64
	for range n {
		start := reader.NextInt64()
		end := reader.NextInt64()

		// Remove shifts that have already ended.
		for heap.Len() > 0 && heap.Top() <= start {
			heap.Pop()
		}

		// Add the current shift.
		heap.Push(end)

		// Keep only k largest end values.
		if heap.Len() > k {
			heap.Pop()
		}

		// At least k couriers are working simultaneously.
		if heap.Len() == k {
			segStart := start

			if lastAccounted > segStart {
				segStart = lastAccounted
			}

			segEnd := heap.Top()

			if segEnd > segStart {
				answer += segEnd - segStart
				lastAccounted = segEnd
			}
		}
	}

	writer.WriteString(strconv.FormatInt(answer, 10))
	writer.WriteByte('\n')
}
