package autumnintern2023

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/autumn-intern-2023/problems/fortune-wheel
// FortuneWheel - problem 3
func FortuneWheel() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// B input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	b, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	n := 2*b - 1
	offset := b - 1

	A := make([][]float64, n)
	for i := range A {
		A[i] = make([]float64, n)
	}

	B := make([]float64, n)

	for d := -b + 1; d <= b-1; d++ {
		row := d + offset
		A[row][row] = 1.0
		B[row] = 1.0

		for x := 1; x <= 10; x++ {
			for y := 1; y <= 10; y++ {
				nd := d + x - y
				if -b < nd && nd < b {
					col := nd + offset
					A[row][col] -= 0.01 // 1/100
				}
			}
		}
	}

	for col := 0; col < n; col++ {
		pivot := col
		for i := col + 1; i < n; i++ {
			if math.Abs(A[i][col]) > math.Abs(A[pivot][col]) {
				pivot = i
			}
		}

		A[col], A[pivot] = A[pivot], A[col]
		B[col], B[pivot] = B[pivot], B[col]

		div := A[col][col]
		for j := col; j < n; j++ {
			A[col][j] /= div
		}
		B[col] /= div

		for i := 0; i < n; i++ {
			if i == col {
				continue
			}
			f := A[i][col]
			if math.Abs(f) < 1e-12 {
				continue
			}
			for j := col; j < n; j++ {
				A[i][j] -= f * A[col][j]
			}
			B[i] -= f * B[col]
		}
	}

	writer.WriteString(strconv.FormatFloat(B[offset], 'f', 9, 64))
	writer.WriteByte('\n')
}
