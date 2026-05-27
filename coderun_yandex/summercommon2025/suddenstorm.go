package summercommon2025

// https://coderun.yandex.ru/selections/2025-summer-common/problems/sudden-storm
// SuddenStorm - problem 10
func SolveSuddenStorm(n int, t int, a []int64, b []int64) []int64 {
	ans := make([]int64, t+1)

	// Difference arrays:
	// diffConst stores constant additions
	// diffSlope stores coefficient additions for time k
	diffConst := make([]int64, t+2)
	diffSlope := make([]int64, t+2)

	for i := 0; i < n; i++ {
		A := int64(a[i])
		B := int64(b[i])

		// Cloud never disappears
		if B == 0 {
			diffConst[0] += A

			continue
		}

		// Last minute when the cloud still contains water
		last := (A - 1) / B

		if last > int64(t) {
			last = int64(t)
		}

		// Skip empty clouds
		if last < 0 {
			continue
		}

		l := 0
		r := int(last)

		// Cloud contribution:
		// A - k*B
		diffConst[l] += A
		diffConst[r+1] -= A

		diffSlope[l] -= B
		diffSlope[r+1] += B
	}

	var curConst int64
	var curSlope int64

	for k := 0; k <= t; k++ {
		curConst += diffConst[k]
		curSlope += diffSlope[k]

		// Current total amount of water in the sky
		ans[k] = curConst + int64(k)*curSlope
	}

	return ans
}
