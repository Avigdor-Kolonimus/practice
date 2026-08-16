package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func isIdentifierChar(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') ||
		c == '_'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAllDigits(s string) bool {
	for i := 0; i < len(s); i++ {
		if !isDigit(s[i]) {
			return false
		}
	}

	return true
}

// https://coderun.yandex.ru/problem/additional-check-cheating
// AdditionalCheckCheating - problem 207
func AdditionalCheckCheating() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N, caseSensitive and canStartWithDigit input
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	n, _ := strconv.Atoi(parts[0])
	caseSensitive := parts[1] == "yes"
	canStartWithDigit := parts[2] == "yes"

	// Keywords input
	keywords := make(map[string]struct{})
	for range n {
		line, _ := reader.ReadString('\n')
		word := strings.TrimSpace(line)

		if !caseSensitive {
			word = strings.ToLower(word)
		}

		keywords[word] = struct{}{}
	}

	count := make(map[string]int)
	first := make(map[string]int)

	best := ""
	bestCount, order := 0, 0
	bestFirst := int(^uint(0) >> 1)
	for {
		line, err := reader.ReadString('\n')

		if len(line) > 0 {
			for i := 0; i < len(line); {
				if !isIdentifierChar(line[i]) {
					i++
					continue
				}

				start := i
				for i < len(line) && isIdentifierChar(line[i]) {
					i++
				}

				word := line[start:i]

				if !canStartWithDigit && isDigit(word[0]) {
					continue
				}

				if isAllDigits(word) {
					continue
				}

				normalized := word

				if !caseSensitive {
					normalized = strings.ToLower(normalized)
				}

				if _, exists := keywords[normalized]; exists {
					continue
				}

				count[normalized]++

				if _, exists := first[normalized]; !exists {
					first[normalized] = order
					order++
				}

				if count[normalized] > bestCount ||
					(count[normalized] == bestCount &&
						first[normalized] < bestFirst) {

					best = normalized
					bestCount = count[normalized]
					bestFirst = first[normalized]
				}
			}
		}

		if err != nil {
			break
		}
	}

	writer.WriteString(best)
	writer.WriteByte('\n')
}
