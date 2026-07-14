package summercommon2026

import (
	"bufio"
	"io"
	"math/bits"
	"os"
	"strconv"
)

const (
	INF int64 = 2e18
)

type Sequence struct {
	offset   int
	startVal int64
	diffs    []int64
}

type MatrixSetUpDeckChairs [2][2]*Sequence

type Arena struct {
	sequences []Sequence
	int64s    []int64
	seqIdx    int
	intIdx    int
}

func NewArena(maxSeqs, maxInts int) *Arena {
	return &Arena{
		sequences: make([]Sequence, maxSeqs),
		int64s:    make([]int64, maxInts),
	}
}

func (a *Arena) Reset() {
	a.seqIdx = 0
	a.intIdx = 0
}

func (a *Arena) AllocSeq(offset int, startVal int64, diffs []int64) *Sequence {
	if a.seqIdx >= len(a.sequences) {
		a.sequences = append(a.sequences, Sequence{})
	}
	s := &a.sequences[a.seqIdx]
	a.seqIdx++
	s.offset = offset
	s.startVal = startVal
	s.diffs = diffs
	return s
}

func (a *Arena) AllocInt64s(size int) []int64 {
	if a.intIdx+size > len(a.int64s) {
		needed := a.intIdx + size - len(a.int64s)
		a.int64s = append(a.int64s, make([]int64, needed)...)
	}
	res := a.int64s[a.intIdx : a.intIdx+size]
	a.intIdx += size
	return res
}

var a []int64
var b []int64

func mergeDiffs(arr1, arr2 []int64, arena *Arena) []int64 {
	n1, n2 := len(arr1), len(arr2)
	if n1 == 0 {
		return arr2
	}
	if n2 == 0 {
		return arr1
	}
	var res []int64
	if arena != nil {
		res = arena.AllocInt64s(n1 + n2)
	} else {
		res = make([]int64, n1+n2)
	}

	_ = arr1[n1-1]
	_ = arr2[n2-1]
	_ = res[n1+n2-1]

	i, j := 0, 0
	for i < n1 && j < n2 {
		v1 := arr1[i]
		v2 := arr2[j]
		if v1 <= v2 {
			res[i+j] = v1
			i++
		} else {
			res[i+j] = v2
			j++
		}
	}
	for i < n1 {
		res[i+j] = arr1[i]
		i++
	}
	for j < n2 {
		res[i+j] = arr2[j]
		j++
	}
	return res
}

func convolveAndShift(A, B *Sequence, w int, arena *Arena) *Sequence {
	if A == nil || B == nil {
		return nil
	}
	resStart := A.startVal + B.startVal
	resDiffs := mergeDiffs(A.diffs, B.diffs, arena)
	if arena != nil {
		return arena.AllocSeq(A.offset+B.offset+w, resStart, resDiffs)
	}
	return &Sequence{
		offset:   A.offset + B.offset + w,
		startVal: resStart,
		diffs:    resDiffs,
	}
}

func elementWiseMin(valid []*Sequence, arena *Arena) *Sequence {
	if len(valid) == 0 {
		return nil
	}
	if len(valid) == 1 {
		return valid[0]
	}

	minK := valid[0].offset
	maxK := valid[0].offset + len(valid[0].diffs)

	for i := 1; i < len(valid); i++ {
		opt := valid[i]
		if opt.offset < minK {
			minK = opt.offset
		}
		endK := opt.offset + len(opt.diffs)
		if endK > maxK {
			maxK = endK
		}
	}

	size := maxK - minK + 1
	var newVals []int64
	if arena != nil {
		newVals = arena.AllocInt64s(size)
	} else {
		newVals = make([]int64, size)
	}

	for i := 0; i < size; i++ {
		newVals[i] = INF
	}

	for _, opt := range valid {
		offsetDiff := opt.offset - minK
		curr := opt.startVal

		_ = newVals[offsetDiff+len(opt.diffs)]

		if curr < newVals[offsetDiff] {
			newVals[offsetDiff] = curr
		}
		for i := 0; i < len(opt.diffs); i++ {
			curr += opt.diffs[i]
			if curr < newVals[offsetDiff+i+1] {
				newVals[offsetDiff+i+1] = curr
			}
		}
	}

	realSize := size
	for realSize > 0 && newVals[realSize-1] == INF {
		realSize--
	}
	if realSize == 0 {
		return nil
	}
	newVals = newVals[:realSize]

	var newDiffs []int64
	if realSize > 1 {
		if arena != nil {
			newDiffs = arena.AllocInt64s(realSize - 1)
		} else {
			newDiffs = make([]int64, realSize-1)
		}

		_ = newVals[realSize-1]
		_ = newDiffs[realSize-2]
		for i := 1; i < realSize; i++ {
			newDiffs[i-1] = newVals[i] - newVals[i-1]
		}
	}

	if arena != nil {
		return arena.AllocSeq(minK, newVals[0], newDiffs)
	}
	return &Sequence{
		offset:   minK,
		startVal: newVals[0],
		diffs:    newDiffs,
	}
}

