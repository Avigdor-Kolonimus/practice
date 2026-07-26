package atolstikov

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/atolstikov/problems/checkered-board
// CheckeredBoard - problem 3
func CheckeredBoard() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and K input
	line, _ := reader.ReadString('\n')
	fields := strings.Fields(line)

	n, _ := strconv.Atoi(fields[0])
	k, _ := strconv.Atoi(fields[1])

	if n*n%k != 0 || (k == 1 && n > 1) || k > n*n {
		writer.WriteString("No")
		writer.WriteByte('\n')

		return
	}

	board := make([]int, 0, n*n)

	// Case 1: n is not divisible by k
	if n%k != 0 {
		/*
			Example:
			n = 3 and k = 2

			1 2 1 | 2 1 2 | 1 2 1

			1 2 1
			2 1 2
			1 2 1
		*/
		for len(board) < n*n {
			for c := 1; c <= k && len(board) < n*n; c++ {
				board = append(board, c)
			}
		}
	} else {
		/*
			Example:
			n = 6 and k = 3

			1 2 3 1 2 3 | ....

			123123
			231231
			312312
			123123
			...
		*/

		// Build the first row
		for len(board) < n {
			for c := 1; c <= k && len(board) < n; c++ {
				board = append(board, c)
			}
		}

		// Build the remaining rows by cyclically shifting the previous row by one position.
		shift := 1
		for len(board) < n*n {
			for j := shift; j < n && len(board) < n*n; j++ {
				board = append(board, board[j])
			}
			for j := 0; j < shift && len(board) < n*n; j++ {
				board = append(board, board[j])
			}
			shift++
		}
	}

	writer.WriteString("Yes")
	writer.WriteByte('\n')
	for i := range n {
		for j := range n {
			if j > 0 {
				writer.WriteByte(' ')
			}
			writer.WriteString(strconv.Itoa(board[i*n+j]))
		}
		writer.WriteByte('\n')
	}
}
