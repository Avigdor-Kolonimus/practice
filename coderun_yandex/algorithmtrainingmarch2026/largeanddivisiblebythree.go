package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func removeOne(cnt []int, digits []int) bool {
	for _, d := range digits {
		if cnt[d] > 0 {
			cnt[d]--

			return true
		}
	}

	return false
}

func removeTwo(cnt []int, digits []int) bool {
	tmp := make([]int, 10)
	copy(tmp, cnt)

	removed := 0
	for _, d := range digits {
		for tmp[d] > 0 && removed < 2 {
			tmp[d]--
			removed++
		}
	}

	if removed < 2 {
		return false
	}

	copy(cnt, tmp)

	return true
}

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/large-and-divisible-by-three
// LargeAndDivisibleByThree - assignment 11
func LargeAndDivisibleByThree() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// string input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	sum := 0
	cnt := make([]int, 10)
	for _, ch := range line {
		d := int(ch - '0')
		cnt[d]++
		sum += d
	}

	r := sum % 3

	if r == 1 {
		tmp := make([]int, 10)
		copy(tmp, cnt)

		if removeOne(tmp, []int{1, 4, 7}) {
			copy(cnt, tmp)
		} else {
			removeTwo(cnt, []int{2, 5, 8})
		}
	}

	if r == 2 {
		tmp := make([]int, 10)
		copy(tmp, cnt)

		if removeOne(tmp, []int{2, 5, 8}) {
			copy(cnt, tmp)
		} else {
			removeTwo(cnt, []int{1, 4, 7})
		}
	}

	var ans strings.Builder
	for d := 9; d >= 0; d-- {
		for i := 0; i < cnt[d]; i++ {
			ans.WriteByte(byte('0' + d))
		}
	}

	writer.WriteString(ans.String())
	writer.WriteByte('\n')
}
