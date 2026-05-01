package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func heapify(arr []int, n int, i int) {
	left, right := 2*i+1, 2*i+2
	largest := left

	if left >= n {
		return
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if arr[largest] <= arr[i] {
		return
	}

	arr[i], arr[largest] = arr[largest], arr[i]

	heapify(arr, n, largest)
}

func heapSort(arr []int) {
	n := len(arr)

	// 1. build heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 2. sorting
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func validateRemovingDuplicatesInput(n int) bool {
	return n >= 0 && n <= 100_000
}

// https://coderun.yandex.ru/problem/pyramid-sorting
// PyramidSorting - problem 252
func PyramidSorting() {
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
	if !validateRemovingDuplicatesInput(n) {
		panic("number N out of range")
	}

	// second line input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		num, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}
		arr[i] = num
	}

	heapSort(arr)

	for i := range n {
		writer.WriteString(strconv.Itoa(arr[i]) + " ")
	}
	writer.WriteByte('\n')
}
