package first2023dataanalytics

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const INF_CORNERS int = 1 << 30

type Point struct {
	x, y int
	t    int
}

// https://coderun.yandex.ru/selections/first-2023-data-analytics/problems/corners
// Corners - problem 2
func Corners() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	points := make([]Point, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &points[i].x, &points[i].y)

		var v, h int
		fmt.Fscan(reader, &v, &h)

		switch {
		case v == 1 && h == 1:
			points[i].t = 0 // UR
		case v == 1 && h == -1:
			points[i].t = 1 // UL
		case v == -1 && h == 1:
			points[i].t = 2 // DR
		default:
			points[i].t = 3 // DL
		}
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].x < points[j].x
	})

	xs := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if i == 0 || points[i].x != points[i-1].x {
			xs = append(xs, points[i].x)
		}
	}

	m := len(xs)

	urPref := make([]int, m)
	drPref := make([]int, m)
	ulSuff := make([]int, m)
	dlSuff := make([]int, m)

	for i := range m {
		urPref[i] = INF_CORNERS
		drPref[i] = -INF_CORNERS
		ulSuff[i] = INF_CORNERS
		dlSuff[i] = -INF_CORNERS
	}

	idx := 0
	minUR, maxDR := INF_CORNERS, -INF_CORNERS
	for i := range m {
		for idx < n && points[idx].x == xs[i] {
			p := points[idx]

			if p.t == 0 && p.y < minUR {
				minUR = p.y
			}
			if p.t == 2 && p.y > maxDR {
				maxDR = p.y
			}

			idx++
		}

		urPref[i] = minUR
		drPref[i] = maxDR
	}

	idx = n - 1
	minUL, maxDL := INF_CORNERS, -INF_CORNERS
	for i := m - 1; i >= 0; i-- {
		for idx >= 0 && points[idx].x == xs[i] {
			p := points[idx]

			if p.t == 1 && p.y < minUL {
				minUL = p.y
			}
			if p.t == 3 && p.y > maxDL {
				maxDL = p.y
			}

			idx--
		}

		ulSuff[i] = minUL
		dlSuff[i] = maxDL
	}

	check := func(l, u int) bool {
		return l < u
	}

	if check(dlSuff[0], ulSuff[0]) {
		writer.WriteString("NO\n")
		return
	}

	for i := range m {
		u := min(urPref[i], ulSuff[i])
		l := max(drPref[i], dlSuff[i])

		if check(l, u) {
			writer.WriteString("NO\n")
			return
		}

		if i+1 < m {
			uMid := min(urPref[i], ulSuff[i+1])
			lMid := max(drPref[i], dlSuff[i+1])

			if check(lMid, uMid) {
				writer.WriteString("NO\n")
				return
			}
		}
	}

	if check(drPref[m-1], urPref[m-1]) {
		writer.WriteString("NO\n")
		return
	}

	writer.WriteString("YES\n")
}
