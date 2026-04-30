package problems

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Block struct {
	width  int
	height int
}

func validatePyramidInput(n int) bool {
	return n >= 1 && n <= 100_000
}

func validatePyramidWHInput(wh int) bool {
	return wh >= 1 && wh <= 1_000_000_000
}

// https://coderun.yandex.ru/problem/pyramid
// Pyramid - problem 70
func Pyramid() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N line
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validatePyramidInput(n) {
		panic("number N out of range")
	}

	// block inputs
	blocks := make([]Block, n)
	for i := range n {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}

		strNum := strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		w, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		if !validatePyramidWHInput(w) {
			panic("number Width out of range")
		}
		h, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}
		if !validatePyramidWHInput(h) {
			panic("number Height out of range")
		}

		blocks[i].height = h
		blocks[i].width = w
	}

	sort.Slice(blocks, func(i, j int) bool {
		if blocks[i].width == blocks[j].width {
			return blocks[i].height > blocks[j].height
		}

		return blocks[i].width > blocks[j].width
	})

	height := blocks[0].height
	for i := 1; i < len(blocks); i++ {
		if blocks[i-1].width != blocks[i].width {
			height += blocks[i].height
		}
	}

	writer.WriteString(strconv.Itoa(height))
	writer.WriteByte('\n')
}
