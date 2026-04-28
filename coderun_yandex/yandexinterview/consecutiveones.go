package yandexinterview

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func validateConsecutiveOnesInput(n int) bool {
	return n >= 1 && n <= 10_000
}

// https://coderun.yandex.ru/selections/yandex-interview/problems/consecutive-ones
// ConsecutiveOnes - problem 3
func ConsecutiveOnes() {
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
	if !validateConsecutiveOnesInput(n) {
		panic("number N out of range")
	}

	// array input
	count, max := 0, 0
	for range n {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		if line == "1" {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 0
		}
	}

	if count > max {
		max = count
	}

	writer.WriteString(strconv.Itoa(max))
	writer.WriteByte('\n')
}
