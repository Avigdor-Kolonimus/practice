package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type City struct {
	index int
	price int
}

// https://coderun.yandex.ru/problem/great-lineland-migration
// GreatLinelandMigration- problem 98
func GreatLinelandMigration() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first  input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// price inputs
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	prices := strings.Fields(line)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		p, err := strconv.Atoi(prices[i])
		if err != nil {
			panic(err)
		}

		a[i] = p
	}

	ans := make([]int, n)
	stack := make([]City, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1].price >= a[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			ans[i] = -1
		} else {
			ans[i] = stack[len(stack)-1].index
		}

		stack = append(stack, City{
			index: i,
			price: a[i],
		})
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			writer.WriteByte(' ')
		}
		writer.WriteString(strconv.Itoa(ans[i]))
	}
	writer.WriteByte('\n')
}
