package mobiledev

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/jolly-delivery
// JollyDelivery - problem 18
func JollyDelivery() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	path := strings.TrimSpace(line)
	n := len(path)

	var dp [2][2][]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			dp[i][j] = make([]int, n+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	var helper func(fromLane, lane, step int) int

	helper = func(fromLane, lane, step int) int {
		if dp[fromLane][lane][step] != -1 {
			return dp[fromLane][lane][step]
		}

		result := 0

		if step == n {
			if lane == 0 {
				result = 1
			}
		} else if lane == 0 {
			cross := 0
			if path[step] == '1' || path[step] == '3' {
				cross = 1
			}

			next1 := cross + helper(0, 0, step+1)
			next2 := int(^uint(0) >> 1)

			if fromLane == 0 {
				next2 = 1 + helper(0, 1, step)
			}

			result = min(next1, next2)
		} else {
			cross := 0
			if path[step] == '2' || path[step] == '3' {
				cross = 1
			}

			next1 := cross + helper(1, 1, step+1)
			next2 := int(^uint(0) >> 1)

			if fromLane == 1 {
				next2 = 1 + helper(1, 0, step)
			}

			result = min(next1, next2)
		}

		dp[fromLane][lane][step] = result
		return result
	}

	answer := helper(0, 0, 0)

	writer.WriteString(strconv.Itoa(answer))
	writer.WriteByte('\n')
}
