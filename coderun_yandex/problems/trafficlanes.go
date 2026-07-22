package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/traffic-lanes
// TrafficLanes - problem 35
func TrafficLanes() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	fields := strings.Fields(line)

	m, _ := strconv.Atoi(fields[0])
	n, _ := strconv.Atoi(fields[1])

	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, m+1)
	}

	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j] + dp[i-1][j-1]
		}
	}

	writer.WriteString(strconv.FormatInt(dp[n][m], 10))
	writer.WriteByte('\n')
}
