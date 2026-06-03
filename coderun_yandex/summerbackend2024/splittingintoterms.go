package summerbackend2024

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateSplittingIntoTermsNInput(p int) bool {
	return p >= 1 && p <= 40
}

func printPartition(writer *bufio.Writer, part []int) {
	for i, v := range part {
		if i > 0 {
			writer.WriteString(" + ")
		}
		writer.WriteString(strconv.Itoa(v))
	}
	writer.WriteByte('\n')
}

func generateLine(writer *bufio.Writer, lines []int, n, max int) {
	if n == 0 {
		printPartition(writer, lines)
		return
	}

	for x := 1; x <= min(n, max); x++ {
		lines = append(lines, x)
		generateLine(writer, lines, n-x, x)
		lines = lines[:len(lines)-1]
	}
}

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/splitting-into-terms
// SplittingIntoTerms - problem 15
func SplittingIntoTerms() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	strNum := strings.Fields(line)
	if len(strNum) != 1 {
		panic("numbers count does not match 1")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	if !validateSplittingIntoTermsNInput(n) {
		panic("number N out of range")
	}

	lines := make([]int, 0, n)
	generateLine(writer, lines, n, n)
}
