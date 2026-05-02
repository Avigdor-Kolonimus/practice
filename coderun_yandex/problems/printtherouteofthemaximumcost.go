package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func maxPathPrintTheRouteOfTheMaximumCost(mat [][]int, rows, col int) (int, [][]int) {
	mpMat := make([][]int, rows)
	for i := 0; i < rows; i++ {
		mpMat[i] = make([]int, col)
	}

	mpMat[0][0] = mat[0][0]
	// only down
	for r := 1; r < rows; r++ {
		mpMat[r][0] = mat[r][0] + mpMat[r-1][0]
	}

	// only right
	for c := 1; c < col; c++ {
		mpMat[0][c] = mat[0][c] + mpMat[0][c-1]
	}

	for r := 1; r < rows; r++ {
		for c := 1; c < col; c++ {
			mpMat[r][c] = mat[r][c] + max(mpMat[r-1][c], mpMat[r][c-1])
		}
	}

	return mpMat[rows-1][col-1], mpMat
}

func findPathPrintTheRouteOfTheMaximumCost(mat [][]int, rows, col int) string {
	path := ""
	sum := 0
	r, c := rows-1, col-1
	for r != 0 || c != 0 {
		if r == 0 {
			path = "R " + path
			c--
			sum += mat[r][c]
		} else if c == 0 {
			path = "D " + path
			r--
			sum += mat[r][c]
		} else if mat[r-1][c] > mat[r][c-1] {
			path = "D " + path
			sum += mat[r-1][c]
			r--
		} else {
			path = "R " + path
			sum += mat[r][c-1]
			c--
		}
	}

	return path
}

// https://coderun.yandex.ru/problem/print-the-route-of-the-maximum-cost
// PrintTheRouteOfTheMaximumCost - problem 3
func PrintTheRouteOfTheMaximumCost() {
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

	// N
	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}

	// M
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// board input
	board := make([][]int, n)
	for i := range n {
		// coordinate input
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

	maxCost, dp := maxPathPrintTheRouteOfTheMaximumCost(board, n, m)
	path := findPathPrintTheRouteOfTheMaximumCost(dp, n, m)

	writer.WriteString(strconv.Itoa(maxCost))
	writer.WriteByte('\n')
	writer.WriteString(path)
	writer.WriteByte('\n')
}
