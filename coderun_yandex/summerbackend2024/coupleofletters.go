package summerbackend2024

import (
	"bufio"
	"os"
	"strings"
)

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/couple-of-letters
// CoupleOfLetters - problem 13
func CoupleOfLetters() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// string input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	str := strings.TrimSpace(line)

	m := make(map[string]int)
	subStr, freq := "", 0

	if len(str) >= 2 {
		for i := 1; i < len(str); i++ {
			ss := str[i-1 : i+1]
			if ss[0] != ' ' && ss[1] != ' ' {
				m[ss]++
			}

			if m[ss] > freq || (m[ss] == freq && ss > subStr) {
				subStr = ss
				freq = m[ss]
			}
		}
	}

	writer.WriteString(subStr)
	writer.WriteByte('\n')
}
