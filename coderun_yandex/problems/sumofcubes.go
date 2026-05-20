package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/sum-of-cubes
// SumOfCubes - problem 27
func SumOfCubes() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	cubes := []int{}
	for i := 1; i*i*i <= n; i++ {
		cubes = append(cubes, i*i*i)
	}

	const INF = int(1e9)

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = INF
	}

	dp[0] = 0
	for x := 1; x <= n; x++ {
		for _, cube := range cubes {
			if cube > x {
				break
			}

			dp[x] = min(dp[x], dp[x-cube]+1)
		}
	}

	writer.WriteString(strconv.Itoa(dp[n]))
	writer.WriteByte('\n')
}
