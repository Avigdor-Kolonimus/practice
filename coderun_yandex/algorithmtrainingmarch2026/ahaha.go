package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/ahaha
// Ahaha - assignment 3
func Ahaha() {
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

	// response input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	response := strings.TrimRight(line, "\r\n")

	cur := 0
	ans := 0
	for i := 0; i < n; i++ {
		if response[i] != 'a' && response[i] != 'h' {
			cur = 0
		} else if cur == 0 {
			cur = 1
		} else if response[i] != response[i-1] {
			cur++
		} else {
			cur = 1
		}

		if cur > ans {
			ans = cur
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
