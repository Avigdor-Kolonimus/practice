package first2023mobiledev

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/segments-with-min-mex
// SegmentsWithMinMex - problem 33
func SegmentsWithMinMex() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	var result int64
	var lastOne int64 = -1
	for r := int64(0); r < int64(n); r++ {
		var x int64
		fmt.Fscan(reader, &x)

		if x != 1 {
			result += r - lastOne
		} else {
			lastOne = r
		}
	}

	if result == 0 {
		n64 := int64(n)
		result = (1 + n64) * n64 / 2
	}

	writer.WriteString(strconv.FormatInt(result, 10))
	writer.WriteByte('\n')
}
