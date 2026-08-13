package problems

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/hack-pentagon
// HackPentagon - problem 49
func HackPentagon() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20)
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	// N input
	line, err := in.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// all 'a'
	base := make([]byte, n)
	for i := range base {
		base[i] = 'a'
	}

	fmt.Fprintln(out, string(base))
	out.Flush()

	// D_0 input
	line, err = in.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	d0, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	if d0 == 0 {
		return
	}

	// Для каждой позиции ставим 'z'.
	answer := make([]byte, n)
	for i := range n {
		query := make([]byte, n)
		copy(query, base)
		query[i] = 'z'

		fmt.Fprintln(out, string(query))
		out.Flush()

		var d int
		fmt.Fscan(in, &d)

		if d == 0 {
			return
		}

		// d = d0 + 25 - 2*p_i
		p := (d0 + 25 - d) / 2

		answer[i] = byte('a' + p)
	}

	fmt.Fprintln(out, string(answer))
	out.Flush()

	var result int
	fmt.Fscan(in, &result)

	if result == 0 {
		return
	}
}
