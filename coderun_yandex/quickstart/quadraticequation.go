package quickstart

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func validateQuadraticEquationInput(n float64) bool {
	return n >= -100 && n <= 100
}

// https://coderun.yandex.ru/selections/quickstart/problems/quadratic-equation
// QuadraticEquation - assignment 5
func QuadraticEquation() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// slice input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}

	strNum := strings.Fields(line)
	if len(strNum) != 3 {
		panic("numbers count does not match 3")
	}

	// A
	a, err := strconv.ParseFloat(strNum[0], 64)
	if err != nil {
		panic(err)
	}

	if a == 0 || !validateQuadraticEquationInput(a) {
		panic("number A out of range")
	}

	// B
	b, err := strconv.ParseFloat(strNum[1], 64)
	if err != nil {
		panic(err)
	}

	if !validateQuadraticEquationInput(b) {
		panic("number B out of range")
	}

	// C
	c, err := strconv.ParseFloat(strNum[2], 64)
	if err != nil {
		panic(err)
	}

	if !validateQuadraticEquationInput(c) {
		panic("number C out of range")
	}

	d := math.Pow(b, 2) - 4*a*c
	switch {
	case d > 0:
		x1 := (-b - math.Sqrt(d)) / (2 * a)
		x2 := (-b + math.Sqrt(d)) / (2 * a)
		if x1 > x2 {
			x1, x2 = x2, x1
		}

		writer.WriteByte('2')
		writer.WriteByte('\n')
		writer.WriteString(fmt.Sprintf("%.6f %.6f", x1, x2))
	case d == 0:
		x := -b / (2 * a)
		writer.WriteByte('1')
		writer.WriteByte('\n')
		writer.WriteString(fmt.Sprintf("%.6f", x))
	default:
		writer.WriteByte('0')
	}
	writer.WriteByte('\n')
}
