package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func validateTourismInput(p int) bool {
	return p >= 1 && p <= 30_000
}

// https://coderun.yandex.ru/selections/hr-tech-interview/problems/tourism
// Tourism - problem 6
func Tourism() {
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

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateTourismInput(n) {
		panic("N out of range")
	}

	diffsPlus := make([]int, n)
	diffsMinus := make([]int, n)
	num, prev, sum, sumMinus := 0, 0, 0, 0
	for i := range n {
		// coordinate input
		line = readLine()
		coordinates := strings.Fields(line)
		if len(coordinates) != 2 {
			panic("coordinates count does not match 2")
		}
		num, err = strconv.Atoi(coordinates[1])
		if err != nil {
			panic(err)
		}

		if num > prev {
			sum += num - prev
		} else if prev > num {
			sumMinus += prev - num
		}

		diffsPlus[i] = sum
		diffsMinus[i] = sumMinus
		prev = num
	}

	// M input
	line = readLine()
	m, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateTourismInput(m) {
		panic("M out of range")
	}

	results := make([]int, m)
	num2 := 0
	for i := range m {
		// start and end input
		line = readLine()
		se := strings.Fields(line)
		if len(se) != 2 {
			panic("start and end count does not match 2")
		}

		num, err = strconv.Atoi(se[0])
		if err != nil {
			panic(err)
		}
		num2, err = strconv.Atoi(se[1])
		if err != nil {
			panic(err)
		}

		if num < num2 {
			results[i] = diffsPlus[num2-1] - diffsPlus[num-1]
		} else {
			results[i] = diffsMinus[num-1] - diffsMinus[num2-1]
		}
	}

	for _, v := range results {
		writer.WriteString(strconv.Itoa(v))
		writer.WriteByte('\n')
	}
}
