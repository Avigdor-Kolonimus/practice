package quickstart

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/quickstart/problems/largest-product-three-numbers
// LargestProductThreeNumbers - assignment 14
func LargestProductThreeNumbers() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// slice input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	strNum := strings.Fields(line)

	max1, max2, max3 := math.MinInt, math.MinInt, math.MinInt
	min1, min2 := math.MaxInt, math.MaxInt
	for _, value := range strNum {
		x, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		switch {
		case x > max1:
			max3 = max2
			max2 = max1
			max1 = x
		case x > max2:
			max3 = max2
			max2 = x
		case x > max3:
			max3 = x
		}

		switch {
		case x < min1:
			min2 = min1
			min1 = x
		case x < min2:
			min2 = x
		}
	}

	res1 := max1 * max2 * max3
	res2 := max1 * min1 * min2
	if res1 > res2 {
		writer.WriteString(strconv.Itoa(max1) + " " + strconv.Itoa(max2) + " " + strconv.Itoa(max3))
	} else {
		writer.WriteString(strconv.Itoa(max1) + " " + strconv.Itoa(min2) + " " + strconv.Itoa(min1))
	}
	writer.WriteByte('\n')
}
