package summerbackend2024

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x, y int64
}

type Line struct {
	p1, p2 Point
	isVert bool
}

func NewLine(p1, p2 Point) Line {
	return Line{p1: p1, p2: p2, isVert: p1.x == p2.x}
}

func (l Line) isBelongsToLine(p Point) bool {
	if l.isVert {
		return p.x == l.p1.x
	}

	return (p.x-l.p1.x)*(l.p2.y-l.p1.y) == (p.y-l.p1.y)*(l.p2.x-l.p1.x)
}

func cross(a, b, c Point) int64 {
	return (b.x-a.x)*(c.y-a.y) - (b.y-a.y)*(c.x-a.x)
}

func buildConvexHull(pts []Point) []Point {
	if len(pts) < 2 {
		return append([]Point{}, pts...)
	}

	lower := []Point{pts[0], pts[1]}
	for i := 2; i < len(pts); i++ {
		c := pts[i]
		for len(lower) >= 2 {
			a := lower[len(lower)-2]
			b := lower[len(lower)-1]
			if cross(a, b, c) <= 0 {
				lower = lower[:len(lower)-1]
			} else {
				break
			}
		}
		lower = append(lower, c)
	}

	upper := []Point{pts[0], pts[1]}
	for i := 2; i < len(pts); i++ {
		c := pts[i]
		for len(upper) >= 2 {
			a := upper[len(upper)-2]
			b := upper[len(upper)-1]
			if cross(a, b, c) >= 0 {
				upper = upper[:len(upper)-1]
			} else {
				break
			}
		}
		upper = append(upper, c)
	}

	hull := append([]Point{}, lower...)
	for i := len(upper) - 2; i > 0; i-- {
		hull = append(hull, upper[i])
	}
	return hull
}

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/four-points
// FourPoints - problem 4
func FourPoints() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	lineStr, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	lineStr = strings.TrimRight(lineStr, "\r\n")
	if lineStr == "" {
		return
	}
	n, _ := strconv.Atoi(lineStr)

	points := make([]Point, n)
	for i := 0; i < n; i++ {
		lineStr, _ = reader.ReadString('\n')
		lineStr = strings.TrimRight(lineStr, "\r\n")
		fields := strings.Fields(lineStr)
		x, _ := strconv.ParseInt(fields[0], 10, 64)
		y, _ := strconv.ParseInt(fields[1], 10, 64)
		points[i] = Point{x: x, y: y}
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i].x != points[j].x {
			return points[i].x < points[j].x
		}
		return points[i].y < points[j].y
	})

	convexHull := buildConvexHull(points)

	if len(convexHull) >= 4 {
		writer.WriteString("Yes\n")
		return
	} else if len(convexHull) == 3 {
		line1 := NewLine(convexHull[0], convexHull[1])
		line2 := NewLine(convexHull[0], convexHull[2])
		line3 := NewLine(convexHull[1], convexHull[2])

		var pointOnLinesOfTriangle []Point
		var pointsInsideTriangle []Point

		for i := 0; i < n; i++ {
			isOnLine1 := line1.isBelongsToLine(points[i])
			isOnLine2 := line2.isBelongsToLine(points[i])
			isOnLine3 := line3.isBelongsToLine(points[i])

			if !isOnLine1 && !isOnLine2 && !isOnLine3 {
				pointsInsideTriangle = append(pointsInsideTriangle, points[i])
			} else if points[i] != convexHull[0] && points[i] != convexHull[1] && points[i] != convexHull[2] {
				pointOnLinesOfTriangle = append(pointOnLinesOfTriangle, points[i])
			}
		}

		var pointOnLine1, pointOnLine2, pointOnLine3 []Point
		for i := 0; i < n; i++ {
			if points[i] == convexHull[0] || points[i] == convexHull[1] || points[i] == convexHull[2] {
				continue
			}
			if line1.isBelongsToLine(points[i]) {
				pointOnLine1 = append(pointOnLine1, points[i])
			}
			if line2.isBelongsToLine(points[i]) {
				pointOnLine2 = append(pointOnLine2, points[i])
			}
			if line3.isBelongsToLine(points[i]) {
				pointOnLine3 = append(pointOnLine3, points[i])
			}
		}

		nonEmptyLines := 0
		if len(pointOnLine1) > 0 {
			nonEmptyLines++
		}
		if len(pointOnLine2) > 0 {
			nonEmptyLines++
		}
		if len(pointOnLine3) > 0 {
			nonEmptyLines++
		}

		if nonEmptyLines >= 2 {
			writer.WriteString("Yes\n")
			return
		}

		if len(pointsInsideTriangle) == 0 {
			writer.WriteString("No\n")
			return
		}

		if len(pointsInsideTriangle) == 1 {
			pInside := pointsInsideTriangle[0]
			if len(pointOnLine1) > 0 {
				for _, p := range pointOnLine1 {
					l := NewLine(pInside, p)
					if !l.isBelongsToLine(convexHull[2]) {
						writer.WriteString("Yes\n")
						return
					}
				}
			}
			if len(pointOnLine2) > 0 {
				for _, p := range pointOnLine2 {
					l := NewLine(pInside, p)
					if !l.isBelongsToLine(convexHull[1]) {
						writer.WriteString("Yes\n")
						return
					}
				}
			}
			if len(pointOnLine3) > 0 {
				for _, p := range pointOnLine3 {
					l := NewLine(pInside, p)
					if !l.isBelongsToLine(convexHull[0]) {
						writer.WriteString("Yes\n")
						return
					}
				}
			}
			writer.WriteString("No\n")
			return
		}

		sort.Slice(pointsInsideTriangle, func(i, j int) bool {
			if pointsInsideTriangle[i].x != pointsInsideTriangle[j].x {
				return pointsInsideTriangle[i].x < pointsInsideTriangle[j].x
			}
			return pointsInsideTriangle[i].y < pointsInsideTriangle[j].y
		})

		newConvexHull := buildConvexHull(pointsInsideTriangle)

		if len(newConvexHull) >= 4 {
			writer.WriteString("Yes\n")
			return
		} else if len(newConvexHull) == 3 {
			line4 := NewLine(newConvexHull[0], newConvexHull[1])
			line5 := NewLine(newConvexHull[0], newConvexHull[2])
			line6 := NewLine(newConvexHull[1], newConvexHull[2])
			cool := false

			for _, p := range pointOnLinesOfTriangle {
				if !line4.isBelongsToLine(p) && !line5.isBelongsToLine(p) && !line6.isBelongsToLine(p) {
					cool = true
					break
				}
			}

			for i := 0; i < 3; i++ {
				if !line4.isBelongsToLine(convexHull[i]) && !line5.isBelongsToLine(convexHull[i]) && !line6.isBelongsToLine(convexHull[i]) {
					cool = true
					break
				}
			}

			if cool {
				writer.WriteString("Yes\n")
				return
			}

			var newPointsInsideTriangle []Point
			for _, p := range pointsInsideTriangle {
				if !line4.isBelongsToLine(p) && !line5.isBelongsToLine(p) && !line6.isBelongsToLine(p) {
					newPointsInsideTriangle = append(newPointsInsideTriangle, p)
				}
			}

			if len(newPointsInsideTriangle) > 0 {
				writer.WriteString("Yes\n")
				return
			}
			writer.WriteString("No\n")
			return

		} else if len(newConvexHull) == 2 {
			line7 := NewLine(newConvexHull[0], newConvexHull[1])
			cool := true

			for i := 0; i < 3; i++ {
				if line7.isBelongsToLine(convexHull[i]) {
					cool = false
					break
				}
			}

			for _, p := range pointOnLinesOfTriangle {
				if !line7.isBelongsToLine(p) {
					cool = true
					break
				}
			}

			if cool {
				writer.WriteString("Yes\n")
				return
			}
			writer.WriteString("No\n")
			return
		}
	} else {
		writer.WriteString("No\n")
	}
}
