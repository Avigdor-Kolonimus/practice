package ababin

import (
	"bufio"
	"os"
	"strconv"
)

const INF int64 = 4_000_000_000_000_000_000

type Pair struct {
	mn int64
	mx int64
}

func addSat(a, b int64) int64 {
	if a >= INF || b >= INF || b > INF-a {
		return INF
	}
	return a + b
}

func mulPair(a, b Pair) Pair {
	if a.mx < 0 || b.mx < 0 {
		return Pair{0, -1}
	}
	return Pair{
		addSat(a.mn, b.mn),
		addSat(a.mx, b.mx),
	}
}

func addPair(a, b Pair) Pair {
	if b.mx < 0 {
		return a
	}
	if a.mx < 0 {
		return b
	}

	mn := a.mn
	if b.mn < mn {
		mn = b.mn
	}

	mx := a.mx
	if b.mx > mx {
		mx = b.mx
	}

	return Pair{mn, mx}
}

func mul3(a, b, c Pair) Pair {
	return mulPair(mulPair(a, b), c)
}

func mul4(a, b, c, d Pair) Pair {
	return mulPair(
		mulPair(a, b),
		mulPair(c, d),
	)
}

func adjacent(a, b uint8) bool {
	x := a ^ b
	return x == 1 || x == 2
}

func (fs *FastScanner) skipSpaces() {
	for fs.pos < len(fs.data) && fs.data[fs.pos] <= ' ' {
		fs.pos++
	}
}

func (fs *FastScanner) nextInt() int {
	fs.skipSpaces()

	x := 0

	for fs.pos < len(fs.data) {
		c := fs.data[fs.pos]

		if c < '0' || c > '9' {
			break
		}

		x = x*10 + int(c-'0')
		fs.pos++
	}

	return x
}

func (fs *FastScanner) nextMask(n int) int64 {
	fs.skipSpaces()

	var mask int64

	for i := 0; i < n; i++ {
		c := fs.data[fs.pos]
		fs.pos++

		var d int64

		switch c {
		case 'a':
			d = 0
		case 'b':
			d = 1
		case 'c':
			d = 3
		default:
			d = 2
		}

		mask |= d << (2 * i)
	}

	return mask
}

type Solver struct {
	n int

	baseDig []Pair
	baseHor []Pair
	baseVer []Pair

	goTo []int32
	cnt  []int32

	dig []Pair
	hor []Pair
	ver []Pair
}

func maxTrieVertices(n, q int) int {
	total := 1
	level := 1

	for i := 1; i <= n; i++ {
		level *= 4

		if level > q {
			level = q
		}

		total += level
	}

	return total
}

func NewSolver(n, q int) *Solver {
	s := &Solver{
		n:       n,
		baseDig: make([]Pair, n+1),
		baseHor: make([]Pair, n+1),
		baseVer: make([]Pair, n+1),
	}

	s.baseDig[0] = Pair{1, 1}
	s.baseHor[0] = Pair{1, 1}
	s.baseVer[0] = Pair{1, 1}

	for i := 1; i <= n; i++ {
		dp := s.baseDig[i-1]
		hp := s.baseHor[i-1]
		vp := s.baseVer[i-1]

		s.baseDig[i] = mul3(hp, dp, vp)

		a := mulPair(hp, hp)

		b := mulPair(
			mulPair(vp, vp),
			mulPair(dp, dp),
		)

		hv := addPair(a, b)

		s.baseHor[i] = hv
		s.baseVer[i] = hv
	}

	maxVertices := maxTrieVertices(n, q)

	s.goTo = make([]int32, 0, maxVertices*4)
	s.cnt = make([]int32, 0, maxVertices)

	s.dig = make([]Pair, 0, maxVertices)
	s.hor = make([]Pair, 0, maxVertices)
	s.ver = make([]Pair, 0, maxVertices)

	s.newVertex(n)

	return s
}

