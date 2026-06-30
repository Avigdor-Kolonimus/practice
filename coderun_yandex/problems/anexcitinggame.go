package problems

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/an-exciting-game
// AnExcitingGame - problem 31
func AnExcitingGame() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N, x and y input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNum := strings.Fields(line)
	if len(strNum) != 3 {
		panic("numbers count does not match 3")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	a, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(strNum[2])
	if err != nil {
		panic(err)
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= n; i++ {
		best := math.MaxInt

		for j := 1; j < i; j++ {
			left := j
			right := i - j

			cost1 := max(a+dp[left], b+dp[right])
			cost2 := max(a+dp[right], b+dp[left])

			cur := min(cost1, cost2)

			if cur < best {
				best = cur
			}
		}

		dp[i] = best
	}

	writer.WriteString(strconv.Itoa(dp[n]))
	writer.WriteByte('\n')
}
