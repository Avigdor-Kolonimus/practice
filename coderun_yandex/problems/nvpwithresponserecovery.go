package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/nvp-with-response-recovery
// NvpWithResponseRecovery - problem 28
func NvpWithResponseRecovery() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// array input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i], err = strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

	}

	// dp
	dp := make([]int, n)
	prev := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		prev[i] = -1

		for j := 0; j < i; j++ {
			if arr[j] < arr[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
	}

	// search end of  LIS
	maxLen := 0
	pos := 0
	for i := 0; i < n; i++ {
		if dp[i] > maxLen {
			maxLen = dp[i]
			pos = i
		}
	}

	// rebuild
	result := ""
	for pos != -1 {
		result = strconv.Itoa(arr[pos]) + " " + result
		pos = prev[pos]
	}

	writer.WriteString(result)
	writer.WriteByte('\n')
}
