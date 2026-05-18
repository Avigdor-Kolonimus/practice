package winterintern2024

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	sizeBoard = 8
)

func checkBoard(board [8][8]bool, i, j int) bool {
	if i < 0 || i >= 8 || j < 0 || j >= 8 {
		return false
	}

	return board[i][j]
}

// https://coderun.yandex.ru/selections/winter-intern-2024/problems/tetramino
// Tetramino - problem 5
func Tetramino() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// board input
	var board [sizeBoard][sizeBoard]bool
	for i := range sizeBoard {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		for j := 0; j < sizeBoard; j++ {
			if line[j] == '.' {
				board[i][j] = true
			}
		}
	}

	result := 0
	for i := range 8 {
		for j := range 8 {
			/*
				***
				 *
			*/
			if checkBoard(board, i, j) && checkBoard(board, i, j+1) && checkBoard(board, i, j+2) && checkBoard(board, i+1, j+1) {
				result++
			}

			/*
				 *
				**
				 *
			*/
			if checkBoard(board, i, j) && checkBoard(board, i, j+1) && checkBoard(board, i-1, j+1) && checkBoard(board, i+1, j+1) {
				result++
			}

			/*
				 *
				***
			*/
			if checkBoard(board, i, j) && checkBoard(board, i, j+1) && checkBoard(board, i-1, j+1) && checkBoard(board, i, j+2) {
				result++
			}

			/*
			*
			**
			*
			 */
			if checkBoard(board, i, j) && checkBoard(board, i-1, j) && checkBoard(board, i+1, j) && checkBoard(board, i, j+1) {
				result++
			}
		}
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
