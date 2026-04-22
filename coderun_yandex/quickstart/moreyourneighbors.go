package quickstart

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/quickstart/problems/more-your-neighbors
// MoreYourNeighbors - assignment 3
func MoreYourNeighbors() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// slice input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	strNum := strings.Fields(line)

	numerics := make([]int, len(strNum))
	for i, s := range strNum {
		digit, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		numerics[i] = digit
	}

	count := 0
	for i := 1; i < len(numerics)-1; i++ {
		if numerics[i] > numerics[i-1] && numerics[i] > numerics[i+1] {
			count++
		}
	}

	writer.WriteString(strconv.Itoa(count))
	writer.WriteByte('\n')
}
