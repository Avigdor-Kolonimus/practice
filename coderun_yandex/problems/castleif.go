package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateCastleIfInput(p int) bool {
	return p >= 1 && p <= 10_000
}

// https://coderun.yandex.ru/problem/castle-if
// CastleIf - problem 58
func CastleIf() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// 5 inputs
	numInputs := make([]int, 5)
	for i := range 5 {
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		if !validateCastleIfInput(n) {
			panic("number out of range")
		}

		numInputs[i] = n
	}

	result := "NO"
	if (numInputs[0] <= numInputs[3] && numInputs[1] <= numInputs[4]) || (numInputs[0] <= numInputs[4] && numInputs[1] <= numInputs[3]) {
		result = "YES"
	}
	if (numInputs[0] <= numInputs[3] && numInputs[2] <= numInputs[4]) || (numInputs[0] <= numInputs[4] && numInputs[2] <= numInputs[3]) {
		result = "YES"
	}
	if (numInputs[2] <= numInputs[3] && numInputs[1] <= numInputs[4]) || (numInputs[2] <= numInputs[4] && numInputs[1] <= numInputs[3]) {
		result = "YES"
	}

	writer.WriteString(result)
	writer.WriteByte('\n')
}
