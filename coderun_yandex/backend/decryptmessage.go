package backend

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func normalize(word string) string {
	key := make([]byte, len(word))
	first := word[0] - 'a'

	for i := range word {
		key[i] = (word[i]-'a'-first+26)%26 + 'a'
	}

	return string(key)
}

// https://coderun.yandex.ru/selections/first-2023-backend/problems/decrypt-message
// DecryptMessage - problem 14
func DecryptMessage() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	words := strings.Split(strings.TrimSpace(line), " ")

	wordMap := make(map[string]string, len(words))
	for _, word := range words {
		wordMap[normalize(word)] = word
	}

	line, _ = reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	for i := 0; i < n; i++ {
		line, _ = reader.ReadString('\n')
		word := strings.TrimSpace(line)

		result := wordMap[normalize(word)]

		writer.WriteString(result)
		writer.WriteByte('\n')
	}
}
