package wintercommon2025

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

const (
	MOD_CABLEWAY int64 = 998244353
	MAX_N              = 3000
)

// prefix[n][k] = S(n,0) + S(n,1) + ... + S(n,k)
var prefix [MAX_N + 1][MAX_N + 1]int64

func init() {
	prefix[0][0] = 1

	for n := 1; n <= MAX_N; n++ {
		for k := 1; k <= n; k++ {
			// S(n,k) recurrence summed over 0..k:
			//
			// P(n,k) =
			// P(n-1,k-1) + (n-1) * P(n-1,k)
			//
			// For k == n, P(n-1,k) means the whole row,
			// i.e. P(n-1,n-1).

			prevK := k
			if prevK > n-1 {
				prevK = n - 1
			}

			value := prefix[n-1][k-1] +
				int64(n-1)*prefix[n-1][prevK]

			prefix[n][k] = value % MOD_CABLEWAY
		}
	}
}

type FastScannerCabelway struct {
	data []byte
	pos  int
	n    int
}

func NewFastScannerCabelway(reader *bufio.Reader) *FastScannerCabelway {
	data, _ := io.ReadAll(reader)

	return &FastScannerCabelway{
		data: data,
		n:    len(data),
	}
}

func (fs *FastScannerCabelway) NextInt() int {
	for fs.pos < fs.n &&
		(fs.data[fs.pos] < '0' || fs.data[fs.pos] > '9') &&
		fs.data[fs.pos] != '-' {
		fs.pos++
	}

	sign := 1

	if fs.data[fs.pos] == '-' {
		sign = -1
		fs.pos++
	}

	res := 0

	for fs.pos < fs.n &&
		fs.data[fs.pos] >= '0' &&
		fs.data[fs.pos] <= '9' {

		res = res*10 + int(fs.data[fs.pos]-'0')
		fs.pos++
	}

	return res * sign
}

func countCycles(n int, next []int, inDeg []bool, outDeg []bool) int {
	visited := make([]bool, n+1)

	for v := 1; v <= n; v++ {
		if inDeg[v] {
			continue
		}

		cur := v

		for cur != 0 && !visited[cur] {
			visited[cur] = true

			if !outDeg[cur] {
				break
			}

			cur = next[cur]
		}
	}

	cycles := 0
	for v := 1; v <= n; v++ {
		if visited[v] {
			continue
		}

		cycles++

		cur := v

		for !visited[cur] {
			visited[cur] = true
			cur = next[cur]
		}
	}

	return cycles
}

func solve(n, q, l, r int, b, c []int) int64 {
	inDeg := make([]bool, n+1)
	outDeg := make([]bool, n+1)
	next := make([]int, n+1)

	for i := 0; i < q; i++ {
		u := b[i]
		v := c[i]

		if outDeg[u] || inDeg[v] {
			return 0
		}

		outDeg[u] = true
		inDeg[v] = true
		next[u] = v
	}

	visited := make([]bool, n+1)
	fixedCycles := 0

	for v := 1; v <= n; v++ {
		if inDeg[v] {
			continue
		}

		cur := v

		for cur != 0 && !visited[cur] {
			visited[cur] = true

			if !outDeg[cur] {
				break
			}

			cur = next[cur]
		}
	}

	for v := 1; v <= n; v++ {
		if visited[v] {
			continue
		}

		fixedCycles++

		cur := v
		for !visited[cur] {
			visited[cur] = true
			cur = next[cur]
		}
	}

	m := n - q

	needL := l - fixedCycles
	needR := r - fixedCycles

	if needL < 0 {
		needL = 0
	}

	if needR > m {
		needR = m
	}

	if needL > needR || needR < 0 || needL > m {
		return 0
	}

	// Sum S(m,k), k = needL..needR.
	ans := prefix[m][needR]

	if needL > 0 {
		ans -= prefix[m][needL-1]
	}

	if ans < 0 {
		ans += MOD_CABLEWAY
	}

	return ans
}

// https://coderun.yandex.ru/selections/2025-winter-common/problems/cableway
// Cableway - problem 16
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	sc := NewFastScannerCabelway(reader)

	t := sc.NextInt()
	for ; t > 0; t-- {
		n := sc.NextInt()
		q := sc.NextInt()
		l := sc.NextInt()
		r := sc.NextInt()

		b := make([]int, q)
		c := make([]int, q)

		for i := 0; i < q; i++ {
			b[i] = sc.NextInt()
		}

		for i := 0; i < q; i++ {
			c[i] = sc.NextInt()
		}

		ans := solve(n, q, l, r, b, c)

		writer.WriteString(strconv.FormatInt(ans, 10))
		writer.WriteByte('\n')
	}
}
