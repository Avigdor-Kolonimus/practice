package assignments

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Interval struct {
	Bi float64
	Ei float64
	Wi float64
}

func validateScheduleforWeightedIntervalsInput(p int) bool {
	return p >= 0 && p <= 100_000
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/schedule-for-weighted-intervals
// ScheduleforWeightedIntervals - assignment 13
func ScheduleforWeightedIntervals() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first  input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateScheduleforWeightedIntervalsInput(n) {
		panic("number N out of range")
	}

	intervals := make([]Interval, 0, n)
	for range n {
		// input
		line, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		parameters := strings.Fields(line)
		cnt := len(parameters)
		if cnt != 3 {
			panic("input does not match 3")
		}

		b, err := strconv.ParseFloat(parameters[0], 64)
		if err != nil {
			panic(err)
		}
		e, err := strconv.ParseFloat(parameters[1], 64)
		if err != nil {
			panic(err)
		}
		w, err := strconv.ParseFloat(parameters[2], 64)
		if err != nil {
			panic(err)
		}

		intervals = append(intervals, Interval{Bi: b, Ei: e, Wi: w})
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Ei == intervals[j].Ei {
			return intervals[i].Bi < intervals[j].Bi
		}
		return intervals[i].Ei < intervals[j].Ei
	})

	dp := make([]float64, n+1)
	ends := make([]float64, n)
	starts := make([]float64, n)
	weights := make([]float64, n)

	for i := range n {
		starts[i] = intervals[i].Bi
		ends[i] = intervals[i].Ei
		weights[i] = intervals[i].Wi
	}

	for i := range n {
		s := starts[i]

		lo := 0
		hi := i
		for lo < hi {
			mid := (lo + hi) / 2
			if ends[mid] <= s {
				lo = mid + 1
			} else {
				hi = mid
			}
		}

		take := dp[lo] + weights[i]
		skip := dp[i]
		dp[i+1] = max(skip, take)
	}

	result := dp[n]

	writer.WriteString(strconv.FormatFloat(result, 'f', 4, 64))
	writer.WriteByte('\n')
}
