package summercommon2026

import (
	"bufio"
	"os"
)

type Matrix struct {
	a00, a01 int
	a10, a11 int
}

func mul(x, y Matrix, mod int) Matrix {
	return Matrix{
		a00: (x.a00*y.a00 + x.a01*y.a10) % mod,
		a01: (x.a00*y.a01 + x.a01*y.a11) % mod,
		a10: (x.a10*y.a00 + x.a11*y.a10) % mod,
		a11: (x.a10*y.a01 + x.a11*y.a11) % mod,
	}
}

// https://coderun.yandex.ru/seasons/2026-summer/tracks/common/problem/the-solar-sequence
// TheSolarSequence - problem 4
func TheSolarSequence() {
	sc := NewFastScanner()
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	a := sc.NextInt()
	b := sc.NextInt()
	mod := sc.NextInt()
	q := int(sc.NextInt())

	pow := make([]Matrix, 60)
	pow[0] = Matrix{
		a00: a % mod,
		a01: b % mod,
		a10: 1,
		a11: 0,
	}

	for i := 1; i < 60; i++ {
		pow[i] = mul(pow[i-1], pow[i-1], mod)
	}

	for ; q > 0; q-- {
		x := sc.NextInt() % mod
		y := sc.NextInt() % mod
		n := sc.NextInt()

		if n == 0 {
			writeInt(out, x)
			continue
		}
		if n == 1 {
			writeInt(out, y)
			continue
		}

		res := Matrix{
			a00: 1,
			a11: 1,
		}

		e := n - 1
		bit := 0

		for e > 0 {
			if e&1 == 1 {
				res = mul(res, pow[bit], mod)
			}
			e >>= 1
			bit++
		}

		ans := (res.a00*y + res.a01*x) % mod
		writeInt(out, ans)
	}
}
