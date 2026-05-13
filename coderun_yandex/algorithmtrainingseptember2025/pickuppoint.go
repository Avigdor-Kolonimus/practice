package algorithmtrainingseptember2025

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func dfsPickUpPoint(
	v, p int,
	a []int,
	g [][]int,
	sub []int,
	total int,
	bestVal *int,
	bestV *int,
) {
	var mx int = 0

	sub[v] = a[v]
	for _, to := range g[v] {
		if to == p {
			continue
		}

		dfsPickUpPoint(to, v, a, g, sub, total, bestVal, bestV)

		sub[v] += sub[to]
		mx = max(mx, sub[to])
	}

	// component above
	up := total - sub[v]
	mx = max(mx, up)

	// residents of the current square
	mx = max(mx, a[v])

	if mx < *bestVal {
		*bestVal = mx
		*bestV = v
	}
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/pick-up-point
// PickUpPoint - assignment 28
func PickUpPoint() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// square input
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	total := 0
	square := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a, err := strconv.Atoi(strNum[i-1])
		if err != nil {
			panic(err)
		}

		square[i] = a
		total += a
	}

	// road input
	g := make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		u, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		v, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	bestVal := int(^uint(0) >> 1)
	bestV := 1
	sub := make([]int, n+1)

	dfsPickUpPoint(1, 0, square, g, sub, total, &bestVal, &bestV)

	writer.WriteString(strconv.Itoa(bestV))
	writer.WriteByte('\n')
}
