package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// https://coderun.yandex.ru/selections/hr-tech-interview/problems/nearest-number
// NearestNumber - problem 2
func NearestNumber() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	readLine := func() string {
		line, err := reader.ReadString('\n')
		if err != nil && !(err == io.EOF && len(line) > 0) {
			panic(err)
		}
		return strings.TrimRight(line, "\r\n")
	}

	// first input
	line := readLine()
	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	arr := make([]int, 0, n)
	// array input
	line = readLine()
	digits := strings.Fields(line)

	for i := range digits {
		num, err := strconv.Atoi(digits[i])
		if err != nil {
			panic(err)
		}
		arr = append(arr, num)
	}

	// last input
	line = readLine()
	x, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	minDiff := abs(x - arr[0])
	res := arr[0]
	for _, v := range arr {
		if diff := abs(x - v); diff < minDiff {
			minDiff = diff
			res = v
		}
	}

	writer.WriteString(strconv.Itoa(res))
	writer.WriteByte('\n')
}
