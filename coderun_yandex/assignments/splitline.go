package assignments

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateSplitLineInput(p int) bool {
	return p >= 1 && p <= 1_000_000
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/split-line
// SplitLine - assignment 8
func SplitLine() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first line
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	// N
	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	if !validateSplitLineInput(n) {
		panic("N out of range")
	}

	// M
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	if !validateSplitLineInput(m) {
		panic("m out of range")
	}

	// line input
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	currentLine := strings.TrimRight(line, "\r\n")
	if len(currentLine) != n {
		panic("line count does not match")
	}

	lines := make(map[string][]int)
	for i := range m {
		// subline input
		line, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		subLine := strings.TrimRight(line, "\r\n")

		lines[subLine] = append(lines[subLine], i+1)
	}

	length := n / m
	result := make([]string, 0, m)
	for i := 0; i < n; i = i + length {
		substr := currentLine[i : i+length]

		num := lines[substr][0]
		result = append(result, strconv.Itoa(num))

		lines[substr] = lines[substr][1:]
	}

	writer.WriteString(strings.Join(result, " "))
	writer.WriteByte('\n')
}
