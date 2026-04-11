package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateMedianNumInput(n int) bool {
	return n >= -1_000 && n <= 1_000
}

func median(a, b, c int) int {
	if (a >= b && a <= c) || (a >= c && a <= b) {
		return a
	}
	if (b >= a && b <= c) || (b >= c && b <= a) {
		return b
	}
	return c
}

// https://coderun.yandex.ru/problem/median-out-of-three
// MedianOutOfThree - problem 1
func MedianOutOfThree() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	strNum := strings.Fields(line)
	if len(strNum) != 3 {
		panic("numbers count does not match 3")
	}

	nums := make([]int, 3)
	for i, v := range strNum {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		if !validateMedianNumInput(n) {
			panic("number out of range")
		}

		nums[i] = n
	}

	m := median(nums[0], nums[1], nums[2])

	writer.WriteString(strconv.Itoa(m))
	writer.WriteByte('\n')
}
