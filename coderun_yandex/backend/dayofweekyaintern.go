package backend

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	days   = [7]string{"Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	months = map[string]int{"January": 13, "February": 14,
		"March": 3, "April": 4, "May": 5, "June": 6, "July": 7, "August": 8, "September": 9, "October": 10, "November": 11, "December": 12}
)

func zellersCongruence(day, month, year int) string {
	if month == 13 || month == 14 {
		year--
	}

	K := year % 100
	J := year / 100

	h := (day + (13*(month+1))/5 + K + K/4 + J/4 + 5*J) % 7

	return days[h]
}

func validateDayofweekYaInternYearInput(n int) bool {
	return n >= 1980 && n <= 2100
}

// https://coderun.yandex.ru/selections/backend/problems/dayofweek-ya-intern
// DayofweekYaIntern - problem 32
func DayofweekYaIntern() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var results []string
	for {
		// Day input
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimRight(line, "\r\n")

		strNum := strings.Fields(line)
		if len(strNum) != 3 {
			panic("numbers count does not match 3")
		}

		day, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}

		month := strNum[1]

		year, err := strconv.Atoi(strNum[2])
		if err != nil {
			panic(err)
		}
		if !validateDayofweekYaInternYearInput(year) {
			panic("number Year out of range")
		}

		results = append(results, zellersCongruence(day, months[month], year))
	}

	for _, res := range results {
		writer.WriteString(res)
		writer.WriteByte('\n')
	}
}
