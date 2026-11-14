package problems

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type BlockNguBuilding struct {
	z0, z1 int64
	area   int64
}

type EventNguBuilding struct {
	z     int64
	index int
	start bool
}

// https://coderun.yandex.ru/problem/ngu-building
// NguBuilding - problem 84
func NguBuilding() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	n, _ := strconv.Atoi(parts[0])
	w, _ := strconv.ParseInt(parts[1], 10, 64)
	l, _ := strconv.ParseInt(parts[2], 10, 64)

	target := w * l

	blocks := make([]BlockNguBuilding, n)
	events := make([]EventNguBuilding, 0, 2*n)

	for i := 0; i < n; i++ {
		var x0, y0, z0, x1, y1, z1 int64

		fmt.Fscan(reader, &x0, &y0, &z0, &x1, &y1, &z1)

		area := (x1 - x0) * (y1 - y0)

		blocks[i] = BlockNguBuilding{
			z0:   z0,
			z1:   z1,
			area: area,
		}

		events = append(events,
			EventNguBuilding{z: z0, index: i, start: true},
			EventNguBuilding{z: z1, index: i, start: false},
		)
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].z < events[j].z
	})

	var square int64
	var count int

	bestCount := n + 1
	var bestZ int64
	found := false

	var prevZ int64
	hasPrev := false
	for i := 0; i < len(events); {
		z := events[i].z

		if hasPrev && prevZ < z && square == target {
			if count < bestCount {
				bestCount = count
				bestZ = prevZ
				found = true
			}
		}

		for i < len(events) && events[i].z == z {
			e := events[i]
			b := blocks[e.index]

			if e.start {
				square += b.area
				count++
			} else {
				square -= b.area
				count--
			}

			i++
		}

		prevZ = z
		hasPrev = true
	}

	if !found {
		writer.WriteString("NO")
		writer.WriteByte('\n')
	} else {
		writer.WriteString("YES")
		writer.WriteByte('\n')
		writer.WriteString(strconv.Itoa(bestCount))
		writer.WriteByte('\n')

		first := true
		for i, b := range blocks {
			if b.z0 <= bestZ && b.z1 > bestZ {
				if !first {
					writer.WriteByte('\n')
				}

				writer.WriteString(strconv.Itoa(i + 1))
				first = false
			}
		}

		writer.WriteByte('\n')
	}
}
