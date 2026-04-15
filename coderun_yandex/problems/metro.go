package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateMetroInput(p int) bool {
	return p >= 1 && p <= 1_000
}

// https://coderun.yandex.ru/problem/metro
// Metro - problem 57
func Metro() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// a input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	a, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateMetroInput(a) {
		panic("a out of range")
	}

	// b input
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	b, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateMetroInput(b) {
		panic("b out of range")
	}

	// n input
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateMetroInput(n) {
		panic("n out of range")
	}

	// m input
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	m, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateMetroInput(m) {
		panic("m out of range")
	}

	min1 := (n + a*(n-1))
	max1 := (n + a*(n+1))

	min2 := (m + b*(m-1))
	max2 := (m + b*(m+1))

	resMin := max(min1, min2)
	resMax := min(max1, max2)

	result := "-1"
	if resMin <= resMax {
		result = strconv.Itoa(resMin) + " " + strconv.Itoa(resMax)
	}

	writer.WriteString(result)
	writer.WriteByte('\n')
}
