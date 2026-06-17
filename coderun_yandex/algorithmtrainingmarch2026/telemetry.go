package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/telemetry
// Telemetry - assignment 20
func Telemetry() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N, M and K input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 3 {
		panic("numbers count does not match 3")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	k, err := strconv.Atoi(strNum[2])
	if err != nil {
		panic(err)
	}

	// command input
	cur := 0
	clipboard := ""
	docs := make([]string, n)
	for i := 0; i < m; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		cmd := strings.TrimRight(line, "\r\n")

		switch cmd {
		case "Next":
			cur = (cur + 1) % n

		case "Backspace":
			if len(docs[cur]) > 0 {
				docs[cur] = docs[cur][:len(docs[cur])-1]
			}

		case "Copy":
			s := docs[cur]
			if len(s) <= k {
				clipboard = s
			} else {
				clipboard = s[len(s)-k:]
			}

		case "Paste":
			docs[cur] += clipboard

		default: // charachter
			docs[cur] += cmd
		}
	}

	result := docs[cur]

	if len(result) == 0 {
		writer.WriteString("Empty")
		writer.WriteByte('\n')

		return
	}

	if len(result) > k {
		result = result[len(result)-k:]
	}

	writer.WriteString(result)
	writer.WriteByte('\n')
}
