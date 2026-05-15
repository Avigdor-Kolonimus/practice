package algorithmtrainingseptember2025

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	MAXV = 10_000
)

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/superstring-gravity
// SuperstringGravity - assignment 27
func SuperstringGravity() {
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

	// A input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	a := make([]int, n+1)
	for i := range n {
		x, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		a[i+1] = x
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

	// B input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNum = strings.Fields(line)
	if len(strNum) != m {
		panic("numbers count does not match m")
	}

	cnt := make([]int64, MAXV+1)
	sumIdx := make([]int64, MAXV+1)
	for i := range m {
		x, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		cnt[x]++
		sumIdx[x] += int64(i + 1)
	}

	// prefix arrays
	pc := make([]int64, MAXV+1)
	ps := make([]int64, MAXV+1)

	pic := make([]int64, MAXV+1)
	pis := make([]int64, MAXV+1)

	for v := 1; v <= MAXV; v++ {
		pc[v] = pc[v-1] + cnt[v]
		ps[v] = ps[v-1] + cnt[v]*int64(v)

		pic[v] = pic[v-1] + sumIdx[v]
		pis[v] = pis[v-1] + sumIdx[v]*int64(v)
	}

	totalCnt := pc[MAXV]
	totalSum := ps[MAXV]

	totalIdxCnt := pic[MAXV]
	totalIdxVal := pis[MAXV]

	ans := int64(0)
	for i := 1; i <= n; i++ {
		x := a[i]

		leftCnt := pc[x]
		leftSum := ps[x]

		rightCnt := totalCnt - leftCnt
		rightSum := totalSum - leftSum

		sumAbs :=
			int64(x)*leftCnt - leftSum +
				rightSum - int64(x)*rightCnt

		leftIdxCnt := pic[x]
		leftIdxVal := pis[x]

		rightIdxCnt := totalIdxCnt - leftIdxCnt
		rightIdxVal := totalIdxVal - leftIdxVal

		sumJAbs :=
			int64(x)*leftIdxCnt - leftIdxVal +
				rightIdxVal - int64(x)*rightIdxCnt

		ans += int64(i)*sumAbs - sumJAbs
	}

	writer.WriteString(strconv.FormatInt(ans, 10))
	writer.WriteByte('\n')
}
