package quickstart

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func validateImprovingAcademicPerformanceInput(p float64) bool {
	return p >= 0 && p <= 1_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000
}

// https://coderun.yandex.ru/selections/quickstart/problems/improving-academic-performance
// ImprovingAcademicPerformance - assignment 16
func ImprovingAcademicPerformance() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// A input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	a, err := strconv.ParseFloat(line, 64)
	if err != nil {
		panic(err)
	}
	if !validateImprovingAcademicPerformanceInput(a) {
		panic("number A out of range")
	}

	// B input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	b, err := strconv.ParseFloat(line, 64)
	if err != nil {
		panic(err)
	}
	if !validateImprovingAcademicPerformanceInput(b) {
		panic("number B out of range")
	}

	// C input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	c, err := strconv.ParseFloat(line, 64)
	if err != nil {
		panic(err)
	}
	if !validateImprovingAcademicPerformanceInput(c) {
		panic("number C out of range")
	}

	if (a + b + c) < 1 {
		panic("sum A+B+C out of range")
	}

	result := int64(math.Ceil((-1 * (c - 3*a - b)) / 3))
	result = max(0, result)

	writer.WriteString(strconv.FormatInt(result, 10))
	writer.WriteByte('\n')
}
