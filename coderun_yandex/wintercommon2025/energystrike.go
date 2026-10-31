package wintercommon2025

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

const eps = 1e-9

type PointEnergyStrike struct {
	x, y float64
}

type Candidate struct {
	p    PointEnergyStrike
	mask int
	gain int
}

func dist2(a, b PointEnergyStrike) float64 {
	dx := a.x - b.x
	dy := a.y - b.y
	return dx*dx + dy*dy
}

func intersections(a PointEnergyStrike, r1 float64, b PointEnergyStrike, r2 float64) []PointEnergyStrike {
	d2 := dist2(a, b)
	d := math.Sqrt(d2)

	if d < 1e-12 || d > r1+r2+eps || d < math.Abs(r1-r2)-eps {
		return nil
	}

	x := (r1*r1 - r2*r2 + d2) / (2 * d)
	h2 := r1*r1 - x*x

	if h2 < -eps {
		return nil
	}
	if h2 < 0 {
		h2 = 0
	}

	h := math.Sqrt(h2)

	mx := a.x + x*(b.x-a.x)/d
	my := a.y + x*(b.y-a.y)/d

	p1 := PointEnergyStrike{
		x: mx + h*(b.y-a.y)/d,
		y: my - h*(b.x-a.x)/d,
	}

	if h < 1e-12 {
		return []PointEnergyStrike{p1}
	}

	p2 := PointEnergyStrike{
		x: mx - h*(b.y-a.y)/d,
		y: my + h*(b.x-a.x)/d,
	}

	return []PointEnergyStrike{p1, p2}
}

func buildStaticCandidates(points []PointEnergyStrike) []PointEnergyStrike {
	n := len(points)
	candidates := make([]PointEnergyStrike, 0, 100)

	for i := 0; i < n; i++ {
		candidates = append(candidates, points[i])
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			candidates = append(
				candidates,
				intersections(points[i], 1, points[j], 1)...,
			)
		}
	}

	return candidates
}

func solveComponent(points []PointEnergyStrike) ([]PointEnergyStrike, bool) {
	n := len(points)
	fullMask := (1 << n) - 1

	static := buildStaticCandidates(points)
	solution := make([]PointEnergyStrike, 0, n)

	var dfs func(int) bool

	dfs = func(mask int) bool {
		if mask == fullMask {
			return true
		}

		u := 0
		for mask&(1<<u) != 0 {
			u++
		}

		candidates := make([]Candidate, 0, 300)

		addCandidate := func(c PointEnergyStrike) {
			if dist2(c, points[u]) > 1+eps {
				return
			}

			for _, old := range solution {
				if dist2(c, old) < 4-eps {
					return
				}
			}

			newMask := mask
			for i := 0; i < n; i++ {
				if dist2(c, points[i]) <= 1+eps {
					newMask |= 1 << i
				}
			}

			if newMask == mask {
				return
			}

			for _, old := range candidates {
				if dist2(c, old.p) < 1e-18 {
					return
				}
			}

			gain := 0
			for x := newMask ^ mask; x != 0; x &= x - 1 {
				gain++
			}

			candidates = append(candidates, Candidate{
				p:    c,
				mask: newMask,
				gain: gain,
			})
		}

		for _, c := range static {
			addCandidate(c)
		}

		k := len(solution)

		for i := 0; i < k; i++ {
			for j := 0; j < n; j++ {
				for _, c := range intersections(points[j], 1, solution[i], 2) {
					addCandidate(c)
				}
			}
		}

		for i := 0; i < k; i++ {
			for j := i + 1; j < k; j++ {
				for _, c := range intersections(solution[i], 2, solution[j], 2) {
					addCandidate(c)
				}
			}
		}

		sort.Slice(candidates, func(i, j int) bool {
			return candidates[i].gain > candidates[j].gain
		})

		for _, c := range candidates {
			solution = append(solution, c.p)

			if dfs(c.mask) {
				return true
			}

			solution = solution[:len(solution)-1]
		}

		return false
	}

	if dfs(0) {
		result := make([]PointEnergyStrike, len(solution))
		copy(result, solution)
		return result, true
	}

	return nil, false
}

func solveEnergyStrike(points []PointEnergyStrike) ([]PointEnergyStrike, bool) {
	n := len(points)
	visited := make([]bool, n)

	answer := make([]PointEnergyStrike, 0, n)

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}

		queue := []int{i}
		visited[i] = true

		component := make([]PointEnergyStrike, 0, n)

		for head := 0; head < len(queue); head++ {
			u := queue[head]
			component = append(component, points[u])

			for v := 0; v < n; v++ {
				if !visited[v] && dist2(points[u], points[v]) <= 16+eps {
					visited[v] = true
					queue = append(queue, v)
				}
			}
		}

		res, ok := solveComponent(component)
		if !ok {
			return nil, false
		}

		answer = append(answer, res...)
	}

	return answer, true
}

// https://coderun.yandex.ru/selections/2025-winter-common/problems/energy-strike
// EnergyStrike - problem 21
func EnergyStrike() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20)
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)

		points := make([]PointEnergyStrike, n)

		for i := 0; i < n; i++ {
			fmt.Fscan(in, &points[i].x, &points[i].y)
		}

		answer, ok := solveEnergyStrike(points)

		if !ok {
			fmt.Fprintln(out, "NO")
			continue
		}

		fmt.Fprintln(out, "YES")
		fmt.Fprintln(out, len(answer))

		for _, c := range answer {
			fmt.Fprintf(out, "%.15f %.15f\n", c.x, c.y)
		}
	}
}
