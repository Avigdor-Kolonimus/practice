package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func checkAverageInWindow(a []int, k int, mid float64) bool {
	n := len(a)

	pref := make([]float64, n+1)
	for i := 0; i < n; i++ {
		pref[i+1] = pref[i] + float64(a[i]) - mid
	}

	minPref := 0.0
	for r := k; r <= n; r++ {
		if pref[r]-minPref >= 0 {
			return true
		}

		if pref[r-k+1] < minPref {
			minPref = pref[r-k+1]
		}
	}

	return false
}

// https://coderun.yandex.ru/problem/average-in-window
// AverageInWindow - problem 443
func AverageInWindow() {
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

	// array input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum = strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	maxA := 0
	a := make([]int, n)
	for i := 0; i < n; i++ {
		x, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		a[i] = x
		if a[i] > maxA {
			maxA = a[i]
		}
	}

	l := 0.0
	r := float64(maxA)
	for iter := 0; iter < 70; iter++ {
		mid := (l + r) / 2

		if checkAverageInWindow(a, k, mid) {
			l = mid
		} else {
			r = mid
		}
	}

	writer.WriteString(strconv.FormatFloat(l, 'f', 6, 64))
	writer.WriteByte('\n')
}
