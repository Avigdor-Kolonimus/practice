package backend

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Server struct {
	a float64
	b float64
}

func validateServerErrorInput(n int) bool {
	return n >= 0 && n <= 100
}

// https://coderun.yandex.ru/selections/backend/problems/server-error
// ServerError - problem 24
func ServerError() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if n < 1 || !validateServerErrorInput(n) {
		panic("number N out of range")
	}

	sum := float64(0)
	servers := make([]Server, n)
	for i := range n {
		// percentages input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		strNum := strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		a, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		if !validateServerErrorInput(a) {
			panic("number A out of range")
		}

		b, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}
		if !validateServerErrorInput(b) {
			panic("number B out of range")
		}

		servers[i] = Server{a: float64(a), b: float64(b)}
		sum += float64(a) * float64(b)
	}

	for i := range n {
		p := (servers[i].a * servers[i].b) / sum
		writer.WriteString(strconv.FormatFloat(p, 'f', 9, 64))
		writer.WriteByte('\n')
	}
}
