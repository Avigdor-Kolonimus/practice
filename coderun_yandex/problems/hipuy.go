package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Heap struct {
	data []int
}

func (h *Heap) Insert(val int) {
	h.data = append(h.data, val)
	idx := len(h.data) - 1
	h.shiftUp(idx)
}

func (h *Heap) shiftUp(idx int) {
	for idx > 0 {
		parent := (idx - 1) / 2

		if h.data[parent] >= h.data[idx] {
			break
		}

		h.data[idx], h.data[parent] = h.data[parent], h.data[idx]
		idx = parent
	}
}

func (h *Heap) Extract() int {
	val := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	if len(h.data) > 1 {
		h.shiftDown(0)
	}

	return val
}

func (h *Heap) shiftDown(idx int) {
	for {
		left, right := 2*idx+1, 2*idx+2
		if left >= len(h.data) {
			break
		}

		// choose larger child
		largest := left
		if right < len(h.data) && h.data[right] > h.data[largest] {
			largest = right
		}

		// if current is already >= largest child, done
		if h.data[idx] >= h.data[largest] {
			break
		}

		// swap and continue
		h.data[idx], h.data[largest] = h.data[largest], h.data[idx]
		idx = largest
	}
}

func validateHipuyInput(n int) bool {
	return n >= 1 && n <= 100_000
}

// https://coderun.yandex.ru/problem/hipuy
// Hipuy - problem 251
func Hipuy() {
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
	if !validateHipuyInput(n) {
		panic("number N out of range")
	}

	// commands input
	heap := &Heap{
		data: make([]int, 0),
	}
	for range n {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		parameters := strings.Fields(line)
		if len(parameters) != 2 {
			maxNum := heap.Extract()

			writer.WriteString(strconv.Itoa(maxNum))
			writer.WriteByte('\n')
		} else {
			num, err := strconv.Atoi(parameters[1])
			if err != nil {
				panic(err)
			}

			heap.Insert(num)
		}
	}
}
