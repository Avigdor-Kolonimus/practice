package algorithmtrainingseptember2025

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type PointBanner struct {
	x int
	y int
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/banner
// Banner - assignment 39
func Banner() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and D line
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	// N
	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}

	// D
	d, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// coordinates input
	points := make([]PointBanner, n)
	set := make(map[PointBanner]bool)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
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

		points[i] = PointBanner{x, y}
		set[points[i]] = true
	}

	var shifts []PointBanner
	for dx := 0; dx*dx <= d; dx++ {
		rem := d - dx*dx

		dy := int(math.Sqrt(float64(rem)))

		if dy*dy != rem {
			continue
		}

		candidates := []PointBanner{
			{dx, dy},
			{dx, -dy},
			{-dx, dy},
			{-dx, -dy},
		}

		shifts = append(shifts, candidates...)
	}

	// remove duplicate
	sort.Slice(shifts, func(i, j int) bool {
		if shifts[i].x == shifts[j].x {
			return shifts[i].y < shifts[j].y
		}

		return shifts[i].x < shifts[j].x
	})

	unique := make([]PointBanner, 0)
	for i, p := range shifts {
		if i == 0 || p != shifts[i-1] {
			unique = append(unique, p)
		}
	}

	shifts = unique

	var ans int
	for _, p := range points {
		for _, s := range shifts {
			if s.x < 0 || (s.x == 0 && s.y < 0) {
				continue
			}

			np := PointBanner{
				p.x + s.x,
				p.y + s.y,
			}

			if set[np] {
				ans++
			}
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
