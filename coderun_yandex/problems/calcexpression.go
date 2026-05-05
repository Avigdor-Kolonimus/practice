package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/calc-expression
// CalcExpression - problem 402
func CalcExpression() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// calculation line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	result := 0
	num := 0
	sign := 1 // +1 or -1

	for i := 0; i < len(line); i++ {
		c := line[i]

		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		} else {
			result += sign * num
			num = 0

			if c == '+' {
				sign = 1
			} else {
				sign = -1
			}
		}
	}

	result += sign * num

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
