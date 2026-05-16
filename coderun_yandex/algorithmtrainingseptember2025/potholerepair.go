package algorithmtrainingseptember2025

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Road struct {
	cnt int
	a   int
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/pothole-repair
// PotholeRepair - assignment 35
func PotholeRepair() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N, M and K input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 3 {
		panic("numbers count does not match 3")
	}

	// N
	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}

	// D
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// K
	k, err := strconv.Atoi(strNum[2])
	if err != nil {
		panic(err)
	}

	// potholes input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum = strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	potholes := make([]int, n)
	for i := 0; i < n; i++ {
		a, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		potholes[i] = a
	}

	// routes input
	diff := make([]int, n+1)
	for i := 0; i < m; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		l, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		l--
		r--

		diff[l]++
		if r+1 < n {
			diff[r+1]--
		}

		cur := 0
		c := make([]int, n)
		for i := 0; i < n; i++ {
			cur += diff[i]
			c[i] = cur
		}
	}

	cur := 0
	c := make([]int, n)
	for i := 0; i < n; i++ {
		cur += diff[i]
		c[i] = cur
	}

	total := 0
	roads := make([]Road, n)
	for i := 0; i < n; i++ {
		total += potholes[i] * c[i]
		roads[i] = Road{cnt: c[i], a: potholes[i]}
	}

	sort.Slice(roads, func(i, j int) bool {
		return roads[i].cnt > roads[j].cnt
	})

	reduction := 0
	for _, road := range roads {
		if k == 0 {
			break
		}

		x := road.a
		if x > k {
			x = k
		}

		reduction += x * road.cnt
		k -= x
	}

	writer.WriteString(strconv.Itoa(total - reduction))
	writer.WriteByte('\n')
}
