package problems

import (
	"bufio"
	"container/heap"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	InfAvto = int(1e9)
)

type ItemAvto struct {
	next int // next occurrence position
	car  int // toy car id
}

type MaxHeapAvto []ItemAvto

func (h MaxHeapAvto) Len() int {
	return len(h)
}

// Car with the farthest next usage has highest priority
func (h MaxHeapAvto) Less(i, j int) bool {
	return h[i].next > h[j].next
}

func (h MaxHeapAvto) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeapAvto) Push(x interface{}) {
	*h = append(*h, x.(ItemAvto))
}

func (h *MaxHeapAvto) Pop() interface{} {
	old := *h
	n := len(old)

	item := old[n-1]
	*h = old[:n-1]

	return item
}

// https://coderun.yandex.ru/problem/avto
// Avto - problem 26
func Avto() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N, K and P input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 3 {
		panic("numbers count does not match 3")
	}

	_, err = strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	k, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	p, err := strconv.Atoi(strNum[2])
	if err != nil {
		panic(err)
	}

	// cars input
	req := make([]int, p)
	for i := 0; i < p; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		car, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		req[i] = car

	}

	// next occurrence
	nextIdx := make([]int, p)
	last := make(map[int]int)

	for i := p - 1; i >= 0; i-- {
		car := req[i]

		if pos, ok := last[car]; ok {
			nextIdx[i] = pos
		} else {
			nextIdx[i] = InfAvto
		}

		last[car] = i
	}

	onFloor := make(map[int]bool)
	curNext := make(map[int]int)

	h := &MaxHeapAvto{}
	heap.Init(h)

	ans := 0

	for i := 0; i < p; i++ {
		car := req[i]

		// Car is already on the floor
		if onFloor[car] {
			curNext[car] = nextIdx[i]
			heap.Push(h, ItemAvto{nextIdx[i], car})

			continue
		}

		// Need to bring a new car
		ans++

		// No free space on the floor
		if len(onFloor) == k {
			// Remove car used farthest in the future
			for {
				top := heap.Pop(h).(ItemAvto)

				// Skip outdated heap entries
				if !onFloor[top.car] {
					continue
				}

				if curNext[top.car] != top.next {
					continue
				}

				delete(onFloor, top.car)
				break
			}
		}

		// Put current car on the floor
		onFloor[car] = true
		curNext[car] = nextIdx[i]

		heap.Push(h, ItemAvto{nextIdx[i], car})
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
