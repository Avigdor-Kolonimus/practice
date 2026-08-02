package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/phi
// Phi - problem 627
func Phi() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	if n == 1 {
		writer.WriteString("1\n")
		return
	}

	result := n
	x := n
	for p := 2; p*p <= x; p++ {
		if x%p == 0 {
			for x%p == 0 {
				x /= p
			}
			result = result / p * (p - 1)
		}
	}

	if x > 1 {
		result = result / x * (x - 1)
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
