package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateKnightMoveInput(n int) bool {
	return n >= 1 && n <= 50
}

// https://coderun.yandex.ru/problem/knight-move
// KnightMove - problem 2
func KnightMove() {
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

		if !validateKnightMoveInput(n) {
			panic("number out of range")
		}

		nums[i] = n
	}

	chessboard := make([][]int, nums[0])
	for i := range chessboard {
		chessboard[i] = make([]int, nums[1])
	}
	chessboard[0][0] = 1

	for i := range nums[0] {
		for j := range nums[1] {
			if i >= 2 && j >= 1 {
				chessboard[i][j] += chessboard[i-2][j-1]
			}

			if i >= 1 && j >= 2 {
				chessboard[i][j] += chessboard[i-1][j-2]
			}
		}
	}

	cntWays := chessboard[nums[0]-1][nums[1]-1]

	writer.WriteString(strconv.Itoa(cntWays))
	writer.WriteByte('\n')
}
