package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/tic-tac-toe
// TicTacToe - problem 436
func TicTacToe() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// board input
	board := make([]string, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}

		board[i] = strings.TrimRight(line, "\r\n")
	}

	dirs := [][2]int{
		{0, 1},  // -
		{1, 0},  // |
		{1, 1},  // ↘
		{1, -1}, // ↙
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			c := board[i][j]

			if c == '.' {
				continue
			}

			for _, d := range dirs {
				ok := true

				for k := 1; k < 5; k++ {
					ni := i + d[0]*k
					nj := j + d[1]*k

					if ni < 0 || ni >= n || nj < 0 || nj >= m || board[ni][nj] != c {
						ok = false
						break
					}
				}

				if ok {
					writer.WriteString("Yes")
					writer.WriteByte('\n')

					return
				}
			}
		}
	}

	writer.WriteString("No")
	writer.WriteByte('\n')
}
