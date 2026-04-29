package summerbackend2024

import (
	"bufio"
	"os"
	"strconv"
)

func validateComboNInput(n int) bool {
	return n >= 4 && n <= 20
}

func validateComboXInput(n int) bool {
	return n >= 1 && n <= 1_000
}

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/combo
// Combo - problem 27
func Combo() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line := mustReadIntArray(reader, 1)
	if !validateComboNInput(line[0]) {
		panic("number N out of range")
	}
	n := line[0]

	// prices input
	prices := mustReadIntArray(reader, n)

	// X input
	line = mustReadIntArray(reader, 1)
	if !validateComboXInput(line[0]) {
		panic("number X out of range")
	}
	x := line[0]

	// items input
	b := mustReadIntArray(reader, 4)

	// K input
	line = mustReadIntArray(reader, 1)
	k := line[0]

	// wanted items input
	rawNeed := mustReadIntArray(reader, k)
	need := make([]int, n)
	for _, item := range rawNeed {
		need[item-1]++
	}

	ans := int(1e9)
	for t := 0; t <= k; t++ {
		cost := t * x

		have := make([]int, n)
		for _, idx := range b {
			have[idx-1] += t
		}

		for i := 0; i < n; i++ {
			if need[i] > have[i] {
				cost += (need[i] - have[i]) * prices[i]
			}
		}

		if cost < ans {
			ans = cost
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
