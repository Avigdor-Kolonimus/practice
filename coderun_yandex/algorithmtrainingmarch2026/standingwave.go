package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type EventStandingWave struct {
	l int
	r int
	x int
}

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/standing-wave
// StandingWave - assignment 2
func StandingWave() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// standing wave input
	events := make([]EventStandingWave, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
		if len(strNum) != 3 {
			panic("numbers count does not match 3")
		}

		l, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}
		x, err := strconv.Atoi(strNum[2])
		if err != nil {
			panic(err)
		}

		events[i] = EventStandingWave{l, r, x}

	}

	// request input
	for i := 0; i < m; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		q, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		ans := 0
		for _, e := range events {
			if q < e.l || q > e.r {
				continue
			}

			if (q-e.l)%2 == 0 {
				ans += e.x
			} else {
				ans -= e.x
			}
		}

		writer.WriteString(strconv.Itoa(ans))
		writer.WriteByte('\n')
	}
}
