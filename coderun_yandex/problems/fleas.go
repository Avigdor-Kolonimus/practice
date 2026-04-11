package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var knightMoves = [][2]int{
	{2, 1}, {2, -1}, {-2, 1}, {-2, -1},
	{1, 2}, {1, -2}, {-1, 2}, {-1, -2},
}

func validateFleasNandMInput(p int) bool {
	return p >= 2 && p <= 250
}

func validateNumFleasInput(q int) bool {
	return q >= 0 && q <= 10_000
}

func knightMovesBFS(n, m, s, t int) [][]int {
	matrix := make([][]int, n)
	for i := range n {
		matrix[i] = make([]int, m)
	}
	matrix[s-1][t-1] = 1

	var stack [][]int
	stack = append(stack, []int{s - 1, t - 1})
	for len(stack) > 0 {
		el := stack[0]
		i, j := el[0], el[1]
		for _, coords := range knightMoves {
			x, y := coords[0], coords[1]
			if i+x >= 0 && i+x < n && j+y >= 0 && j+y < m && matrix[i+x][j+y] == 0 {
				stack = append(stack, []int{i + x, j + y})
				matrix[i+x][j+y] = 1 + matrix[i][j]
			}
		}
		stack = stack[1:]
	}

	return matrix
}

// https://coderun.yandex.ru/problem/fleas
// Fleas - problem 14
func Fleas() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first line
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	strNum := strings.Fields(line)
	if len(strNum) != 5 {
		panic("numbers count does not match 5")
	}

	// N
	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	if !validateFleasNandMInput(n) {
		panic("N out of range")
	}

	// M
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	if !validateFleasNandMInput(m) {
		panic("M out of range")
	}

	// S,T
	s, err := strconv.Atoi(strNum[2])
	if err != nil {
		panic(err)
	}
	t, err := strconv.Atoi(strNum[3])
	if err != nil {
		panic(err)
	}

	// Q
	q, err := strconv.Atoi(strNum[4])
	if err != nil {
		panic(err)
	}
	if !validateNumFleasInput(q) {
		panic("Q out of range")
	}

	matrix := knightMovesBFS(n, m, s, t)

	sumWays := 0
	for range q {
		// coordinate input
		line, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		strNum := strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		x, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		if matrix[x-1][y-1] > 0 {
			sumWays += matrix[x-1][y-1] - 1
		} else {
			sumWays = -1
			break
		}
	}

	writer.WriteString(strconv.Itoa(sumWays))
	writer.WriteByte('\n')
}
