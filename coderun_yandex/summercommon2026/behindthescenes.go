package summercommon2026

import (
	"bufio"
	"os"
)

// https://coderun.yandex.ru/seasons/2026-summer/tracks/common/problem/behind-the-scenes
// BehindTheScenes - problem 3
func BehindTheScenes() {
	w := bufio.NewWriterSize(os.Stdout, 32)
	defer w.Flush()

	sc := NewFastScanner()
	n := sc.NextInt()

	arrived := make([]bool, n+2)
	waiting := 0
	best := 0
	next := 1
	for i := 0; i < n; i++ {
		x := sc.NextInt()

		arrived[x] = true
		waiting++

		for next <= n && arrived[next] {
			waiting--
			next++
		}

		if waiting > best {
			best = waiting
		}
	}

	if best == 0 {
		w.WriteString("0\n")

		return
	}

	var buf [20]byte
	pos := len(buf)
	for best > 0 {
		pos--
		buf[pos] = byte(best%10) + '0'
		best /= 10
	}

	w.Write(buf[pos:])
	w.WriteByte('\n')
}
