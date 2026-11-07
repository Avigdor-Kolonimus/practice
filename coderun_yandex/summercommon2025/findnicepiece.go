package summercommon2025

import (
	"bufio"
	"math/bits"
	"os"
	"strconv"
)

func calculateAnswerFindNicePiece(n int, s string) int {
	mod := int64(998244353)

	powMod := func(a, b int64) int64 {
		res := int64(1)

		for b > 0 {
			if b&1 == 1 {
				res = res * a % mod
			}
			a = a * a % mod
			b >>= 1
		}

		return res
	}

	k := n * n / 4

	inv2 := powMod(2, mod-2)

	invPow := make([]int64, n/2+1)
	invPow[0] = 1

	for i := 1; i <= n/2; i++ {
		invPow[i] = invPow[i-1] * inv2 % mod
	}

	type Interval struct {
		l      int
		length int
		inv    int64
	}

	intervals := make([]Interval, 0)

	for length := 2; length <= n; length += 2 {
		for l := 0; l+length <= n; l++ {
			intervals = append(intervals, Interval{
				l:      l,
				length: length,
				inv:    invPow[length/2],
			})
		}
	}

	pal := make([][]bool, n+1)
	for length := 2; length <= n; length += 2 {
		size := 1 << length
		pal[length] = make([]bool, size)

		for mask := 0; mask < size; mask++ {
			rev := 0

			for j := 0; j < length; j++ {
				rev = (rev << 1) | ((mask >> j) & 1)
			}

			pal[length][mask] = mask == rev
		}
	}

	ones := 0
	for i := range n {
		if s[i] == '1' {
			ones |= 1 << i
		}
	}

	size := 1 << n
	f := make([]int64, size)

	for mask := 0; mask < size; mask++ {
		var diag int64
		var rhs int64

		for _, it := range intervals {
			length := it.length
			l := it.l

			pattern := (mask >> l) & ((1 << length) - 1)

			if pattern == 0 {
				diag += it.inv
				if diag >= mod {
					diag -= mod
				}
				continue
			}

			if !pal[length][pattern] {
				continue
			}

			cnt := bits.OnesCount(uint(pattern)) / 2

			coeff := it.inv

			if cnt&1 == 1 {
				coeff = mod - coeff
			}

			prev := mask ^ (pattern << l)

			rhs = (rhs + coeff*f[prev]) % mod
		}

		if mask == 0 {
			rhs = (rhs + int64(k)) % mod
		}

		denom := (int64(k) - diag + mod) % mod

		f[mask] = rhs * powMod(denom, mod-2) % mod
	}

	var answer int64
	for mask := 0; mask < size; mask++ {
		term := f[mask]

		if bits.OnesCount(uint(mask&ones))&1 == 1 {
			term = (mod - term) % mod
		}

		answer = (answer + term) % mod
	}

	return int(answer)
}

// ввод/вывод
// не изменяйте сигнатуру метода
// https://coderun.yandex.ru/selections/2025-summer-common/problems/find-nice-piece
// FindNicePiece - problem 24
func FindNicePieceSolve() {
	input := NewFastScanner(os.Stdin)

	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	s := input.readToken()
	n := len(s)

	answer := calculateAnswerFindNicePiece(n, s)
	output.WriteString(strconv.Itoa(answer))
	output.WriteByte('\n')
}
