package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	DayMinutes = 1440
)

// https://coderun.yandex.ru/problem/buying-tickets
// BuyingTickets - problem 37
func BuyingTickets() {
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

	diff := make([]int, DayMinutes+1)
	for i := 0; i < n; i++ {
		// input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		parameters := strings.Fields(line)
		cnt := len(parameters)
		if cnt != 4 {
			panic("input does not match 4")
		}

		h1, err := strconv.Atoi(parameters[0])
		if err != nil {
			panic(err)
		}
		m1, err := strconv.Atoi(parameters[1])
		if err != nil {
			panic(err)
		}
		h2, err := strconv.Atoi(parameters[2])
		if err != nil {
			panic(err)
		}
		m2, err := strconv.Atoi(parameters[3])
		if err != nil {
			panic(err)
		}

		start := h1*60 + m1
		end := h2*60 + m2

		// 24/7
		if start == end {
			diff[0]++
			diff[DayMinutes]--
			continue
		}

		if start < end {
			// [start, end)
			diff[start]++
			diff[end]--
		} else {
			// after noon

			// [start, 1440)
			diff[start]++
			diff[DayMinutes]--

			// [0, end)
			diff[0]++
			diff[end]--
		}
	}

	current := 0
	answer := 0
	for minute := 0; minute < DayMinutes; minute++ {
		current += diff[minute]

		if current == n {
			answer++
		}
	}

	writer.WriteString(strconv.Itoa(answer))
	writer.WriteByte('\n')
}
