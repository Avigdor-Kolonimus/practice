package problems

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	eps = 1e-9
)

func isZero(x float64) bool {
	return math.Abs(x) < eps
}

func solveSingle(p, q, r float64) string {
	// 0x + 0y = r
	if isZero(p) && isZero(q) {
		if isZero(r) {
			return "5"
		}

		return "0"
	}

	// x = const
	if isZero(q) {
		x := r / p

		return fmt.Sprintf("3 %.5f", x)
	}

	// y = const
	if isZero(p) {
		y := r / q

		return fmt.Sprintf("4 %.5f", y)
	}

	// y = kx + b
	k := -p / q
	b := r / q

	return fmt.Sprintf("1 %.5f %.5f", k, b)
}

// https://coderun.yandex.ru/problem/system-of-linear-equations-2
// SystemOfLinearQquations2 - problem 188
func SystemOfLinearQquations2() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	nums := make([]int, 6)
	for i := range 6 {
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		x, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		nums[i] = x
	}

	a := float64(nums[0])
	b := float64(nums[1])
	c := float64(nums[2])
	d := float64(nums[3])
	e := float64(nums[4])
	f := float64(nums[5])

	det := a*d - b*c

	if !isZero(det) {
		x := (e*d - b*f) / det
		y := (a*f - e*c) / det
		writer.WriteString(fmt.Sprintf("2 %.5f %.5f\n", x, y))

		return
	}

	row1Zero := isZero(a) && isZero(b)
	row2Zero := isZero(c) && isZero(d)

	if row1Zero && row2Zero {
		if isZero(e) && isZero(f) {
			writer.WriteString("5\n")
		} else {
			writer.WriteString("0\n")
		}

		return
	}

	if row1Zero && !isZero(e) {
		writer.WriteString("0\n")
		return
	}

	if row2Zero && !isZero(f) {
		writer.WriteString("0\n")
		return
	}

	if row1Zero {
		writer.WriteString(solveSingle(c, d, f))
		writer.WriteByte('\n')

		return
	}

	if row2Zero {
		writer.WriteString(solveSingle(a, b, e))
		writer.WriteByte('\n')

		return
	}

	ok1 := isZero(a*f - c*e)
	ok2 := isZero(b*f - d*e)

	if ok1 && ok2 {
		writer.WriteString(solveSingle(a, b, e))
		writer.WriteByte('\n')

	} else {
		writer.WriteString("0\n")
	}
}
