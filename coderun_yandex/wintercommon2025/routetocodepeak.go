package wintercommon2025

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

type Point struct {
	x, y, z int
	idx     int
}

type Edge struct {
	from, to int
	weight   int
}

func absRouteToCodePeak(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

func union(parent, rank []int, x, y int) {
	if rank[x] < rank[y] {
		parent[x] = y
	} else if rank[x] > rank[y] {
		parent[y] = x
	} else {
		parent[y] = x
		rank[x]++
	}
}

func solveRouteToCodePeak(points []Point) int64 {
	N := len(points)

	edges := make([]Edge, 0, 3*N)

	sortedByX := make([]Point, N)
	copy(sortedByX, points)
	sort.Slice(sortedByX, func(i, j int) bool {
		return sortedByX[i].x < sortedByX[j].x
	})
	for i := 0; i < N-1; i++ {
		cost := min(
			absRouteToCodePeak(sortedByX[i].x-sortedByX[i+1].x),
			min(
				absRouteToCodePeak(sortedByX[i].y-sortedByX[i+1].y),
				absRouteToCodePeak(sortedByX[i].z-sortedByX[i+1].z),
			),
		)
		edges = append(edges, Edge{
			from:   sortedByX[i].idx,
			to:     sortedByX[i+1].idx,
			weight: cost,
		})
	}

	sortedByY := make([]Point, N)
	copy(sortedByY, points)
	sort.Slice(sortedByY, func(i, j int) bool {
		return sortedByY[i].y < sortedByY[j].y
	})
	for i := 0; i < N-1; i++ {
		cost := min(
			absRouteToCodePeak(sortedByY[i].x-sortedByY[i+1].x),
			min(
				absRouteToCodePeak(sortedByY[i].y-sortedByY[i+1].y),
				absRouteToCodePeak(sortedByY[i].z-sortedByY[i+1].z),
			),
		)
		edges = append(edges, Edge{
			from:   sortedByY[i].idx,
			to:     sortedByY[i+1].idx,
			weight: cost,
		})
	}

	sortedByZ := make([]Point, N)
	copy(sortedByZ, points)
	sort.Slice(sortedByZ, func(i, j int) bool {
		return sortedByZ[i].z < sortedByZ[j].z
	})
	for i := 0; i < N-1; i++ {
		cost := min(
			absRouteToCodePeak(sortedByZ[i].x-sortedByZ[i+1].x),
			min(
				absRouteToCodePeak(sortedByZ[i].y-sortedByZ[i+1].y),
				absRouteToCodePeak(sortedByZ[i].z-sortedByZ[i+1].z),
			),
		)
		edges = append(edges, Edge{
			from:   sortedByZ[i].idx,
			to:     sortedByZ[i+1].idx,
			weight: cost,
		})
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	parent := make([]int, N)
	rank := make([]int, N)
	for i := 0; i < N; i++ {
		parent[i] = i
		rank[i] = 0
	}

	var totalCost int64
	edgesUsed := 0

	for _, edge := range edges {
		if edgesUsed == N-1 {
			break
		}
		fromRoot := find(parent, edge.from)
		toRoot := find(parent, edge.to)
		if fromRoot != toRoot {
			union(parent, rank, fromRoot, toRoot)
			totalCost += int64(edge.weight)
			edgesUsed++
		}
	}

	return totalCost
}

// https://coderun.yandex.ru/selections/2025-winter-common/problems/route-to-code-peak
// RouteToCodePeak - problem 5
func RouteToCodePeak() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	firstLine := mustReadIntArray(reader, 1)
	n := firstLine[0]

	points := make([]Point, n)
	for i := range n {
		// X,Y and Z input
		parts := mustReadIntArray(reader, 3)
		x, y, z := parts[0], parts[1], parts[2]
		points[i] = Point{x: x, y: y, z: z, idx: i}
	}

	answer := solveRouteToCodePeak(points)

	writer.WriteString(strconv.FormatInt(answer, 10))
	writer.WriteByte('\n')
}
