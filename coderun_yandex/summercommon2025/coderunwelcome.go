package summercommon2025

import (
	"math"
)

// https://coderun.yandex.ru/selections/2025-summer-common/problems/coderun-welcome
// CoderunWelcome - problem 1
func SolveCoderunWelcome(n, m int) int {
	a := min(n, m)
	s := n + m

	bound := min(s, 2*a+1)

	return int(math.Floor(math.Sqrt(float64(bound))))
}
