package backend

import (
	"bufio"
	"container/heap"
	"io"
	"os"
	"strconv"
	"strings"
)

type ItemDiversityImprovement struct {
	cat   int
	count int
}

type MaxHeapDiversityImprovement []ItemDiversityImprovement

func (h MaxHeapDiversityImprovement) Len() int { return len(h) }
func (h MaxHeapDiversityImprovement) Less(i, j int) bool {
	return h[i].count > h[j].count
}
func (h MaxHeapDiversityImprovement) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeapDiversityImprovement) Push(x interface{}) {
	*h = append(*h, x.(ItemDiversityImprovement))
}

func (h *MaxHeapDiversityImprovement) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Cooldown struct {
	release int
	item    ItemDiversityImprovement
}

func build(freq map[int]int, n int, dist int) ([]int, bool) {
	if dist <= 1 {
		res := make([]int, 0, n)
		for cat, cnt := range freq {
			for i := 0; i < cnt; i++ {
				res = append(res, cat)
			}
		}
		return res, true
	}

	h := &MaxHeapDiversityImprovement{}
	heap.Init(h)

	for cat, cnt := range freq {
		heap.Push(h, ItemDiversityImprovement{cat: cat, count: cnt})
	}

	queue := make([]Cooldown, 0)
	head := 0

	result := make([]int, 0, n)

	for pos := 0; pos < n; pos++ {
		for head < len(queue) && queue[head].release <= pos {
			heap.Push(h, queue[head].item)
			head++
		}

		if h.Len() == 0 {
			return nil, false
		}

		cur := heap.Pop(h).(ItemDiversityImprovement)
		result = append(result, cur.cat)

		cur.count--
		if cur.count > 0 {
			queue = append(queue, Cooldown{
				release: pos + dist,
				item:    cur,
			})
		}
	}

	return result, true
}

// https://coderun.yandex.ru/selections/backend/problems/diversity-improvement
// DiversityImprovement - problem 47
func DiversityImprovement() {
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

	// product and category inputs
	freq := make(map[int]int)
	productsByCat := make(map[int][]int)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum := strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		p, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		c, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		productsByCat[c] = append(productsByCat[c], p)
		freq[c]++
	}

	lo, hi := 1, n+1
	for lo+1 < hi {
		mid := (lo + hi) / 2

		if _, ok := build(freq, n, mid); ok {
			lo = mid
		} else {
			hi = mid
		}
	}

	categories, _ := build(freq, n, lo)
	ptr := make(map[int]int)

	for i, cat := range categories {
		if i > 0 {
			writer.WriteByte(' ')
		}

		idx := ptr[cat]
		writer.WriteString(strconv.Itoa(productsByCat[cat][idx]))
		ptr[cat]++
	}
	writer.WriteByte('\n')
}
