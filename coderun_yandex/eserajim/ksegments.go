package eserajim

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func validateKSegmentsNInput(p int) bool {
	return p >= 1 && p <= 100_000
}

func validateKSegmentsKInput(k, n int) bool {
	return k >= 1 && k <= n
}

func can(x []int, k int, L int) bool {
	n := len(x)
	cnt := 0
	i := 0

	for i < n {
		cnt++
		end := x[i] + L
		i++
		for i < n && x[i] <= end {
			i++
		}
	}

	return cnt <= k
}

// https://coderun.yandex.ru/selections/eserajim/problems/k-segments
// KSegments - problem 5
func KSegments() {
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
	if !validateKSegmentsNInput(n) {
		panic("number N out of range")
	}

	k, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	if !validateKSegmentsKInput(k, n) {
		panic("number K out of range")
	}

	// points input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNum = strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	// points
	points := make([]int, n)
	for i, v := range strNum {
		x, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		points[i] = x
	}

	sort.Ints(points)

	left, right := 0, points[n-1]-points[0]
	ans := right
	for left <= right {
		mid := (left + right) / 2
		if can(points, k, mid) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