func multiply(X, Y MatrixSetUpDeckChairs, arena *Arena) MatrixSetUpDeckChairs {
	var M MatrixSetUpDeckChairs
	var opts [4]*Sequence
	for u := 0; u < 2; u++ {
		for v := 0; v < 2; v++ {
			optCount := 0
			for mid := 0; mid < 2; mid++ {
				for midPrime := 0; midPrime < 2; midPrime++ {
					w := 0
					if mid == 1 && midPrime == 0 {
						w = 1
					}
					xOpt := X[u][mid]
					yOpt := Y[midPrime][v]
					if xOpt != nil && yOpt != nil {
						opts[optCount] = convolveAndShift(xOpt, yOpt, w, arena)
						optCount++
					}
				}
			}
			M[u][v] = elementWiseMin(opts[:optCount], arena)
		}
	}
	return M
}

// ---- Disjoint sparse table (0-indexed positions 0..n-1) ----

var leaf []MatrixSetUpDeckChairs  // leaf[i], i in [0,n)
var pre [][]MatrixSetUpDeckChairs // pre[bit][i]
var suf [][]MatrixSetUpDeckChairs // suf[bit][i]
var LOG int

func buildSparse(n int) {
	LOG = bits.Len(uint(n)) + 1
	leaf = make([]MatrixSetUpDeckChairs, n)
	for i := 0; i < n; i++ {
		var M MatrixSetUpDeckChairs
		M[0][0] = &Sequence{offset: 0, startVal: a[i], diffs: nil}
		M[1][1] = &Sequence{offset: 0, startVal: b[i], diffs: nil}
		leaf[i] = M
	}

	pre = make([][]MatrixSetUpDeckChairs, LOG)
	suf = make([][]MatrixSetUpDeckChairs, LOG)

	for bit := 0; bit < LOG; bit++ {
		pre[bit] = make([]MatrixSetUpDeckChairs, n)
		suf[bit] = make([]MatrixSetUpDeckChairs, n)
		blockSize := 1 << (bit + 1)
		half := 1 << bit

		for blockStart := 0; blockStart < n; blockStart += blockSize {
			mid := blockStart + half
			if mid >= n {
				break
			}
			end := blockStart + blockSize - 1
			if end >= n {
				end = n - 1
			}

			// prefix: leaf[i] * leaf[i+1] * ... * leaf[mid-1], built right-to-left
			pre[bit][mid-1] = leaf[mid-1]
			for i := mid - 2; i >= blockStart; i-- {
				pre[bit][i] = multiply(leaf[i], pre[bit][i+1], nil)
			}

			// suffix: leaf[mid] * ... * leaf[i], built left-to-right
			suf[bit][mid] = leaf[mid]
			for i := mid + 1; i <= end; i++ {
				suf[bit][i] = multiply(suf[bit][i-1], leaf[i], nil)
			}
		}
	}
}

// lo, hi are 0-indexed, lo <= hi
func querySparse(lo, hi int, arena *Arena) MatrixSetUpDeckChairs {
	if lo == hi {
		return leaf[lo]
	}
	bit := bits.Len(uint(lo^hi)) - 1
	return multiply(pre[bit][lo], suf[bit][hi], arena)
}

// https://coderun.yandex.ru/seasons/2026-summer/tracks/common/problem/set-up-deck-chairs
// SetUpDeckChairs - problem 14
func SetUpDeckChairs() {
	inputBytes, _ := io.ReadAll(os.Stdin)
	pos := 0

	nextInt64 := func() int64 {
		for pos < len(inputBytes) && (inputBytes[pos] < '0' || inputBytes[pos] > '9') {
			pos++
		}
		if pos >= len(inputBytes) {
			return 0
		}
		var res int64 = 0
		for pos < len(inputBytes) && inputBytes[pos] >= '0' && inputBytes[pos] <= '9' {
			res = res*10 + int64(inputBytes[pos]-'0')
			pos++
		}
		return res
	}

	nextInt := func() int {
		return int(nextInt64())
	}

	n := nextInt()
	if n == 0 {
		return
	}
	q := nextInt()

	a = make([]int64, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt64()
	}

	b = make([]int64, n)
	for i := 0; i < n; i++ {
		b[i] = nextInt64()
	}

	buildSparse(n)

	arena := NewArena(20000, 1500000)

	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	for i := 0; i < q; i++ {
		l := nextInt()
		r := nextInt()
		C := nextInt64()

		arena.Reset()
		finalM := querySparse(l-1, r-1, arena)
		ans := -1

		for u := 0; u < 2; u++ {
			for v := 0; v < 2; v++ {
				opt := finalM[u][v]
				if opt != nil {
					curr := opt.startVal
					if curr <= C {
						k := opt.offset
						if ans == -1 || k < ans {
							ans = k
						}
						continue
					}
					for idx := 0; idx < len(opt.diffs); idx++ {
						curr += opt.diffs[idx]
						if curr <= C {
							k := opt.offset + idx + 1
							if ans == -1 || k < ans {
								ans = k
							}
							break
						}
					}
				}
			}
		}

		writer.WriteString(strconv.Itoa(ans))
		writer.WriteByte('\n')
	}
}
