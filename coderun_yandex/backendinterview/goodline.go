package backendinterview

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateGoodLineInput(n int) bool {
	return n >= 1 && n <= 26
}

// https://coderun.yandex.ru/selections/backend-interview/problems/good-line
// GoodLine - assignment 6
func GoodLine() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateGoodLineInput(n) {
		panic("number N out of range")
	}

	// a input
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	prev, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	count := 0
	for range n - 1 {
		line, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		cur, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		count += min(prev, cur)
		prev = cur
	}

	writer.WriteString(strconv.Itoa(count))
	writer.WriteByte('\n')
}
