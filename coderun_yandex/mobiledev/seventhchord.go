package mobiledev

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
)

const (
	maxA    = 100000
	maxRoot = 316
)

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/seventh-chord
// SeventhChord - problem 12
func SeventhChord() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	data, _ := io.ReadAll(reader)
	pos := 0

	readInt := func() int {
		for pos < len(data) && data[pos] <= ' ' {
			pos++
		}

		x := 0
		for pos < len(data) && data[pos] >= '0' && data[pos] <= '9' {
			x = x*10 + int(data[pos]-'0')
			pos++
		}
		return x
	}

	n := readInt()

	counts := make([]int, maxRoot+1)
	maxIndex := 0
	for range n {
		x := readInt()
		root := int(math.Sqrt(float64(x)))

		counts[root]++

		if root > maxIndex {
			maxIndex = root
		}
	}

	septaccords := 0
	for k := maxIndex; k >= 1; k-- {
		septaccords += counts[k]

		if septaccords >= k {
			writer.WriteString(strconv.Itoa(k))
			writer.WriteByte('\n')
			return
		}
	}

	writer.WriteString("0\n")
}
