package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func binarySearchDiplomas(w, h, n int) int {
	left, right := 0, max(w, h)*n

	for left < right {
		mid := (left + right) / 2

		if canFit(mid, w, h, n) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func canFit(x, w, h, n int) bool {
	return (x/w)*(x/h) >= n
}

func validateDiplomasInput(p int) bool {
	return p >= 1 && p <= 1_000_000_000
}

// https://coderun.yandex.ru/problem/diplomas
// Diplomas - problem 225
func Diplomas() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// W, H and N input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNum := strings.Fields(line)
	if len(strNum) != 3 {
		panic("numbers count does not match 3")
	}

	w, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	if !validateDiplomasInput(w) {
		panic("number w out of range")
	}

	h, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	if !validateDiplomasInput(h) {
		panic("number h out of range")
	}

	n, err := strconv.Atoi(strNum[2])
	if err != nil {
		panic(err)
	}
	if !validateDiplomasInput(n) {
		panic("number n out of range")
	}

	result := binarySearchDiplomas(w, h, n)

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
