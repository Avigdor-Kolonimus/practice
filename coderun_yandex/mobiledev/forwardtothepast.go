package mobiledev

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	layout = "02.01.2006"
)

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/forward-to-the-past
// ForwardToThePast - problem 15
func ForwardToThePast() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)

	start := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC)

	fmt.Fprintf(writer, "? %s\n", end.Format(layout))
	writer.Flush()

	var response string
	var total int
	fmt.Fscan(reader, &response)

	if response == "!" {
		return
	}
	fmt.Fscan(reader, &total)

	targetRank := (total + 1) / 2

	low, high := 0, int(end.Sub(start).Hours()/24)
	for low < high {
		mid := low + (high-low)/2
		date := start.AddDate(0, 0, mid)

		fmt.Fprintf(writer, "? %s\n", date.Format(layout))
		writer.Flush()

		fmt.Fscan(reader, &response)
		if response == "!" {
			return
		}

		var k int
		fmt.Fscan(reader, &k)

		if k >= targetRank {
			high = mid
		} else {
			low = mid + 1
		}
	}

	answer := start.AddDate(0, 0, low)
	fmt.Fprintf(writer, "! %s\n", answer.Format(layout))
	writer.Flush()
}
