package newyearadventures

import (
	"bufio"
	"io"
	"math"
	"os"
	"strings"
)

// https://coderun.yandex.ru/selections/new-year-adventures/problems/cha_cha
// ChaCha - problem 1
func ChaCha() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// score input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	sum := 0
	minVal := 25
	for i := 0; i < len(line); i++ {
		val := int('Z' - line[i])

		sum += val

		if val < minVal {
			minVal = val
		}
	}

	avg := float64(sum) / float64(len(line))

	res := int(math.Round(avg))

	if res > minVal+1 {
		res = minVal + 1
	}

	ans := 'Z' - byte(res)
	writer.WriteByte(ans)
	writer.WriteByte('\n')
}
