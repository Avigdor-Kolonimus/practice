package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/number-of-triangles
// NumberOfTriangles - problem 29
func NumberOfTriangles() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	total := 0
	if n%2 == 0 {
		total = n * (n + 2) * (2*n + 1) / 8
	} else {
		total = (n*(n+2)*(2*n+1) - 1) / 8
	}

	writer.WriteString(strconv.Itoa(total))
	writer.WriteByte('\n')
}
