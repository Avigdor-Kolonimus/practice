package summercommon2025

import (
	"bufio"
	"os"
	"strconv"
)

func validateStairsInput(n int) bool {
	return n >= 1 && n <= 100_000
}

func solveStairs(n int, a []int) []int {
	prev := -1000000
	result := make([]int, 0, n+1)
	result = append(result, 1)
	for _, num := range a {
		if num < 0 {
			num = -num
		}

		if prev <= -num {
			prev = -num
		} else if prev <= num {
			prev = num
		} else {
			return []int{0}
		}

		result = append(result, prev)
	}

	return result
}

// https://coderun.yandex.ru/selections/2025-summer-common/problems/stairs
// Stairs - problem 4
func Stairs() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line := mustReadIntArray(reader, 1)
	n := line[0]
	if !validateStairsInput(n) {
		panic("number N out of range")
	}

	// stairs input
	stairs := make([]int, n)
	line = mustReadIntArray(reader, n)
	for i := range n {
		stairs[i] = line[i]
	}

	answer := solveStairs(n, stairs)
	for _, v := range answer {
		writer.WriteString(strconv.Itoa(v) + " ")
	}
	writer.WriteByte('\n')
}
