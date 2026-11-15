package mgustokashin

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

type Point struct {
	x, y int
}

type Triangle struct {
	mask   uint64
	coords [6]int
}

var (
	n, m, k int
	s       int

	mini    [][3]Point
	tris    []Triangle
	buckets [][]int

	memo   map[uint64]int
	choice map[uint64]int

	full uint64
)

func cross(a, b, c Point) int {
	return (b.x-a.x)*(c.y-a.y) -
		(b.y-a.y)*(c.x-a.x)
}

func inside(p, a, b, c Point) bool {
	c1 := cross(a, b, p)
	c2 := cross(b, c, p)
	c3 := cross(c, a, p)

	return (c1 >= 0 && c2 >= 0 && c3 >= 0) ||
		(c1 <= 0 && c2 <= 0 && c3 <= 0)
}

func getMask(a, b, c Point) uint64 {
	var mask uint64

	for i := 0; i < s; i++ {
		t := mini[i]

		if inside(t[0], a, b, c) &&
			inside(t[1], a, b, c) &&
			inside(t[2], a, b, c) {
			mask |= uint64(1) << i
		}
	}

	return mask
}

func check(dx, dy int) bool {
	return dx == 0 || dy == 0 || abs(dx) == abs(dy)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dist2(a, b Point) int {
	dx := a.x - b.x
	dy := a.y - b.y

	return dx*dx + dy*dy
}

func dfs(ms uint64) int {
	if ms == full {
		return 0
	}

	if result, ok := memo[ms]; ok {
		return result
	}

	bit := bits.TrailingZeros64(^ms & full)

	best := 1000
	bestIdx := -1

	for _, idx := range buckets[bit] {
		mask := tris[idx].mask

		if mask&ms != 0 {
			continue
		}

		cur := 1 + dfs(ms|mask)

		if cur < best {
			best = cur
			bestIdx = idx
		}
	}

	memo[ms] = best

	if bestIdx >= 0 {
		choice[ms] = bestIdx
	}

	return best
}

// https://coderun.yandex.ru/selections/mgustokashin/problems/echpochmak-packing
// EchpochmakPacking - problem 10
func EchpochmakPacking() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	fmt.Fscan(reader, &n, &m, &k)

	s = 4 * n * m

	if s == 64 {
		full = ^uint64(0)
	} else {
		full = (uint64(1) << s) - 1
	}

	mini = make([][3]Point, s)
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			bl := Point{2 * i, 2 * j}
			br := Point{2*i + 2, 2 * j}
			tl := Point{2 * i, 2*j + 2}
			tr := Point{2*i + 2, 2*j + 2}
			cc := Point{2*i + 1, 2*j + 1}

			base := 4 * (j*n + i)

			mini[base] = [3]Point{bl, br, cc}
			mini[base+1] = [3]Point{br, tr, cc}
			mini[base+2] = [3]Point{tr, tl, cc}
			mini[base+3] = [3]Point{tl, bl, cc}
		}
	}

	var occupied uint64
	for i := 0; i < k; i++ {
		var a, b, c Point

		fmt.Fscan(reader,
			&a.x, &a.y,
			&b.x, &b.y,
			&c.x, &c.y,
		)

		a.x *= 2
		a.y *= 2
		b.x *= 2
		b.y *= 2
		c.x *= 2
		c.y *= 2

		occupied |= getMask(a, b, c)
	}

	pts := make([]Point, 0, (n+1)*(m+1))
	for y := 0; y <= m; y++ {
		for x := 0; x <= n; x++ {
			pts = append(pts, Point{2 * x, 2 * y})
		}
	}

	seen := make(map[uint64]struct{})
	for i := 0; i < len(pts); i++ {
		for j := i + 1; j < len(pts); j++ {
			for l := j + 1; l < len(pts); l++ {
				a := pts[i]
				b := pts[j]
				c := pts[l]

				if cross(a, b, c) == 0 {
					continue
				}

				if !check(b.x-a.x, b.y-a.y) ||
					!check(c.x-b.x, c.y-b.y) ||
					!check(a.x-c.x, a.y-c.y) {
					continue
				}

				d1 := dist2(a, b)
				d2 := dist2(b, c)
				d3 := dist2(c, a)

				if d1 != d2 && d2 != d3 && d1 != d3 {
					continue
				}

				mask := getMask(a, b, c)

				if mask == 0 {
					continue
				}

				if _, ok := seen[mask]; ok {
					continue
				}

				seen[mask] = struct{}{}

				tris = append(tris, Triangle{
					mask: mask,
					coords: [6]int{
						a.x / 2, a.y / 2,
						b.x / 2, b.y / 2,
						c.x / 2, c.y / 2,
					},
				})
			}
		}
	}

	buckets = make([][]int, s)
	for i, t := range tris {
		mask := t.mask

		for mask != 0 {
			bit := bits.TrailingZeros64(mask)

			buckets[bit] = append(buckets[bit], i)

			mask &= mask - 1
		}
	}

	memo = make(map[uint64]int)
	choice = make(map[uint64]int)

	need := dfs(occupied)

	answer := make([][6]int, 0, need)
	cur := occupied
	for cur != full {
		idx, ok := choice[cur]

		if !ok {
			fmt.Fprintln(writer, -1)
			return
		}

		answer = append(answer, tris[idx].coords)
		cur |= tris[idx].mask
	}

	fmt.Fprintln(writer, need)

	for _, coords := range answer {
		fmt.Fprintf(writer,
			"%d %d %d %d %d %d\n",
			coords[0], coords[1],
			coords[2], coords[3],
			coords[4], coords[5],
		)
	}
}
