package quickstart

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func validateOpenCalculatorXYZInput(n int) bool {
	return n >= 0 && n <= 9
}

func validateOpenCalculatorNInput(n int) bool {
	return n >= 0 && n <= 10_000
}

// https://coderun.yandex.ru/selections/quickstart/problems/open-calculator
// OpenCalculator - assignment 6
func OpenCalculator() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// slice input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}

	strNum := strings.Fields(line)
	if len(strNum) != 3 {
		panic("numbers count does not match 3")
	}

	set := make(map[string]struct{})
	for _, value := range strNum {
		x, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		if !validateOpenCalculatorXYZInput(x) {
			panic("number out of range")
		}

		if _, ok := set[value]; !ok {
			set[value] = struct{}{}
		}
	}

	// N input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateOpenCalculatorNInput(n) {
		panic("number N out of range")
	}

	for _, numb := range strings.TrimSpace(line) {
		if _, ok := set[string(numb)]; !ok {
			set[string(numb)] = struct{}{}
		}
	}

	writer.WriteString(strconv.Itoa(len(set) - 3))
	writer.WriteByte('\n')
}
