package problems

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/bypass
// Bypass - problem 241
func Bypass() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// Values input
	line, _ := reader.ReadString('\n')
	values := strings.Fields(line)

	unique := make(map[int]struct{})
	for _, value := range values {
		n, _ := strconv.Atoi(value)

		if n == 0 {
			break
		}

		unique[n] = struct{}{}
	}

	result := make([]int, 0, len(unique))
	for n := range unique {
		result = append(result, n)
	}

	sort.Ints(result)

	for _, n := range result {
		writer.WriteString(strconv.Itoa(n))
		writer.WriteByte('\n')
	}
}
