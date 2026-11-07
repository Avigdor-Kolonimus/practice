package summercommon2025

import (
	"bufio"
	"os"
	"strconv"
)

func calculateAnswerDangerCore(n int, a []int) []int {
	// Calculate MEX of the entire array.
	used := make([]bool, n+1)

	mx := 0
	for _, v := range a {
		if v <= n {
			used[v] = true
		}
		if v > mx {
			mx = v
		}
	}

	mex := 0
	for mex <= n && used[mex] {
		mex++
	}

	seen := make([]int, n+1)
	stamp := 0

	check := func(x int) bool {
		l, r := n, -1

		for i, v := range a {
			if v > x {
				if i < l {
					l = i
				}
				r = i
			}
		}

		if r == -1 {
			return x >= 1
		}

		stamp++

		for i := l; i <= r; i++ {
			v := a[i]
			if v <= x {
				seen[v] = stamp
			}
		}

		for v := 0; v <= x; v++ {
			if seen[v] != stamp {
				return true
			}
		}

		return false
	}

	lo, hi := -1, n

	for hi-lo > 1 {
		mid := lo + (hi-lo)/2

		if check(mid) {
			hi = mid
		} else {
			lo = mid
		}
	}

	minMax := hi

	maxMax := mx

	if n == 1 {
		maxMax = mex
	} else if mex > maxMax {
		maxMax = mex
	}

	minMin := 0
	if n == 1 {
		minMin = mex
	}

	maxMin := mex

	return []int{
		minMax,
		maxMax,
		minMin,
		maxMin,
	}
}

// ввод/вывод
// не изменяйте сигнатуру метода
// https://coderun.yandex.ru/selections/2025-summer-common/problems/danger-core
// DangerCore - problem 23
func DangerCoreSolve() {
	input := NewFastScanner(os.Stdin)

	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	t := input.readInt()
	for test := 0; test < t; test++ {
		n := input.readInt()
		a := input.readIntArray(n)

		answer := calculateAnswerDangerCore(n, a)
		for i := range len(answer) {
			output.WriteString(strconv.Itoa(answer[i]))
			output.WriteByte(' ')
		}

		output.WriteString("\n")
	}
}
