package backendinterview

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateThreeBlocksRowInput(n int) bool {
	return n >= 0 && n <= 35
}

// https://coderun.yandex.ru/selections/backend-interview/problems/three-blocks-row
// ThreeBlocksRow - assignment 7
func ThreeBlocksRow() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateThreeBlocksRowInput(n) {
		panic("number N out of range")
	}

	var dp0 int = 1
	var dp1 int = 1
	var dp11 int = 0

	for range n - 1 {
		// ...0
		dp0New := dp0 + dp1 + dp11

		// ..01
		dp1New := dp0

		// ..011
		dp11New := dp1

		dp0, dp1, dp11 = dp0New, dp1New, dp11New
	}

	writer.WriteString(strconv.Itoa(dp0 + dp1 + dp11))
	writer.WriteByte('\n')
}
