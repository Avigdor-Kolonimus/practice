package quickstart

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateCalculateTagsInput(n int) bool {
	return n >= 1 && n <= 35
}

// https://coderun.yandex.ru/selections/quickstart/problems/calculate-tags
// CalculateTags - assignment 4
func CalculateTags() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateCalculateTagsInput(n) {
		panic("number N out of range")
	}

	sum := 1
	if n != 1 {
		prev, cur := 1, 1
		sum = 2

		for i := 3; i <= n; i++ {
			prev, cur = cur, prev+cur
			sum += cur
		}
	}

	writer.WriteString(strconv.Itoa(sum))
	writer.WriteByte('\n')
}
