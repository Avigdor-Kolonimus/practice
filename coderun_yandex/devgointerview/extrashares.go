package devgointerview

import (
	"bufio"
	"container/heap"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Promo struct {
	start  int
	end    int
	profit int
	id     int
}

type Event struct {
	x      int
	typ    int // 0 = end, 1 = start
	profit int
	id     int
}

type Item struct {
	profit int
	id     int
}

type MaxHeap []Item

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].profit > h[j].profit
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)

	item := old[n-1]
	*h = old[:n-1]

	return item
}

// https://coderun.yandex.ru/selections/dev-go-interview/problems/extra-shares
// ExtraShares - problem 15
func ExtraShares() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	events := make([]Event, 0, 2*n)
	for i := 0; i < n; i++ {
		// promo info input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum := strings.Fields(line)
		if len(strNum) != 3 {
			panic("numbers count does not match 3")
		}

		start, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}
		profit, err := strconv.Atoi(strNum[2])
		if err != nil {
			panic(err)
		}

		events = append(events, Event{
			x:      start,
			typ:    1,
			profit: profit,
			id:     i,
		})

		events = append(events, Event{
			x:      end,
			typ:    0,
			profit: profit,
			id:     i,
		})
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].x == events[j].x {
			return events[i].typ < events[j].typ
		}
		return events[i].x < events[j].x
	})

	active := make(map[int]bool)

	h := &MaxHeap{}
	heap.Init(h)

	useful := make([]bool, n)

	i := 0

	for i < len(events) {
		x := events[i].x

		// process all events at coordinate x
		for i < len(events) && events[i].x == x {
			e := events[i]

			if e.typ == 0 {
				delete(active, e.id)
			} else {
				active[e.id] = true

				heap.Push(h, Item{
					profit: e.profit,
					id:     e.id,
				})
			}

			i++
		}

		// remove inactive heap top
		for h.Len() > 0 && !active[(*h)[0].id] {
			heap.Pop(h)
		}

		// current maximum profit promo
		if h.Len() > 0 {
			useful[(*h)[0].id] = true
		}
	}

	ans := 0

	for i := 0; i < n; i++ {
		if !useful[i] {
			ans++
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
