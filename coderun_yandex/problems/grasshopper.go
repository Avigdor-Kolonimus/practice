package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/grasshopper
// Grasshopper - problem 104
func Grasshopper() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and K line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}

	k, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if i-j >= 1 {
				dp[i] += dp[i-j]
			}
		}
	}

	writer.WriteString(strconv.Itoa(dp[n]))
	writer.WriteByte('\n')
}
