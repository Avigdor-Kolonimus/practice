package backend

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func countBooksWillBeRead(countDays int64) int64 {
	return countDays * (countDays + 1) / 2
}

func countBooksWillBeAvailable(countDays, startDay, countBooksPerDay int64) int64 {
	var countBooks int64
	copyStartDay := startDay

	for copyStartDay < 6 && copyStartDay-startDay < countDays {
		countBooks += countBooksPerDay
		copyStartDay++
	}

	usedDays := 7 - startDay
	if usedDays > countDays {
		usedDays = countDays
	}

	countDays -= usedDays

	countWeeks := countDays / 7
	countBooks += countWeeks * 5 * countBooksPerDay

	copyStartDay = 1

	for copyStartDay < countDays%7 && copyStartDay < 6 {
		countBooks += countBooksPerDay
		copyStartDay++
	}

	return countBooks
}

type PointLibrary struct {
	x int64
	y int64
}

func upperLine(countBooksPerDay, countBooksAtHome, startDay int64) (float64, float64) {
	var A, B PointLibrary

	if startDay < 7 {
		A = PointLibrary{
			x: 6 - startDay,
			y: countBooksPerDay*(6-startDay) + countBooksAtHome,
		}

		B = PointLibrary{
			x: 13 - startDay,
			y: countBooksPerDay*(11-startDay) + countBooksAtHome,
		}

		k := float64(B.y-A.y) / float64(B.x-A.x)

		b := float64(B.x*A.y-A.x*B.y) /
			float64(B.x-A.x)

		return k, b
	}

	k := float64(countBooksPerDay) * 5.0 / 7.0
	b := float64(countBooksAtHome) + k

	return k, b
}

func maxCountOfBook(countBooksPerDay, countBooksAtHome, startDay int64) int64 {
	k, b := upperLine(
		countBooksPerDay,
		countBooksAtHome,
		startDay,
	)

	D := (1-2*k)*(1-2*k) + 8*b

	sqrtD := math.Sqrt(D)

	result := (-1 + 2*k + sqrtD) / 2

	return int64(result) + 10
}

// https://coderun.yandex.ru/selections/backend/problems/library
// Library - problem 18
func Library() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// K, M and d line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 3 {
		panic("numbers count does not match 3")
	}

	k, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	d, err := strconv.Atoi(strNum[2])
	if err != nil {
		panic(err)
	}

	for i := int64(1); i < 1_000_000; i++ {
		need := countBooksWillBeRead(i)

		available := countBooksWillBeAvailable(i, int64(d), int64(k)) + int64(m)

		if need > available {
			fmt.Fprintln(writer, i-1)
			return
		}
	}

	left := int64(1_000_000)
	right := maxCountOfBook(int64(k), int64(m), int64(d))

	for left < right {
		mid := (left + right + 1) / 2

		need := countBooksWillBeRead(mid)

		available := countBooksWillBeAvailable(mid, int64(d), int64(k)) + int64(m)

		if need <= available {
			left = mid
		} else {
			right = mid - 1
		}
	}

	writer.WriteString(strconv.FormatInt(right, 10))
	writer.WriteByte('\n')
}
