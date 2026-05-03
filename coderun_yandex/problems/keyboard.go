package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func validateKeyboardNInput(n int) bool {
	return n >= 1 && n <= 1_000
}

func validateKeyboardKeyInput(key int) bool {
	return key >= 1 && key <= 100_000
}

// https://coderun.yandex.ru/problem/keyboard
// Keyboard - problem 69
func Keyboard() {
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
	if !validateKeyboardNInput(n) {
		panic("number N out of range")
	}

	// keys input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strKeys := strings.Fields(line)
	keys := make([]int, n)
	for i := 0; i < n; i++ {
		keys[i], err = strconv.Atoi(strKeys[i])
		if err != nil {
			panic(err)
		}
		if !validateKeyboardKeyInput(keys[i]) {
			panic("key out of range")
		}
	}

	// K input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	k, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateKeyboardKeyInput(k) {
		panic("number K out of range")
	}

	// commands input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strCommands := strings.Fields(line)
	commands := make([]int, k)
	for i := 0; i < k; i++ {
		commands[i], err = strconv.Atoi(strCommands[i])
		if err != nil {
			panic(err)
		}

		keys[commands[i]-1]--
	}

	for i := 0; i < n; i++ {
		if keys[i] < 0 {
			writer.WriteString("YES")
		} else {
			writer.WriteString("NO")
		}
		writer.WriteByte('\n')
	}
}
