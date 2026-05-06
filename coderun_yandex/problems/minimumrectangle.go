package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/minimum-rectangle
// MinimumRectangle - problem 247
func MinimumRectangle() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// K line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	k, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// cell inputs
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
	y, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	minX, maxX := x, x
	minY, maxY := y, y
	for i := 1; i < k; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}
		x, err = strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		y, err = strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}

	writer.WriteString(strconv.Itoa(minX) + " " + strconv.Itoa(minY) + " " + strconv.Itoa(maxX) + " " + strconv.Itoa(maxY))
	writer.WriteByte('\n')
}
