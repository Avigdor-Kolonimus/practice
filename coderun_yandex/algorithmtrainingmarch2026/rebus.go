package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/rebus
// Rebus - assignment 5
func Rebus() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	parts := strings.Fields(line)

	var ans strings.Builder
	for _, p := range parts {
		l := 0
		for l < len(p) && p[l] == '\'' {
			l++
		}

		r := 0
		for r < len(p)-l && p[len(p)-1-r] == '\'' {
			r++
		}

		word := p[l : len(p)-r]
		ans.WriteString(word[l : len(word)-r])
	}

	writer.WriteString(ans.String())
	writer.WriteByte('\n')
}
