package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func gcdNumber(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// https://coderun.yandex.ru/problem/number
// Number - problem 575
func Number() {
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

	sum := 0
	for x := n; x > 0; x /= 10 {
		sum += x % 10
	}

	g := gcdNumber(n, sum)

	// k = n / gcd(n, sum)
	if n%g != 0 {
		writer.WriteString("Epic fail")
		writer.WriteByte('\n')

		return
	}

	k := n / g

	writer.WriteString("I got it")
	writer.WriteByte('\n')

	for i := 0; i < k; i++ {
		writer.WriteString(strconv.Itoa(n))
	}
	writer.WriteByte('\n')
}
