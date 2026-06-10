package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/beauty-above-all
// BeautyAboveAll - problem 212
func BeautyAboveAll() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and K input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("input does not match 2")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	k, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// colors input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum = strings.Fields(line)
	if len(strNum) != n {
		panic("input does not match n")
	}

	a := make([]int, n)
	for i := range n {
		ai, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		a[i] = ai
	}

	cnt := make([]int, k+1)

	distinct := 0
	l := 0

	bestL, bestR := 0, n-1
	bestLen := n + 1

	for r := 0; r < n; r++ {
		color := a[r]

		if cnt[color] == 0 {
			distinct++
		}
		cnt[color]++

		for distinct == k {
			if r-l+1 < bestLen {
				bestLen = r - l + 1
				bestL = l
				bestR = r
			}

			leftColor := a[l]
			cnt[leftColor]--

			if cnt[leftColor] == 0 {
				distinct--
			}

			l++
		}
	}

	writer.WriteString(strconv.Itoa(bestL + 1))
	writer.WriteByte(' ')
	writer.WriteString(strconv.Itoa(bestR + 1))
	writer.WriteByte('\n')
}
