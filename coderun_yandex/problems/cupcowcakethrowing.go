package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/cup-cowcake-throwing
// CupCowcakeThrowing - problem 60
func CupCowcakeThrowing() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// lenghts input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	maxVal := 0
	lenghts := make([]int, n)
	for i := 0; i < n; i++ {
		x, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		if x > maxVal {
			maxVal = x
		}

		lenghts[i] = x
	}

	best := -1
	prefixMax := lenghts[0]
	for i := 1; i < n-1; i++ {
		if prefixMax == maxVal && lenghts[i]%10 == 5 && lenghts[i] > lenghts[i+1] {

			if lenghts[i] > best {
				best = lenghts[i]
			}
		}

		if lenghts[i] > prefixMax {
			prefixMax = lenghts[i]
		}
	}

	if best == -1 {
		writer.WriteByte('0')
		writer.WriteByte('\n')

		return
	}

	place := 1
	for _, x := range lenghts {
		if x > best {
			place++
		}
	}

	writer.WriteString(strconv.Itoa(place))
	writer.WriteByte('\n')
}
