package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/arrangement-laptops
// Arrangementlaptops - problem 181
func Arrangementlaptops() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	tokens := strings.Fields(line)
	if len(tokens) != 4 {
		panic("invalid input")
	}
	a, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}
	c, err := strconv.Atoi(tokens[2])
	if err != nil {
		panic(err)
	}
	d, err := strconv.Atoi(tokens[3])
	if err != nil {
		panic(err)
	}

	sides := [][2]int{
		{(a + c), max(b, d)},
		{(a + d), max(b, c)},
		{(b + c), max(a, d)},
		{(b + d), max(a, c)},
	}

	// calculation
	t1, t2 := sides[0][0], sides[0][1]
	for i := 1; i < len(sides); i++ {
		if t1*t2 > sides[i][0]*sides[i][1] {
			t1, t2 = sides[i][0], sides[i][1]
		}
	}

	// enough one output
	writer.WriteString(strconv.Itoa(t1) + " " + strconv.Itoa(t2))
	writer.WriteByte('\n')
}
