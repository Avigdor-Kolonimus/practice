package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func validateSubstringNInput(n int) bool {
	return n >= 1 && n <= 100_000
}

func validateSubstringKInput(n, k int) bool {
	return k >= 1 && k <= n
}

// https://coderun.yandex.ru/problem/substring
// Substring - problem 75
func Substring() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and K input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	parameters := strings.Fields(line)
	if len(parameters) != 2 {
		panic("input does not match 2")
	}

	// N
	n, err := strconv.Atoi(parameters[0])
	if err != nil {
		panic(err)
	}

	if !validateSubstringNInput(n) {
		panic("number N out of range")
	}

	// K
	k, err := strconv.Atoi(parameters[1])
	if err != nil {
		panic(err)
	}

	if !validateSubstringKInput(n, k) {
		panic("number K out of range")
	}

	// string input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	left := 0
	bestLen := 0
	bestL := 0
	cnt := make([]int, 26)
	for right := 0; right < len(line); right++ {
		c := line[right] - 'a'
		cnt[c]++

		for cnt[c] > k {
			cnt[line[left]-'a']--
			left++
		}

		if right-left+1 > bestLen {
			bestLen = right - left + 1
			bestL = left
		}
	}

	writer.WriteString(strconv.Itoa(bestLen))
	writer.WriteByte(' ')
	writer.WriteString(strconv.Itoa(bestL + 1))
	writer.WriteByte('\n')
}
