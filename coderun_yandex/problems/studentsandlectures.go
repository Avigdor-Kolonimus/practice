package problems

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/students-and-lectures
// StudentsAndLectures - problem 48
func StudentsAndLectures() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M line
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

	// pleasures input
	pleasures := make([][]int, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
		if len(strNum) != m {
			panic("numbers count does not match m")
		}

		pleasures[i] = make([]int, m)
		for j := 0; j < m; j++ {
			pleasures[i][j], err = strconv.Atoi(strNum[j])
			if err != nil {
				panic(err)
			}
		}
	}

	maxDelta := make([]int, m)
	for j := 0; j < m; j++ {
		maxDelta[j] = math.MinInt
	}

	base := make([]int, m)
	sumPos := make([]int, m)

	// displeasure input
	displeasure := make([][]int, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
		if len(strNum) != m {
			panic("numbers count does not match m")
		}

		displeasure[i] = make([]int, m)
		for j := 0; j < m; j++ {
			displeasure[i][j], err = strconv.Atoi(strNum[j])
			if err != nil {
				panic(err)
			}

			base[j] += displeasure[i][j]
			delta := pleasures[i][j] - displeasure[i][j]

			if delta > 0 {
				sumPos[j] += delta
			}
			if delta > maxDelta[j] {
				maxDelta[j] = delta
			}
		}
	}

	ans := 0
	for j := 0; j < m; j++ {
		ans += base[j]
		if sumPos[j] > 0 {
			ans += sumPos[j]
		} else {
			ans += maxDelta[j]
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
