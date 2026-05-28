package devgointerview

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/dev-go-interview/problems/rover-with-two-compartments
// RoverWithTwoCompartments - problem 3
func RoverWithTwoCompartments() {
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

	// weights input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	weights := make([]int, n)
	prefOdd := make([]int, n+1)
	prefEven := make([]int, n+1)
	for i := 0; i < n; i++ {
		w, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		weights[i] = w

		prefOdd[i+1] = prefOdd[i]
		prefEven[i+1] = prefEven[i]

		if i%2 == 1 {
			prefEven[i+1] += w
		} else {
			prefOdd[i+1] += w
		}
	}

	// q input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	q, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// queries input
	for i := 0; i < q; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		l, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		odd := prefOdd[r] - prefOdd[l-1]
		even := prefEven[r] - prefEven[l-1]

		var first, second int
		if l%2 == 1 {
			first = odd
			second = even
		} else {
			first = even
			second = odd
		}

		if first == second {
			writer.WriteString("YES")
		} else {
			writer.WriteString("NO")
		}
		writer.WriteByte('\n')
	}
}
