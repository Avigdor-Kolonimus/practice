package winterintern2024

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/winter-intern-2024/problems/lap
// Lap - problem 2
func Lap() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// n, t, s input
	line, _ := reader.ReadString('\n')
	fields := strings.Fields(line)
	if len(fields) < 3 {
		return
	}
	n, _ := strconv.Atoi(fields[0])
	t, _ := strconv.ParseInt(fields[1], 10, 64)
	s, _ := strconv.ParseInt(fields[2], 10, 64)

	// speeds input
	line, _ = reader.ReadString('\n')
	vFields := strings.Fields(line)

	v1, _ := strconv.ParseInt(vFields[0], 10, 64)
	var totalOvertakes int64 = 0
	for i := 1; i < n; i++ {
		vi, _ := strconv.ParseInt(vFields[i], 10, 64)

		if v1 > vi {
			diffV := v1 - vi
			distDiff := diffV * t

			if distDiff >= s {
				totalOvertakes += (distDiff - 1) / s
			}
		}
	}

	writer.WriteString(strconv.FormatInt(totalOvertakes, 10))
	writer.WriteByte('\n')
}
