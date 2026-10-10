package mobiledev

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/numbers-game
// NumbersGame - problem 24
func NumbersGame() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	fields := strings.Fields(line)

	stack := make([]int, 0, len(fields))
	lastDeleteNum := -1
	result := 0
	for _, field := range fields {
		num, _ := strconv.Atoi(field)

		if num == lastDeleteNum {
			result++
			continue
		}

		if len(stack) < 2 {
			stack = append(stack, num)
			continue
		}

		last := stack[len(stack)-1]

		if num != last {
			stack = append(stack, num)
			continue
		}

		prev := stack[len(stack)-2]

		if prev == last {
			stack = stack[:len(stack)-2]
			lastDeleteNum = last
			result += 3
		} else {
			stack = append(stack, num)
		}
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
