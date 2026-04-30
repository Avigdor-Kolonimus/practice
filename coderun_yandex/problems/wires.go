package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func canCut(wires []int, k, l int) bool {
	total := 0
	for _, w := range wires {
		total += w / l
		if total >= k {
			return true
		}
	}

	return false
}

func validateWiresNKInput(p int) bool {
	return p >= 1 && p <= 10_000
}

func validateWiresLInput(l int) bool {
	return l >= 100 && l <= 10_000_000
}

// https://coderun.yandex.ru/problem/wires
// Wires - problem 230
func Wires() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and K line
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	if !validateWiresNKInput(n) {
		panic("number N out of range")
	}

	k, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	if !validateWiresNKInput(k) {
		panic("number K out of range")
	}

	// lenghts inputs
	maxL := 0
	wires := make([]int, n)
	for i := range n {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		l, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		if !validateWiresLInput(l) {
			panic("number L out of range")
		}

		wires[i] = l
		maxL = max(l, maxL)
	}

	low, high, result := 1, maxL, 0
	for low <= high {
		mid := (low + high) / 2
		if canCut(wires, k, mid) {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
