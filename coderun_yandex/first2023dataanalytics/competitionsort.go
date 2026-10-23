package first2023dataanalytics

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Participant struct {
	time int
	val  int64
}

// https://coderun.yandex.ru/selections/first-2023-data-analytics/problems/competition-sort
// CompetitionSort - problem 37
func CompetitionSort() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	fields := strings.Fields(line)

	n, _ := strconv.Atoi(fields[0])
	k, _ := strconv.Atoi(fields[1])

	fixed := make([]Participant, n)
	for i := 0; i < n; i++ {
		var id, time int
		var team string
		fmt.Fscan(reader, &id, &time, &team)

		val := int64(-1)
		if team == "SKI" {
			val = 1
		}

		fixed[i] = Participant{
			time: time,
			val:  val,
		}
	}

	sort.Slice(fixed, func(i, j int) bool {
		return fixed[i].time < fixed[j].time
	})

	var base int64
	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		points := int64(n - i)
		base += fixed[i].val * points
		prefix[i+1] = prefix[i] + fixed[i].val
	}

	weights := make([]int64, k)
	for i := 0; i < k; i++ {
		var id, time int
		var team string
		fmt.Fscan(reader, &id, &time, &team)

		L := sort.Search(n, func(j int) bool {
			return fixed[j].time >= time
		})

		R := n - L

		weights[i] = int64(1+R) + prefix[L]
	}

	sort.Slice(weights, func(i, j int) bool {
		return weights[i] > weights[j]
	})

	answer := base
	var sum int64
	for m := 1; m <= k; m++ {
		sum += weights[m-1]

		extra := int64(m) * int64(m-1) / 2
		score := base + sum + extra

		if score > answer {
			answer = score
		}
	}

	fmt.Fprintln(writer, answer)
}
