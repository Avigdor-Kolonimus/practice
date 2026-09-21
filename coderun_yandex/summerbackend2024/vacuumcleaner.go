package summerbackend2024

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var (
	matrix [][]float64
)

func search(diameter float64, start, finish int) bool {
	used := make([]bool, len(matrix))
	queue := make([]int, 0, len(matrix))
	queue = append(queue, start)
	used[start] = true

	head := 0
	for head < len(queue) {
		v := queue[head]
		head++

		if v == finish {
			return false
		}

		for to, distance := range matrix[v] {
			if used[to] {
				continue
			}
			if distance >= diameter || (v >= n && to >= n && distance == 0) {
				continue
			}
			used[to] = true
			queue = append(queue, to)
		}
	}
	return true
}

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/vacuum-cleaner
// VacuumCleaner - problem 6
func VacuumCleaner() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20)
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	fmt.Fscan(in, &n)
	var m int
	fmt.Fscan(in, &m)

	var w, h int
	fmt.Fscan(in, &w, &h)

	type mebelT struct{ x, y, r float64 }
	mebel := make([]mebelT, n)
	for i := 0; i < n; i++ {
		var x, y, r int
		fmt.Fscan(in, &x, &y, &r)
		mebel[i] = mebelT{float64(x), float64(y), float64(r)}
	}

	size := n + 4
	matrix = make([][]float64, size)
	for i := range matrix {
		matrix[i] = make([]float64, size)
	}

	down := n
	top := n + 1
	left := n + 2
	right := n + 3

	matrix[down][top] = float64(h)
	matrix[top][down] = float64(h)
	matrix[left][right] = float64(w)
	matrix[right][left] = float64(w)

	for i := 0; i < n; i++ {
		xi, yi, rm := mebel[i].x, mebel[i].y, mebel[i].r

		v := math.Max(0, yi-rm)
		matrix[i][down] = v
		matrix[down][i] = v

		v = math.Max(0, float64(h)-yi-rm)
		matrix[i][top] = v
		matrix[top][i] = v

		v = math.Max(0, xi-rm)
		matrix[i][left] = v
		matrix[left][i] = v

		v = math.Max(0, float64(w)-xi-rm)
		matrix[i][right] = v
		matrix[right][i] = v

		for j := 0; j < n; j++ {
			dx := mebel[i].x - mebel[j].x
			dy := mebel[i].y - mebel[j].y
			r := mebel[i].r + mebel[j].r
			d := math.Max(0, math.Hypot(dx, dy)-r)
			matrix[i][j] = d
			matrix[j][i] = d
		}
	}

	prov := [][2]int{
		{left, right},
		{left, down},
		{left, top},
		{right, top},
		{right, down},
		{down, top},
	}
	dist := make([]float64, 6)
	for i := 0; i < len(prov); i++ {
		dLeft, dRight := 0, max(w, h)
		for dLeft+1 < dRight {
			dMid := (dLeft + dRight) / 2
			if search(float64(dMid), prov[i][0], prov[i][1]) {
				dLeft = dMid
			} else {
				dRight = dMid
			}
		}

		dd := int(float64(dLeft) * 1e4)
		if dd%10 >= 5 {
			dd += 10
		}
		dd -= dd % 10
		dist[i] = float64(dd) / 1e4
	}

	for q := 0; q < m; q++ {
		var r, s int
		fmt.Fscan(in, &r, &s)

		temp := []int{}
		twoR := float64(r * 2)

		if s == 1 {
			temp = append(temp, 1)
			if dist[5] >= twoR && dist[4] >= twoR && dist[1] >= twoR {
				temp = append(temp, 2)
			}
			if dist[1] >= twoR && dist[5] >= twoR && dist[3] >= twoR && dist[0] >= twoR {
				temp = append(temp, 3)
			}
			if dist[0] >= twoR && dist[1] >= twoR && dist[2] >= twoR {
				temp = append(temp, 4)
			}
		}

		if s == 2 {
			temp = append(temp, 2)
			if dist[5] >= twoR && dist[1] >= twoR && dist[4] >= twoR {
				temp = append(temp, 1)
			}
			if dist[0] >= twoR && dist[5] >= twoR && dist[2] >= twoR && dist[4] >= twoR {
				temp = append(temp, 4)
			}
			if dist[0] >= twoR && dist[3] >= twoR && dist[4] >= twoR {
				temp = append(temp, 3)
			}
		}

		if s == 3 {
			temp = append(temp, 3)
			if dist[5] >= twoR && dist[1] >= twoR && dist[0] >= twoR && dist[3] >= twoR {
				temp = append(temp, 1)
			}
			if dist[5] >= twoR && dist[2] >= twoR && dist[3] >= twoR {
				temp = append(temp, 4)
			}
			if dist[0] >= twoR && dist[4] >= twoR && dist[3] >= twoR {
				temp = append(temp, 2)
			}
		}

		if s == 4 {
			temp = append(temp, 4)
			if dist[1] >= twoR && dist[0] >= twoR && dist[2] >= twoR {
				temp = append(temp, 1)
			}
			if dist[5] >= twoR && dist[0] >= twoR && dist[4] >= twoR && dist[2] >= twoR {
				temp = append(temp, 2)
			}
			if dist[5] >= twoR && dist[3] >= twoR && dist[2] >= twoR {
				temp = append(temp, 3)
			}
		}

		sort.Ints(temp)
		for _, v := range temp {
			fmt.Fprint(out, v)
		}
		fmt.Fprintln(out)
	}
}
