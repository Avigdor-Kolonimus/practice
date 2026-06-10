package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/domino-placement
// DominoPlacement - assignment 1
func DominoPlacement() {
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
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		board[i] = strings.TrimRight(line, "\r\n")
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] != '.' {
				continue
			}

			// horizantal -
			if j+1 < m && board[i][j+1] == '.' {
				ans++
			}

			// vertical |
			if i+1 < n && board[i+1][j] == '.' {
				ans++
			}
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
