package quickstart

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func validateGcdAndLcmInput(p int) bool {
	return p >= 1 && p <= 1_000_000_000
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a, b, g int) int {
	// overflow
	return (a / g) * b
}

// https://coderun.yandex.ru/selections/quickstart/problems/gcd-and-lcm
// GcdAndLcm - assignment 10
func GcdAndLcm() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// slice input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("input does not match 2")
	}

	// A
	a, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	if !validateGcdAndLcmInput(a) {
		panic("number A out of range")
	}

	// B
	b, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	if !validateGcdAndLcmInput(b) {
		panic("number B out of range")
	}

	g := gcd(a, b)
	l := lcm(a, b, g)

	writer.WriteString(strconv.Itoa(g))
	writer.WriteByte(' ')
	writer.WriteString(strconv.Itoa(l))
	writer.WriteByte('\n')
}
