package first2023mobiledev

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type DSU struct {
	parent map[int64]int64
	size   map[int64]int
}

func NewDSU() *DSU {
	return &DSU{
		parent: make(map[int64]int64),
		size:   make(map[int64]int),
	}
}

func (d *DSU) add(x int64) {
	if _, ok := d.parent[x]; !ok {
		d.parent[x] = x
		d.size[x] = 1
	}
}

func (d *DSU) find(x int64) int64 {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) union(a, b int64) {
	d.add(a)
	d.add(b)

	rootA := d.find(a)
	rootB := d.find(b)

	if rootA == rootB {
		return
	}

	if d.size[rootA] < d.size[rootB] {
		rootA, rootB = rootB, rootA
	}

	d.parent[rootB] = rootA
	d.size[rootA] += d.size[rootB]
}

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/resource-downloading
// ResourceDownloading - problem 7
func ResourceDownloading() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	dsu := NewDSU()
	for range n {
		// U and V input
		line, _ = reader.ReadString('\n')
		parts := strings.Fields(line)

		u, _ := strconv.ParseInt(parts[0], 10, 64)
		v, _ := strconv.ParseInt(parts[1], 10, 64)

		dsu.union(u, v)
	}

	// T input
	line, _ = reader.ReadString('\n')
	t, _ := strconv.Atoi(strings.TrimSpace(line))

	for range t {
		// X and K input
		line, _ = reader.ReadString('\n')
		parts := strings.Fields(line)

		x, _ := strconv.ParseInt(parts[0], 10, 64)
		k, _ := strconv.Atoi(parts[1])

		line, _ = reader.ReadString('\n')
		warehouses := strings.Fields(line)

		xRoot := x

		if _, ok := dsu.parent[x]; ok {
			xRoot = dsu.find(x)
		}

		result := make([]string, 0, k)
		for j := 0; j < k; j++ {
			y, _ := strconv.ParseInt(warehouses[j], 10, 64)

			yRoot := y

			if _, ok := dsu.parent[y]; ok {
				yRoot = dsu.find(y)
			}

			if xRoot == yRoot {
				result = append(result, warehouses[j])
			}
		}

		writer.WriteString(strconv.Itoa(len(result)))

		for _, warehouse := range result {
			writer.WriteByte(' ')
			writer.WriteString(warehouse)
		}

		writer.WriteByte('\n')
	}
}
