package problems

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
	"strings"
)

type FreeHeap []int

func (h FreeHeap) Len() int           { return len(h) }
func (h FreeHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h FreeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *FreeHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *FreeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Train struct {
	departure int
	track     int
}

type BusyHeap []Train

func (h BusyHeap) Len() int {
	return len(h)
}

func (h BusyHeap) Less(i, j int) bool {
	return h[i].departure < h[j].departure
}

func (h BusyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *BusyHeap) Push(x interface{}) {
	*h = append(*h, x.(Train))
}

func (h *BusyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func readLineDeadEnds(reader *bufio.Reader) []int {
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	result := make([]int, len(parts))

	for i, part := range parts {
		result[i], _ = strconv.Atoi(part)
	}

	return result
}

// https://coderun.yandex.ru/problem/dead-ends
// DeadEnds - problem 24
func DeadEnds() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	first := readLineDeadEnds(reader)

	k := first[0]
	n := first[1]

	free := &FreeHeap{}
	heap.Init(free)

	for track := 1; track <= k; track++ {
		heap.Push(free, track)
	}

	busy := &BusyHeap{}
	heap.Init(busy)

	answer := make([]int, n)

	for i := range n {
		line := readLineDeadEnds(reader)

		arrival := line[0]
		departure := line[1]

		for busy.Len() > 0 && (*busy)[0].departure < arrival {
			train := heap.Pop(busy).(Train)
			heap.Push(free, train.track)
		}

		if free.Len() == 0 {
			writer.WriteString("0 ")
			writer.WriteString(strconv.Itoa(i + 1))
			writer.WriteByte('\n')

			return
		}

		track := heap.Pop(free).(int)
		answer[i] = track
		heap.Push(busy, Train{
			departure: departure,
			track:     track,
		})
	}

	for i, track := range answer {
		if i > 0 {
			writer.WriteByte(' ')
		}

		writer.WriteString(strconv.Itoa(track))
	}

	writer.WriteByte('\n')
}
