package problems

import (
	"bufio"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

var replacer = strings.NewReplacer(
	")", "",
	"(", "",
	",", "",
	"-", "",
	"+", "",
)

func parsePhoneNumber(str string) string {
	number := replacer.Replace(str)

	if len(number) != 11 {
		number = "7495" + number
	}

	_, size := utf8.DecodeRuneInString(number)
	number = "7" + number[size:]

	return number
}

// https://coderun.yandex.ru/problem/phone-numbers
// PhoneNumbers - problem 54
func PhoneNumbers() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	readLine := func() string {
		line, err := reader.ReadString('\n')
		if err != nil && !(err == io.EOF && len(line) > 0) {
			panic(err)
		}
		return strings.TrimRight(line, "\r\n")
	}

	// first input
	line := readLine()
	number := parsePhoneNumber(line)

	numbers := make([]string, 3)
	for i := range 3 {
		line = readLine()
		numbers[i] = parsePhoneNumber(line)

	}

	result := "NO"
	for _, value := range numbers {
		if number == value {
			result = "YES"
		}

		writer.WriteString(result)
		writer.WriteByte('\n')
		result = "NO"
	}
}
