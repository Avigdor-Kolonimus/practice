package summerbackend2024

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/divisors-number
// DivisorsNumber - problem 12
func DivisorsNumber() {
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

	divs := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j += i {
			divs[j]++
		}
	}

	bestNum := 1
	bestCnt := 1
	for i := 1; i <= n; i++ {
		if divs[i] > bestCnt || (divs[i] == bestCnt && i > bestNum) {
			bestCnt = divs[i]
			bestNum = i
		}
	}

	writer.WriteString(strconv.Itoa(bestNum))
	writer.WriteByte('\n')
	writer.WriteString(strconv.Itoa(bestCnt))
	writer.WriteByte('\n')
}
