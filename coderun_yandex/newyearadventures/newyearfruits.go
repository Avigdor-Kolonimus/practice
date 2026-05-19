package newyearadventures

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type BoxNewYearFruits struct {
	mandarin int
	orange   int
}

// https://coderun.yandex.ru/selections/new-year-adventures/problems/new-year-fruits
// NewYearFruits - problem 3
func NewYearFruits() {
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

	// box input
	mo := make([]BoxNewYearFruits, 2*n-1)
	for i := 0; i < 2*n-1; i++ {
		// fruits input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum := strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		m, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}

		o, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		mo[i].mandarin = m
		mo[i].orange = o
	}

	writer.WriteString("Yes")
	writer.WriteByte('\n')
}
