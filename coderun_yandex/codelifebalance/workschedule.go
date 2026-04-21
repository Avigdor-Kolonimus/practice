package codelifebalance

import (
	"bufio"
	"container/heap"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Task struct {
	d, w int
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

func validateWorkScheduleInput(p int) bool {
	return p >= 1 && p <= 200_000
}

// https://coderun.yandex.ru/selections/code-life-balance/problems/work-schedule
// WorkSchedule - assignment 9
func WorkSchedule() {
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
	if !validateWorkScheduleInput(n) {
		panic("number N out of range")
	}

	// tasks
	tasks := make([]Task, n)
	total := 0
	for i := range n {
		// N input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		strNum := strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		// D
		d, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		if !validateWorkScheduleInput(d) {
			panic("number D out of range")
		}

		// W
		w, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}
		if !validateWorkScheduleInput(w) {
			panic("number W out of range")
		}

		tasks[i] = Task{d: d, w: w}
		total += tasks[i].w
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].d < tasks[j].d
	})

	h := &MinHeap{}
	heap.Init(h)

	sumDone := 0
	for _, t := range tasks {
		heap.Push(h, t.w)
		sumDone += t.w

		if h.Len() > t.d {
			removed := heap.Pop(h).(int)
			sumDone -= removed
		}
	}

	writer.WriteString(strconv.Itoa(total - sumDone))
	writer.WriteByte('\n')
}
