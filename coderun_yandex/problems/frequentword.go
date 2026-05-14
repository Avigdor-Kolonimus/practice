package problems

import (
	"bufio"
	"os"
)

// https://coderun.yandex.ru/problem/frequent-word
// FrequentWord - problem 50
func FrequentWord() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	ans := ""
	maxCount := 0
	count := make(map[string]int)
	for scanner.Scan() {
		word := scanner.Text()
		count[word]++

		if count[word] > maxCount || (count[word] == maxCount && word < ans) {
			maxCount = count[word]
			ans = word
		}
	}

	writer.WriteString(ans)
	writer.WriteByte('\n')
}
