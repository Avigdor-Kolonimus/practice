package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type PointCellColoring struct {
	x int
	y int
}

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/cell-coloring
// CellColoring - assignment 15
func CellColoring() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// path input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	x, y := 0, 0
	visited := make(map[PointCellColoring]int)
	visited[PointCellColoring{x, y}] = 1
	for _, c := range s {
		switch c {
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x--
		case 'R':
			x++
		}

		visited[PointCellColoring{x, y}]++
	}

	ans := 0
	for _, cnt := range visited {
		if cnt > 1 {
			ans++
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
