package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/angry-pigs
// AngryPigs - problem 203
func AngryPigs() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	used := make(map[int]struct{})
	for i := 0; i < n; i++ {
		// coordinates input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum := strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		x, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}

		used[x] = struct{}{}
	}

	writer.WriteString(strconv.Itoa(len(used)))
	writer.WriteByte('\n')
}
