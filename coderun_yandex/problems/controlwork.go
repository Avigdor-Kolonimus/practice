package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/control-work
// ControlWork - problem 92
func ControlWork() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	readInt := func() int {
		line, _ := reader.ReadString('\n')
		n, _ := strconv.Atoi(strings.TrimSpace(line))
		return n
	}

	n := readInt()
	k := readInt()
	row := readInt()
	seat := readInt()

	p := (row-1)*2 + seat - 1

	left := p - k
	right := p + k

	getRow := func(p int) int {
		return p/2 + 1
	}

	getSeat := func(p int) int {
		return p%2 + 1
	}

	if left < 0 && right >= n {
		writer.WriteString("-1\n")
		return
	}

	if left < 0 {
		p = right
	} else if right >= n {
		p = left
	} else {
		if getRow(p)-getRow(left) < getRow(right)-getRow(p) {
			p = left
		} else {
			p = right
		}
	}

	writer.WriteString(strconv.Itoa(getRow(p)))
	writer.WriteByte(' ')
	writer.WriteString(strconv.Itoa(getSeat(p)))
	writer.WriteByte('\n')
}
