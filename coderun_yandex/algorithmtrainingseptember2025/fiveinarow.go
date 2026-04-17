package algorithmtrainingseptember2025

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateFiveInARowInput(p int) bool {
	return p >= 1 && p <= 1_000
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/five-in-a-row
// FiveInARow - assignment 7
func FiveInARow() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first line
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	// N
	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	if !validateFiveInARowInput(n) {
		panic("N out of range")
	}

	// M
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	if !validateFiveInARowInput(m) {
		panic("m out of range")
	}

	grid := make([][]int, n)
	for i := range n {
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		line = strings.TrimRight(line, "\r\n")
		if len(line) != m {
			panic("line count does not match")
		}

		grid[i] = make([]int, m)
		for j, r := range line {
			switch r {
			case 'X':
				grid[i][j] = 1
			case 'O':
				grid[i][j] = -1
			case '.':
				grid[i][j] = 0
			default:
				panic("illegal rune")
			}
		}
	}

	result := "No"
	for i := range n {
		for j := range m {
			if grid[i][j] == 0 {
				continue
			}

			ch := grid[i][j]

			// -
			if j+4 < m {
				if grid[i][j+1] == ch && grid[i][j+2] == ch && grid[i][j+3] == ch && grid[i][j+4] == ch {
					result = "Yes"
					break
				}
			}

			// |
			if i+4 < n {
				if grid[i+1][j] == ch && grid[i+2][j] == ch && grid[i+3][j] == ch && grid[i+4][j] == ch {
					result = "Yes"
					break
				}
			}

			// \
			if i+4 < n && j+4 < m {
				if grid[i+1][j+1] == ch && grid[i+2][j+2] == ch && grid[i+3][j+3] == ch && grid[i+4][j+4] == ch {
					result = "Yes"
					break
				}
			}

			// /
			if i+4 < n && j-4 >= 0 {
				if grid[i+1][j-1] == ch && grid[i+2][j-2] == ch && grid[i+3][j-3] == ch && grid[i+4][j-4] == ch {
					result = "Yes"
					break
				}
			}
		}
	}
	writer.WriteString(result)
	writer.WriteByte('\n')
}
