package problems

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://coderun.yandex.ru/problem/sasha-and-circles
// SashaAndCircles - problem 568
func SashaAndCircles() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var h [4]Point
	var totalX, totalY int64

	for i := 0; i < 4; i++ {
		fmt.Fscan(reader, &h[i].x, &h[i].y)
		totalX += h[i].x
		totalY += h[i].y
	}

	var p [4]Point
	for i := 0; i < 4; i++ {
		p[i].x = totalX - 3*h[i].x
		p[i].y = totalY - 3*h[i].y
	}

	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			if p[i] == p[j] {
				fmt.Fprintln(writer, "NO")
				return
			}
		}
	}

	dist2 := func(p Point) int64 {
		return p.x*p.x + p.y*p.y
	}

	r2 := dist2(p[0])

	if r2 == 0 {
		fmt.Fprintln(writer, "NO")
		return
	}

	for i := 1; i < 4; i++ {
		if dist2(p[i]) != r2 {
			fmt.Fprintln(writer, "NO")
			return
		}
	}

	fmt.Fprintln(writer, "YES")

	for i := 0; i < 4; i++ {
		x := float64(p[i].x) / 3.0
		y := float64(p[i].y) / 3.0

		writer.WriteString(strconv.FormatFloat(x, 'f', 12, 64))
		writer.WriteByte(' ')
		writer.WriteString(strconv.FormatFloat(y, 'f', 12, 64))
		writer.WriteByte('\n')
	}
}
