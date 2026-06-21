package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/turtles
// Turtles - problem 202
func Turtles() {
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

	// turtels input
	ans := 0
	seen := make(map[[2]int]struct{})
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum := strings.Fields(line)

		a, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		if a+b != n-1 {
			continue
		}

		if a < 0 || a >= n || b < 0 || b >= n {
			continue
		}

		key := [2]int{a, b}

		if _, ok := seen[key]; !ok {
			seen[key] = struct{}{}
			ans++
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
