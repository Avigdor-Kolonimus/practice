package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func validateSnowballsInput(t int) bool {
	return t >= 1 && t <= 10
}

// https://coderun.yandex.ru/problem/snowballs
// Snowballs - problem 618
func Snowballs() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// T line
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	t, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateSnowballsInput(t) {
		panic("number T out of range")
	}

	// heaps input
	for range t {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}

		strNum := strings.Fields(line)
		if len(strNum) != 3 {
			panic("numbers count does not match 3")
		}

		heapA, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		heapB, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}
		heapC, err := strconv.Atoi(strNum[2])
		if err != nil {
			panic(err)
		}

		x := (heapA % 3) ^ (heapB % 3) ^ (heapC % 3)
		if x > 0 {
			writer.WriteByte('1')
		} else {
			writer.WriteByte('0')
		}
		writer.WriteByte('\n')
	}
}
