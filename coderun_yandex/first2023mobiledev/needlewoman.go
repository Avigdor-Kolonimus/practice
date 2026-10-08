package first2023mobiledev

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/needlewoman
// Needlewoman - problem 19
func Needlewoman() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	line, _ = reader.ReadString('\n')
	s := strings.TrimSpace(line)

	for p := 1; p <= n; p++ {
		chars := make([]byte, p)
		for i := 0; i < p; i++ {
			chars[i] = '#'
		}

		ok := true

		for i := 0; i < n; i++ {
			if s[i] == '#' {
				continue
			}

			pos := i % p

			if chars[pos] == '#' {
				chars[pos] = s[i]
			} else if chars[pos] != s[i] {
				ok = false
				break
			}
		}

		if ok {
			writer.WriteString(strconv.Itoa(p))
			writer.WriteByte('\n')
			return
		}
	}
}
