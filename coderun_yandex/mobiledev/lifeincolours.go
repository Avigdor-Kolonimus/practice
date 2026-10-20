package mobiledev

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Interval struct {
	l, r int64
}

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/life-in-colours
// LifeInColours - problem 7
func LifeInColours() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var k int
	var t int64
	fmt.Fscan(reader, &k, &t)

	intervals := make([]Interval, 0, k)

	for i := 0; i < k; i++ {
		var start, end, period int64
		fmt.Fscan(reader, &start, &end, &period)

		l := end
		r := start + period - 1

		if l <= r {
			intervals = append(intervals, Interval{l, r})
		}
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].l == intervals[j].l {
			return intervals[i].r < intervals[j].r
		}
		return intervals[i].l < intervals[j].l
	})

	var covered int64 = 0
	for _, interval := range intervals {
		if interval.l > covered+1 {
			fmt.Fprintln(writer, "No")
			return
		}

		if interval.r > covered {
			covered = interval.r
		}

		if covered >= t {
			fmt.Fprintln(writer, "Yes")
			return
		}
	}

	fmt.Fprintln(writer, "No")
}
