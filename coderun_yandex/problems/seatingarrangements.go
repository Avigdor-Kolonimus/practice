package problems

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Student struct {
	index    int
	distance int
}

func countAmountOfVariants(n, d int, students []int) (int, []int) {
	indexedStudents := make([]Student, n)

	for i, x := range students {
		indexedStudents[i] = Student{
			index:    i,
			distance: x,
		}
	}

	sort.Slice(indexedStudents, func(i, j int) bool {
		return indexedStudents[i].distance < indexedStudents[j].distance
	})

	places := make([]int, n)
	places[indexedStudents[0].index] = 1

	variant := 1
	start := 0

	for end := 1; end < n; end++ {
		previous := indexedStudents[start]
		current := indexedStudents[end]

		if current.distance-previous.distance <= d {
			variant++
			places[current.index] = variant
		} else {
			places[current.index] = places[previous.index]
			start++
		}
	}

	return variant, places
}

// https://coderun.yandex.ru/problem/seating-arrangements
// SeatingArrangements - problem 235
func SeatingArrangements() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and D line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	d, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// students coordinate input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum = strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	students := make([]int, n)
	for i := 0; i < n; i++ {
		x, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}
		students[i] = x
	}

	variant, places := countAmountOfVariants(n, d, students)

	writer.WriteString(strconv.Itoa(variant))
	writer.WriteByte('\n')

	for i := 0; i < n; i++ {
		if i > 0 {
			writer.WriteByte(' ')
		}
		writer.WriteString(strconv.Itoa(places[i]))
	}
	writer.WriteByte('\n')
}
