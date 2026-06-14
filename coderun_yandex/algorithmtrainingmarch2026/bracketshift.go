package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func isValidBracketShift(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]

		switch c {
		case '(', '[', '{':
			stack = append(stack, c)

		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if (c == ')' && top != '(') || (c == ']' && top != '[') || (c == '}' && top != '{') {
				return false
			}
		}
	}

	return len(stack) == 0
}

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/bracket-shift
// BracketShift - assignment 21
func BracketShift() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// bracket input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	brackets := strings.TrimRight(line, "\r\n")

	n := len(brackets)

	if n == 0 {
		writer.WriteString("YES")
		writer.WriteByte('\n')

		return
	}

	t := brackets + brackets
	for start := 0; start < n; start++ {
		if isValidBracketShift(t[start : start+n]) {
			writer.WriteString("YES")
			writer.WriteByte('\n')

			return
		}
	}

	writer.WriteString("NO")
	writer.WriteByte('\n')
}
