package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/number-different-numbers
// NumberDifferentNumbers - problem 64
func NumberDifferentNumbers() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// numbers line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)

	// count
	uniqueNumbers := make(map[string]struct{})
	for _, value := range strNum {
		uniqueNumbers[value] = struct{}{}
	}

	writer.WriteString(strconv.Itoa(len(uniqueNumbers)))
	writer.WriteByte('\n')
}
