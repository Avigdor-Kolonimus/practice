package devgointerview

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/dev-go-interview/problems/request-counter
// RequestCounter - problem 5
func RequestCounter() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and K input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	k, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// request input
	data := make([]int, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		d, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		data[i] = d
	}

	// diff
	diff := make([]int, n)
	diff[0] = data[0]
	for i := 1; i < n; i++ {
		if data[i] >= data[i-1] {
			diff[i] = data[i] - data[i-1]
		} else {
			diff[i] = data[i]
		}
	}

	// window
	window := 0
	for i := 0; i < k; i++ {
		window += diff[i]
	}

	ans := window
	for i := k; i < n; i++ {
		window += diff[i] - diff[i-k]
		if window > ans {
			ans = window
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
