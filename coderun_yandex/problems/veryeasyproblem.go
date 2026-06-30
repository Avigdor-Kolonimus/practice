package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/very-easy-problem
// VeryEasyProblem - problem 228
func VeryEasyProblem() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N, x and y input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNum := strings.Fields(line)
	if len(strNum) != 3 {
		panic("numbers count does not match 3")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	x, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(strNum[2])
	if err != nil {
		panic(err)
	}

	minTime := min(x, y)
	if n == 1 {
		writer.WriteString(strconv.Itoa(minTime))
		writer.WriteByte('\n')

		return
	}

	left := 0
	right := (n - 1) * minTime
	for left < right {
		mid := (left + right) / 2

		copies := mid/x + mid/y

		if copies >= n-1 {
			right = mid
		} else {
			left = mid + 1
		}
	}

	writer.WriteString(strconv.Itoa(minTime + left))
	writer.WriteByte('\n')
}
