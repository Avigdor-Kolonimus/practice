package problems

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func upperCountControlAccent(s string) int {
	cnt := 0
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			cnt++
		}
	}

	return cnt
}

// https://coderun.yandex.ru/problem/control-accent
// ControlAccent - problem 205
func ControlAccent() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// dictionaries input
	variants := make(map[string]bool)
	base := make(map[string]bool)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		variants[line] = true
		base[strings.ToLower(line)] = true
	}

	ans := 0
	word := ""
	for {
		if _, err := fmt.Fscan(reader, &word); err != nil {
			break
		}

		// Word exists in the dictionary.
		if base[strings.ToLower(word)] {
			if !variants[word] {
				ans++
			}

			continue
		}

		// Unknown word: it is correct only if it has exactly one uppercase letter.
		if upperCountControlAccent(word) != 1 {
			ans++
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
