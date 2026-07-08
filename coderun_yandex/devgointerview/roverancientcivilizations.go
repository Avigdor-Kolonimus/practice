package devgointerview

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/dev-go-interview/problems/rover-ancient-civilizations
// RoverAncientCivilizations - problem 14
func RoverAncientCivilizations() {
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

	// sequence input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		sequence[i], err = strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}
	}

	pi := make([]int, n)
	for i := 1; i < n; i++ {
		j := pi[i-1]
		for j > 0 && sequence[i] != sequence[j] {
			j = pi[j-1]
		}

		if sequence[i] == sequence[j] {
			j++
		}

		pi[i] = j
	}

	p := n - pi[n-1]

	if p == n {
		writer.WriteString("0\n")

		return
	}

	for i := p; i < n; i++ {
		if sequence[i] != sequence[i-p] {
			writer.WriteString("0\n")

			return
		}
	}

	writer.WriteString(strconv.Itoa(p))
	writer.WriteByte('\n')
}
