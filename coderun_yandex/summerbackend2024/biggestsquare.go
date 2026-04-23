package summerbackend2024

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func mustReadLine(reader *bufio.Reader) string {
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(line)
}

func mustReadIntArray(reader *bufio.Reader, size int) []int {
	rawNumbers := strings.Split(mustReadLine(reader), " ")
	if len(rawNumbers) != size {
		panic("len must be eq size")
	}

	result := make([]int, 0, size)
	for _, rawNumber := range rawNumbers {
		number, err := strconv.Atoi(rawNumber)
		if err != nil {
			panic(err)
		}
		result = append(result, number)
	}

	return result
}

func validateBiggestSquareInput(n int) bool {
	return n >= 1 && n <= 1_000
}

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/biggest-square
// BiggestSquare - problem 26
func BiggestSquare() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	firstLine := mustReadIntArray(reader, 2)
	if !validateBiggestSquareInput(firstLine[0]) {
		panic("number N out of range")
	}
	if !validateBiggestSquareInput(firstLine[1]) {
		panic("number M out of range")
	}
	n, m := firstLine[0], firstLine[1]

	var x, xi, xj int
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, m)
		nums := mustReadIntArray(reader, m)
		for j := 0; j < m; j++ {
			if nums[j] == 1 {
				dp[i][j] = 1
				x, xi, xj = 1, i, j
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if dp[i][j] == 1 {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1])

				if dp[i][j] >= x {
					x = dp[i][j]
					xi = i - x + 1
					xj = j - x + 1
				}
			}
		}
	}

	writer.WriteString(strconv.Itoa(x))
	writer.WriteByte('\n')
	writer.WriteString(strconv.Itoa(xi+1) + " " + strconv.Itoa(xj+1))
	writer.WriteByte('\n')
}
