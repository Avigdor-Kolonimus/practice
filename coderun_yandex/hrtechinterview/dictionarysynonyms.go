package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/hr-tech-interview/problems/dictionary-synonyms
// DictionarySynonyms - problem 1
func DictionarySynonyms() {
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

	// first input
	line := readLine()

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	synonyms := make(map[string]string, n*2)

	for range n {
		// dictionary input
		line = readLine()

		words := strings.Fields(line)
		if len(words) != 2 {
			panic("words count does not match 2")
		}

		word1, word2 := words[0], words[1]

		synonyms[word1] = word2
		synonyms[word2] = word1
	}

	// last input
	line = readLine()

	writer.WriteString(synonyms[line])
	writer.WriteByte('\n')
}
