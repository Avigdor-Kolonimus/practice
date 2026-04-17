package quickstart

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/quickstart/problems/triangle
// Triangle - assignment 2
func Triangle() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	a, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// second input
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	b, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// third input
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	c, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	result := "YES"
	if a+b <= c {
		result = "NO"
	}
	if a+c <= b {
		result = "NO"
	}
	if b+c <= a {
		result = "NO"
	}

	writer.WriteString(result)
	writer.WriteByte('\n')
}
