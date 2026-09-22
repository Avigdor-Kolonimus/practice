package backend

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type Poly struct {
	a, b, c float64
}

func (p Poly) integral(x float64) float64 {
	return p.a*x*x*x/3 + p.b*x*x/2 + p.c*x
}

func value(p Poly, x float64) float64 {
	return p.a*x*x + p.b*x + p.c
}

// Integral |p(x)| between [l, r].
func absIntegral(p Poly, l, r float64) float64 {
	if l >= r {
		return 0
	}

	if p.a == 0 {
		if p.b == 0 {
			return math.Abs(p.c) * (r - l)
		}

		root := -p.c / p.b

		if root <= l || root >= r {
			return math.Abs(p.integral(r) - p.integral(l))
		}

		return math.Abs(p.integral(root)-p.integral(l)) +
			math.Abs(p.integral(r)-p.integral(root))
	}

	d := p.b*p.b - 4*p.a*p.c

	if d <= 0 {
		mid := (l + r) / 2
		v := value(p, mid)

		res := p.integral(r) - p.integral(l)

		if v < 0 {
			res = -res
		}

		return res
	}

	sqrtD := math.Sqrt(d)

	q := -0.5 * (p.b + math.Copysign(sqrtD, p.b))

	x1 := q / p.a
	x2 := p.c / q

	if x1 > x2 {
		x1, x2 = x2, x1
	}

	roots := make([]float64, 0, 2)

	if x1 > l && x1 < r {
		roots = append(roots, x1)
	}

	if x2 > l && x2 < r {
		roots = append(roots, x2)
	}

	if len(roots) == 0 {
		mid := (l + r) / 2
		res := p.integral(r) - p.integral(l)

		if value(p, mid) < 0 {
			res = -res
		}

		return res
	}

	res := 0.0
	prev := l
	for _, root := range roots {
		mid := (prev + root) / 2

		part := p.integral(root) - p.integral(prev)

		if value(p, mid) < 0 {
			part = -part
		}

		res += part
		prev = root
	}

	mid := (prev + r) / 2

	part := p.integral(r) - p.integral(prev)

	if value(p, mid) < 0 {
		part = -part
	}

	res += part

	return res
}

func readInts(reader *bufio.Reader) []int64 {
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	res := make([]int64, len(parts))

	for i, s := range parts {
		res[i], _ = strconv.ParseInt(s, 10, 64)
	}

	return res
}

// https://coderun.yandex.ru/selections/first-2023-backend/problems/area-between-curves
// AreaBetweenCurves - problem 12
func AreaBetweenCurves() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// n and m input
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	// f input
	lInt := readInts(reader)
	l := make([]float64, n+1)
	for i := 0; i <= n; i++ {
		l[i] = float64(lInt[i])
	}

	f := make([]Poly, n)
	for i := 0; i < n; i++ {
		parts := readInts(reader)

		f[i] = Poly{
			a: float64(parts[0]),
			b: float64(parts[1]),
			c: float64(parts[2]),
		}
	}

	// g input
	rInt := readInts(reader)
	r := make([]float64, m+1)
	for i := 0; i <= m; i++ {
		r[i] = float64(rInt[i])
	}

	g := make([]Poly, m)
	for i := 0; i < m; i++ {
		parts := readInts(reader)

		g[i] = Poly{
			a: float64(parts[0]),
			b: float64(parts[1]),
			c: float64(parts[2]),
		}
	}

	i, j := 0, 0
	answer := 0.0
	for i < n && j < m {
		left := math.Max(l[i], r[j])
		right := math.Min(l[i+1], r[j+1])

		if left < right {
			p := Poly{
				a: f[i].a - g[j].a,
				b: f[i].b - g[j].b,
				c: f[i].c - g[j].c,
			}

			answer += absIntegral(p, left, right)
		}

		if l[i+1] < r[j+1] {
			i++
		} else if r[j+1] < l[i+1] {
			j++
		} else {
			i++
			j++
		}
	}

	writer.WriteString(strconv.FormatFloat(answer, 'f', 10, 64))
	writer.WriteByte('\n')
}
