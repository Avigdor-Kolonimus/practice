package newyearadventures

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func primeFactorsCount(x int) int {
	cnt := 0
	for d := 2; d*d <= x; d++ {
		if x%d == 0 {
			cnt++
			for x%d == 0 {
				x /= d
			}
		}
	}

	if x > 1 {
		cnt++
	}

	return cnt
}

// https://coderun.yandex.ru/selections/new-year-adventures/problems/postcard-equation
// PostcardEquation - problem 4
func PostcardEquation() {
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

	ans := 0
	for k := 0; k <= 30; k++ {
		x := n - k
		if x <= 0 {
			continue
		}

		if primeFactorsCount(x) == k {
			ans++
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
