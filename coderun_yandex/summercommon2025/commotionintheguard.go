package summercommon2025

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

func validateCommotionInTheGuardInput(n int) bool {
	return n >= 1 && n <= 100_000
}

// https://coderun.yandex.ru/selections/2025-summer-common/problems/commotion-in-the-guard
// CommotionInTheGuard - problem 3
func CommotionInTheGuard() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line := mustReadIntArray(reader, 1)
	n := line[0]
	if !validateCommotionInTheGuardInput(n) {
		panic("number N out of range")
	}

	// M input
	line = mustReadIntArray(reader, 1)
	m := line[0]
	if !validateCommotionInTheGuardInput(line[0]) {
		panic("number M out of range")
	}

	// swap inputs
	swaps := make([]int, 2*m)
	for i := 0; i < 2*m; i += 2 {
		line := mustReadIntArray(reader, 2)
		swaps[i], swaps[i+1] = line[0], line[1]
	}

	// position
	position := make([]int, 2*n+1)
	for i := 1; i <= 2*n; i++ {
		position[i] = i
	}

	answer := n
	for i := 0; i < m; i++ {
		x := swaps[2*i]
		y := swaps[2*i+1]

		if x > y {
			x, y = y, x
		}

		if x <= n && y > n {
			if position[x] <= n {
				answer--
			}
			if position[y] <= n {
				answer++
			}
		}

		// swap
		position[x], position[y] = position[y], position[x]

		writer.WriteString(strconv.Itoa(answer))
		writer.WriteByte(' ')
	}
	writer.WriteByte('\n')
}
