package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type PointTriangle struct {
	x int
	y int
}

type VecTriangle struct {
	x int
	y int
}

// https://coderun.yandex.ru/problem/triangles
// Triangles - problem 221
func Triangles() {
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

	// points input
	p := make([]PointTriangle, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum := strings.Fields(line)

		x, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		p[i] = PointTriangle{x, y}
	}

	ans := 0
	for i := 0; i < n; i++ {
		distCnt := make(map[int]int)

		type Info struct {
			dist int
			dx   int
			dy   int
		}

		arr := make([]Info, 0, n-1)

		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			dx := p[j].x - p[i].x
			dy := p[j].y - p[i].y

			d := dx*dx + dy*dy

			distCnt[d]++

			arr = append(arr, Info{
				dist: d,
				dx:   dx,
				dy:   dy,
			})
		}

		for _, cnt := range distCnt {
			ans += (cnt * (cnt - 1)) / 2
		}

		vecs := make(map[VecTriangle]struct{}, len(arr))

		for _, v := range arr {
			vecs[VecTriangle{v.dx, v.dy}] = struct{}{}
		}

		degenerate := 0
		for _, v := range arr {
			if _, ok := vecs[VecTriangle{-v.dx, -v.dy}]; ok {
				degenerate++
			}
		}

		ans -= degenerate / 2
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
