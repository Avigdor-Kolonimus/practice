package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/boring-lecture
// BoringLecture - problem 96
func BoringLecture() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n := len(line)
	count := make(map[rune]int64)

	for i, ch := range line {
		// (i+1)⋅(n−i)
		cnt := int64(i+1) * int64(n-i)
		count[ch] += cnt
	}

	for c := 'a'; c <= 'z'; c++ {
		if count[c] > 0 {
			// a: 44
			writer.WriteString(string(c) + ": " + strconv.FormatInt(count[c], 10))
			writer.WriteByte('\n')
		}
	}
}
