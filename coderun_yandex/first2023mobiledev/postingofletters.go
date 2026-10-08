package first2023mobiledev

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/posting-of-letters
// PostingOfLetters - problem 22
func PostingOfLetters() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	s, _ := strconv.ParseInt(parts[0], 10, 64)
	n, _ := strconv.Atoi(parts[1])

	line, _ = reader.ReadString('\n')
	parts = strings.Fields(line)

	var left, right int64
	for i := range n {
		x, _ := strconv.ParseInt(parts[i], 10, 64)

		if i == 0 {
			left = x
			right = x
			continue
		}

		if x < left {
			left = x
		}
		if x > right {
			right = x
		}
	}

	distance := right - left

	option1 := abs(s-left) + distance
	option2 := abs(s-right) + distance

	answer := option1
	if option2 < answer {
		answer = option2
	}

	writer.WriteString(strconv.FormatInt(answer, 10))
	writer.WriteByte('\n')
}
