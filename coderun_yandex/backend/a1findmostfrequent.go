package backend

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/backend/problems/a-1-find-most-frequent
// A1FindMostFrequent - problem 42
func A1FindMostFrequent() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// sum inputs
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	a := make([]int, n)
	entries := make(map[int]int, n)
	for i := 0; i < n; i++ {
		x, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}
		a[i] = x
		entries[x] += 1
	}

	maxElementValue := 0
	maxElementEntries := 0
	for elementValue, elementEntries := range entries {
		if (elementEntries > maxElementEntries) || (elementEntries == maxElementEntries && elementValue > maxElementValue) {
			maxElementEntries = elementEntries
			maxElementValue = elementValue
		}
	}

	writer.WriteString(strconv.Itoa(maxElementValue))
	writer.WriteByte('\n')
}
