package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/rectangle-sum
// RectangleSum - problem 95
func RectangleSum() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N, M and K line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 3 {
		panic("numbers count does not match 3")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	k, err := strconv.Atoi(strNum[2])
	if err != nil {
		panic(err)
	}

	// prefex
	pref := make([][]int, n+1)
	for i := range pref {
		pref[i] = make([]int, m+1)
	}

	// rectangle inputs
	rectangle := make([][]int, n)
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

		rectangle[i] = make([]int, m)
		for j := 0; j < m; j++ {
			rectangle[i][j], err = strconv.Atoi(strNum[j])
			if err != nil {
				panic(err)
			}

			pref[i+1][j+1] = rectangle[i][j] +
				pref[i][j+1] +
				pref[i+1][j] -
				pref[i][j]
		}
	}

	// queries
	for ; k > 0; k-- {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
		if len(strNum) != 4 {
			panic("numbers count does not match 4")
		}

		x1, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		y1, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}
		x2, err := strconv.Atoi(strNum[2])
		if err != nil {
			panic(err)
		}
		y2, err := strconv.Atoi(strNum[3])
		if err != nil {
			panic(err)
		}

		res := pref[x2][y2] -
			pref[x1-1][y2] -
			pref[x2][y1-1] +
			pref[x1-1][y1-1]

		writer.WriteString(strconv.Itoa(res))
		writer.WriteByte('\n')
	}
}
