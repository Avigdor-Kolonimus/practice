package algorithmtrainingseptember2025

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type PairPlanningTheSeries struct {
	s int
	a int
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/planning-the-series
// PlanningTheSeries - assignment 37
func PlanningTheSeries() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// series input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	series := make([]int, n)
	for i := 0; i < n; i++ {
		s, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		series[i] = s
	}

	// importance coefficient input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNum = strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	total := 0
	impCoef := make([]int, n)
	for i := 0; i < n; i++ {
		a, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		impCoef[i] = a
		total += a
	}

	pairs := make([]PairPlanningTheSeries, n)

	for i := 0; i < n; i++ {
		pairs[i] = PairPlanningTheSeries{series[i], impCoef[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].s < pairs[j].s
	})

	pref := 0
	e := 0
	for _, p := range pairs {
		pref += p.a

		if pref*2 >= total {
			e = p.s

			break
		}
	}

	cost := 0
	for i := 0; i < n; i++ {
		diff := series[i] - e
		if diff < 0 {
			diff = -diff
		}

		cost += diff * impCoef[i]
	}

	writer.WriteString(strconv.Itoa(e))
	writer.WriteByte(' ')
	writer.WriteString(strconv.Itoa(cost))
	writer.WriteByte('\n')
}