func (s *Solver) newVertex(h int) int {
	id := len(s.cnt)

	s.goTo = append(s.goTo, -1, -1, -1, -1)
	s.cnt = append(s.cnt, 0)

	s.dig = append(s.dig, s.baseDig[h])
	s.hor = append(s.hor, s.baseHor[h])
	s.ver = append(s.ver, s.baseVer[h])

	return id
}

func (s *Solver) getState(v int32, h int) (Pair, Pair, Pair, int32) {
	if v < 0 {
		return s.baseDig[h],
			s.baseVer[h],
			s.baseHor[h],
			0
	}

	i := int(v)

	return s.dig[i],
		s.ver[i],
		s.hor[i],
		s.cnt[i]
}

func (s *Solver) getDig(v int32, h int) Pair {
	if v < 0 {
		return s.baseDig[h]
	}

	return s.dig[int(v)]
}

func (s *Solver) calculate(v, h, corner int) {
	base := v << 2

	c0 := s.goTo[base+corner]
	c1 := s.goTo[base+(corner^1)]
	c2 := s.goTo[base+(corner^2)]
	c3 := s.goTo[base+(corner^3)]

	d0, v0, h0, cnt0 := s.getState(c0, h-1)
	d1, v1, h1, cnt1 := s.getState(c1, h-1)
	d2, v2, h2, cnt2 := s.getState(c2, h-1)
	d3, _, _, cnt3 := s.getState(c3, h-1)

	s.cnt[v] = cnt0 + cnt1 + cnt2 + cnt3

	digV := Pair{0, -1}

	if cnt0 == 0 {
		tmp := mul3(h2, d3, v1)
		digV = addPair(digV, tmp)
	}

	if cnt3 == 0 {
		tmp := mul3(v2, d0, h1)
		digV = addPair(digV, tmp)
	}

	horV := mul4(v0, d2, d3, v1)

	if cnt2+cnt3 == 0 {
		tmp := mulPair(h0, h1)
		horV = addPair(horV, tmp)
	}

	verV := mul4(h0, d1, d3, h2)

	if cnt1+cnt3 == 0 {
		tmp := mulPair(v0, v2)
		verV = addPair(verV, tmp)
	}

	s.dig[v] = digV
	s.hor[v] = horV
	s.ver[v] = verV
}

func (s *Solver) update(mask int64) {
	var path [20]int
	var corners [20]int

	v := 0
	corner := 0
	x := mask

	for depth := 0; depth < s.n; depth++ {
		path[depth] = v
		corners[depth] = corner

		digit := int(x & 3)

		v = int(s.goTo[(v<<2)+digit])

		x >>= 2
		corner = digit
	}

	s.cnt[v] ^= 1

	for depth := s.n - 1; depth >= 0; depth-- {
		s.calculate(
			path[depth],
			s.n-depth,
			corners[depth],
		)
	}
}

func digitAt(mask int64, i int) uint8 {
	return uint8((mask >> (2 * i)) & 3)
}

func (s *Solver) isTwoCycle(m1, m2 int64) bool {
	h := s.n

	aLast := digitAt(m1, h-1)
	bLast := digitAt(m2, h-1)

	if !adjacent(aLast, bLast) {
		return false
	}

	samePrefix := true

	for i := 0; i < h-1; i++ {
		if digitAt(m1, i) != digitAt(m2, i) {
			samePrefix = false
			break
		}
	}

	if samePrefix {
		return true
	}

	aSuf, bSuf := 0, 0

	for aSuf < h &&
		digitAt(m1, h-1-aSuf) == aLast {
		aSuf++
	}

	for bSuf < h &&
		digitAt(m2, h-1-bSuf) == bLast {
		bSuf++
	}

	if aSuf == bSuf && aSuf < h {
		cut := h - aSuf - 1

		samePref2 := true

		for i := 0; i < cut; i++ {
			if digitAt(m1, i) != digitAt(m2, i) {
				samePref2 = false
				break
			}
		}

		if samePref2 {
			pos := h - aSuf - 1

			if digitAt(m1, h-1) == digitAt(m2, pos) &&
				digitAt(m2, h-1) == digitAt(m1, pos) {
				return true
			}
		}
	}

	return false
}

