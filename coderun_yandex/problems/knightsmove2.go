package problems

import (
	"bufio"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func validateKnightMove2Input(n int) bool {
	return n >= 1 && n <= 50
}

// https://coderun.yandex.ru/problem/knights-move-2
// KnightMove2 - problem 32
func KnightMove2() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	strNum := strings.Fields(line)

	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	nums := make([]int, 2)

	for i, v := range strNum {

		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		if !validateKnightMove2Input(n) {
			panic("number out of range")
		}

		nums[i] = n
	}

	n := nums[0]
	m := nums[1]

	chessboard := make([][]*big.Int, n)

	for i := range chessboard {

		chessboard[i] = make([]*big.Int, m)

		for j := range chessboard[i] {
			chessboard[i][j] = big.NewInt(0)
		}
	}

	chessboard[0][0] = big.NewInt(1)

	moves := [][2]int{
		{2, 1},
		{1, 2},
		{-1, 2},
		{2, -1},
	}

	maxSum := n + m - 2
	for sum := 0; sum <= maxSum; sum++ {
		for i := 0; i < n; i++ {
			j := sum - i
			if j < 0 || j >= m {
				continue
			}

			cur := chessboard[i][j]
			if cur.Sign() == 0 {
				continue
			}

			for _, mv := range moves {
				ni := i + mv[0]
				nj := j + mv[1]

				if ni >= 0 && ni < n && nj >= 0 && nj < m && ni+nj > i+j {
					chessboard[ni][nj].Add(chessboard[ni][nj], cur)
				}
			}
		}
	}

	writer.WriteString(chessboard[n-1][m-1].String())
	writer.WriteByte('\n')
}
