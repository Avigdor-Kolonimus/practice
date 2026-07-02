package summercommon2026

import (
	"bufio"
	"os"
)

type Pair struct {
	k int
	v int
}

type Pairs []Pair

func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Less(i, j int) bool { return p[i].k < p[j].k }
func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func radixSortPairs(a []Pair) []Pair {
	n := len(a)
	if n <= 1 {
		return a
	}

	tmp := make([]Pair, n)

	const bits = 8
	const mask = 255

	src := a
	dst := tmp

	for shift := 0; shift < 32; shift += bits {

		var cnt [256]int

		for i := 0; i < n; i++ {
			cnt[(uint32(src[i].k)>>shift)&mask]++
		}

		sum := 0
		for i := 0; i < 256; i++ {
			c := cnt[i]
			cnt[i] = sum
			sum += c
		}

		for i := 0; i < n; i++ {
			v := src[i]
			b := (uint32(v.k) >> shift) & mask
			dst[cnt[b]] = v
			cnt[b]++
		}

		src, dst = dst, src
	}

	return src
}

// https://coderun.yandex.ru/seasons/2026-summer/tracks/common/problem/summer-bike-tour
// SummerBikeTour - problem 2
func SummerBikeTour() {
	in := NewFastScanner()
	out := bufio.NewWriterSize(os.Stdout, 4<<20)
	defer out.Flush()

	m := in.NextInt()

	last := make(map[int]int, 6_000_000)
	for i := 0; i < m; i++ {
		l := in.NextInt()

		for j := 0; j < l; j++ {
			k := in.NextInt()
			v := in.NextInt()
			last[k] = v
		}
	}

	pairs := make([]Pair, 0, len(last))
	for k, v := range last {
		pairs = append(pairs, Pair{k, v})
	}

	pairs = radixSortPairs(pairs)

	for _, p := range pairs {
		writeInt(out, p.k)
		out.WriteByte(' ')
		writeInt(out, p.v)
		out.WriteByte('\n')
	}
}
