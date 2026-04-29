package summercommon2025

import (
	"bufio"
	"os"
	"strconv"
)

func validateErasedLeaderboardNInput(n int) bool {
	return n >= 1 && n <= 1_000
}

func validateErasedLeaderboardMInput(n int) bool {
	return n >= 0 && n <= 100
}

func solveErasedLeaderboard(n int, m int, p []int) []int {
	cur := 0
	a := make([]int, n)
	for i := range n {
		if p[i] != -1 && cur+m > p[i] {
			return []int{-1}
		} else if p[i] == -1 {
			a[i] = m
		} else {
			a[i] = p[i] - cur
		}
		cur += a[i]
	}

	return a
}

// https://coderun.yandex.ru/selections/2025-summer-common/problems/erased-leaderboard
// ErasedLeaderboard - problem 5
func ErasedLeaderboard() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	line := mustReadIntArray(reader, 2)
	n, m := line[0], line[1]
	if !validateErasedLeaderboardNInput(n) {
		panic("number N out of range")
	}
	if !validateErasedLeaderboardMInput(m) {
		panic("number M out of range")
	}

	// stairs input
	coins := make([]int, n)
	line = mustReadIntArray(reader, n)
	for i := range n {
		coins[i] = line[i]
	}

	answer := solveErasedLeaderboard(n, m, coins)
	for _, v := range answer {
		writer.WriteString(strconv.Itoa(v) + " ")
	}
	writer.WriteByte('\n')
}
