package devgointerview

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/dev-go-interview/problems/reserve-periods
// ReservePeriods - problem 10
func ReservePeriods() {
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

	// Ai input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum = strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	a := make([]int, n)
	for i := range n {
		a[i], err = strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}
	}

	// If k == 0, every subarray is valid.
	if k == 0 {
		ans := n*n + 1/2
		writer.WriteString(strconv.Itoa(ans))
		writer.WriteByte('\n')

		return
	}

	ans, sum, l := 0, 0, 0
	for r := range n {
		sum += a[r]

		// Move the left boundary while the sum is still >= k.
		for l <= r && sum-a[l] >= k {
			sum -= a[l]
			l++
		}

		// If [l, r] has sum >= k, then
		// [0, r], [1, r], ..., [l, r] are valid.
		if sum >= k {
			ans += l + 1
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
