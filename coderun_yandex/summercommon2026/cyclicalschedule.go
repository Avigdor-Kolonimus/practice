package summercommon2026

import (
	"bufio"
	"os"
)

const (
	MaxN = 2_000_005
	Log  = 21
)

type Task struct {
	s int
	e int
}

type Tasks []Task

var (
	tasks [MaxN]Task
	up    [Log][MaxN]int
)

var tmp [MaxN]Task

func radixSort(a []Task, n int) {
	const B = 8
	const M = 1 << B

	var cnt [M]int

	for shift := 0; shift < 64; shift += B {

		for i := range M {
			cnt[i] = 0
		}

		for i := range n {
			key := (uint64(uint32(a[i].e)) << 32) |
				uint64(uint32(a[i].s))

			cnt[(key>>shift)&(M-1)]++
		}

		sum := 0
		for i := 0; i < M; i++ {
			c := cnt[i]
			cnt[i] = sum
			sum += c
		}

		for i := range n {
			key := (uint64(uint32(a[i].e)) << 32) |
				uint64(uint32(a[i].s))

			pos := cnt[(key>>shift)&(M-1)]
			tmp[pos] = a[i]
			cnt[(key>>shift)&(M-1)]++
		}

		copy(a, tmp[:n])
	}
}

// https://coderun.yandex.ru/seasons/2026-summer/tracks/common/problem/cyclical-schedule
// CyclicalSchedule - problem 9
func CyclicalSchedule() {
	fs := NewFastScanner()

	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	tc := fs.NextInt()

	for ; tc > 0; tc-- {

		n := fs.NextInt()
		T := fs.NextInt()

		for i := range n {
			s := fs.NextInt()
			d := fs.NextInt()

			e := s + d

			tasks[i].s = s
			tasks[i].e = e

			tasks[i+n].s = s + T
			tasks[i+n].e = e + T
		}

		m := n << 1

		radixSort(tasks[:m], m)

		ptr := 0
		for i := range m {
			for ptr < m && tasks[ptr].s < tasks[i].e {
				ptr++
			}
			up[0][i] = ptr
		}
		up[0][m] = m

		for k := 1; k < Log; k++ {
			prev := up[k-1][:]
			cur := up[k][:]
			for i := 0; i <= m; i++ {
				cur[i] = prev[prev[i]]
			}
		}

		best := 0

		for i := range m {
			if tasks[i].s >= T {
				break
			}

			limit := tasks[i].s + T

			cur := i
			cnt := 1

			for k := Log - 1; k >= 0; k-- {
				nxt := up[k][cur]
				if nxt < m && tasks[nxt].e <= limit {
					cur = nxt
					cnt += 1 << k
				}
			}

			if cnt > best {
				best = cnt
			}
		}

		writeInt(out, best)
	}
}
