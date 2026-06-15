package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	m       int
	grid    []string
	visited [][]bool
)

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

type Point struct {
	x int
	y int
}

func bfsCountingSquares(sx, sy int) {
	q := []Point{{sx, sy}}
	visited[sx][sy] = true

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for d := 0; d < 4; d++ {
			nx := cur.x + dx[d]
			ny := cur.y + dy[d]

			if nx < 0 || nx >= n || ny < 0 || ny >= m {
				continue
			}

			if visited[nx][ny] || grid[nx][ny] != '#' {
				continue
			}

			visited[nx][ny] = true
			q = append(q, Point{nx, ny})
		}
	}
}

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/counting-squares
// CountingSquares - assignment 19
func CountingSquares() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	n, err = strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	m, err = strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// field input
	field := make([]string, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}

		field[i] = strings.TrimRight(line, "\r\n")
	}

	visited = make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '#' && !visited[i][j] {
				ans++
				bfsCountingSquares(i, j)
			}
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
