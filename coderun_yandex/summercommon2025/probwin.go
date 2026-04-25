package summercommon2025

import (
	"bufio"
	"os"
	"strconv"
)

const (
	PRIME int = 1000000007
)

func binPow(residue, power int) int {
	result := 1
	residue %= PRIME

	for power > 0 {
		if power%2 == 1 {
			result = (result * residue) % PRIME
		}
		residue = (residue * residue) % PRIME
		power /= 2
	}

	return result
}

func solveProbWin(_, k, n int) int {
	var answer int = 1

	if n < k {
		n, k = k, n
	}

	for i := 1; i <= k; i++ {
		answer = (answer * i) % PRIME
		answer = (answer * binPow(n+i, PRIME-2)) % PRIME
	}

	return answer
}

// https://coderun.yandex.ru/selections/2025-summer-common/problems/prob-win
// ProbWin - problem 16
func ProbWin() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// A, K and N input
	line := mustReadIntArray(reader, 3)

	answer := solveProbWin(line[0], line[1], line[2])

	writer.WriteString(strconv.Itoa(answer))
	writer.WriteByte('\n')
}
