package problems

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/queen-amount
// QueenAmount - problem 432
func QueenAmount() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	readLine := func() string {
		line, err := reader.ReadString('\n')
		if err != nil && !(err == io.EOF && len(line) > 0) {
			panic(err)
		}
		return strings.TrimRight(line, "\r\n")
	}

	// input
	line := readLine()
	nums := strings.Fields(line)

	arr := make([]int, 4)
	for i, num := range nums {
		input, err := strconv.Atoi(num)
		if err != nil {
			break
		}

		arr[i] = input
	}
	sort.Ints(arr)

	sum := 0
	i := 0
	for i = 0; i < 4; i++ {
		sum += arr[i]
		if sum > 4 {
			break
		}
	}

	ans := 4 - i
	if sum < 2 {
		ans = 1
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
