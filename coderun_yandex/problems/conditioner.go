package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func validateConditionerInput(t int) bool {
	return t >= -50 && t <= 50
}

// https://coderun.yandex.ru/problem/conditioner
// Conditioner - problem 102
func Conditioner() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// temperatures line
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	tokens := strings.Fields(line)

	tRoom, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic(err)
	}
	if !validateConditionerInput(tRoom) {
		panic("number tRoom out of range")
	}

	tConditioner, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}
	if !validateConditionerInput(tConditioner) {
		panic("number tConditioner out of range")
	}

	// command input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	mode := strings.TrimRight(line, "\r\n")

	switch mode {
	case "freeze":
		writer.WriteString(strconv.Itoa(min(tRoom, tConditioner)))

	case "heat":
		writer.WriteString(strconv.Itoa(max(tRoom, tConditioner)))

	case "auto":
		writer.WriteString(strconv.Itoa(tConditioner))

	case "fan":
		writer.WriteString(strconv.Itoa(tRoom))
	}

	writer.WriteByte('\n')
}
