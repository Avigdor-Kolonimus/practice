package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	Lunch = 13 * 3600
	Work  = 14 * 3600
	End   = 18 * 3600
	Start = 9 * 3600
)

func parseTime(s string) int {
	h := int(s[0]-'0')*10 + int(s[1]-'0')
	m := int(s[3]-'0')*10 + int(s[4]-'0')
	sec := int(s[6]-'0')*10 + int(s[7]-'0')

	return h*3600 + m*60 + sec
}

// https://coderun.yandex.ru/problem/buratino
// Buratino - problem 30
func Buratino() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	times := make([]int, n)
	rates := make([]int, n)
	for i := range n {
		line, _ = reader.ReadString('\n')
		fields := strings.Fields(line)

		times[i] = parseTime(fields[0])
		rates[i], _ = strconv.Atoi(fields[1])
	}

	dp := make([]int, End+1)
	timeIndex := n - 1
	for t := End - 1; t >= Start; t-- {
		for times[timeIndex] > t {
			timeIndex--
		}

		if Lunch <= t && t < Work {
			dp[t] = dp[Work]
			continue
		}

		rate := rates[timeIndex]

		if End-rate < t && t <= End ||
			Lunch-rate < t && t <= Lunch {
			dp[t] = dp[t+1]
			continue
		}

		if dp[t+rate]+1 > dp[t+1] {
			dp[t] = dp[t+rate] + 1
		} else {
			dp[t] = dp[t+1]
		}
	}

	writer.WriteString(strconv.Itoa(dp[Start]))
	writer.WriteByte('\n')
}
