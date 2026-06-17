package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/dance-school
// DanceSchool - assignment 13
func DanceSchool() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	_, err = strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// students input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	pref, ans := 0, 0
	cnt := make(map[int]int)
	cnt[0] = 1
	for _, ch := range line {
		if ch == 'a' {
			pref++
		} else {
			pref--
		}

		ans += cnt[pref]
		cnt[pref]++
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
