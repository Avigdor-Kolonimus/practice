package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type coordSpeleologistWay struct {
	Y int
	X int
	Z int
}

type entrySpeleologistWay struct {
	start coordSpeleologistWay
	dist  int
}

func bfsExitSpeleologistWay(n int, s coordSpeleologistWay, cave [][][]int) int {
	queue := []entrySpeleologistWay{{coordSpeleologistWay{Y: s.Y, X: s.X, Z: s.Z}, 0}}
	moves := []coordSpeleologistWay{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}, {-1, 0, 0}, {0, -1, 0}, {0, 0, -1}}

	visited := make([][][]bool, n)
	for z := 0; z < n; z++ {
		visited[z] = make([][]bool, n)
		for i := 0; i < n; i++ {
			visited[z][i] = make([]bool, n)
		}
	}
	visited[s.Z][s.Y][s.X] = true

	front := 0
	for front < len(queue) {
		cur := queue[front]
		front++
		if cur.start.Z == 0 {
			return cur.dist
		}

		for _, move := range moves {
			nextZ, nextY, nextX := cur.start.Z+move.Z, cur.start.Y+move.Y, cur.start.X+move.X
			if nextZ >= n || nextZ < 0 || nextX >= n || nextX < 0 || nextY >= n || nextY < 0 || visited[nextZ][nextY][nextX] {
				continue
			}

			if cave[nextZ][nextY][nextX] == 0 {
				if nextZ == 0 {
					return cur.dist + 1
				}

				visited[nextZ][nextY][nextX] = true
				queue = append(queue, entrySpeleologistWay{coordSpeleologistWay{nextY, nextX, nextZ}, cur.dist + 1})
			}
		}
	}
	return 0
}

// https://coderun.yandex.ru/problem/speleologist-way
// SpeleologistWay - problem 15
func SpeleologistWay() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// cave input
	var speleolog coordSpeleologistWay
	cave := make([][][]int, n)
	for z := 0; z < n; z++ {
		cave[z] = make([][]int, n)
		_, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		for i := 0; i < n; i++ {
			cave[z][i] = make([]int, n)
			line, err = reader.ReadString('\n')
			if err != nil && err != io.EOF {
				panic(err)
			}
			for j := 0; j < n; j++ {
				if line[j] == 'S' {
					if z == 0 {
						writer.WriteByte('0')
						writer.WriteByte('\n')
						return
					}
					speleolog = coordSpeleologistWay{Y: i, X: j, Z: z}
				}
				if line[j] == '.' {
					cave[z][i][j] = 0
				}
				if line[j] == '#' {
					cave[z][i][j] = 1
				}
			}
		}
	}

	result := bfsExitSpeleologistWay(n, speleolog, cave)

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
