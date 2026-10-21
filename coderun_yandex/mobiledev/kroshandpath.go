package mobiledev

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/krosh-and-path
// KroshAndPath - problem 11
func KroshAndPath() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	nk := strings.Fields(line)
	n, _ := strconv.Atoi(nk[0])
	k, _ := strconv.Atoi(nk[1])

	a := make([]int, n)

	line, _ = reader.ReadString('\n')
	values := strings.Fields(line)

	maxA := 0
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(values[i])
		if a[i] > maxA {
			maxA = a[i]
		}
	}

	canReach := func(level int) bool {
		if a[0] > level || a[n-1] > level {
			return false
		}

		dp := make([]bool, n)
		dp[0] = true

		reachable := 1

		for i := 1; i < n; i++ {
			if i-k-1 >= 0 && dp[i-k-1] {
				reachable--
			}

			if a[i] <= level && reachable > 0 {
				dp[i] = true
				reachable++
			}
		}

		return dp[n-1]
	}

	left := a[0]
	if a[n-1] > left {
		left = a[n-1]
	}
	right := maxA

	for left < right {
		mid := left + (right-left)/2

		if canReach(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	writer.WriteString(strconv.Itoa(left))
	writer.WriteByte('\n')
}
