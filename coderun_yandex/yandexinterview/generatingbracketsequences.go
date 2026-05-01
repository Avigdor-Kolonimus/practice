package yandexinterview

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type State struct {
	str   []byte
	open  int
	close int
}

func validateGeneratingBracketSequencesInput(n int) bool {
	return n >= 0 && n <= 11
}

// https://coderun.yandex.ru/selections/yandex-interview/problems/generating-bracket-sequences
// GeneratingBracketSequences - problem 4
func GeneratingBracketSequences() {
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
	if !validateGeneratingBracketSequencesInput(n) {
		panic("number N out of range")
	}

	stack := []State{{str: []byte{}, open: 0, close: 0}}

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(cur.str) == 2*n {
			writer.Write(cur.str)
			writer.WriteByte('\n')
			continue
		}

		// queue <- )
		if cur.close < cur.open {
			next := append([]byte{}, cur.str...)
			next = append(next, ')')
			stack = append(stack, State{next, cur.open, cur.close + 1})
		}

		// queue <- (
		if cur.open < n {
			next := append([]byte{}, cur.str...)
			next = append(next, '(')
			stack = append(stack, State{next, cur.open + 1, cur.close})
		}
	}
}
