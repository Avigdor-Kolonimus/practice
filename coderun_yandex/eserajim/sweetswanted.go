package eserajim

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func validateSweetsWantedInput(p int) bool {
	return p >= 1 && p <= 100_000
}

// https://coderun.yandex.ru/selections/eserajim/problems/sweets-wanted
// SweetsWanted - problem 1
func SweetsWanted() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateSweetsWantedInput(n) {
		panic("number N out of range")
	}

	// candy jar input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	// points
	candyJar := make([]int, n)
	for i, v := range strNum {
		x, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		candyJar[i] = x
	}

	sort.Ints(candyJar)

	n = len(candyJar)
	prefixSum := make([]int, n+1)

	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + candyJar[i]
	}

	totalSum := prefixSum[n]
	maxVal := candyJar[n-1]

	ans := int(1 << 62)

	for i := 0; i < n; i++ {
		x := candyJar[i]

		countL := i + 1
		sumL := prefixSum[i+1]

		costL := x*countL - sumL

		countR := n - (i + 1)
		sumR := totalSum - sumL

		costR := maxVal*countR - sumR

		if costL+costR < ans {
			ans = costL + costR
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
