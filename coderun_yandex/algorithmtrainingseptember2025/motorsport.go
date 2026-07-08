package algorithmtrainingseptember2025

import (
	"bufio"
	"os"
	"sort"
)

type FastScanner struct {
	r []byte
	i int
}

func NewFastScanner() *FastScanner {
	data, _ := bufio.NewReaderSize(os.Stdin, 1<<20).ReadBytes(0)
	data = append(data, 0)
	return &FastScanner{r: data}
}

func (fs *FastScanner) skipSpaces() {
	for {
		c := fs.r[fs.i]
		if c != ' ' && c != '\n' && c != '\r' && c != '\t' {
			return
		}
		fs.i++
	}
}

func (fs *FastScanner) Int64() int64 {
	fs.skipSpaces()

	sign := int64(1)
	if fs.r[fs.i] == '-' {
		sign = -1
		fs.i++
	}

	var x int64
	for {
		c := fs.r[fs.i]
		if c < '0' || c > '9' {
			break
		}
		x = x*10 + int64(c-'0')
		fs.i++
	}
	return x * sign
}

type FastOutput struct {
	buf []byte
}

func NewFastOutput(capacity int) *FastOutput {
	return &FastOutput{
		buf: make([]byte, 0, capacity),
	}
}

func (o *FastOutput) WriteInt(x int) {
	if x == 0 {
		o.buf = append(o.buf, '0')
		return
	}

	if x < 0 {
		o.buf = append(o.buf, '-')
		x = -x
	}

	var tmp [20]byte
	n := 0

	for x > 0 {
		tmp[n] = byte(x%10) + '0'
		x /= 10
		n++
	}

	for i := n - 1; i >= 0; i-- {
		o.buf = append(o.buf, tmp[i])
	}
}

func (o *FastOutput) Space() {
	o.buf = append(o.buf, ' ')
}

func (o *FastOutput) NL() {
	o.buf = append(o.buf, '\n')
}

func (o *FastOutput) Flush() {
	os.Stdout.Write(o.buf)
}

