package algorithmtrainingseptember2025

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	negINF = math.MinInt / 4
)

func validateCollectCoinsInput(n int) bool {
	return n >= 1 && n <= 10_000
}

func isWall(ch byte) bool {
	return ch == 'W'
}

func coinValue(ch byte) int {
	if ch == 'C' {
		return 1
	}
	return 0
}

func parseRow(line string) []byte {
	line = strings.TrimRight(line, "\r\n")
	line = strings.TrimSpace(line)
	row := []byte(line)

	if len(row) >= 3 {
		return row[:3]
	}

	res := make([]byte, 3)
	copy(res, row)
	for i := len(row); i < 3; i++ {
		res[i] = 'W'
	}

	return res
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/collect-coins
// CollectCoins - assignment 16
func CollectCoins() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	// N
	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateCollectCoinsInput(n) {
		panic("number N out of range")
	}

	dp := make([]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = negINF
	}

	result := 0

	if n > 0 {
		line, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		row := parseRow(line)

		// initialization first line
		allWalls := true
		for c := 0; c < 3; c++ {
			if !isWall(row[c]) {
				dp[c] = coinValue(row[c])
				if dp[c] > result {
					result = dp[c]
				}
				allWalls = false
			}
		}

		// next lines
		for i := 1; !allWalls && i < n; i++ {
			line, err = reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			nextRow := parseRow(line)

			nextDP := make([]int, 3)
			for j := 0; j < 3; j++ {
				nextDP[j] = negINF
			}

			for c := 0; c < 3; c++ {
				if dp[c] == negINF {
					continue
				}

				for d := -1; d <= 1; d++ {
					nc := c + d
					if nc < 0 || nc >= 3 {
						continue
					}
					if isWall(nextRow[nc]) {
						continue
					}

					cand := dp[c] + coinValue(nextRow[nc])
					if cand > nextDP[nc] {
						nextDP[nc] = cand
					}
				}
			}

			dp = nextDP

			for c := 0; c < 3; c++ {
				if dp[c] > result {
					result = dp[c]
				}
			}
		}
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
