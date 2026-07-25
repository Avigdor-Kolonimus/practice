package atolstikov

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

// https://coderun.yandex.ru/selections/atolstikov/problems/lab-work
// LabWork - problem 4
func LabWork() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	readInt64 := func() int64 {
		n := int64(0)
		c, _ := reader.ReadByte()
		// skip whitespace/newlines
		for c == ' ' || c == '\n' || c == '\r' || c == '\t' {
			c, _ = reader.ReadByte()
		}
		neg := false
		if c == '-' {
			neg = true
			c, _ = reader.ReadByte()
		}
		for c >= '0' && c <= '9' {
			n = n*10 + int64(c-'0')
			c, _ = reader.ReadByte()
		}
		if neg {
			n = -n
		}
		return n
	}

	N := readInt64()
	X := readInt64()
	K := readInt64()

	A := make([]int64, N)
	var totalSum int64 = 0
	for i := int64(0); i < N; i++ {
		A[i] = readInt64()
		totalSum += A[i]
	}

	var F int64 = 0
	var rem []int64

	if X > 0 {
		rem = make([]int64, 0, N)
		for i := int64(0); i < N; i++ {
			F += A[i] / X
			r := A[i] % X
			if r > 0 {
				rem = append(rem, r)
			}
		}
		sort.Slice(rem, func(i, j int) bool { return rem[i] > rem[j] })
	}

	remLen := int64(len(rem))
	prefix := make([]int64, remLen+1)
	for i := int64(0); i < remLen; i++ {
		prefix[i+1] = prefix[i] + rem[i]
	}

	// maxReduction(D)
	maxReduction := func(D int64) int64 {
		if X == 0 {
			return 0
		}
		if D <= F {
			return D * X
		}
		extra := D - F
		if extra > remLen {
			extra = remLen
		}
		return F*X + prefix[extra]
	}

	feasible := func(D int64) bool {
		reduction := maxReduction(D)
		need := totalSum - reduction
		if need <= 0 {
			return true
		}
		if K == 0 {
			return false
		}
		requiredD := (need + K - 1) / K
		return D >= requiredD
	}

	// upper bound
	var hi int64 = 1
	if K > 0 {
		if totalSum > hi {
			hi = totalSum
		}
	}
	if X > 0 {
		dFull := F + remLen
		if dFull > hi {
			hi = dFull
		}
	}
	if hi < 1 {
		hi = 1
	}

	lo := int64(1)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if feasible(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	writer.WriteString(strconv.FormatInt(lo, 10))
	writer.WriteByte('\n')
}
