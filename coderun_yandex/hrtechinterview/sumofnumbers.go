package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func validateSumOfNumbersNInput(p int) bool {
	return p >= 1 && p <= 100_000
}

func validateSumOfNumbersKInput(p int) bool {
	return p >= 1 && p <= 1_000_000_000
}

// https://coderun.yandex.ru/selections/hr-tech-interview/problems/sum-of-numbers
// SumOfNumbers - problem 7
func SumOfNumbers() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	readLine := func() string {
		line, err := reader.ReadString('\n')
		if err != nil && !(err == io.EOF && len(line) > 0) {
			panic(err)
		}
		return strings.TrimRight(line, "\r\n")
	}

	// first input
	line := readLine()
	nk := strings.Fields(line)
	if len(nk) != 2 {
		panic("nk count does not match 2")
	}

	n, err := strconv.Atoi(nk[0])
	if err != nil {
		panic(err)
	}
	if !validateSumOfNumbersNInput(n) {
		panic("N out of range")
	}

	k, err := strconv.Atoi(nk[1])
	if err != nil {
		panic(err)
	}
	if !validateSumOfNumbersKInput(k) {
		panic("K out of range")
	}

	// numbers input
	line = readLine()
	numbers := strings.Fields(line)
	if len(numbers) != n {
		panic("numbers count does not match")
	}

	prefixCount := make(map[int]int)
	prefixCount[0] = 1

	sum := 0
	result := 0

	for i := range n {
		num, err := strconv.Atoi(numbers[i])
		if err != nil {
			panic(err)
		}

		sum += num

		if count, ok := prefixCount[sum-k]; ok {
			result += count
		}

		prefixCount[sum]++
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
