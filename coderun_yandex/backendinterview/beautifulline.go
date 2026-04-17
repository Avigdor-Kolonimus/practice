package backendinterview

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateBeautifulLineInput(k int) bool {
	return k >= 0 && k <= 1_000_000_000
}

func maxSubstringLength(s string, target byte, k int) int {
	n := len(s)
	left := 0
	replacements := 0
	maxLen := 0

	for right := range n {
		if s[right] != target {
			replacements++
		}

		for replacements > k {
			if s[left] != target {
				replacements--
			}
			left++
		}

		if currentLen := right - left + 1; currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}

// https://coderun.yandex.ru/selections/backend-interview/problems/beautiful-line
// BeautifulLine - assignment 9
func BeautifulLine() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// K input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	k, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateBeautifulLineInput(k) {
		panic("number K out of range")
	}

	// string input
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	maxBeauty := 0
	for ch := byte('a'); ch <= byte('z'); ch++ {
		beauty := maxSubstringLength(line, ch, k)
		if beauty > maxBeauty {
			maxBeauty = beauty
		}
	}

	writer.WriteString(strconv.Itoa(maxBeauty))
	writer.WriteByte('\n')
}
