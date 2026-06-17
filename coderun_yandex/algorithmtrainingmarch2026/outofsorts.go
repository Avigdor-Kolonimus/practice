package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/out-of-sorts
// OutOfSorts - assignment 12
func OutOfSorts() {
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

	// dishes input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	dishes := make([]int, n+1)
	for i := 1; i <= n; i++ {
		x, err := strconv.Atoi(strNum[i-1])
		if err != nil {
			panic(err)
		}

		dishes[x] = i
	}

	forbidden := make([]bool, n)
	for i := 1; i <= n; i++ {
		k := (i - dishes[i]) % n
		if k < 0 {
			k += n
		}

		forbidden[k] = true
	}

	for k := 0; k < n; k++ {
		if !forbidden[k] {
			writer.WriteString(strconv.Itoa(k))
			writer.WriteByte('\n')

			return
		}
	}

	writer.WriteString(strconv.Itoa(-1))
	writer.WriteByte('\n')
}
