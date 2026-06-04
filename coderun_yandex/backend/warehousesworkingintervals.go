package backend

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type Interval struct {
	Start string
	End   string
}

type Key struct {
	ID   int
	Type string
}

type Record struct {
	ID    int
	Start string
	End   string
	Type  string
}

// https://coderun.yandex.ru/selections/backend/problems/warehouses-working-intervals
// WarehousesWorkingIntervals - problem 41
func WarehousesWorkingIntervals() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	groups := make(map[Key][]Interval)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}

		var id int
		fmt.Sscanf(parts[0], "%d", &id)

		dates := strings.Split(parts[1], " ")
		start := dates[0]
		end := dates[1]

		typ := parts[2]

		if typ == "NULL" {
			for _, t := range []string{"KGT", "COLD", "OTHER"} {
				k := Key{id, t}
				groups[k] = append(groups[k], Interval{start, end})
			}
		} else {
			k := Key{id, typ}
			groups[k] = append(groups[k], Interval{start, end})
		}
	}

	var result []Record
	for k, intervals := range groups {
		sort.Slice(intervals, func(i, j int) bool {
			return intervals[i].Start < intervals[j].Start
		})

		merged := make([]Interval, 0)
		for _, cur := range intervals {
			if len(merged) == 0 {
				merged = append(merged, cur)
				continue
			}

			last := &merged[len(merged)-1]

			if cur.Start <= last.End {
				if cur.End > last.End {
					last.End = cur.End
				}
			} else {
				merged = append(merged, cur)
			}
		}

		for _, iv := range merged {
			result = append(result, Record{
				ID:    k.ID,
				Start: iv.Start,
				End:   iv.End,
				Type:  k.Type,
			})
		}
	}

	order := map[string]int{
		"KGT":   0,
		"COLD":  1,
		"OTHER": 2,
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].ID != result[j].ID {
			return result[i].ID < result[j].ID
		}

		if result[i].Type != result[j].Type {
			return order[result[i].Type] < order[result[j].Type]
		}

		return result[i].Start < result[j].Start
	})

	for _, r := range result {
		ans := fmt.Sprintf(
			"%d,%s %s,%s",
			r.ID,
			r.Start,
			r.End,
			r.Type,
		)

		writer.WriteString(ans)
		writer.WriteByte('\n')
	}
}
