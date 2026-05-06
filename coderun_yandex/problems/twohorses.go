package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var moves = [8][2]int{
	{2, 1}, {2, -1}, {-2, 1}, {-2, -1},
	{1, 2}, {1, -2}, {-1, 2}, {-1, -2},
}

type State struct {
	rx, ry int
	gx, gy int
}

func inside(x, y int) bool {
	return x >= 0 && x < 8 && y >= 0 && y < 8
}

func color(x, y int) int {
	return (x + y) % 2
}

func parse(s string) (int, int) {
	x := int(s[0] - 'a')
	y := int(s[1] - '1')

	return x, y
}

// https://coderun.yandex.ru/problem/two-horses
// TwoHorses - problem 43
func TwoHorses() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// horses input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	parameters := strings.Fields(line)
	if len(parameters) != 2 {
		panic("input does not match 2")
	}

	// horse 1
	rx, ry := parse(parameters[0])

	// horse 2
	gx, gy := parse(parameters[1])

	// black and white
	result := -1
	if color(rx, ry) != color(gx, gy) {
		writer.WriteString(strconv.Itoa(result))
		writer.WriteByte('\n')
		return
	}

	visited := make(map[State]bool)
	queue := []State{{rx, ry, gx, gy}}
	visited[queue[0]] = true

	steps := 0
	for len(queue) > 0 {
		next := []State{}

		for _, cur := range queue {
			// meet up
			if cur.rx == cur.gx && cur.ry == cur.gy {
				writer.WriteString(strconv.Itoa(steps))
				writer.WriteByte('\n')
				return
			}

			// moves
			for _, m1 := range moves {
				nrx, nry := cur.rx+m1[0], cur.ry+m1[1]
				if !inside(nrx, nry) {
					continue
				}

				for _, m2 := range moves {
					ngx, ngy := cur.gx+m2[0], cur.gy+m2[1]
					if !inside(ngx, ngy) {
						continue
					}

					ns := State{nrx, nry, ngx, ngy}
					if !visited[ns] {
						visited[ns] = true
						next = append(next, ns)
					}
				}
			}
		}

		queue = next
		steps++
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
