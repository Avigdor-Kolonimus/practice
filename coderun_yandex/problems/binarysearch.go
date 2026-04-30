package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func validateBinarySearchNKInput(p int) bool {
	return p >= 1 && p <= 100_000
}

// https://coderun.yandex.ru/problem/binary-search
// BinarySearch - problem 224
func BinarySearch() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and K line
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	if !validateBinarySearchNKInput(n) {
		panic("number N out of range")
	}

	k, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	if !validateBinarySearchNKInput(k) {
		panic("number K out of range")
	}

	// n-array inputs
	nArray := make([]int, n)
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}

	strNum = strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match N")
	}
	for i := range n {
		l, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		nArray[i] = l
	}

	// k-array inputs
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}

	strNum = strings.Fields(line)
	if len(strNum) != k {
		panic("numbers count does not match K")
	}
	for i := range k {
		l, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		low, high := 0, n-1
		result := "NO"
		for low <= high {
			mid := (low + high) / 2
			nNum := nArray[mid]

			if l == nNum {
				result = "YES"
				break
			} else if l > nNum {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

		writer.WriteString(result)
		writer.WriteByte('\n')
	}
}
