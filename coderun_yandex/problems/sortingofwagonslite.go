package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/sorting-of-wagons-lite
// SortingOfWagonsLite - problem 250
func SortingOfWagonsLite() {
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

	// stack
	targetVan := 1
	stack := make([]int, 0, n)

	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	for i := 0; i < n; i++ {
		inp, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		if inp != targetVan {
			stack = append(stack, inp)
		} else {
			targetVan++
			for len(stack) > 0 && stack[len(stack)-1] == targetVan {
				targetVan++
				stack = stack[:len(stack)-1]
			}
		}
	}

	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] != targetVan {
			writer.WriteString("NO")
			writer.WriteByte('\n')
			return
		} else {
			targetVan++
		}
	}

	writer.WriteString("YES")
	writer.WriteByte('\n')
}
