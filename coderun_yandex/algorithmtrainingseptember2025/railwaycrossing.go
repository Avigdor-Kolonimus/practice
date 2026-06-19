package algorithmtrainingseptember2025

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type IntervalRailwayCrossing struct {
	l float64
	r float64
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/railway-crossing
// RailwayCrossing - assignment 36
func RailwayCrossing() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N, M and X input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 3 {
		panic("numbers count does not match 3")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	x, err := strconv.Atoi(strNum[2])
	if err != nil {
		panic(err)
	}

	// intervals input
	intervals := make([]IntervalRailwayCrossing, 0, n)
	for i := 0; i < n; i++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum := strings.Fields(line)
		if len(strNum) != 3 {
			panic("numbers count does not match 3")
		}

		a, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}
		v, err := strconv.Atoi(strNum[2])
		if err != nil {
			panic(err)
		}

		var l, r float64
		xf := float64(x)
		vf := float64(v)

		if a < b { // move right
			enter := (xf - float64(b)) / vf
			leave := (xf - float64(a)) / vf

			if leave < 0 {
				continue
			}

			if enter < 0 {
				enter = 0
			}

			l, r = enter, leave
		} else { // move left
			enter := (float64(b) - xf) / vf
			leave := (float64(a) - xf) / vf

			if leave < 0 {
				continue
			}

			if enter < 0 {
				enter = 0
			}

			l, r = enter, leave
		}

		if l <= r {
			intervals = append(intervals, IntervalRailwayCrossing{l, r})
		}
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].l == intervals[j].l {
			return intervals[i].r < intervals[j].r
		}
		return intervals[i].l < intervals[j].l
	})

	merged := make([]IntervalRailwayCrossing, 0, len(intervals))
	const eps = 1e-12
	for _, cur := range intervals {
		if len(merged) == 0 {
			merged = append(merged, cur)
			continue
		}

		last := &merged[len(merged)-1]

		if cur.l <= last.r+eps {
			if cur.r > last.r {
				last.r = cur.r
			}
		} else {
			merged = append(merged, cur)
		}
	}

	// car time input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum = strings.Fields(line)
	if len(strNum) != m {
		panic("numbers count does not match m")
	}

	for i := 0; i < m; i++ {
		ti, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		t := float64(ti)

		pos := sort.Search(len(merged), func(j int) bool {
			return merged[j].l > t
		}) - 1

		ans := t
		if pos >= 0 &&
			merged[pos].l-eps <= t &&
			t <= merged[pos].r+eps {
			ans = merged[pos].r
		}

		writer.WriteString(strconv.FormatFloat(ans, 'f', 9, 64))
		writer.WriteByte('\n')
	}
}
