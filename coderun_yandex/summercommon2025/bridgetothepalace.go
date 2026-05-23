package summercommon2025

import (
	"sort"
)

// https://coderun.yandex.ru/selections/2025-summer-common/problems/bridge-to-the-palace
// BridgeToThePalace - problem 2
func SolveBridgeToThePalace(n int, a []int) int {
	sort.Ints(a)

	l := 0
	maxKeep := 0
	for r := 0; r < n; r++ {
		for a[r]-a[l] > n-1 {
			l++
		}

		if r-l+1 > maxKeep {
			maxKeep = r - l + 1
		}
	}

	return n - maxKeep
}
