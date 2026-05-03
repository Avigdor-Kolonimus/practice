package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func validateMedianUnionNInput(n int) bool {
	return n >= 2 && n <= 100
}

func validateMedianUnionLInput(l int) bool {
	return l >= 1 && l <= 300
}

func kth(a, b []int, k int) int {
	i, j := 0, 0

	for {
		if i == len(a) {
			return b[j+k]
		}
		if j == len(b) {
			return a[i+k]
		}

		if k == 0 {
			if a[i] < b[j] {
				return a[i]
			}
			return b[j]
		}

		step := (k + 1) / 2

		ni := min(i+step, len(a)) - 1
		nj := min(j+step, len(b)) - 1

		if a[ni] <= b[nj] {
			k -= (ni - i + 1)
			i = ni + 1
		} else {
			k -= (nj - j + 1)
			j = nj + 1
		}
	}
}

// https://coderun.yandex.ru/problem/median-union
// MedianUnion - problem 80
func MedianUnion() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and L input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strKeys := strings.Fields(line)
	if len(strKeys) != 2 {
		panic("invalid input")
	}

	n, err := strconv.Atoi(strKeys[0])
	if err != nil {
		panic(err)
	}
	if !validateMedianUnionNInput(n) {
		panic("number N out of range")
	}

	l, err := strconv.Atoi(strKeys[1])
	if err != nil {
		panic(err)
	}
	if !validateMedianUnionLInput(l) {
		panic("number L out of range")
	}

	// list input
	seqList := make([][]int, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		strNum := strings.Fields(line)
		if len(strNum) != l {
			panic("numbers count does not match l")
		}

		seqList[i] = make([]int, l)
		for j := 0; j < l; j++ {
			seqList[i][j], err = strconv.Atoi(strNum[j])
			if err != nil {
				panic(err)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			median := kth(seqList[i], seqList[j], l-1)

			writer.WriteString(strconv.Itoa(median))
			writer.WriteByte('\n')
		}
	}
}
