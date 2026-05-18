package mgustokashin

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type PointBuildASquare struct {
	x, y int
}

// https://coderun.yandex.ru/selections/mgustokashin/problems/build-a-square
// BuildASquare - problem 7
func BuildASquare() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// coordinate inputs
	points := make([]PointBuildASquare, n)
	exists := make(map[PointBuildASquare]bool)
	for i := range n {
		line, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum := strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		x, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		points[i].x = x
		points[i].y = y
		exists[points[i]] = true
	}

	best := []PointBuildASquare{
		{0, 0},
		{0, 0},
	}
	bestCnt := 2
	if n == 1 {
		writer.WriteString(strconv.Itoa(3))
		writer.WriteByte('\n')
		x := points[0].x
		y := points[0].y

		writer.WriteString(strconv.Itoa(x+1) + " " + strconv.Itoa(y))
		writer.WriteByte('\n')
		writer.WriteString(strconv.Itoa(x) + " " + strconv.Itoa(y+1))
		writer.WriteByte('\n')
		writer.WriteString(strconv.Itoa(x+1) + " " + strconv.Itoa(y+1))
		writer.WriteByte('\n')

		return
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {

			x1, y1 := points[i].x, points[i].y
			x2, y2 := points[j].x, points[j].y

			dx := x2 - x1
			dy := y2 - y1

			// first square
			c1 := PointBuildASquare{x1 - dy, y1 + dx}
			d1 := PointBuildASquare{x2 - dy, y2 + dx}

			cnt := 0
			missing := []PointBuildASquare{}

			if !exists[c1] {
				cnt++
				missing = append(missing, c1)
			}

			if !exists[d1] {
				cnt++
				missing = append(missing, d1)
			}

			if cnt < bestCnt {
				bestCnt = cnt
				copy(best, missing)
			}

			// second square
			c2 := PointBuildASquare{x1 + dy, y1 - dx}
			d2 := PointBuildASquare{x2 + dy, y2 - dx}

			cnt = 0
			missing = []PointBuildASquare{}

			if !exists[c2] {
				cnt++
				missing = append(missing, c2)
			}

			if !exists[d2] {
				cnt++
				missing = append(missing, d2)
			}

			if cnt < bestCnt {
				bestCnt = cnt
				copy(best, missing)
			}
		}
	}

	writer.WriteString(strconv.Itoa(bestCnt))
	writer.WriteByte('\n')

	for i := 0; i < bestCnt; i++ {
		writer.WriteString(strconv.Itoa(best[i].x) + " " + strconv.Itoa(best[i].y))
		writer.WriteByte('\n')
	}
}
