package first2023mobiledev

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/flight
// Flight - problem 27
func Flight() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// Departure time input
	line, _ := reader.ReadString('\n')
	departureTime := strings.Split(strings.TrimSpace(line), ":")

	// Arrival time input
	line, _ = reader.ReadString('\n')
	arrivalTime := strings.Split(strings.TrimSpace(line), ":")

	// Time difference input
	line, _ = reader.ReadString('\n')
	timeDiff, _ := strconv.Atoi(strings.TrimSpace(line))

	depH, _ := strconv.Atoi(departureTime[0])
	depM, _ := strconv.Atoi(departureTime[1])
	arrH, _ := strconv.Atoi(arrivalTime[0])
	arrM, _ := strconv.Atoi(arrivalTime[1])

	depTotalMin := depH*60 + depM
	arrTotalMin := arrH*60 + arrM

	day := 24 * 60

	arrMinTimeDiff := (arrTotalMin - timeDiff*60 + day) % day
	dur := (arrMinTimeDiff - depTotalMin + day) % day

	hours := dur / 60
	minutes := dur % 60

	if dur == 0 && depTotalMin != arrMinTimeDiff {
		fmt.Fprintln(writer, "24:00")
	} else {
		fmt.Fprintf(writer, "%d:%02d\n", hours, minutes)
	}
}
