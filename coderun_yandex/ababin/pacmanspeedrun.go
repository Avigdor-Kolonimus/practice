package ababin

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type State struct {
	r, c int
	a, b int
}

type Directions struct {
	d [4]byte
}

func repeat(sb *strings.Builder, ch byte, count int) {
	for range count {
		sb.WriteByte(ch)
	}
}

func swapInt(x, y *int) {
	*x, *y = *y, *x
}

func swapByte(x, y *byte) {
	*x, *y = *y, *x
}

func solve(r, c, a, b int) string {
	s := State{
		r: r,
		c: c,
		a: a,
		b: b,
	}

	d := Directions{
		d: [4]byte{'L', 'U', 'R', 'D'},
	}

	var result strings.Builder
	if s.a == 1 || (s.b == s.c && s.a != s.r) {
		swapInt(&s.a, &s.b)
		swapInt(&s.r, &s.c)

		swapByte(&d.d[0], &d.d[1])
		swapByte(&d.d[2], &d.d[3])
	}

	if s.r == s.a {
		if s.b > 2 {
			repeat(&result, d.d[3], s.r-1)
			result.WriteByte(d.d[2])

			s.a = s.b - 1
			s.b = 1

			oldR := s.r
			s.r = s.c - 1
			s.c = oldR

			tmp := d.d[0]
			d.d[0] = d.d[3]
			d.d[3] = d.d[2]
			d.d[2] = d.d[1]
			d.d[1] = tmp

		} else {
			if s.b != 1 {
				if s.c == 2 {
					result.WriteByte(d.d[2])
					result.WriteByte(d.d[3])

					s.a--
					s.b = 1
					s.r--

					swapByte(&d.d[0], &d.d[2])

					if s.a == 1 {
						result.WriteByte(d.d[2])

						s.r = 2
						s.c = 1
						s.a = 2
						s.b = 1

						d.d[3] = d.d[0]
					}

				} else {
					repeat(&result, d.d[2], s.c-1)
					repeat(&result, d.d[3], s.r-1)

					result.WriteByte(d.d[0])

					s.a = s.b
					s.b = 1

					oldR := s.r
					s.r = s.c - 1
					s.c = oldR - 1

					tmp := d.d[0]
					d.d[0] = d.d[3]
					d.d[3] = d.d[2]
					d.d[2] = d.d[1]
					d.d[1] = tmp

					if s.a == s.r {
						result.WriteByte(d.d[1])
					} else {
						s.a = s.r - s.a + 1
						swapByte(&d.d[1], &d.d[3])
					}
				}
			}
		}
	} else {
		for l := 1; l < s.b; l++ {
			idx := 1 + (l%2)*2

			repeat(&result, d.d[idx], s.r-1)
			result.WriteByte(d.d[2])
		}

		if s.b%2 == 0 {
			s.a = s.r - s.a + 1
			swapByte(&d.d[1], &d.d[3])
		}

		s.c = s.c - s.b + 1
		s.b = 1
	}

	start := 1 + (s.a%2)*2

	for l := start; l < s.a; l++ {
		idx := (l % 2) * 2

		repeat(&result, d.d[idx], s.c-1)
		result.WriteByte(d.d[3])
	}

	if s.a%2 != 0 {
		for l := 1; l < s.c; l++ {
			idx := 1 + (l%2)*2

			result.WriteByte(d.d[idx])
			result.WriteByte(d.d[2])
		}

		if s.c%2 == 0 {
			result.WriteByte(d.d[1])
		}

		result.WriteByte(d.d[3])
		result.WriteByte(d.d[3])
	}

	diff := s.r - s.a
	start = diff - diff%2

	for l := start; l > 0; l-- {
		idx := (l % 2) * 2

		repeat(&result, d.d[idx], s.c-2)
		result.WriteByte(d.d[3])
	}

	if diff%2 == 0 {
		if s.c-1 != 0 {
			repeat(&result, d.d[0], s.c-1)
		}
	} else {
		for l := 1; l < s.c; l++ {
			idx := 1 + (l%2)*2

			result.WriteByte(d.d[idx])
			result.WriteByte(d.d[0])
		}

		if s.c%2 != 0 {
			result.WriteByte(d.d[3])
		}
	}

	if diff != 0 {
		repeat(&result, d.d[1], diff)
	}

	return result.String()
}

// https://coderun.yandex.ru/selections/ababin/problems/pacman-speedrun
// PacmanSpeedrun - problem 2
func PacmanSpeedrun() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// T input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	t, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// R,C,A and B input
	for ; t > 0; t-- {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		fields := strings.Fields(line)

		r, _ := strconv.Atoi(fields[0])
		c, _ := strconv.Atoi(fields[1])
		a, _ := strconv.Atoi(fields[2])
		b, _ := strconv.Atoi(fields[3])

		writer.WriteString(solve(r, c, a, b))
		writer.WriteByte('\n')
	}
}
