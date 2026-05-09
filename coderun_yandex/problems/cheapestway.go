package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func cheapestWay(matrix [][]int, high int, len int) int {
	mpMat := make([][]int, high)
	for i := range mpMat {
		mpMat[i] = make([]int, len)
	}

	mpMat[0][0] = matrix[0][0]
	for c := 1; c < len; c++ {
		mpMat[0][c] = mpMat[0][c-1] + matrix[0][c]
	}

	for r := 1; r < high; r++ {
		mpMat[r][0] = matrix[r][0] + mpMat[r-1][0]
	}

	for r := 1; r < high; r++ {
		for c := 1; c < len; c++ {
			mpMat[r][c] = matrix[r][c] + min(mpMat[r-1][c], mpMat[r][c-1])
		}
	}

	return mpMat[high-1][len-1]
}

// https://coderun.yandex.ru/problem/cheapest-way
// CheapestWay - problem 2
func CheapestWay() {
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

	board := make([][]int, n)
	for i := 0; i < n; i++ {
		// row input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
		if len(strNum) != m {
			panic("numbers count does not match m")
		}

		board[i] = make([]int, m)
		for j := 0; j < m; j++ {
			board[i][j], err = strconv.Atoi(strNum[j])
			if err != nil {
				panic(err)
			}
		}
	}

	ans := cheapestWay(board, n, m)

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
