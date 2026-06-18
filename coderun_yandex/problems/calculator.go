package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/calculator
// Calculator - problem 105
func Calculator() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	dp := make([]int, n+1)
	prev := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		prev[i] = i - 1

		if i%2 == 0 && dp[i/2]+1 < dp[i] {
			dp[i] = dp[i/2] + 1
			prev[i] = i / 2
		}

		if i%3 == 0 && dp[i/3]+1 < dp[i] {
			dp[i] = dp[i/3] + 1
			prev[i] = i / 3
		}
	}

	path := make([]int, 0)
	for cur := n; cur != 0; cur = prev[cur] {
		path = append(path, cur)

		if cur == 1 {
			break
		}
	}

	for l, r := 0, len(path)-1; l < r; l, r = l+1, r-1 {
		path[l], path[r] = path[r], path[l]
	}

	writer.WriteString(strconv.Itoa(dp[n]))
	writer.WriteByte('\n')

	for i, x := range path {
		if i > 0 {
			writer.WriteByte(' ')
		}
		writer.WriteString(strconv.Itoa(x))
	}
	writer.WriteByte('\n')
}
