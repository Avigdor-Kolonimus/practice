package summercommon2026

import (
	"bufio"
	"os"
)

// https://coderun.yandex.ru/seasons/2026-summer/tracks/common/problem/bike-rental
// BikeRental - problem 1
func BikeRental() {
	var a, f, s, n, t, i int
	var cur, best int64

	sc := NewFastScanner()
	w := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer w.Flush()

	n = sc.NextInt()
	t = sc.NextInt()

	diff := make([]int, t+1)
	for i = 0; i < n; i++ {
		a = sc.NextInt()
		f = sc.NextInt()
		s = sc.NextInt()

		diff[a] += s
		diff[f] -= s
	}

	for i = 0; i <= t; i++ {
		cur += int64(diff[i])
		if cur > best {
			best = cur
		}
	}

	writeInt64(w, best)
}
