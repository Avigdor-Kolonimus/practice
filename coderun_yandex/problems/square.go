package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func can(n, m, t, w int64) bool {
	used := n*m - (n-2*w)*(m-2*w)

	return used <= t
}

// https://coderun.yandex.ru/problem/square
// Square - problem 78
func Square() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		panic(err)
	}

	// M input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	m, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		panic(err)
	}

	// T input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	t, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		panic(err)
	}

	var left, ans int64
	right := min(n, m) / 2
	for left <= right {
		mid := (left + right) / 2

		if can(n, m, t, mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	writer.WriteString(strconv.FormatInt(ans, 10))
	writer.WriteByte('\n')
}
