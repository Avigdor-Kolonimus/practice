package problems

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// https://coderun.yandex.ru/selections/hr-tech-interview/problems/correct-bracket-sequence
// CorrectBracketSequence - problem 5
func CorrectBracketSequence() {
	result := "yes"
	d := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	readLine := func() string {
		line, err := reader.ReadString('\n')
		if err != nil && !(err == io.EOF && len(line) > 0) {
			panic(err)
		}
		return strings.TrimRight(line, "\r\n")
	}

	// input
	line := readLine()

	st := make([]rune, 0, len(line)/2)
	for _, x := range line {
		if x == '(' || x == '[' || x == '{' {
			st = append(st, x)
		} else {
			if len(st) == 0 || st[len(st)-1] != d[x] {
				result = "no"
				break
			}
			st = st[:len(st)-1]
		}
	}

	if len(st) > 0 {
		result = "no"
	}

	writer.WriteString(result)
	writer.WriteByte('\n')
}
