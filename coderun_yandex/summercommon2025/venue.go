package summercommon2025

import (
	"bufio"
	"os"
	"strconv"
)

type Point struct {
	X int
	Y int
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateAnswer(_ int, points []Point) (int, int) {
	dx, dy := points[0].X, points[0].Y
	k := 0
	for i := range points {
		points[i].X -= dx
		points[i].Y -= dy
		k = max(k, abs(points[i].X), abs(points[i].Y))
	}

	for i := range points {
		if points[i].X != 0 {
			k = gcd(k, abs(points[i].X))
		}
		if points[i].Y != 0 {
			k = gcd(k, abs(points[i].Y))
		}
		if k == 1 {
			break
		}
	}
	if k > 1 {
		for i := range points {
			points[i].X /= k
			points[i].Y /= k
		}
	}

	oddOddIdx := -1
	for i := range points {
		if abs(points[i].X)%2+abs(points[i].Y)%2 == 1 {
			return 1, i + 1
		}
		if points[i].X%2 != 0 && points[i].Y%2 != 0 {
			oddOddIdx = i
		}
	}

	return 1, oddOddIdx + 1
}

// ввод/вывод
// не изменяйте сигнатуру метода
// https://coderun.yandex.ru/selections/2025-summer-common/problems/venue
// Venue - problem 21
func Venue() {
	input := NewFastScanner(os.Stdin)

	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	n := input.readInt()

	points := make([]Point, n)
	for i := 0; i < n; i++ {
		x := input.readInt()
		y := input.readInt()
		points[i] = Point{X: x, Y: y}
	}

	first, second := calculateAnswer(n, points)
	output.WriteString(strconv.Itoa(first))
	output.WriteByte(' ')
	output.WriteString(strconv.Itoa(second))
	output.WriteByte('\n')
}
