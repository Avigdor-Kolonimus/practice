package problems

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/maxim-triangle
// MaximTriangle - problem 196
func MaximTriangle() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	left, right := 30.0, 4000.0

	line, _ = reader.ReadString('\n')
	prev, _ := strconv.ParseFloat(strings.TrimSpace(line), 64)

	for i := 1; i < n; i++ {
		line, _ = reader.ReadString('\n')
		fields := strings.Fields(line)

		curr, _ := strconv.ParseFloat(fields[0], 64)
		direction := fields[1]

		mid := (prev + curr) / 2

		if (curr > prev && direction == "closer") ||
			(curr < prev && direction == "further") {
			if mid > left {
				left = mid
			}
		} else if (curr > prev && direction == "further") ||
			(curr < prev && direction == "closer") {
			if mid < right {
				right = mid
			}
		}

		prev = curr
	}

	ans := fmt.Sprintf("%.10f %.10f\n", left, right)
	writer.WriteString(ans)
}
