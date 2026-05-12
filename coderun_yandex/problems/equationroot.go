package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/equation-root
// EquationRoot - problem 55
func EquationRoot() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// A input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	a, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// B input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	b, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// C input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	c, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// sqrt(...) cannot be negative
	if c < 0 {
		writer.WriteString("NO SOLUTION")
		writer.WriteByte('\n')

		return
	}

	// sqrt(b) = c
	if a == 0 {
		if c*c == b {
			writer.WriteString("MANY SOLUTIONS")
		} else {
			writer.WriteString("NO SOLUTION")
		}
		writer.WriteByte('\n')

		return
	}

	val := c*c - b

	if val%a != 0 {
		writer.WriteString("NO SOLUTION")
		writer.WriteByte('\n')

		return
	}

	writer.WriteString(strconv.Itoa(val / a))
	writer.WriteByte('\n')
}
