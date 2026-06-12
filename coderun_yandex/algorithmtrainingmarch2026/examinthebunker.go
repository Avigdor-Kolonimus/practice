package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func absExamInTheBunker(x int64) int64 {
	if x < 0 {
		return -x
	}

	return x
}

func relaxExamInTheBunker(ans, x int64) int64 {
	x = absExamInTheBunker(x)

	if x < ans {
		ans = x
	}

	return ans
}

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/exam-in-the-bunker
// ExamInTheBunker - assignment 6
func ExamInTheBunker() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	ans := int64(math.MaxInt64)

	// equal
	for d := int64(1); d*d <= int64(n); d++ {
		if int64(n)%d == 0 {
			ans = relaxExamInTheBunker(ans, d-int64(n)/d)
		}
	}

	// r odd
	for d := int64(1); d*d <= int64(n); d++ {
		if int64(n)%d != 0 {
			continue
		}

		a := d
		b := int64(n) / d

		if b%2 == 1 {
			r := 2 * a
			m := (b + 1) / 2
			ans = relaxExamInTheBunker(ans, r-m)
		}

		if a%2 == 1 {
			r := 2 * b
			m := (a + 1) / 2
			ans = relaxExamInTheBunker(ans, r-m)
		}
	}

	// r even
	for _, N := range []int64{2*int64(n) - 1, 2*int64(n) + 1} {
		limit := int64(math.Sqrt(float64(N)))

		for d := int64(1); d <= limit; d++ {
			if N%d != 0 {
				continue
			}

			e := N / d

			r := d
			m := (e + 1) / 2
			ans = relaxExamInTheBunker(ans, r-m)

			r = e
			m = (d + 1) / 2
			ans = relaxExamInTheBunker(ans, r-m)
		}
	}

	writer.WriteString(strconv.FormatInt(ans, 10))
	writer.WriteByte('\n')
}
