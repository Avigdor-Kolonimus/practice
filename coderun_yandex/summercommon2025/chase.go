package summercommon2025

import (
	"bufio"
	"os"
	"strconv"
)

func validateChaseInput(n int) bool {
	return n >= 1 && n <= 100_000
}

func solveChase(n int, _ int, a []int) []int {
	haveStone := make([]bool, n+1)
	for _, x := range a {
		if x >= 0 && x <= n {
			haveStone[x] = true
		}
	}
	haveStone[n] = true

	const INF = int(1e9)
	dp := make([]int, n+2)

	for i := 0; i <= n; i++ {
		dp[i] = INF
	}
	dp[n] = 0

	for i := n - 1; i >= 0; i-- {
		if i+1 <= n && haveStone[i+1] {
			if dp[i] > dp[i+1]+1 {
				dp[i] = dp[i+1] + 1
			}
		}

		if i+2 <= n && haveStone[i+2] {
			if dp[i] > dp[i+2]+1 {
				dp[i] = dp[i+2] + 1
			}
		}
	}

	if dp[0] == INF {
		return []int{-1}
	}

	res := []int{}
	i := 0

	for i < n {
		if i+1 <= n && haveStone[i+1] && dp[i] == dp[i+1]+1 {
			res = append(res, 1)
			i += 1
		} else {
			res = append(res, 2)
			i += 2
		}
	}

	answer := make([]int, 0, len(res)+1)
	answer = append(answer, len(res))
	answer = append(answer, res...)

	return answer
}

// https://coderun.yandex.ru/selections/2025-summer-common/problems/chase
// Chase - problem 12
func Chase() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and K input
	line := mustReadIntArray(reader, 2)
	n, k := line[0], line[1]
	if !validateChaseInput(n) {
		panic("number N out of range")
	}

	// coordinates input
	coordinates := make([]int, n)
	line = mustReadIntArray(reader, n)
	for i := range coordinates {
		coordinates[i] = line[i]
	}

	answer := solveChase(n, k, coordinates)

	for _, v := range answer {
		writer.WriteString(strconv.Itoa(v))
		writer.WriteByte('\n')
	}
}
