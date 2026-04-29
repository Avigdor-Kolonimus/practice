package summercommon2025

import (
	"bufio"
	"os"
	"strconv"
)

func validateSheepWolfInput(n int) bool {
	return n >= 1 && n <= 1_000_000
}

func solveSheepWolf(n int, a []int) int {
	ans := 0
	twoDiffLen := 1
	oneSameLen := 1

	id1 := a[0]
	id2 := -1

	for i := 1; i < n; i++ {
		if a[i] == a[i-1] {
			twoDiffLen++
			oneSameLen++
		} else {
			if id2 == -1 {
				id2 = a[i]
				twoDiffLen++
			} else if a[i] == id1 || a[i] == id2 {
				twoDiffLen++
			} else {
				if twoDiffLen > ans {
					ans = twoDiffLen
				}
				twoDiffLen = oneSameLen + 1
				id1 = a[i-1]
				id2 = a[i]
			}
			oneSameLen = 1
		}
	}

	if twoDiffLen > ans {
		ans = twoDiffLen
	}

	if id2 == -1 {
		ans = 0
	}

	return ans
}

// https://coderun.yandex.ru/selections/2025-summer-common/problems/sheep-wolf
// SheepWolf - problem 11
func SheepWolf() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line := mustReadIntArray(reader, 1)
	n := line[0]
	if !validateSheepWolfInput(n) {
		panic("number N out of range")
	}

	// ids input
	ids := make([]int, n)
	line = mustReadIntArray(reader, n)
	for i := range ids {
		ids[i] = line[i]
	}

	answer := solveSheepWolf(n, ids)

	writer.WriteString(strconv.Itoa(answer))
	writer.WriteByte('\n')
}
