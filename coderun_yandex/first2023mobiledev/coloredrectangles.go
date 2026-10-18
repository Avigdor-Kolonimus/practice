package first2023mobiledev

import (
	"bufio"
	"fmt"
	"os"
)

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/colored-rectangles
// ColoredRectangles - problem 13
func ColoredRectangles() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	type Rect struct {
		x1, y1, x2, y2 int
	}

	rects := make([]Rect, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader,
			&rects[i].x1, &rects[i].y1,
			&rects[i].x2, &rects[i].y2)
	}

	const width = 1280
	const height = 1920

	covered := make([][]bool, width)
	for x := 0; x < width; x++ {
		covered[x] = make([]bool, height)
	}

	answer := 0

	for i := n - 1; i >= 0; i-- {
		r := rects[i]
		visible := false

		for x := r.x1; x <= r.x2; x++ {
			for y := r.y1; y <= r.y2; y++ {
				if !covered[x][y] {
					visible = true
				}
				covered[x][y] = true
			}
		}

		if visible {
			answer++
		}
	}

	fmt.Fprintln(writer, answer)
}
