package problems

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/conveyor
// Conveyor - problem 17
func Conveyor() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	for ; n > 0; n-- {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum := strings.Fields(line)
		k, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}

		a := make([]float64, k)
		for i := 0; i < k; i++ {
			a[i], err = strconv.ParseFloat(strNum[i+1], 64)
			if err != nil {
				panic(err)
			}

		}

		target := make([]float64, k)
		copy(target, a)

		sort.Float64s(target)

		stack := make([]float64, 0, k)
		pos := 0

		for _, x := range a {
			for len(stack) > 0 && pos < k && stack[len(stack)-1] == target[pos] {
				stack = stack[:len(stack)-1]
				pos++
			}

			if pos < k && x == target[pos] {
				pos++
			} else {
				stack = append(stack, x)
			}
		}

		for len(stack) > 0 && pos < k && stack[len(stack)-1] == target[pos] {
			stack = stack[:len(stack)-1]
			pos++
		}

		if pos == k {
			writer.WriteByte('1')
		} else {
			writer.WriteByte('0')
		}
		writer.WriteByte('\n')
	}
}
