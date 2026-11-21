package problems

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const secondsPerDay = 24 * 60 * 60

func getTimestamp(h, m, s int) int {
	return h*3600 + m*60 + s
}

func getRTT(a, c int) int {
	if c < a {
		c += secondsPerDay
	}

	return c - a
}

func getPathDelay(rtt int) int {
	return (rtt + 1) / 2
}

func add(a, d int) int {
	return (a + d) % secondsPerDay
}

func getTime(timestamp int) (int, int, int) {
	h := timestamp / 3600
	m := (timestamp % 3600) / 60
	s := timestamp % 60

	return h, m, s
}

func parseTimeSNTP(line string) int {
	parts := strings.Split(strings.TrimSpace(line), ":")

	h, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	s, _ := strconv.Atoi(parts[2])

	return getTimestamp(h, m, s)
}

// https://coderun.yandex.ru/problem/sntp
// SNTP - problem 94
func SNTP() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	a := parseTimeSNTP(line)

	line, _ = reader.ReadString('\n')
	b := parseTimeSNTP(line)

	line, _ = reader.ReadString('\n')
	c := parseTimeSNTP(line)

	rtt := getRTT(a, c)
	pathDelay := getPathDelay(rtt)

	result := add(b, pathDelay)

	h, m, s := getTime(result)

	ans := fmt.Sprintf("%02d:%02d:%02d\n", h, m, s)
	writer.WriteString(ans)
}
