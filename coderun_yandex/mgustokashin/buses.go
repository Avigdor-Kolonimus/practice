package mgustokashin

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func toMin(s string) int {
	return int(s[0]-'0')*600 +
		int(s[1]-'0')*60 +
		int(s[3]-'0')*10 +
		int(s[4]-'0')
}

// https://coderun.yandex.ru/selections/mgustokashin/problems/buses
// Buses - problem 5
func Buses() {
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

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// route input
	stock := make([]int, n+1)
	inDeg := make([]int, n+1)
	outDeg := make([]int, n+1)
	buckets := make([][]int, 4320)
	for i := 0; i < m; i++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)

		f, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		g, err := strconv.Atoi(strNum[2])
		if err != nil {
			panic(err)
		}

		xs := strNum[1]
		ys := strNum[3]

		t := toMin(xs)
		y := toMin(ys)

		dt := (y - t + 1440) % 1440
		a := t + dt // 0..2879

		outDeg[f]++
		inDeg[g]++

		// day 0
		buckets[t] = append(buckets[t], -f)
		buckets[a] = append(buckets[a], g)

		// day 1
		buckets[t+1440] = append(buckets[t+1440], -f)
		buckets[a+1440] = append(buckets[a+1440], g)
	}

	for city := 1; city <= n; city++ {
		if inDeg[city] != outDeg[city] {
			writer.WriteString(strconv.Itoa(-1))
			writer.WriteByte('\n')

			return
		}
	}

	ans := 0
	for time := 0; time < 4320; time++ {
		// input
		for _, id := range buckets[time] {
			if id > 0 {
				stock[id]++
			}
		}

		// output
		for _, id := range buckets[time] {
			if id < 0 {
				city := -id

				if stock[city] > 0 {
					stock[city]--
				} else {
					ans++
				}
			}
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
