package devgointerview

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/dev-go-interview/problems/income-during-rush-hour
// IncomeDuringRushHour - problem 6
func IncomeDuringRushHour() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and K input
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
	k, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// time and amount input
	time := make([]int, n)
	amount := make([]int, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		t, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		a, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		time[i] = t
		amount[i] = a
	}

	cur, ans, l := 0, 0, 0

	for r := 0; r < n; r++ {
		cur += amount[r]

		for time[r]-time[l] >= k {
			cur -= amount[l]
			l++
		}

		if cur > ans {
			ans = cur
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
