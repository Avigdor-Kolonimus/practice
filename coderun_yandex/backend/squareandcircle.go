package backend

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	x float64
	y float64
}

func quarterSquare(xb, yb, r float64) float64 {
	x1 := 0.0
	x2 := r
	r2 := r * r
	square := 0.0

	if yb < r {
		x3 := math.Sqrt(r2 - yb*yb)
		if x3 >= xb {
			return xb * yb
		}
		x1 = x3
		square = yb * x1
	}

	if xb < x2 {
		x2 = xb
	}

	result := 0.5*(x2*math.Sqrt(r2-x2*x2)+r2*math.Asin(x2/r)-x1*math.Sqrt(r2-x1*x1)-r2*math.Asin(x1/r)) + square

	return result
}

func upperHalfSquare(x0, y0, r float64) float64 {
	return quarterSquare(1-x0, 1-y0, r) +
		quarterSquare(x0, 1-y0, r)
}

func lowerHalfSquare(x0, y0, r float64) float64 {
	return quarterSquare(1-x0, y0, r) +
		quarterSquare(x0, y0, r)
}

// https://coderun.yandex.ru/selections/backend/problems/square-and-circle
// SquareAndCircle - problem 11
func SquareAndCircle() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var r float64
	fmt.Fscan(in, &n, &r)

	ans := 0.0

	for i := 0; i < n; i++ {
		var x, y float64
		fmt.Fscan(in, &x, &y)
		ans += upperHalfSquare(x, y, r)
		ans += lowerHalfSquare(x, y, r)
	}

	fmt.Printf("%.10f\n", ans)
}
