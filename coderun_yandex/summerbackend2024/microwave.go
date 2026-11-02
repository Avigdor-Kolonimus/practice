package summerbackend2024

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	dist int64
	v    int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Item))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func dijkstra(a, b, c int64, limit int64) []int64 {
	dist := make([]int64, int(a))
	for i := range dist {
		dist[i] = 1<<63 - 1
	}

	dist[0] = 0

	pq := &PriorityQueue{
		{dist: 0, v: 0},
	}

	heap.Init(pq)
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Item)

		if cur.dist != dist[cur.v] {
			continue
		}

		if cur.dist > limit {
			break
		}

		next := int((int64(cur.v) + b) % a)
		newDist := cur.dist + b

		if newDist < dist[next] && newDist <= limit {
			dist[next] = newDist
			heap.Push(pq, Item{
				dist: newDist,
				v:    next,
			})
		}

		next = int((int64(cur.v) + c) % a)
		newDist = cur.dist + c

		if newDist < dist[next] && newDist <= limit {
			dist[next] = newDist
			heap.Push(pq, Item{
				dist: newDist,
				v:    next,
			})
		}
	}

	return dist
}

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/microwave
// Microwave - problem 19
func Microwave() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 64)

	line, _ = reader.ReadString('\n')
	fields := strings.Fields(line)

	a, _ := strconv.ParseInt(fields[0], 10, 64)
	b, _ := strconv.ParseInt(fields[1], 10, 64)
	c, _ := strconv.ParseInt(fields[2], 10, 64)

	g := gcd(gcd(a, b), c)

	a /= g
	b /= g
	c /= g

	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}

	limit := (n - 1) / g

	dist := dijkstra(a, b, c, limit)

	var ans int64
	for _, d := range dist {
		if d <= limit {
			ans += (limit-d)/a + 1
		}
	}

	writer.WriteString(strconv.FormatInt(ans, 10))
	writer.WriteByte('\n')
}
