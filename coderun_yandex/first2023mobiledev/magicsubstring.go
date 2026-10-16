package first2023mobiledev

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func findMagicSubstring(t, s string) int {
	n, m := len(t), len(s)

	if m == 0 || m > n {
		return -1
	}

	var diff [26]int

	for i := 0; i < m; i++ {
		diff[t[i]-'a']++
		diff[s[i]-'a']--
	}

	nonZero := 0
	absSum := 0

	for _, v := range diff {
		if v != 0 {
			nonZero++
		}
		if v < 0 {
			absSum -= v
		} else {
			absSum += v
		}
	}

	if nonZero == 2 && absSum == 2 {
		return 0
	}

	for i := m; i < n; i++ {
		left := t[i-m] - 'a'
		right := t[i] - 'a'

		diff[left]--
		diff[right]++

		nonZero = 0
		absSum = 0

		for _, v := range diff {
			if v != 0 {
				nonZero++
			}
			if v < 0 {
				absSum -= v
			} else {
				absSum += v
			}
		}

		if nonZero == 2 && absSum == 2 {
			return i - m + 1
		}
	}

	return -1
}

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/magic-substring
// MagicSubstring - problem 5
func MagicSubstring() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	t := strings.TrimSpace(line)

	line, _ = reader.ReadString('\n')
	s := strings.TrimSpace(line)

	result := findMagicSubstring(t, s)

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
