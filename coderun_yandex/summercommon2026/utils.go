package summercommon2026

import (
	"bufio"
	"os"
)

type FastScanner struct {
	buf []byte
	idx int
	n   int
}

func NewFastScanner() *FastScanner {
	return &FastScanner{
		buf: make([]byte, 1<<20),
	}
}

// Refill the buffer. Returns false on EOF.
func (fs *FastScanner) refill() bool {
	fs.n, _ = os.Stdin.Read(fs.buf)
	fs.idx = 0
	return fs.n > 0
}

func (fs *FastScanner) NextInt() int {
	for {
		if fs.idx >= fs.n {
			fs.n, _ = os.Stdin.Read(fs.buf)
			fs.idx = 0
		}
		if fs.buf[fs.idx] >= '0' && fs.buf[fs.idx] <= '9' {
			break
		}
		fs.idx++
	}

	x := 0
	for {
		if fs.idx >= fs.n {
			n, _ := os.Stdin.Read(fs.buf)
			if n == 0 {
				break
			}
			fs.n = n
			fs.idx = 0
		}

		c := fs.buf[fs.idx]
		if c < '0' || c > '9' {
			break
		}

		x = x*10 + int(c-'0')
		fs.idx++
	}

	return x
}

func writeInt(w *bufio.Writer, x int) {
	if x == 0 {
		w.WriteByte('0')
		w.WriteByte('\n')
		return
	}

	var buf [20]byte
	pos := len(buf)

	for x > 0 {
		pos--
		buf[pos] = byte(x%10) + '0'
		x /= 10
	}

	w.Write(buf[pos:])
	w.WriteByte('\n')
}

func writeInt64(w *bufio.Writer, x int64) {
	var buf [20]byte
	i := len(buf)

	for {
		i--
		buf[i] = byte(x%10) + '0'
		x /= 10
		if x == 0 {
			break
		}
	}

	w.Write(buf[i:])
	w.WriteByte('\n')
}
