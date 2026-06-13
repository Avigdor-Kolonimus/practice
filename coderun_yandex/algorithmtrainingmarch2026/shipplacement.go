package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"os"
	"strings"
)

const (
	N = 10
)

var (
	field [N]string
	used  [N][N]bool
	cnt   [5]int
)

func inside(x, y int) bool {
	return x >= 0 && x < N && y >= 0 && y < N
}

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/ship-placement
// ShipPlacement - assignment 24
func ShipPlacement() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// field input
	for i := 0; i < N; i++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		field[i] = strings.TrimRight(line, "\r\n")
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if field[i][j] != '#' || used[i][j] {
				continue
			}

			right := j+1 < N && field[i][j+1] == '#'
			down := i+1 < N && field[i+1][j] == '#'

			if right && down {
				writer.WriteString("NO")
				writer.WriteByte('\n')

				return
			}

			var cells [][2]int
			if right {
				k := j
				for k < N && field[i][k] == '#' {
					cells = append(cells, [2]int{i, k})
					k++
				}
			} else if down {
				k := i
				for k < N && field[k][j] == '#' {
					cells = append(cells, [2]int{k, j})
					k++
				}
			} else {
				cells = append(cells, [2]int{i, j})
			}

			for _, p := range cells {
				x, y := p[0], p[1]

				if used[x][y] {
					writer.WriteString("NO")
					writer.WriteByte('\n')

					return
				}

				used[x][y] = true
			}

			for _, p := range cells {
				x, y := p[0], p[1]

				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						nx := x + dx
						ny := y + dy

						if !inside(nx, ny) {
							continue
						}

						if field[nx][ny] != '#' {
							continue
						}

						ok := false

						for _, q := range cells {
							if q[0] == nx && q[1] == ny {
								ok = true
								break
							}
						}

						if !ok {
							writer.WriteString("NO")
							writer.WriteByte('\n')

							return
						}
					}
				}
			}

			size := len(cells)

			if size < 1 || size > 4 {
				writer.WriteString("NO")
				writer.WriteByte('\n')

				return
			}

			cnt[size]++
		}
	}

	if cnt[1] != 4 || cnt[2] != 3 || cnt[3] != 2 || cnt[4] != 1 {
		writer.WriteString("NO")
	} else {
		writer.WriteString("YES")
	}
	writer.WriteByte('\n')
}
