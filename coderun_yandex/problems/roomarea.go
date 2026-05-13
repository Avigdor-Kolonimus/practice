package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type PointRoomArea struct {
	x, y int
}

// https://coderun.yandex.ru/problem/room-area
// RoomArea - problem 38
func RoomArea() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// rectangle inputs
	rectangle := make([][]int, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		rectangle[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if line[j] == '*' {
				rectangle[i][j] = 1
			}
		}
	}

	// coordinate input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	x, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	x--
	y--

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	queue := []PointRoomArea{{x, y}}
	visited[x][y] = true

	area := 0

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		area++

		for k := 0; k < 4; k++ {
			nx := cur.x + dx[k]
			ny := cur.y + dy[k]

			if !visited[nx][ny] && rectangle[nx][ny] == 0 {
				visited[nx][ny] = true
				queue = append(queue, PointRoomArea{nx, ny})
			}
		}
	}

	writer.WriteString(strconv.Itoa(area))
	writer.WriteByte('\n')
}
