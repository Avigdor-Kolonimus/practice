package mobiledev

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const size = 1001

func readInts(reader *bufio.Reader) []int {
	line, _ := reader.ReadString('\n')
	fields := strings.Fields(line)

	values := make([]int, len(fields))
	for i, field := range fields {
		values[i], _ = strconv.Atoi(field)
	}

	return values
}

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/keeping-in-touch
// KeepingInTouch - problem 19
func KeepingInTouch() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	matrix := make([][]bool, size)
	for i := range matrix {
		matrix[i] = make([]bool, size)
	}

	for i := 0; i < n; i++ {
		values := readInts(reader)
		x, y, r := values[0], values[1], values[2]

		for dx := -r; dx <= r; dx++ {
			for dy := -r; dy <= r; dy++ {
				nx, ny := x+dx, y+dy

				if nx >= 0 && nx < size &&
					ny >= 0 && ny < size {
					matrix[nx][ny] = true
				}
			}
		}
	}

	start := readInts(reader)
	end := readInts(reader)

	sx, sy := start[0], start[1]
	ex, ey := end[0], end[1]

	if !matrix[sx][sy] || !matrix[ex][ey] {
		writer.WriteString("0\n")
		return
	}

	visited := make([][]bool, size)
	for i := range visited {
		visited[i] = make([]bool, size)
	}

	queue := make([]int, 0, size*size)
	queue = append(queue, sx*size+sy)
	visited[sx][sy] = true

	dx := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}

	for head := 0; head < len(queue); head++ {
		curr := queue[head]
		x, y := curr/size, curr%size

		if x == ex && y == ey {
			writer.WriteString("1\n")
			return
		}

		for d := 0; d < 8; d++ {
			nx := x + dx[d]
			ny := y + dy[d]

			if nx < 0 || nx >= size ||
				ny < 0 || ny >= size {
				continue
			}

			if visited[nx][ny] || !matrix[nx][ny] {
				continue
			}

			visited[nx][ny] = true
			queue = append(queue, nx*size+ny)
		}
	}

	writer.WriteString("0\n")
}
