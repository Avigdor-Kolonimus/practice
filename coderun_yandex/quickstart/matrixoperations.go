package quickstart

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func validateMatrixOperationsInput(p int) bool {
	return p >= 1 && p <= 100
}

func validateMatrixOperationsMatrixElementInput(p int) bool {
	return p >= 0 && p <= 10
}

func multMatrix(a, b [][]int) [][]int {
	n := len(a)
	m := len(a[0])
	k := len(b[0])

	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, k)
		for j := 0; j < k; j++ {
			for t := 0; t < m; t++ {
				c[i][j] += a[i][t] * b[t][j]
			}
		}
	}

	return c
}

func transMatrix(a [][]int) [][]int {
	n := len(a)
	m := len(a[0])

	d := make([][]int, m)
	for i := 0; i < m; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			d[i][j] = a[j][i]
		}
	}

	return d
}

// https://coderun.yandex.ru/selections/quickstart/problems/matrix-operations
// MatrixOperations - assignment 9
func MatrixOperations() {
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
		panic("input does not match 3")
	}

	// N
	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	if !validateMatrixOperationsInput(n) {
		panic("number N out of range")
	}

	// M
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	if !validateMatrixOperationsInput(m) {
		panic("number M out of range")
	}

	// K
	k, err := strconv.Atoi(strNum[2])
	if err != nil {
		panic(err)
	}
	if !validateMatrixOperationsInput(k) {
		panic("number K out of range")
	}

	// Matrix A input
	matrixA := make([][]int, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		strNum = strings.Fields(line)
		if len(strNum) != m {
			panic("input does not match m")
		}

		matrixA[i] = make([]int, m)
		for j := 0; j < m; j++ {
			curr, err := strconv.Atoi(strNum[j])
			if err != nil {
				panic(err)
			}
			if !validateMatrixOperationsMatrixElementInput(curr) {
				panic("number Ai out of range")
			}

			matrixA[i][j] = curr
		}
	}

	// Matrix B input
	matrixB := make([][]int, m)
	for i := 0; i < m; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		strNum = strings.Fields(line)
		if len(strNum) != k {
			panic("input does not match k")
		}

		matrixB[i] = make([]int, k)
		for j := 0; j < k; j++ {
			curr, err := strconv.Atoi(strNum[j])
			if err != nil {
				panic(err)
			}
			if !validateMatrixOperationsMatrixElementInput(curr) {
				panic("number Bi out of range")
			}

			matrixB[i][j] = curr
		}
	}

	matrixC := multMatrix(matrixA, matrixB)
	matrixC = transMatrix(matrixC)

	for i, v := range matrixC {
		for j, d := range v {
			writer.WriteString(strconv.Itoa(d))
			if j < len(v)-1 {
				writer.WriteByte(' ')
			}
		}
		if i < len(matrixC)-1 {
			writer.WriteByte('\n')
		}
	}
}
