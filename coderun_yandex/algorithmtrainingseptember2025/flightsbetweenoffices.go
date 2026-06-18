package algorithmtrainingseptember2025

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func parseTime(s string) int {
	h := int(s[0]-'0')*10 + int(s[1]-'0')
	m := int(s[3]-'0')*10 + int(s[4]-'0')

	return h*60 + m
}

func needBuses(arr, dep []int) int {
	cur := 0
	minCur := 0
	for t := 0; t < 1440; t++ {
		cur += arr[t]
		cur -= dep[t]

		if cur < minCur {
			minCur = cur
		}
	}

	return -minCur
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/flights-between-offices
// FlightsBetweenOffices - assignment 31
func FlightsBetweenOffices() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// office 1 input
	depFrom1 := make([]int, 1440)
	arrTo2 := make([]int, 1440)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		trip := strings.TrimRight(line, "\r\n")

		dep := parseTime(trip[:5])
		arr := parseTime(trip[6:])

		depFrom1[dep]++
		arrTo2[arr]++
	}

	// M input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	m, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// office 2 input
	depFrom2 := make([]int, 1440)
	arrTo1 := make([]int, 1440)
	for i := 0; i < m; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		trip := strings.TrimRight(line, "\r\n")

		dep := parseTime(trip[:5])
		arr := parseTime(trip[6:])

		depFrom2[dep]++
		arrTo1[arr]++
	}

	// Office 1:
	// arrivals = 2 -> 1 trips
	// departures = 1 -> 2 trips
	need1 := needBuses(arrTo1, depFrom1)

	// Office 2:
	// arrivals = 1 -> 2 trips
	// departures = 2 -> 1 trips
	need2 := needBuses(arrTo2, depFrom2)

	writer.WriteString(strconv.Itoa(need1 + need2))
	writer.WriteByte('\n')
}
