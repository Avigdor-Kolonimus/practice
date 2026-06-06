package mgustokashin

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type PointRunManhattan struct {
	x int
	y int
}

// https://coderun.yandex.ru/selections/mgustokashin/problems/run-manhattan
// RunManhattan - problem 2
func RunManhattan() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// T, D and N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)

	t, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	d, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	n, err := strconv.Atoi(strNum[2])
	if err != nil {
		panic(err)
	}

	// coordinates input
	uMin, uMax := 0, 0
	vMin, vMax := 0, 0
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)

		x, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		uMin -= t
		uMax += t
		vMin -= t
		vMax += t

		u := x + y
		v := x - y

		uMin = max(uMin, u-d)
		uMax = min(uMax, u+d)

		vMin = max(vMin, v-d)
		vMax = min(vMax, v+d)
	}

	var ans []PointRunManhattan
	for u := uMin; u <= uMax; u++ {
		for v := vMin; v <= vMax; v++ {
			if (u+v)%2 != 0 {
				continue
			}

			x := (u + v) / 2
			y := (u - v) / 2

			ans = append(ans, PointRunManhattan{x, y})
		}
	}

	writer.WriteString(strconv.Itoa(len(ans)))
	writer.WriteByte('\n')
	for _, p := range ans {
		writer.WriteString(fmt.Sprintf("%d %d", p.x, p.y))
		writer.WriteByte('\n')
	}
}
