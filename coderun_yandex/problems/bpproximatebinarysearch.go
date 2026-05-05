package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func lowerBound(a []int, x int) int {
	l, r := 0, len(a)
	for l < r {
		m := (l + r) / 2
		if a[m] < x {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}

func absBpproximateBinarySearch(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func validateBpproximateBinarySearchInput(n int) bool {
	return n > 0 && n < 100_001
}

// https://coderun.yandex.ru/problem/bpproximate-binary-search
// BpproximateBinarySearch - problem 76
func BpproximateBinarySearch() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and K input
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
	if !validateBpproximateBinarySearchInput(n) {
		panic("number N out of range")
	}

	k, err := strconv.Atoi(strKeys[1])
	if err != nil {
		panic(err)
	}
	if !validateBpproximateBinarySearchInput(k) {
		panic("number K out of range")
	}

	// nArray input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNums := strings.Fields(line)

	nArray := make([]int, n)
	for i := 0; i < n; i++ {
		num, err := strconv.Atoi(strNums[i])
		if err != nil {
			panic(err)
		}

		nArray[i] = num
	}

	// kArray input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strQueries := strings.Fields(line)

	result := 0
	for i := 0; i < k; i++ {
		num, err := strconv.Atoi(strQueries[i])
		if err != nil {
			panic(err)
		}

		pos := lowerBound(nArray, num)

		// candidate
		switch pos {
		case 0:
			result = nArray[0]
		case n:
			result = nArray[n-1]
		default:
			left := nArray[pos-1]
			right := nArray[pos]

			if absBpproximateBinarySearch(left-num) <= absBpproximateBinarySearch(right-num) {
				result = left
			} else {
				result = right
			}
		}

		writer.WriteString(strconv.Itoa(result))
		writer.WriteByte('\n')
	}
}
