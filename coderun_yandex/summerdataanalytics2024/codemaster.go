package summerdataanalytics2024

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/2024-summer-data-analytics/problems/codemaster
// Codemaster - problem 17
func Codemaster() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	data, _ := io.ReadAll(reader)
	fields := strings.Fields(string(data))
	pos := 0

	readInt := func() int {
		v, _ := strconv.Atoi(fields[pos])
		pos++
		return v
	}

	n := readInt()
	values := make([]int, n)
	maxValue := 0

	for i := 0; i < n; i++ {
		values[i] = readInt()
		if values[i] > maxValue {
			maxValue = values[i]
		}
	}

	if maxValue == 0 {
		writer.WriteString("0\n")
		return
	}

	freq := make([]int, maxValue+1)
	for _, v := range values {
		freq[v]++
	}

	for i := maxValue - 1; i >= 0; i-- {
		freq[i] += freq[i+1]
	}

	kMax := n
	for (kMax > 0) && (kMax > maxValue/(kMax)) {
		kMax--
	}

	answer := 0
	for k := 1; k <= kMax; k++ {
		if freq[k*k] >= k {
			answer = k
		} else {
			break
		}
	}

	writer.WriteString(strconv.Itoa(answer))
	writer.WriteByte('\n')
}