func gcd(a, b int64) int64 {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

type Time struct {
	num int64
	den int64
}

func NewTime(n, d int64) Time {
	if d < 0 {
		n = -n
		d = -d
	}

	g := gcd(n, d)

	return Time{
		num: n / g,
		den: d / g,
	}
}

func (a Time) Less(b Time) bool {
	return a.num*b.den < b.num*a.den
}

func (a Time) Equal(b Time) bool {
	return a.num == b.num && a.den == b.den
}

type Event struct {
	t    Time
	kind int8 // 0 wall, 1 finish, 2 collision
	a    int
	b    int
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/motorsport
// Motorsport - assignment 40
func Motorsport() {
	sc := NewFastScanner()

	N := int(sc.Int64())
	L := sc.Int64()
	W := sc.Int64()

	x := make([]int64, N)
	y := make([]int64, N)
	vx := make([]int64, N)
	vy := make([]int64, N)

	for i := 0; i < N; i++ {
		x[i] = sc.Int64()
		y[i] = sc.Int64()
		vx[i] = sc.Int64()
		vy[i] = sc.Int64()
	}

	wallTime := make([]*Time, N)

	for i := 0; i < N; i++ {
		if vy[i] > 0 {
			t := NewTime(W-y[i], vy[i])
			wallTime[i] = &t
		} else if vy[i] < 0 {
			t := NewTime(y[i], -vy[i])
			wallTime[i] = &t
		}
	}

	finishTime := make([]*Time, N)

	for i := 0; i < N; i++ {
		if vx[i] <= 0 {
			continue
		}

		dx := L - x[i]
		if dx <= 0 {
			continue
		}

		t := NewTime(dx, vx[i])

		yNum := y[i]*vx[i] + vy[i]*dx
		wv := W * vx[i]

		if yNum <= 0 || yNum >= wv {
			continue
		}

		if wallTime[i] != nil {
			if !t.Less(*wallTime[i]) {
				continue
			}
		}

		tmp := t
		finishTime[i] = &tmp
	}

	terminalTime := make([]*Time, N)

	for i := 0; i < N; i++ {
		wt := wallTime[i]
		ft := finishTime[i]

		if ft != nil {
			if wt == nil || ft.Less(*wt) {
				terminalTime[i] = ft
			} else {
				terminalTime[i] = wt
			}
		} else {
			terminalTime[i] = wt
		}
	}

	events := make([]Event, 0, N*N/2+3*N)

	for i := 0; i < N; i++ {
		if finishTime[i] != nil {
			events = append(events, Event{
				t:    *finishTime[i],
				kind: 1,
				a:    i,
				b:    -1,
			})
		}
	}

	for i := 0; i < N; i++ {
		if wallTime[i] == nil {
			continue
		}

		if finishTime[i] != nil {
			if wallTime[i].Less(*finishTime[i]) {
				events = append(events, Event{
					t:    *wallTime[i],
					kind: 0,
					a:    i,
					b:    -1,
				})
			}
		} else {
			events = append(events, Event{
				t:    *wallTime[i],
				kind: 0,
				a:    i,
				b:    -1,
			})
		}
	}

	if N >= 2 {
		for i := 0; i < N-1; i++ {
			for j := i + 1; j < N; j++ {

				dx := x[j] - x[i]
				dy := y[j] - y[i]

				dvx := vx[i] - vx[j]
				dvy := vy[i] - vy[j]

				if dvx == 0 && dvy == 0 {
					continue
				}

				var tt Time
				ok := false

				if dvx == 0 {
					if dx != 0 {
						continue
					}
					if dvy == 0 {
						continue
					}

					tt = NewTime(dy, dvy)
					ok = true

				} else if dvy == 0 {

					if dy != 0 {
						continue
					}

					tt = NewTime(dx, dvx)
					ok = true

				} else {

					if dx*dvy != dy*dvx {
						continue
					}

					tt = NewTime(dx, dvx)
					ok = true
				}

				if !ok || tt.num <= 0 {
					continue
				}

				if terminalTime[i] != nil &&
					terminalTime[i].Less(tt) {
					continue
				}

				if terminalTime[j] != nil &&
					terminalTime[j].Less(tt) {
					continue
				}

				events = append(events, Event{
					t:    tt,
					kind: 2,
					a:    i,
					b:    j,
				})
			}
		}
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].t.Equal(events[j].t) {
			return events[i].kind < events[j].kind
		}
		return events[i].t.Less(events[j].t)
	})
	alive := make([]bool, N)
	for i := range alive {
		alive[i] = true
	}

	var bestTime *Time
	var winners []int

	markCollide := make([]int, N)
	markWall := make([]int, N)
	markFinish := make([]int, N)
	markTouched := make([]int, N)

	stamp := 0
	idx := 0

	for idx < len(events) {
		t := events[idx].t

		if bestTime != nil && bestTime.Less(t) {
			break
		}

		j := idx
		for j < len(events) && events[j].t.Equal(t) {
			j++
		}

		stamp++

		touched := make([]int, 0, 32)

		touch := func(v int) {
			if markTouched[v] != stamp {
				markTouched[v] = stamp
				touched = append(touched, v)
			}
		}

		for k := idx; k < j; k++ {
			e := events[k]

			switch e.kind {

			case 0: // wall
				if alive[e.a] {
					markWall[e.a] = stamp
					touch(e.a)
				}

			case 1: // finish
				if alive[e.a] {
					markFinish[e.a] = stamp
					touch(e.a)
				}

			default: // collision
				if alive[e.a] && alive[e.b] {
					markCollide[e.a] = stamp
					markCollide[e.b] = stamp
					touch(e.a)
					touch(e.b)
				}
			}
		}

		crashList := make([]int, 0, 32)
		finishList := make([]int, 0, 32)

		for _, v := range touched {
			if !alive[v] {
				continue
			}

			coll := markCollide[v] == stamp
			wall := markWall[v] == stamp
			fin := markFinish[v] == stamp

			if fin && !coll && !wall {
				finishList = append(finishList, v)
			}

			if coll || wall {
				crashList = append(crashList, v)
			}
		}

		for _, v := range crashList {
			alive[v] = false
		}

		for _, v := range finishList {
			alive[v] = false
		}

		if len(finishList) > 0 {
			if bestTime == nil {
				tmp := t
				bestTime = &tmp
				winners = append(winners, finishList...)
			} else {
				winners = append(winners, finishList...)
			}
		}

		idx = j
	}

	for i := range winners {
		winners[i]++
	}

	sort.Ints(winners)

	out := NewFastOutput(64 + len(winners)*8)

	out.WriteInt(len(winners))
	out.NL()

	for i, v := range winners {
		if i > 0 {
			out.Space()
		}
		out.WriteInt(v)
	}

	out.NL()
	out.Flush()
}
