package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/nop-with-response-recovery
// NopWithResponseRecovery - problem 6
func NopWithResponseRecovery() {
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

	// row input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}
	nArray := make([]int, n)
	for i := 0; i < n; i++ {
		nArray[i], err = strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}
	}

	// M input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	m, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// row input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum = strings.Fields(line)
	if len(strNum) != m {
		panic("numbers count does not match m")
	}
	mArray := make([]int, m)
	for i := 0; i < m; i++ {
		mArray[i], err = strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}
	}

	// dp
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// fill dp
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if nArray[i-1] == mArray[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// Восстановление LCS
	lcs := make([]int, 0)
	i, j := n, m
	for i > 0 && j > 0 {
		if nArray[i-1] == mArray[j-1] {
			lcs = append(lcs, nArray[i-1])
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	for k := len(lcs) - 1; k >= 0; k-- {
		writer.WriteString(strconv.Itoa(lcs[k]))
		if k > 0 {
			writer.WriteByte(' ')
		}
	}
	writer.WriteByte('\n')
}
