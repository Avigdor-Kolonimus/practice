package problems

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Date struct {
	day   int
	month int
	year  int
}

type Person struct {
	id    int
	start int
	end   int
}

func isLeap(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

func daysInMonth(month, year int) int {
	days := [...]int{
		0, 31, 28, 31, 30, 31,
		30, 31, 31, 30, 31, 30, 31,
	}

	if month == 2 && isLeap(year) {
		return 29
	}

	return days[month]
}

func toDays(d Date) int {
	result := 0

	for year := 1; year < d.year; year++ {
		if isLeap(year) {
			result += 366
		} else {
			result += 365
		}
	}

	for month := 1; month < d.month; month++ {
		result += daysInMonth(month, d.year)
	}

	return result + d.day - 1
}

func addYears(d Date, years int) Date {
	d.year += years

	// february 29th -> february 28th in non-leap years
	if d.month == 2 && d.day == 29 && !isLeap(d.year) {
		d.day = 28
	}

	return d
}

// https://coderun.yandex.ru/problem/coevals
// Coevals - problem 237
func Coevals() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	// People input
	people := make([]Person, 0, n)
	for id := 1; id <= n; id++ {
		var birth, death Date

		fmt.Fscan(
			reader,
			&birth.day,
			&birth.month,
			&birth.year,
			&death.day,
			&death.month,
			&death.year,
		)

		// At the age of 18, a person begins to participate
		startDate := addYears(birth, 18)

		// Participation is no longer allowed on the day of death.
		// Participation is also not allowed on the day of one's 80th birthday.
		deathDay := toDays(death)
		end80 := toDays(addYears(birth, 80))

		end := deathDay
		if end80 < end {
			end = end80
		}

		start := toDays(startDate)

		if start < end {
			people = append(people, Person{
				id:    id,
				start: start,
				end:   end,
			})
		}
	}

	if len(people) == 0 {
		writer.WriteString("0\n")
		return
	}

	for i := 1; i < len(people); i++ {
		j := i

		for j > 0 && people[j].start < people[j-1].start {
			people[j], people[j-1] = people[j-1], people[j]
			j--
		}
	}

	active := make(map[int]bool)

	ends := make(map[int]int)
	printed := false
	i := 0
	for i < len(people) {
		currentStart := people[i].start

		for id, end := range ends {
			if end <= currentStart {
				delete(ends, id)
				delete(active, id)
			}
		}

		j := i
		for j < len(people) && people[j].start == currentStart {
			id := people[j].id

			active[id] = true
			ends[id] = people[j].end

			j++
		}

		nextStart := -1
		if j < len(people) {
			nextStart = people[j].start
		}

		minEnd := -1

		for _, end := range ends {
			if minEnd == -1 || end < minEnd {
				minEnd = end
			}
		}

		maximal := false
		if len(active) > 0 {
			if nextStart == -1 || minEnd <= nextStart {
				maximal = true
			}
		}

		if maximal {
			ids := make([]int, 0, len(active))

			for id := range active {
				ids = append(ids, id)
			}

			sort.Ints(ids)

			for k, id := range ids {
				if k > 0 {
					writer.WriteByte(' ')
				}

				writer.WriteString(strconv.Itoa(id))
			}

			writer.WriteByte('\n')

			printed = true
		}

		i = j
	}

	if !printed {
		writer.WriteString("0\n")
	}
}
