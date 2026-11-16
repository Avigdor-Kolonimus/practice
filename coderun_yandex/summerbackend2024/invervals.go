package summerbackend2024

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Segment struct {
	a, b int64
}

type Event struct {
	x     int64
	delta int64
}

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/invervals
// Invervals - problem 21
func Invervals() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	var l, r int64

	fmt.Fscan(in, &n, &m)
	fmt.Fscan(in, &l, &r)

	blue := make([]Segment, n)
	red := make([]Segment, m)

	for i := range blue {
		fmt.Fscan(in, &blue[i].a, &blue[i].b)
	}

	L := int64(-1 << 60)
	R := int64(1 << 60)

	for i := range red {
		fmt.Fscan(in, &red[i].a, &red[i].b)

		L = max(L, l-red[i].a)
		R = min(R, r-red[i].b)
	}

	events := make([]Event, 0, 4*n*m)

	var current int64
	var slope int64
	addEvent := func(x, delta int64) {
		if x > L {
			slope -= delta

			if x <= R {
				events = append(events, Event{
					x:     x,
					delta: delta,
				})
			}
		}
	}

	for _, b := range blue {
		for _, c := range red {
			a, bb := b.a, b.b
			cc, d := c.a, c.b

			left := max(a, cc+L)
			right := min(bb, d+L)

			if right > left {
				current += right - left
			}

			addEvent(bb-cc, 1)
			addEvent(a-cc, -1)
			addEvent(bb-d, -1)
			addEvent(a-d, 1)
		}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].x < events[j].x
	})

	ans := current
	pos := L
	for i := 0; i < len(events); {
		x := events[i].x

		current += slope * (x - pos)
		pos = x

		if current < ans {
			ans = current
		}

		for i < len(events) && events[i].x == x {
			slope += events[i].delta
			i++
		}
	}

	current += slope * (R - pos)

	if current < ans {
		ans = current
	}

	fmt.Fprintln(out, ans)
}
