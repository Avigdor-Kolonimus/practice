package summercommon2025

// https://coderun.yandex.ru/selections/2025-summer-common/problems/binary-forest
// BinaryForest - problem 13
func SolveBinaryForest(n int, a []int, m int, b []int) int {
	build := func(arr []int) []int {
		totalOnes := 0
		for _, x := range arr {
			totalOnes += x
		}

		zeros := 0
		ones := totalOnes

		// dp[k] = maximum number of ones
		// that can be taken after choosing k zeros
		dp := make([]int, 0, len(arr)+1)

		dp = append(dp, totalOnes)

		for _, x := range arr {
			if x == 0 {
				zeros++
				dp = append(dp, ones)
			} else {
				ones--
			}
		}

		return dp
	}

	dpA := build(a)
	dpB := build(b)

	limit := min(len(dpA), len(dpB))

	ans := 0

	for zeros := 0; zeros < limit; zeros++ {
		cur := zeros + min(dpA[zeros], dpB[zeros])

		if cur > ans {
			ans = cur
		}
	}

	return ans
}
