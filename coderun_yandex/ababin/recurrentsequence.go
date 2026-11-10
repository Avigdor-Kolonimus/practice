package ababin

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

type FastScanner struct {
	data []byte
	pos  int
}

func NewFastScanner() *FastScanner {
	data, _ := io.ReadAll(os.Stdin)
	return &FastScanner{data: data}
}

func (fs *FastScanner) NextInt() int {
	for fs.pos < len(fs.data) && fs.data[fs.pos] <= ' ' {
		fs.pos++
	}

	sign := 1
	if fs.pos < len(fs.data) && fs.data[fs.pos] == '-' {
		sign = -1
		fs.pos++
	}

	n := 0
	for fs.pos < len(fs.data) {
		c := fs.data[fs.pos]
		if c < '0' || c > '9' {
			break
		}
		n = n*10 + int(c-'0')
		fs.pos++
	}

	return n * sign
}

func modNorm(x int64, p int) int {
	pp := int64(p)
	x %= pp
	if x < 0 {
		x += pp
	}
	return int(x)
}

func modPow(a int, e int64, p int) int {
	mod := int64(p)
	base := int64(a % p)
	res := int64(1)

	for e > 0 {
		if e&1 == 1 {
			res = res * base % mod
		}
		base = base * base % mod
		e >>= 1
	}
	return int(res)
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int64) int64 {
	return a / gcd(a, b) * b
}

func egcd(a, b int64) (int64, int64, int64) {
	if b == 0 {
		return a, 1, 0
	}
	g, x1, y1 := egcd(b, a%b)
	return g, y1, x1 - (a/b)*y1
}

func invMod(a, m int64) (int64, bool) {
	if m == 1 {
		return 0, true
	}

	g, x, _ := egcd(a, m)
	if g != 1 {
		return 0, false
	}

	x %= m
	if x < 0 {
		x += m
	}
	return x, true
}

type CRT2 struct {
	m1, m2 int64
	g      int64
	l      int64
	m2g    int64
	inv    int64
}

func NewCRT2(m1, m2 int64) CRT2 {
	g := gcd(m1, m2)
	m1g := m1 / g
	m2g := m2 / g

	inv, _ := invMod(m1g%m2g, m2g)

	return CRT2{
		m1:  m1,
		m2:  m2,
		g:   g,
		l:   m1g * m2,
		m2g: m2g,
		inv: inv,
	}
}

func (c CRT2) Solve(a1, a2 int64) (int64, bool) {
	diff := a2 - a1
	if diff%c.g != 0 {
		return 0, false
	}

	t := (diff / c.g) % c.m2g
	if t < 0 {
		t += c.m2g
	}

	t = t * c.inv % c.m2g

	x := (a1 + c.m1*t) % c.l
	if x < 0 {
		x += c.l
	}

	return x, true
}

func buildSPFRecurrentSequence(n int) []int {
	spf := make([]int, n+1)

	for i := 2; i <= n; i++ {
		if spf[i] != 0 {
			continue
		}

		spf[i] = i
		if i*i <= n {
			for j := i * i; j <= n; j += i {
				if spf[j] == 0 {
					spf[j] = i
				}
			}
		}
	}

	return spf
}

type Factor struct {
	p, cnt int
}

func factorize(x int, spf []int) []Factor {
	res := make([]Factor, 0)

	for x > 1 {
		p := spf[x]
		cnt := 0

		for x%p == 0 {
			x /= p
			cnt++
		}

		res = append(res, Factor{p, cnt})
	}

	return res
}

func orderFp(g, p int, factors []Factor) int {
	if g%p == 0 {
		return 0
	}

	order := p - 1

	for _, f := range factors {
		for order%f.p == 0 {
			if modPow(g, int64(order/f.p), p) == 1 {
				order /= f.p
			} else {
				break
			}
		}
	}

	return order
}

