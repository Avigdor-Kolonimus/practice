package mgustokashin

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func id(ch byte) int {
	if ch >= 'a' && ch <= 'z' {
		return int(ch - 'a')
	}

	return int(ch-'A') + 26
}

// https://coderun.yandex.ru/selections/mgustokashin/problems/maya-script
// MayaScript - problem 3
func MayaScript() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// g and S input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)

	g, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	s, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// symbol Word input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	symbolW := strings.TrimRight(line, "\r\n")

	// badges Word input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	badgesW := strings.TrimRight(line, "\r\n")

	var need, cur [52]int
	for i := 0; i < g; i++ {
		need[id(symbolW[i])]++
		cur[id(badgesW[i])]++
	}

	same := 0
	for i := 0; i < 52; i++ {
		if need[i] == cur[i] {
			same++
		}
	}

	ans := 0
	if same == 52 {
		ans++
	}

	update := func(c int, delta int) {
		before := (cur[c] == need[c])

		cur[c] += delta

		after := (cur[c] == need[c])

		if before && !after {
			same--
		} else if !before && after {
			same++
		}
	}

	for r := g; r < s; r++ {
		update(id(badgesW[r-g]), -1)
		update(id(badgesW[r]), 1)

		if same == 52 {
			ans++
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
