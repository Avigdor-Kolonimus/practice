package summerdataanalytics2024

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/2024-summer-data-analytics/problems/darkstore
// Darkstore - problem 16
func Darkstore() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	data, _ := io.ReadAll(reader)
	fields := strings.Fields(string(data))
	pos := 0

	readInt := func() int64 {
		n, _ := strconv.ParseInt(fields[pos], 10, 64)
		pos++
		return n
	}

	n := int(readInt())
	m := int(readInt())

	matrix := make([][]int64, n)
	for i := range n {
		matrix[i] = make([]int64, m)
		for j := range m {
			matrix[i][j] = readInt()
		}
	}

	kadane := func(arr []int64) int64 {
		curMax := arr[0]
		globalMax := arr[0]

		for i := 1; i < len(arr); i++ {
			if curMax+arr[i] > arr[i] {
				curMax += arr[i]
			} else {
				curMax = arr[i]
			}

			if curMax > globalMax {
				globalMax = curMax
			}
		}

		return globalMax
	}

	globalMax := matrix[0][0]

	for top := 0; top < n; top++ {
		colSum := make([]int64, m)

		for bottom := top; bottom < n; bottom++ {
			for col := 0; col < m; col++ {
				colSum[col] += matrix[bottom][col]
			}

			sum := kadane(colSum)
			if sum > globalMax {
				globalMax = sum
			}
		}
	}

	writer.WriteString(strconv.FormatInt(globalMax, 10))
	writer.WriteByte('\n')
}
