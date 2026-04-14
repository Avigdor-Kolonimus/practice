package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func validatePolyglotsInput(n int) bool {
	return n >= 1 && n <= 1_000
}

// https://coderun.yandex.ru/selections/hr-tech-interview/problems/polyglots
// Polyglots - problem 3
func Polyglots() {
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

	allLang := make(map[string]struct{})
	intersection := make(map[string]struct{})

	// first input
	line := readLine()
	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validatePolyglotsInput(n) {
		panic("N out of range")
	}

	for i := range n {
		line = readLine()
		m, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		localIntersection := make(map[string]struct{})

		if i == 0 {
			for range m {
				line = readLine()

				allLang[line] = struct{}{}
				intersection[line] = struct{}{}
			}
		} else {
			for range m {
				line = readLine()

				allLang[line] = struct{}{}
				if _, ok := intersection[line]; ok {
					localIntersection[line] = struct{}{}
				}
			}

			intersection = localIntersection
		}
	}

	writer.WriteString(strconv.Itoa(len(intersection)))
	writer.WriteByte('\n')

	for i := range intersection {
		writer.WriteString(i)
		writer.WriteByte('\n')
	}

	writer.WriteString(strconv.Itoa(len(allLang)))
	writer.WriteByte('\n')

	for i := range allLang {
		writer.WriteString(i)
		writer.WriteByte('\n')
	}
}
