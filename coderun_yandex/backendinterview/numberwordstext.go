package backendinterview

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/backend-interview/problems/number-words-text
// NumberWordsText - assignment 1
func NumberWordsText() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	words := make(map[string]struct{})
	for {
		// Text input
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		for _, w := range strings.Fields(line) {
			words[w] = struct{}{}
		}
	}

	writer.WriteString(strconv.Itoa(len(words)))
	writer.WriteByte('\n')
}