func tonelliShanks(n, p int) (int, bool) {
	if n == 0 {
		return 0, true
	}

	if p == 2 {
		return n % p, true
	}

	if modPow(n, int64((p-1)/2), p) != 1 {
		return 0, false
	}

	if p%4 == 3 {
		return modPow(n, int64((p+1)/4), p), true
	}

	q := p - 1
	s := 0

	for q&1 == 0 {
		q >>= 1
		s++
	}

	z := 2
	for modPow(z, int64((p-1)/2), p) != p-1 {
		z++
	}

	c := modPow(z, int64(q), p)
	x := modPow(n, int64((q+1)/2), p)
	t := modPow(n, int64(q), p)
	m := s

	for t != 1 {
		i := 1
		tt := int64(t) * int64(t) % int64(p)

		for i < m && tt != 1 {
			tt = tt * tt % int64(p)
			i++
		}

		if i == m {
			return 0, false
		}

		b := modPow(c, int64(1<<(m-i-1)), p)

		x = int(int64(x) * int64(b) % int64(p))
		t = int(int64(t) * int64(b) % int64(p) * int64(b) % int64(p))
		c = int(int64(b) * int64(b) % int64(p))
		m = i
	}

	return x, true
}

type F2 struct {
	u, v int
}

type FieldContext struct {
	p, a, b int
}

func (c FieldContext) Mul(x, y F2) F2 {
	p := int64(c.p)
	a := int64(c.a)
	b := int64(c.b)

	xu, xv := int64(x.u), int64(x.v)
	yu, yv := int64(y.u), int64(y.v)

	vv := xv * yv % p

	u := (xu*yu + vv*b) % p
	v := (xu*yv + xv*yu + vv*a) % p

	return F2{int(u), int(v)}
}

func (c FieldContext) Neg(x F2) F2 {
	u, v := 0, 0

	if x.u != 0 {
		u = c.p - x.u
	}
	if x.v != 0 {
		v = c.p - x.v
	}

	return F2{u, v}
}

func (c FieldContext) Scale(x F2, k int) F2 {
	p := int64(c.p)
	kk := int64(k)

	return F2{
		int(int64(x.u) * kk % p),
		int(int64(x.v) * kk % p),
	}
}

func (c FieldContext) Conj(x F2) F2 {
	p := int64(c.p)

	u := (int64(x.u) + int64(x.v)*int64(c.a)) % p
	v := 0

	if x.v != 0 {
		v = c.p - x.v
	}

	return F2{int(u), v}
}

func (c FieldContext) Norm(x F2) int {
	p := int64(c.p)
	u, v := int64(x.u), int64(x.v)
	a, b := int64(c.a), int64(c.b)

	val := u * u % p
	val = (val + u*v%p*a) % p
	val = (val - v*v%p*b) % p

	if val < 0 {
		val += p
	}

	return int(val)
}

func (c FieldContext) Pow(base F2, e int64) F2 {
	res := F2{1, 0}

	for e > 0 {
		if e&1 == 1 {
			res = c.Mul(res, base)
		}
		base = c.Mul(base, base)
		e >>= 1
	}

	return res
}

type UInt64IntMap struct {
	keys []uint64
	vals []int32
	mask uint64
}

func NewUInt64IntMap(capacityNeeded int) *UInt64IntMap {
	capacity := 1
	for capacity < capacityNeeded {
		capacity <<= 1
	}

	return &UInt64IntMap{
		keys: make([]uint64, capacity),
		vals: make([]int32, capacity),
		mask: uint64(capacity - 1),
	}
}

func (m *UInt64IntMap) Put(k uint64, v int32) {
	i := k & m.mask

	for m.keys[i] != 0 && m.keys[i] != k {
		i = (i + 1) & m.mask
	}

	m.keys[i] = k
	m.vals[i] = v
}

func (m *UInt64IntMap) Get(k uint64) int32 {
	i := k & m.mask

	for {
		kk := m.keys[i]

		if kk == 0 {
			return -1
		}
		if kk == k {
			return m.vals[i]
		}

		i = (i + 1) & m.mask
	}
}

func packKey(u, v int) uint64 {
	return (uint64(u) << 20) | uint64(v)
}

