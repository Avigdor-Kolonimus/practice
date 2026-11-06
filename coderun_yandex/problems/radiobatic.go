package problems

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int64
}

type Cell struct {
	x, y int64
}

var (
	n      int
	points []Point
)

func checkRadioBatic(r float64) (bool, []int) {
	col := make([]int, n)

	D := 2.0 * r
	lim := D * D

	if D <= 1.0 {
		return true, col
	}

	cellSize := D / 2.0

	mp := make(map[Cell][]int, n)
	cells := make([]Cell, n)
	for i := 0; i < n; i++ {
		cx := int64(math.Floor(float64(points[i].x) / cellSize))
		cy := int64(math.Floor(float64(points[i].y) / cellSize))

		c := Cell{cx, cy}
		cells[i] = c

		mp[c] = append(mp[c], i)

		if len(mp[c]) >= 3 {
			return false, nil
		}
	}

	for s := 0; s < n; s++ {
		if col[s] != 0 {
			continue
		}
		break
	}

	for i := range col {
		col[i] = -1
	}

	queue := make([]int, n)
	for s := 0; s < n; s++ {
		if col[s] != -1 {
			continue
		}

		head, tail := 0, 0

		queue[tail] = s
		tail++
		col[s] = 0

		for head < tail {
			v := queue[head]
			head++

			cx := cells[v].x
			cy := cells[v].y

			for dx := int64(-2); dx <= 2; dx++ {
				for dy := int64(-2); dy <= 2; dy++ {
					c := Cell{
						cx + dx,
						cy + dy,
					}

					neighbors, exists := mp[c]
					if !exists {
						continue
					}

					for _, u := range neighbors {
						if u == v {
							continue
						}

						ddx := points[v].x - points[u].x
						ddy := points[v].y - points[u].y

						d2 := float64(ddx*ddx + ddy*ddy)

						if d2 < lim {
							switch col[u] {
							case -1:
								col[u] = col[v] ^ 1

								queue[tail] = u
								tail++
							case col[v]:
								return false, nil
							default:
								if col[u] == col[v] {
									return false, nil
								}
							}
						}
					}
				}
			}
		}
	}

	return true, col
}

// https://coderun.yandex.ru/problem/radio-batic
// RadioBatic - problem 40
func RadioBatic() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ = strconv.Atoi(strings.TrimSpace(line))

	points = make([]Point, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &points[i].x, &points[i].y)
	}

	left := 0.0
	right := 300000.0
	best := make([]int, n)
	for i := range best {
		best[i] = 0
	}

	for it := 0; it < 80; it++ {
		mid := (left + right) / 2.0

		ok, col := checkRadioBatic(mid)

		if ok {
			left = mid
			best = col
		} else {
			right = mid
		}
	}

	writer.WriteString(strconv.FormatFloat(left, 'f', 15, 64))
	writer.WriteByte('\n')

	for i := 0; i < n; i++ {
		if i > 0 {
			writer.WriteByte(' ')
		}

		writer.WriteString(strconv.Itoa(best[i] + 1))
	}

	writer.WriteByte('\n')
}
