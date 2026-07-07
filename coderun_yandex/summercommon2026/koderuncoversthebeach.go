package summercommon2026

import (
	"bufio"
	"os"
)

const (
	BUF_SIZE = 1 << 22 // 4 MB
	MAXN     = 500_000
)

var (
	aBuf   = make([]int64, MAXN)
	valBuf = make([]int64, MAXN)
	deqBuf = make([]int, MAXN)
)

// https://coderun.yandex.ru/seasons/2026-summer/tracks/common/problem/koderun-covers-the-beach
// KoderunCoversTheBeach - problem 6
func KoderunCoversTheBeach() {
	fs := NewFastScanner()

	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	t := fs.NextInt()
	for ; t > 0; t-- {
		n := fs.NextInt()
		k := fs.NextInt()

		a := aBuf[:n]
		for i := range n {
			a[i] = int64(fs.NextInt())
		}

		m := n - k + 1
		val := valBuf[:m]

		sum, xr := int64(0), int64(0)
		for i := 0; i < k; i++ {
			sum += a[i]
			xr ^= a[i]
		}

		val[0] = sum - xr
		for i := 1; i < m; i++ {
			sum += a[i+k-1] - a[i-1]
			xr ^= a[i-1]
			xr ^= a[i+k-1]
			val[i] = sum - xr
		}

		deq := deqBuf[:m]
		head, tail, right := 0, 0, -1
		for i := range n {
			l := i - k + 1
			l = max(l, 0)

			r := i
			if r >= m {
				r = m - 1
			}

			for right < r {
				right++

				for head < tail && val[deq[tail-1]] >= val[right] {
					tail--
				}

				deq[tail] = right
				tail++
			}

			for head < tail && deq[head] < l {
				head++
			}

			if i > 0 {
				out.WriteByte(' ')
			}
			writeInt64(out, val[deq[head]])
		}
		out.WriteByte('\n')
	}
}
