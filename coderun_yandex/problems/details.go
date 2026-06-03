package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/details
// Details - problem 186
func Details() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N, K and M input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 3 {
		panic("numbers count does not match 3")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	k, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(strNum[2])
	if err != nil {
		panic(err)
	}

	if k < m {
		writer.WriteByte('0')
		writer.WriteByte('\n')

		return
	}

	ans := 0
	for n >= k {
		blanks := n / k

		ans += blanks * (k / m)

		n = n%k + blanks*(k%m)
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