// https://coderun.yandex.ru/selections/ababin/problems/recurrent-sequence
// RecurrentSequence - problem 6
func RecurrentSequence() {
	in := NewFastScanner()

	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	out := make([]byte, 0, 1<<20)

	writeAnswer := func(x int64) {
		out = strconv.AppendInt(out, x, 10)
		out = append(out, '\n')

		if len(out) >= 1<<20 {
			writer.Write(out)
			out = out[:0]
		}
	}

	p := in.NextInt()
	a := in.NextInt()
	b := in.NextInt()
	q := in.NextInt()

	if p == 2 {
		for i := 0; i < q; i++ {
			x := in.NextInt()
			y := in.NextInt()

			switch {
			case x == 0 && y == 1:
				writeAnswer(3)
			case x == 1 && y == 1:
				writeAnswer(1)
			case x == 1 && y == 0:
				writeAnswer(2)
			default:
				writeAnswer(-1)
			}
		}

		writer.Write(out)
		return
	}

	spf := buildSPFRecurrentSequence(p + 1)
	facPm1 := factorize(p-1, spf)
	facPp1 := factorize(p+1, spf)

	inv := make([]int, p)
	inv[1] = 1
	for i := 2; i < p; i++ {
		pi := p / i
		ri := p % i

		val := (int64(p) - int64(pi)*int64(inv[ri])%int64(p)) % int64(p)
		inv[i] = int(val)
	}

	D := modNorm(int64(a)*int64(a)+4*int64(b), p)

	if D == 0 {
		r := int(int64(a) * int64(inv[2]) % int64(p))
		ordR := orderFp(r, p, facPm1)

		logR := make([]int32, p)
		for i := range logR {
			logR[i] = -1
		}

		cur := 1
		for e := 0; e < ordR; e++ {
			logR[cur] = int32(e)
			cur = int(int64(cur) * int64(r) % int64(p))
		}

		period := int64(ordR) * int64(p)
		invOrdModP := inv[ordR%p]

		for i := 0; i < q; i++ {
			x := in.NextInt()
			y := in.NextInt()

			// z = y - x*r = r^i.
			z := modNorm(int64(y)-int64(x)*int64(r), p)

			if z == 0 {
				writeAnswer(-1)
				continue
			}

			k := int(logR[z])
			if k < 0 {
				writeAnswer(-1)
				continue
			}

			t := int(int64(x) * int64(r) % int64(p) *
				int64(inv[z]) % int64(p))

			delta := modNorm(int64(t)-int64(k), p)
			mult := int(int64(delta) * int64(invOrdModP) % int64(p))

			ans := int64(k) + int64(ordR)*int64(mult)

			if ans == 0 {
				ans = period
			}

			writeAnswer(ans)
		}

		writer.Write(out)
		return
	}

	leg := modPow(D, int64((p-1)/2), p)

	if leg == 1 {
		sqrtD, ok := tonelliShanks(D, p)
		if !ok {
			for i := 0; i < q; i++ {
				_ = in.NextInt()
				_ = in.NextInt()
				writeAnswer(-1)
			}

			writer.Write(out)
			return
		}

		inv2 := inv[2]

		lambda := int(
			int64(modNorm(int64(a)+int64(sqrtD), p)) *
				int64(inv2) % int64(p),
		)

		mu := int(
			int64(modNorm(int64(a)-int64(sqrtD), p)) *
				int64(inv2) % int64(p),
		)

		ordL := orderFp(lambda, p, facPm1)
		ordM := orderFp(mu, p, facPm1)

		period := lcm(int64(ordL), int64(ordM))

		logL := make([]int32, p)
		logM := make([]int32, p)

		for i := 0; i < p; i++ {
			logL[i] = -1
			logM[i] = -1
		}

		cur := 1
		for e := 0; e < ordL; e++ {
			logL[cur] = int32(e)
			cur = int(int64(cur) * int64(lambda) % int64(p))
		}

		cur = 1
		for e := 0; e < ordM; e++ {
			logM[cur] = int32(e)
			cur = int(int64(cur) * int64(mu) % int64(p))
		}

		crt := NewCRT2(int64(ordL), int64(ordM))
		for i := 0; i < q; i++ {
			x := in.NextInt()
			y := in.NextInt()

			v1 := modNorm(int64(y)-int64(mu)*int64(x), p)
			v2 := modNorm(int64(y)-int64(lambda)*int64(x), p)

			if v1 == 0 || v2 == 0 {
				writeAnswer(-1)
				continue
			}

			e1 := int64(logL[v1])
			e2 := int64(logM[v2])

			if e1 < 0 || e2 < 0 {
				writeAnswer(-1)
				continue
			}

			ans, ok := crt.Solve(e1, e2)
			if !ok {
				writeAnswer(-1)
				continue
			}

			if ans == 0 {
				ans = period
			}

			writeAnswer(ans)
		}

		writer.Write(out)
		return
	}

	ctx := FieldContext{p: p, a: a, b: b}
	r := F2{0, 1}

	primeSet := make(map[int]bool)
	for _, f := range facPm1 {
		primeSet[f.p] = true
	}
	for _, f := range facPp1 {
		primeSet[f.p] = true
	}

	ordR := int64(p-1) * int64(p+1)
	for pr := range primeSet {
		for ordR%int64(pr) == 0 {
			candidate := ordR / int64(pr)

			if ctx.Pow(r, candidate) == (F2{1, 0}) {
				ordR = candidate
			} else {
				break
			}
		}
	}

	pow2 := make([]F2, 64)
	pow2[0] = r
	for i := 1; i < 64; i++ {
		pow2[i] = ctx.Mul(pow2[i-1], pow2[i-1])
	}

	powR := func(e int64) F2 {
		res := F2{1, 0}
		bit := 0

		for e > 0 {
			if e&1 == 1 {
				res = ctx.Mul(res, pow2[bit])
			}
			e >>= 1
			bit++
		}

		return res
	}

	hasNeg := false
	var halfOrd int64

	if ordR&1 == 0 {
		halfOrd = ordR / 2
		mid := powR(halfOrd)

		if mid == (F2{p - 1, 0}) {
			hasNeg = true
		}
	}

	baseNorm := (p - b%p) % p
	ord1 := orderFp(baseNorm, p, facPm1)
	logNorm := make([]int32, p)
	for i := range logNorm {
		logNorm[i] = -1
	}

	curNorm := 1
	for e := 0; e < ord1; e++ {
		logNorm[curNorm] = int32(e)
		curNorm = int(int64(curNorm) * int64(baseNorm) % int64(p))
	}

	baseUnit := ctx.Pow(r, int64(p-1))
	ord2 := p + 1
	for _, f := range facPp1 {
		for ord2%f.p == 0 {
			candidate := ord2 / f.p

			if ctx.Pow(baseUnit, int64(candidate)) == (F2{1, 0}) {
				ord2 = candidate
			} else {
				break
			}
		}
	}

	mapUnit := NewUInt64IntMap(ord2*2 + 8)

	curU := F2{1, 0}
	for e := 0; e < ord2; e++ {
		mapUnit.Put(packKey(curU.u, curU.v), int32(e))
		curU = ctx.Mul(curU, baseUnit)
	}

	crt := NewCRT2(int64(ord1), int64(ord2))
	for i := 0; i < q; i++ {
		x := in.NextInt()
		y := in.NextInt()

		h := F2{
			u: modNorm(int64(y)-int64(a)*int64(x), p),
			v: modNorm(int64(x), p),
		}

		if h.u == 0 && h.v == 0 {
			writeAnswer(-1)
			continue
		}

		nrm := ctx.Norm(h)
		if nrm == 0 {
			writeAnswer(-1)
			continue
		}

		e1 := int64(logNorm[nrm])
		if e1 < 0 {
			writeAnswer(-1)
			continue
		}

		ch := ctx.Conj(h)
		ch2 := ctx.Mul(ch, ch)
		uElem := ctx.Scale(ch2, inv[nrm])

		e2 := int64(mapUnit.Get(packKey(uElem.u, uElem.v)))
		if e2 < 0 {
			writeAnswer(-1)
			continue
		}

		k0, ok := crt.Solve(e1, e2)
		if !ok {
			writeAnswer(-1)
			continue
		}

		powk := powR(k0)
		var ans int64 = -1
		if powk == h {
			ans = k0
		} else if hasNeg {
			nh := ctx.Neg(h)

			if powk == nh {
				k := k0 + halfOrd
				if k >= ordR {
					k -= ordR
				}
				ans = k
			}
		}

		if ans < 0 {
			writeAnswer(-1)
		} else {
			if ans == 0 {
				ans = ordR
			}
			writeAnswer(ans)
		}
	}

	writer.Write(out)
}
