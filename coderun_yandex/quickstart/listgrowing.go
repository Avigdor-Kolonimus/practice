package quickstart

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/quickstart/problems/list-growing
// ListGrowing - assignment 8
func ListGrowing() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// slice input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	strNum := strings.Fields(line)

	prev, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}

	isMonotonic := "YES"
	for i := 1; i < len(strNum); i++ {
		curr, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		if curr <= prev {
			isMonotonic = "NO"
			break
		}
		prev = curr
	}

	writer.WriteString(isMonotonic)
	writer.WriteByte('\n')
}
