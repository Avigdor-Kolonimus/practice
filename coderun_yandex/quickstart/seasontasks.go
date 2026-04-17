package quickstart

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/quickstart/problems/season-tasks
// SeasonTasks - assignment 1
func SeasonTasks() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	// A
	a, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}

	// B
	b, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	writer.WriteString(strconv.Itoa(a + b))
	writer.WriteByte('\n')
}
