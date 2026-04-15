package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/word-appearance-number
// WordAppearanceNumber - problem 51
func WordAppearanceNumber() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	hash := make(map[string]int)

	for {
		line, err := reader.ReadString('\n')
		if line != "" {
			words := strings.Fields(line)
			for i := range words {
				writer.WriteString(strconv.Itoa(hash[words[i]]))
				hash[words[i]]++
				writer.WriteString(" ")
			}
		}

		if err != nil {
			writer.WriteByte('\n')
			break
		}
	}
}
