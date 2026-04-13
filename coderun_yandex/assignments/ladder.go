package assignments

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateLadderInput(n int) bool {
	return n >= 1 && n <= 150
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/ladder
// Ladder - assignment 17
func Ladder() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateLadderInput(n) {
		panic("number N out of range")
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := n; j >= i; j-- {
			dp[j] += dp[j-i]
		}
	}

	writer.WriteString(strconv.Itoa(dp[n]))
	writer.WriteByte('\n')
}
