package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func checkPairedKPrefixes(words []string, k int) bool {
	cnt := make(map[string]int)

	for _, w := range words {
		cnt[w[:k]]++
	}

	for _, c := range cnt {
		if c%2 != 0 {
			return false
		}
	}

	return true
}

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/paired-k-prefixes
// PairedKPrefixes - assignment 8
func PairedKPrefixes() {
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

	// words input
	words := make([]string, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}

		words[i] = strings.TrimRight(line, "\r\n")
	}

	l := len(words[0])
	left, right := 0, l
	ans := 0
	for left <= right {
		mid := (left + right) / 2

		if checkPairedKPrefixes(words, mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
