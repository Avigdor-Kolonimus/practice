package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/postfix-entry
// PostfixEntry - problem 249
func PostfixEntry() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// postfix line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	tokens := strings.Fields(line)

	// calculation
	stack := []int{}
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" {
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			var res int
			switch token {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			}

			stack = append(stack, res)
		} else {
			num, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}

			stack = append(stack, num)
		}
	}

	writer.WriteString(strconv.Itoa(stack[0]))
	writer.WriteByte('\n')
}
