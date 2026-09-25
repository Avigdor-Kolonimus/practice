package first2023mobiledev

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/hardest-letter
// HardestLetter - problem 4
func HardestLetter() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	// S input
	line, _ = reader.ReadString('\n')
	s := strings.TrimSpace(line)

	line, _ = reader.ReadString('\n')
	values := strings.Fields(line)

	var prev, maxTime int
	var answer byte
	for i := range n {
		a, _ := strconv.Atoi(values[i])
		curTime := a - prev

		if curTime >= maxTime {
			maxTime = curTime
			answer = s[i]
		}

		prev = a
	}

	writer.WriteByte(answer)
	writer.WriteByte('\n')
}
