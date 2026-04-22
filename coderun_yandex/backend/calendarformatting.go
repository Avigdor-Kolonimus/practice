package backend

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func validateCalendarFormattingInput(n int) bool {
	return n >= 28 && n <= 31
}

// https://coderun.yandex.ru/selections/backend/problems/calendar-formatting
// CalendarFormatting - problem 43
func CalendarFormatting() {
	dayOfWeek := map[string]int{"Monday": 1, "Tuesday": 2, "Wednesday": 3, "Thursday": 4, "Friday": 5, "Saturday": 6, "Sunday": 7}
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// days and weekday input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	days, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	if !validateCalendarFormattingInput(days) {
		panic("number nDays out of range")
	}

	weekday := strNum[1]

	day := 1
	start := dayOfWeek[weekday]
	for i := 1; day <= days; i++ {
		if i < start {
			writer.WriteString(".. ")
			continue
		}

		if day < 10 {
			writer.WriteString(".")
		}
		writer.WriteString(strconv.Itoa(day))
		writer.WriteByte(' ')

		if i%7 == 0 {
			writer.WriteByte('\n')
		}
		day++
	}
}
