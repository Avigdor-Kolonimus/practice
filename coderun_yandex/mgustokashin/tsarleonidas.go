package mgustokashin

import (
	"bufio"
	"container/list"
	"os"
	"strconv"
	"strings"
)

const (
	INF = int(1e9)
)

type State struct {
	x   int
	y   int
	dir int
}

var dx = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var dy = []int{-1, 0, 1, -1, 1, -1, 0, 1}

// https://coderun.yandex.ru/selections/mgustokashin/problems/tsar-leonidas
// TsarLeonidas - problem 8
func TsarLeonidas() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	readInts := func() []int {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		fields := strings.Fields(line)

		res := make([]int, len(fields))
		for i, s := range fields {
			res[i], _ = strconv.Atoi(s)
		}
		return res
	}

	// H and W input
	tmp := readInts()
	H, W := tmp[0], tmp[1]

	//  field input
	field := make([][]byte, H)
	for i := range H {
		line, _ := reader.ReadString('\n')
		field[i] = []byte(strings.TrimSpace(line))
	}

	// start input
	tmp = readInts()
	sx, sy := tmp[0]-1, tmp[1]

	// target input
	tmp = readInts()
	tx, ty := tmp[0]-1, tmp[1]

	sy = H - sy
	ty = H - ty

	id := func(x, y, d int) int {
		return ((y*W)+x)*8 + d
	}

	dist := make([]int, H*W*8)
	for i := range dist {
		dist[i] = INF
	}

	deq := list.New()

	for d := range 8 {
		v := id(sx, sy, d)
		dist[v] = 1
		deq.PushFront(State{sx, sy, d})
	}

	for deq.Len() > 0 {
		e := deq.Front()
		deq.Remove(e)

		cur := e.Value.(State)
		curID := id(cur.x, cur.y, cur.dir)

		// move forward
		nx := cur.x + dy[cur.dir]
		ny := cur.y + dx[cur.dir]

		if nx >= 0 && nx < W &&
			ny >= 0 && ny < H &&
			field[ny][nx] == '.' {

			to := id(nx, ny, cur.dir)
			if dist[to] > dist[curID] {
				dist[to] = dist[curID]
				deq.PushFront(State{nx, ny, cur.dir})
			}
		}

		// turn
		for nd := 0; nd < 8; nd++ {
			if nd == cur.dir {
				continue
			}

			to := id(cur.x, cur.y, nd)
			if dist[to] > dist[curID]+1 {
				dist[to] = dist[curID] + 1
				deq.PushBack(State{cur.x, cur.y, nd})
			}
		}
	}

	ans := INF
	for d := range 8 {
		if dist[id(tx, ty, d)] < ans {
			ans = dist[id(tx, ty, d)]
		}
	}

	if ans == INF {
		writer.WriteString("-1\n")
	} else {
		writer.WriteString(strconv.Itoa(ans))
		writer.WriteByte('\n')
	}
}
