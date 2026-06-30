package problems

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func check(s string) bool {
	stack := make([]string, 0)

	i := 0
	n := len(s)

	for i < n {
		if s[i] != '<' {
			return false
		}

		j := i
		for j < n && s[j] != '>' {
			j++
		}
		if j == n {
			return false
		}

		if i+1 >= j {
			return false
		}

		if s[i+1] == '/' {
			if i+2 > j {
				return false
			}

			name := s[i+2 : j]

			if len(stack) == 0 || stack[len(stack)-1] != name {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			for k := i + 1; k < j; k++ {
				if s[k] < 'a' || s[k] > 'z' {
					return false
				}
			}

			stack = append(stack, s[i+1:j])
		}

		i = j + 1
	}

	return len(stack) == 0
}

func fixRobust(s string) string {
	counts := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		counts[s[i]]++
	}

	counts['/'] = 1

	chars := make([]byte, 0, len(counts))
	for c := range counts {
		chars = append(chars, c)
	}

	res := []byte(s)

	for i := 0; i < len(s); i++ {
		oldCnt := counts[s[i]]

		for _, newC := range chars {
			newCnt := counts[newC]

			res[i] = newC

			if ((oldCnt%2 == 1 || s[i] == '/') &&
				(newCnt%2 == 1 || newC == '/')) &&
				check(string(res)) {
				return string(res)
			}
		}

		res[i] = s[i]
	}

	return s
}

// https://coderun.yandex.ru/problem/corrupted-xml
// CorruptedXML - problem 21
func CorruptedXML() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// XML input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	writer.WriteString(fixRobust(line))
	writer.WriteByte('\n')
}
