package quickstart

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func validateSymmetricSequenceInput(p int) bool {
	return p >= 1 && p <= 100
}

func isPalindrome(a []int) bool {
	l, r := 0, len(a)-1
	for l < r {
		if a[l] != a[r] {
			return false
		}

		l++
		r--
	}

	return true
}

// https://coderun.yandex.ru/selections/quickstart/problems/symmetric-sequence
// SymmetricSequence - assignment 13
func SymmetricSequence() {
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
	if !validateSymmetricSequenceInput(n) {
		panic("number N out of range")
	}

	// slice input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("input does not match n")
	}

	matrix := make([]int, n)
	for i, v := range strNum {
		curr, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		if curr > 9 || curr < 1 {
			panic("number Ai out of range")
		}

		matrix[i] = curr
	}

	for i := 0; i < n; i++ {
		if isPalindrome(matrix[i:]) {
			writer.WriteString(strconv.Itoa(i))
			writer.WriteByte('\n')

			// выводим перевёрнутый префикс
			for j := i - 1; j >= 0; j-- {
				writer.WriteString(strconv.Itoa(matrix[j]))
				writer.WriteByte(' ')
			}
			break
		}
	}
	writer.WriteByte('\n')
}
