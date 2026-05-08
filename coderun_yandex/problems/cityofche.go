package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/city-of-che
// CityOfChe - problem 211
func CityOfChe() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and R line
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

	r, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// distances input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum = strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	d := make([]int, n)
	for i := 0; i < n; i++ {
		x, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		d[i] = x
	}

	j, ans := 0, 0
	for i := 0; i < n; i++ {
		for j < n && d[j]-d[i] <= r {
			j++
		}

		if j < n {
			ans += (n - j)
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
