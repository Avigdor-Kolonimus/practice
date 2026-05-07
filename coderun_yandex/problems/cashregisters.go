package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	InfTime = int(1e9)
)

// https://coderun.yandex.ru/problem/cash-registers
// CashRegisters - problem 82
func CashRegisters() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first  input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	aTime := make([]int, n)
	bTime := make([]int, n)
	cTime := make([]int, n)
	for i := 0; i < n; i++ {
		// input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		parameters := strings.Fields(line)
		cnt := len(parameters)
		if cnt != 3 {
			panic("input does not match 3")
		}

		a, err := strconv.Atoi(parameters[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(parameters[1])
		if err != nil {
			panic(err)
		}
		c, err := strconv.Atoi(parameters[2])
		if err != nil {
			panic(err)
		}

		aTime[i] = a
		bTime[i] = b
		cTime[i] = c
	}

	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = InfTime
	}

	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = min(dp[i], dp[i-1]+aTime[i-1])

		if i >= 2 {
			dp[i] = min(dp[i], dp[i-2]+bTime[i-2])
		}

		if i >= 3 {
			dp[i] = min(dp[i], dp[i-3]+cTime[i-3])
		}
	}

	writer.WriteString(strconv.Itoa(dp[n]))
	writer.WriteByte('\n')
}
