package devgointerview

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/dev-go-interview/problems/balance-moments
// BalanceMoments - problem 4
func BalanceMoments() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// str input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	s := strings.TrimRight(line, "\r\n")

	// model input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	target := line[0]

	total := 0
	for i := range s {
		if s[i] == target {
			total++
		}
	}

	pref := 0
	l, r := -1, -1

	for i := range s {
		suf := total - pref
		if s[i] == target {
			suf--
		}

		if pref == suf {
			if l == -1 {
				l = i
			}
			r = i
		}

		if s[i] == target {
			pref++
		}
	}

	if l != -1 {
		writer.WriteString(strconv.Itoa(l))
		writer.WriteByte(' ')
		writer.WriteString(strconv.Itoa(r))
		writer.WriteByte('\n')
	}
}
