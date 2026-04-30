package yandexinterview

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/yandex-interview/problems/anagrams
// Anagrams - problem 2
func Anagrams() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// str1 input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	str1 := strings.TrimRight(line, "\r\n")

	// str2 input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	str2 := strings.TrimRight(line, "\r\n")

	result := 1
	if len(str1) != len(str2) {
		result = 0
	} else {
		cntChar := make([]int, 26)
		for i := 0; i < len(str1); i++ {
			cntChar[str1[i]-'a']++
			cntChar[str2[i]-'a']--
		}

		for i := range 26 {
			if cntChar[i] != 0 {
				result = 0
				break
			}
		}
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
