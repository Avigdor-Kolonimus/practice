package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type StateToyMaze struct {
	x, y int
	dist int
}

var dxToyMaze = []int{-1, 1, 0, 0}
var dyToyMaze = []int{0, 0, -1, 1}

func move(x, y, dir int, grid [][]int, n, m int) (int, int, bool) {
	nx, ny := x, y

	for {
		tx := nx + dxToyMaze[dir]
		ty := ny + dyToyMaze[dir]

		// wall of the maze
		if tx < 0 || tx >= n || ty < 0 || ty >= m {
			return nx, ny, false
		}

		// barrier
		if grid[tx][ty] == 1 {
			return nx, ny, false
		}

		// hole
		if grid[tx][ty] == 2 {
			return tx, ty, true
		}

		nx = tx
		ny = ty
	}
}

// https://coderun.yandex.ru/problem/toy-maze
// ToyMaze - problem 45
func ToyMaze() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	tokens := strings.Fields(line)
	if len(tokens) != 2 {
		panic("invalid input")
	}

	n, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}

	// graph input
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		tokens := strings.Fields(line)
		if len(tokens) != m {
			panic("invalid input")
		}

		graph[i] = make([]int, m)
		for j := 0; j < m; j++ {
			x, err := strconv.Atoi(tokens[j])
			if err != nil {
				panic(err)
			}

			graph[i][j] = x
		}
	}

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	queue := []StateToyMaze{{0, 0, 0}}
	visited[0][0] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for dir := 0; dir < 4; dir++ {
			nx, ny, escaped := move(cur.x, cur.y, dir, graph, n, m)

			// the ball fell into the hole
			if escaped {
				writer.WriteString(strconv.Itoa(cur.dist + 1))
				writer.WriteByte('\n')

				return
			}

			if !visited[nx][ny] {
				visited[nx][ny] = true
				queue = append(queue, StateToyMaze{
					x:    nx,
					y:    ny,
					dist: cur.dist + 1,
				})
			}
		}
	}
}
