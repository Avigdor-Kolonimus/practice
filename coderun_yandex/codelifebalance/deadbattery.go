package codelifebalance

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func validateDeadBatteryNInput(p int) bool {
	return p >= 1 && p <= 100_000
}

func validateDeadBatteryAInput(p int) bool {
	return p >= 1 && p <= 100
}

// https://coderun.yandex.ru/selections/code-life-balance/problems/dead-battery
// DeadBattery - assignment 3
func DeadBattery() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateDeadBatteryNInput(n) {
		panic("number N out of range")
	}

	// battery input
	bu := make([]int, n)
	appUsage := 0
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		a, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		if !validateDeadBatteryAInput(a) {
			panic("number Ai out of range")
		}

		bu[i] = a
		appUsage += a
	}

	writer.WriteString(strconv.Itoa(100 / appUsage))
	writer.WriteByte('\n')
}
