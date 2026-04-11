package assignments

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func validatePlusesMinusesQuestionsNandMInput(p int) bool {
	return p >= 0 && p <= 1_000
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/pluses-minuses-questions
// PlusesMinusesQuestions- assignment 6
func PlusesMinusesQuestions() {
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
	if !validatePlusesMinusesQuestionsNandMInput(n) {
		panic("N out of range")
	}

	// M
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	if !validatePlusesMinusesQuestionsNandMInput(m) {
		panic("M out of range")
	}

	// matrix input
	lineSum := make([]int, n)
	colSum := make([]int, m)
	matrix := make([][]rune, n)
	for i := range n {
		matrix[i] = make([]rune, m)

		// line input
		line, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		lCount := len(line)
		if lCount != m {
			writer.WriteString(strconv.Itoa(lCount))
			writer.WriteString(line)
			panic("matrix line count does not match m")
		}

		for j, r := range line {
			matrix[i][j] = r
			switch r {
			case '+':
				lineSum[i]++
				colSum[j]++
			case '?':
				lineSum[i]++
				colSum[j]--
			case '-':
				lineSum[i]--
				colSum[j]--
			default:
				panic("illegal rune")
			}
		}
	}

	maxDiff := math.MinInt64
	for i := range n {
		for j := range m {
			diff := lineSum[i] - colSum[j]

			if matrix[i][j] == '?' {
				diff -= 2
			}

			maxDiff = max(maxDiff, diff)
		}
	}

	writer.WriteString(strconv.Itoa(maxDiff))
	writer.WriteByte('\n')
}
