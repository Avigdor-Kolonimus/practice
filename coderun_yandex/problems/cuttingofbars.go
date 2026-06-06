package problems

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/cutting-of-bars
// CuttingOfBars - problem 36
func CuttingOfBars() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// L and N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	l, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	n, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// timbers input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum = strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	timbers := make([]int, n+2)
	timbers[0] = 0
	timbers[n+1] = l
	for i := 0; i < n; i++ {
		t, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		timbers[i+1] = t
	}

	m := n + 2
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	const INF = math.MaxInt
	for length := 2; length < m; length++ {
		for i := 0; i+length < m; i++ {
			j := i + length

			best := INF

			for k := i + 1; k < j; k++ {
				cost := dp[i][k] + dp[k][j] + (timbers[j] - timbers[i])
				if cost < best {
					best = cost
				}
			}

			dp[i][j] = best
		}
	}

	writer.WriteString(strconv.Itoa(dp[0][m-1]))
	writer.WriteByte('\n')
}
