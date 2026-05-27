package summercommon2025

import (
	"sort"
)

func computeBinoms(n, m int) []int64 {
	res := []int64{}

	c := int64(1)

	for k := 0; k <= (n+1)/2; k++ {
		res = append(res, c)

		if n-k == 0 {
			break
		}

		next := c * int64(n-k) / int64(k+1)

		if next > int64(m) {
			break
		}

		c = next
	}

	return res
}

// https://coderun.yandex.ru/selections/2025-summer-common/problems/digital-civilization-legacy
// DigitalCivilizationLegacy - problem 18
func SolutionDigitalCivilizationLegacy(n, m int) int64 {
	binoms := computeBinoms(n, m)

	pref := make([]int64, len(binoms))
	copy(pref, binoms)

	for i := 1; i < len(pref); i++ {
		pref[i] += pref[i-1]
	}

	lo := int64(0)
	hi := binoms[len(binoms)-1] + 1

	for hi-lo > 1 {
		mid := (lo + hi) / 2

		idx := sort.Search(len(binoms), func(i int) bool {
			return binoms[i] > mid
		}) - 1

		total := int64(0)

		if idx < 0 {
			total = int64(n+1) * mid
		} else {
			leftSum := pref[idx]

			if n%2 == 0 && idx == n/2 {
				total = 2*leftSum - binoms[idx]
			} else {
				rest := int64((n + 1) - 2*(idx+1))
				total = 2*leftSum + rest*mid
			}
		}

		if total >= int64(m) {
			hi = mid
		} else {
			lo = mid
		}
	}

	return lo + 1
}
