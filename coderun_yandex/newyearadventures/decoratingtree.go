package newyearadventures

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateDecoratingTreeInput(k int) bool {
	return k >= 1 && k <= 25
}

// https://coderun.yandex.ru/selections/new-year-adventures/problems/decorating-tree
// DecoratingTree - problem 2
func DecoratingTree() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// K input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	k, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateDecoratingTreeInput(k) {
		panic("number K out of range")
	}

	result := 2
	for i := 2; i <= k; i++ {
		result *= 2
	}

	writer.WriteString(strconv.Itoa(result - 1))
	writer.WriteByte('\n')
}
