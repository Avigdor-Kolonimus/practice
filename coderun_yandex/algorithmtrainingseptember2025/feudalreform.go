package algorithmtrainingseptember2025

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func absFeudalReform(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/feudal-reform
// FeudalReform - assignment 26
func FeudalReform() {
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

	// children input
	children := make([][]int, n)
	for v := 1; v < n; v++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		p, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		children[p] = append(children[p], v)
	}

	// value input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	a := make([]int, n)
	for i := 0; i < n; i++ {
		ai, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		a[i] = ai
	}

	ans := 0
	for v := 0; v < n; v++ {
		sumChildren := 0
		for _, c := range children[v] {
			sumChildren += a[c]
		}

		x := sumChildren - a[v]
		ans += absFeudalReform(x)
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
