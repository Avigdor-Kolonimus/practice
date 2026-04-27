package wintercommon2025

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type EdgeLabyrinthUnderMountain struct {
	a, b, w int
}

func findLabyrinthUnderMountain(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = findLabyrinthUnderMountain(parent, parent[x])
	}
	return parent[x]
}

func unionLabyrinthUnderMountain(parent, rank, size []int, x, y int) {
	if rank[x] < rank[y] {
		parent[x] = y
		size[y] += size[x]
	} else if rank[x] > rank[y] {
		parent[y] = x
		size[x] += size[y]
	} else {
		parent[y] = x
		size[x] += size[y]
		rank[x]++
	}
}

func solveLabyrinthUnderMountain(n int, edges []EdgeLabyrinthUnderMountain) []int {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	parent := make([]int, n)
	rank := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = -1
	}

	result[0] = 0

	maxReached := 1

	for _, e := range edges {
		if e.a == e.b {
			continue
		}

		rootA := findLabyrinthUnderMountain(parent, e.a)
		rootB := findLabyrinthUnderMountain(parent, e.b)

		if rootA != rootB {
			newSize := size[rootA] + size[rootB]
			unionLabyrinthUnderMountain(parent, rank, size, rootA, rootB)

			for k := maxReached + 1; k <= newSize; k++ {
				result[k-1] = e.w
			}
			if newSize > maxReached {
				maxReached = newSize
			}
		}

		if maxReached == n {
			break
		}
	}

	return result
}

// https://coderun.yandex.ru/selections/2025-winter-common/problems/labyrinth-under-mountain
// LabyrinthUnderMountain - problem 10
func LabyrinthUnderMountain() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// T input
	firstLine := mustReadIntArray(reader, 1)
	t := firstLine[0]

	for test := 0; test < t; test++ {
		// N and M input
		line := mustReadIntArray(reader, 2)
		n, m := line[0], line[1]

		// a, b, c input
		a := make([]int, m)
		b := make([]int, m)
		c := make([]int, m)

		for i := 0; i < m; i++ {
			line = mustReadIntArray(reader, 1)
			a[i] = line[0]
		}
		for i := 0; i < m; i++ {
			line = mustReadIntArray(reader, 1)
			b[i] = line[0]
		}
		for i := 0; i < m; i++ {
			line = mustReadIntArray(reader, 1)
			c[i] = line[0]
		}

		edges := make([]EdgeLabyrinthUnderMountain, m)
		for i := 0; i < m; i++ {
			edges[i] = EdgeLabyrinthUnderMountain{a: a[i] - 1, b: b[i] - 1, w: c[i]} // 0-indexed
		}

		result := solveLabyrinthUnderMountain(n, edges)

		sb := strings.Builder{}
		for i, v := range result {
			if i > 0 {
				sb.WriteByte(' ')
			}
			sb.WriteString(strconv.Itoa(v))
		}
		sb.WriteByte('\n')
		writer.WriteString(sb.String())
	}
}
