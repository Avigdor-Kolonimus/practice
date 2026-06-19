package problems

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/pin
// Pin - problem 106
func Pin() {
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

	// pin input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNums := strings.Fields(line)

	x := make([]int, n)
	for i := 0; i < n; i++ {
		xi, err := strconv.Atoi(strNums[i])
		if err != nil {
			panic(err)
		}

		x[i] = xi
	}

	sort.Ints(x)

	d := make([]int, n)
	for i := 1; i < n; i++ {
		d[i] = x[i] - x[i-1]
	}

	dp := make([]int, n)
	dp[1] = d[1]
	if n >= 3 {
		dp[2] = d[1] + d[2]
	}

	for i := 3; i < n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + d[i]
	}

	writer.WriteString(strconv.Itoa(dp[n-1]))
	writer.WriteByte('\n')
}
