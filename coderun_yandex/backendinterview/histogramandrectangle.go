package backendinterview

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateHistogramAndRectangleNInput(n int) bool {
	return n >= 1 && n <= 1_000_000
}

func validateHistogramAndRectangleHInput(h int) bool {
	return h >= 0 && h <= 1_000_000_000
}

// https://coderun.yandex.ru/selections/backend-interview/problems/histogram-and-rectangle
// HistogramAndRectangle - assignment 10
func HistogramAndRectangle() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) == 0 {
		panic("empty input")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	if !validateHistogramAndRectangleNInput(n) {
		panic("number n out of range")
	}

	// input columns
	h := make([]int, n+2)
	for i := 0; i < n; i++ {
		hi, err := strconv.Atoi(strNum[i+1])
		if err != nil {
			panic(err)
		}
		if !validateHistogramAndRectangleHInput(hi) {
			panic("number hi out of range")
		}

		h[i+1] = hi
	}

	stack := []int{}
	maxArea, l := 0, 0

	for i := range h {
		for len(stack) > 0 && h[i] < h[stack[len(stack)-1]] {
			height := h[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			if len(stack) != 0 {
				l = stack[len(stack)-1]
			}

			width := i - l - 1
			maxArea = max(maxArea, width*height)

		}
		stack = append(stack, i)
	}

	writer.WriteString(strconv.Itoa(maxArea))
	writer.WriteByte('\n')
}
