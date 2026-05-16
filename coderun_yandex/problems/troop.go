package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type PointTroop struct {
	x, y int
}

func bfsTroop(sx, sy int, id int, grid [][]int, comp [][]int, n, m int) {
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	queue := []PointTroop{{sx, sy}}
	comp[sx][sy] = id

	for head := 0; head < len(queue); head++ {
		cur := queue[head]

		for d := 0; d < 4; d++ {
			nx := cur.x + dx[d]
			ny := cur.y + dy[d]

			if nx < 0 || ny < 0 || nx >= n || ny >= m {
				continue
			}

			if comp[nx][ny] != -1 {
				continue
			}

			if grid[nx][ny] != grid[cur.x][cur.y] {
				continue
			}

			comp[nx][ny] = id
			queue = append(queue, PointTroop{nx, ny})
		}
	}
}

// https://coderun.yandex.ru/problem/troop
// Troop - problem 41
func Troop() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 3")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}

	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// height board inputs
	heightBoard := make([][]int, n)
	for i := range heightBoard {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
		if len(strNum) != m {
			panic("numbers count does not match m")
		}

		heightBoard[i] = make([]int, m)
		for j := 0; j < m; j++ {
			h, err := strconv.Atoi(strNum[j])
			if err != nil {
				panic(err)
			}

			heightBoard[i][j] = h
		}
	}

	comp := make([][]int, n)
	for i := 0; i < n; i++ {
		comp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			comp[i][j] = -1
		}
	}

	componentID := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if comp[i][j] == -1 {
				bfsTroop(i, j, componentID, heightBoard, comp, n, m)
				componentID++
			}
		}
	}

	hasLower := make([]bool, componentID)

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for d := 0; d < 4; d++ {
				ni := i + dx[d]
				nj := j + dy[d]

				if ni < 0 || nj < 0 || ni >= n || nj >= m {
					continue
				}

				if heightBoard[ni][nj] < heightBoard[i][j] {
					hasLower[comp[i][j]] = true
				}
			}
		}
	}

	answer := 0
	for i := 0; i < componentID; i++ {
		if !hasLower[i] {
			answer++
		}
	}

	writer.WriteString(strconv.Itoa(answer))
	writer.WriteByte('\n')
}
