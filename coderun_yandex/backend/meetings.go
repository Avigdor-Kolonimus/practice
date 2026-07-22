package backend

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Meeting struct {
	day      int
	start    int
	duration int
	timeStr  string
	names    []string
}

func parseTime(s string) int {
	var h, m int
	fmt.Sscanf(s, "%d:%d", &h, &m)
	return h*60 + m
}

func overlap(s1, e1, s2, e2 int) bool {
	return s1 < e2 && s2 < e1
}

// https://coderun.yandex.ru/selections/backend/problems/meetings
// Meetings - problem 19
func Meetings() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	calendar := make(map[string][]Meeting)
	for ; n > 0; n-- {
		var cmd string
		fmt.Fscan(reader, &cmd)

		if cmd == "APPOINT" {
			var day int
			var timeStr string
			var duration, k int

			fmt.Fscan(reader, &day, &timeStr, &duration, &k)

			names := make([]string, k)
			for i := 0; i < k; i++ {
				fmt.Fscan(reader, &names[i])
			}

			start := parseTime(timeStr)
			end := start + duration

			conflicts := make([]string, 0)

			for _, name := range names {
				ok := true
				for _, meet := range calendar[name] {
					if meet.day != day {
						continue
					}
					if overlap(meet.start, meet.start+meet.duration, start, end) {
						ok = false
						break
					}
				}
				if !ok {
					conflicts = append(conflicts, name)
				}
			}

			if len(conflicts) > 0 {
				fmt.Fprintln(writer, "FAIL")
				for _, name := range conflicts {
					fmt.Fprint(writer, name, " ")
				}
				fmt.Fprintln(writer)
				continue
			}

			fmt.Fprintln(writer, "OK")

			meeting := Meeting{
				day:      day,
				start:    start,
				duration: duration,
				timeStr:  timeStr,
				names:    append([]string(nil), names...),
			}

			for _, name := range names {
				calendar[name] = append(calendar[name], meeting)
			}

		} else { // PRINT
			var day int
			var name string

			fmt.Fscan(reader, &day, &name)

			list := make([]Meeting, 0)

			for _, meet := range calendar[name] {
				if meet.day == day {
					list = append(list, meet)
				}
			}

			sort.Slice(list, func(i, j int) bool {
				return list[i].start < list[j].start
			})

			for _, meet := range list {
				fmt.Fprint(writer, meet.timeStr, " ", meet.duration)
				for _, p := range meet.names {
					fmt.Fprint(writer, " ", p)
				}
				fmt.Fprintln(writer)
			}
		}
	}
}
