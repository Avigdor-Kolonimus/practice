package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/final-grade
// FinalGrade - assignment 17
func FinalGrade() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// grades input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	sum := 0
	worst := 0
	n := len(s)
	for i, c := range line {
		v := int(c - 'A')
		sum += v

		// first, 'Z' > 'A' (ASCII)
		if i == 0 || v > worst {
			worst = v
		}
	}

	//ceil
	q := sum / n
	r := sum % n

	rounded := q
	if 2*r > n {
		rounded++
	}

	limit := worst - 1
	if limit < 0 {
		limit = 0
	}

	if rounded < limit {
		rounded = limit
	}

	writer.WriteByte(byte('A' + rounded))
	writer.WriteByte('\n')
}
