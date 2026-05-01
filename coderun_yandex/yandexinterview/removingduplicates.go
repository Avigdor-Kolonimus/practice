package yandexinterview

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func validateRemovingDuplicatesInput(n int) bool {
	return n >= 0 && n <= 1_000_000
}

// https://coderun.yandex.ru/selections/yandex-interview/problems/removing-duplicates
// RemovingDuplicates - problem 6
func RemovingDuplicates() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateRemovingDuplicatesInput(n) {
		panic("number N out of range")
	}

	lastNumber := ""
	for range n {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		if lastNumber != line {
			lastNumber = line
			writer.WriteString(line + " ")
		}
	}
}
