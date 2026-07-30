package atolstikov

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	SEC int64 = 1_000_000_000
)

// https://coderun.yandex.ru/selections/atolstikov/problems/frequency-limitation
// FrequencyLimitation - problem 2
func FrequencyLimitation() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// X / Y input
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	X, _ := strconv.Atoi(parts[0])
	Y, _ := strconv.Atoi(parts[2])

	// N input
	line, _ = reader.ReadString('\n')
	N, _ := strconv.Atoi(strings.TrimSpace(line))

	// request input
	line, _ = reader.ReadString('\n')
	parts = strings.Fields(line)
	ans := make([]int64, N)
	if X == 1 {
		interval := int64(Y) * SEC
		var last int64 = -interval

		for i := range N {
			t, _ := strconv.ParseInt(parts[i], 10, 64)

			ans[i] = max(t, last+interval)
			last = ans[i]
		}
	} else {
		for i := range N {
			t, _ := strconv.ParseInt(parts[i], 10, 64)

			if i < X {
				ans[i] = t
			} else {
				ans[i] = max(t, ans[i-X]+SEC)
			}
		}
	}

	for i := range N {
		writer.WriteString(strconv.FormatInt(ans[i], 10))
		if i+1 != N {
			writer.WriteByte(' ')
		}
	}
	writer.WriteByte('\n')
}
