package first2023mobiledev

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/standup-order
// StandupOrder - problem 34
func StandupOrder() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	k, _ := strconv.Atoi(strings.TrimSpace(line))

	line, _ = reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	a := make([]int64, n)
	var suffixSum int64
	for i := 0; i < n; {
		line, _ = reader.ReadString('\n')
		for _, field := range strings.Fields(line) {
			a[i], _ = strconv.ParseInt(field, 10, 64)
			suffixSum += a[i]
			i++
			if i == n {
				break
			}
		}
	}

	answer := suffixSum + int64(k)
	for i := 1; i <= n; i++ {
		suffixSum -= a[i-1]
		gain := int64(i+1)*int64(k) + suffixSum
		if gain > answer {
			answer = gain
		}
	}

	writer.WriteString(strconv.FormatInt(answer, 10))
	writer.WriteByte('\n')
}
