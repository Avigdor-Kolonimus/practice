package problems

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Interval struct {
	start int64
	end   int64
}

// https://coderun.yandex.ru/problem/observation-students
// ObservationStudents - problem 232
func ObservationStudents() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	//  N and M input
	line, _ := reader.ReadString('\n')
	first := strings.Fields(strings.TrimSpace(line))

	n, _ := strconv.ParseInt(first[0], 10, 64)
	m, _ := strconv.Atoi(first[1])

	// intervals input
	intervals := make([]Interval, m)
	for i := range m {
		line, _ := reader.ReadString('\n')
		values := strings.Fields(strings.TrimSpace(line))

		b, _ := strconv.ParseInt(values[0], 10, 64)
		e, _ := strconv.ParseInt(values[1], 10, 64)

		intervals[i] = Interval{
			start: b,
			end:   e,
		}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	var covered int64
	left := intervals[0].start
	right := intervals[0].end
	for i := 1; i < m; i++ {
		cur := intervals[i]

		if cur.start <= right+1 {
			if cur.end > right {
				right = cur.end
			}
		} else {
			covered += right - left + 1

			left = cur.start
			right = cur.end
		}
	}

	covered += right - left + 1

	answer := n - covered

	writer.WriteString(strconv.FormatInt(answer, 10))
	writer.WriteByte('\n')
}
