package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/rover-search
// RoverSearch - assignment 9
func RoverSearch() {
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

	// coordinates input
	l := math.MinInt
	r := math.MaxInt
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum := strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		x, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		d, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		left := x - d
		right := x + d

		if left > l {
			l = left
		}

		if right < r {
			r = right
		}
	}

	if l > r {
		writer.WriteString(strconv.Itoa(-1))
	} else {
		writer.WriteString(strconv.Itoa(r))
	}
	writer.WriteByte('\n')
}
