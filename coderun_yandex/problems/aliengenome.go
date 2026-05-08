package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/alien-genome
// AlienGenome - problem 200
func AlienGenome() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first genome line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	s := strings.TrimRight(line, "\r\n")

	// second genome line
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	t := strings.TrimRight(line, "\r\n")

	pairs := make(map[string]bool)
	for i := 0; i+1 < len(t); i++ {
		p := t[i : i+2]
		pairs[p] = true
	}

	ans := 0
	for i := 0; i+1 < len(s); i++ {
		p := s[i : i+2]

		if pairs[p] {
			ans++
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
