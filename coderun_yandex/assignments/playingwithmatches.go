package assignments

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validatePlayingWithMatchesInput(n int) bool {
	return n >= 1 && n <= 10_000
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/playing-with-matches
// PlayingWithMatches - assignment 18
func PlayingWithMatches() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validatePlayingWithMatchesInput(n) {
		panic("number N out of range")
	}

	result := 1
	if n%4 == 0 {
		result = 2
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
