package backend

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateRocksAndJewelsInput(n int) bool {
	return n >= 0 && n <= 100
}

// https://coderun.yandex.ru/selections/backend/problems/rocks-and-jewels
// RocksAndJewels - problem 52
func RocksAndJewels() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// J input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	j := strings.TrimRight(line, "\r\n")

	if !validateRocksAndJewelsInput(len(j)) {
		panic("number J out of range")
	}

	// S input
	line, err = reader.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		panic(err)
	}
	s := strings.TrimRight(line, "\r\n")

	if !validateRocksAndJewelsInput(len(s)) {
		panic("number S out of range")
	}

	m := make(map[byte]bool, len(j))
	for i := range j {
		m[j[i]] = true
	}

	count := 0
	for i := range s {
		if m[s[i]] {
			count++
		}
	}

	writer.WriteString(strconv.Itoa(count))
	writer.WriteByte('\n')
}
