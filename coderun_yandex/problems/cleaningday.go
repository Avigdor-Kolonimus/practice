package problems

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func checkCleaningDay(h []int, n, r, c, x int) bool {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		// Skip the current student.
		dp[i] = dp[i-1]

		// Form a group ending at position i if the inconvenience
		// does not exceed x.
		if i >= c && h[i-1]-h[i-c] <= x {
			if dp[i-c]+1 > dp[i] {
				dp[i] = dp[i-c] + 1
			}
		}
	}

	// Check if we can form at least r groups.
	return dp[n] >= r
}

// https://coderun.yandex.ru/problem/cleaning-day
// CleaningDay - problem 79
func CleaningDay() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N, R and C input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	tokens := strings.Fields(line)
	if len(tokens) != 3 {
		panic("invalid input")
	}
	n, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic(err)
	}
	r, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}
	c, err := strconv.Atoi(tokens[2])
	if err != nil {
		panic(err)
	}

	// heights input
	heights := make([]int, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		h, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		heights[i] = h
	}

	sort.Ints(heights)

	left, right := 0, heights[n-1]-heights[0]

	for left < right {
		mid := (left + right) / 2

		if checkCleaningDay(heights, n, r, c, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	writer.WriteString(strconv.Itoa(left))
	writer.WriteByte('\n')
}