func (s *Solver) findTwoWords() (int64, int64) {
	var nodes [64]int32
	var masks [64]int64
	var depths [64]int

	top := 1
	nodes[0] = 0

	var m1, m2 int64
	found := 0

	for top > 0 {
		top--

		v := nodes[top]
		mask := masks[top]
		depth := depths[top]

		if depth == s.n {
			if found == 0 {
				m1 = mask
			} else {
				m2 = mask
			}

			found++

			if found == 2 {
				return m1, m2
			}

			continue
		}

		base := int(v) << 2

		for digit := 0; digit < 4; digit++ {
			child := s.goTo[base+digit]

			if child < 0 || s.cnt[int(child)] == 0 {
				continue
			}

			nodes[top] = child
			masks[top] = mask | (int64(digit) << (2 * depth))
			depths[top] = depth + 1

			top++
		}
	}

	return m1, m2
}

func (s *Solver) answer() Pair {
	h := s.n
	v := 0

	if s.cnt[0] == 0 {
		mx := s.baseDig[h-1].mx

		if mx > INF/4 {
			mx = INF
		} else {
			mx *= 4
		}

		return Pair{2, mx}
	}

	res := Pair{0, -1}

	cycleLen2 := s.cnt[0] <= 1

	if s.cnt[0] == 2 {
		m1, m2 := s.findTwoWords()

		if s.isTwoCycle(m1, m2) {
			cycleLen2 = true
		}
	}

	if cycleLen2 {
		res = Pair{2, 2}
	}

	for h != 0 {
		base := v << 2

		v0 := s.goTo[base]
		v1 := s.goTo[base+1]
		v2 := s.goTo[base+2]
		v3 := s.goTo[base+3]

		d0 := s.getDig(v0, h-1)
		d1 := s.getDig(v1, h-1)
		d2 := s.getDig(v2, h-1)
		d3 := s.getDig(v3, h-1)

		if d0.mx > 0 &&
			d1.mx > 0 &&
			d2.mx > 0 &&
			d3.mx > 0 {

			tmp := mul4(d0, d1, d2, d3)
			res = addPair(res, tmp)
		}

		h--

		curCnt := s.cnt[v]

		if v0 >= 0 && s.cnt[int(v0)] == curCnt {
			v = int(v0)
			continue
		}

		if v1 >= 0 && s.cnt[int(v1)] == curCnt {
			v = int(v1)
			continue
		}

		if v2 >= 0 && s.cnt[int(v2)] == curCnt {
			v = int(v2)
			continue
		}

		if v3 >= 0 && s.cnt[int(v3)] == curCnt {
			v = int(v3)
			continue
		}

		break
	}

	return res
}

// https://coderun.yandex.ru/selections/ababin/problems/easy-production
// EasyProduction - problem 8
func EasyProduction() {
	fs := NewFastScanner()

	n := fs.nextInt()
	q := fs.nextInt()

	masks := make([]int64, q)
	for i := range q {
		masks[i] = fs.nextMask(n)
	}

	s := NewSolver(n, q)

	for _, mask := range masks {
		v := 0
		x := mask

		for j := range n {
			digit := int(x & 3)
			idx := (v << 2) + digit

			next := s.goTo[idx]

			if next < 0 {
				nv := s.newVertex(n - j - 1)

				s.goTo[idx] = int32(nv)
				v = nv
			} else {
				v = int(next)
			}

			x >>= 2
		}
	}

	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	out := make([]byte, 0, q*24)
	for _, mask := range masks {
		s.update(mask)

		ans := s.answer()

		if ans.mn > ans.mx {
			out = append(out, '-', '1', '\n')
		} else {
			out = strconv.AppendInt(out, ans.mn, 10)
			out = append(out, ' ')
			out = strconv.AppendInt(out, ans.mx, 10)
			out = append(out, '\n')
		}
	}

	_, _ = writer.Write(out)
}
