package eserajim

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const MOD int64 = 1000000007
const MAX = 17

type Seg struct {
	a uint64
	b uint64
	c int
}

type Matrix struct {
	n int
	a [MAX][MAX]int64
}

func mul(A, B Matrix) Matrix {
	var C Matrix
	C.n = A.n
	for i := 0; i < A.n; i++ {
		for k := 0; k < A.n; k++ {
			if A.a[i][k] == 0 {
				continue
			}
			for j := 0; j < A.n; j++ {
				C.a[i][j] = (C.a[i][j] + A.a[i][k]*B.a[k][j]) % MOD
			}
		}
	}
	return C
}

func power(A Matrix, e uint64) Matrix {
	var R Matrix
	R.n = A.n
	for i := 0; i < A.n; i++ {
		R.a[i][i] = 1
	}
	for e > 0 {
		if e&1 == 1 {
			R = mul(R, A)
		}
		A = mul(A, A)
		e >>= 1
	}
	return R
}

func apply(M Matrix, v [MAX]int64) [MAX]int64 {
	var res [MAX]int64
	for i := 0; i < M.n; i++ {
		var cur int64
		for j := 0; j < M.n; j++ {
			cur = (cur + M.a[i][j]*v[j]) % MOD
		}
		res[i] = cur
	}
	return res
}

// https://coderun.yandex.ru/selections/eserajim/problems/spies
// Spies - problem 3
func Spies() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and K input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}

	tmp, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	k := uint64(tmp)

	// segments input
	seg := make([]Seg, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		strNum = strings.Fields(line)
		if len(strNum) != 3 {
			panic("numbers count does not match 3")
		}

		tmp, err = strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		a := uint64(tmp)

		tmp, err = strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}
		b := uint64(tmp)

		c, err := strconv.Atoi(strNum[2])
		if err != nil {
			panic(err)
		}

		seg[i] = Seg{
			a: a, b: b, c: c,
		}
	}

	var dp [MAX]int64
	dp[0] = 1
	for i := 0; i < n; i++ {
		var len uint64
		if i == n-1 {
			len = k - seg[i].a
		} else {
			len = seg[i].b - seg[i].a
		}

		var T Matrix
		T.n = seg[i].c + 1

		for from := 0; from <= seg[i].c; from++ {
			for d := -1; d <= 1; d++ {
				to := from + d
				if to >= 0 && to <= seg[i].c {
					T.a[to][from] = 1
				}
			}
		}

		P := power(T, len)
		dp = apply(P, dp)

		if i+1 < n {
			lim := seg[i].c
			if seg[i+1].c < lim {
				lim = seg[i+1].c
			}
			for y := lim + 1; y < MAX; y++ {
				dp[y] = 0
			}
		}
	}

	writer.WriteString(strconv.FormatInt(dp[0], 10))
	writer.WriteByte('\n')
}
