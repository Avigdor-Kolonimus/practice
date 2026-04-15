package problems

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/hr-tech-interview/problems/largest-product-two-numbers
// LargestProductTwoNumbers - problem 4
func LargestProductTwoNumbers() {
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

	// input
	line := readLine()
	nums := strings.Fields(line)

	if len(nums) < 2 || len(nums) > 100_000 {
		panic("the slice must contain at least 2 and at most 100,000 numbers")
	}

	max1, max2 := math.MinInt, math.MinInt
	min1, min2 := math.MaxInt64, math.MaxInt64

	for _, num := range nums {
		input, err := strconv.Atoi(num)
		if err != nil {
			break
		}

		if input >= max1 {
			max2, max1 = max1, input
		} else if input > max2 {
			max2 = input
		}

		if input < min1 {
			min2, min1 = min1, input
		} else if input < min2 {
			min2 = input
		}
	}

	result := strconv.Itoa(max2) + " " + strconv.Itoa(max1)
	if max2*max1 < min1*min2 {
		result = strconv.Itoa(min1) + " " + strconv.Itoa(min2)
	}

	writer.WriteString(result)
	writer.WriteByte('\n')
}
