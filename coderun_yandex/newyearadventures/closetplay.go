package newyearadventures

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func checkClosetPlay(a, b [26]int) bool {
	for i := 0; i < 26; i++ {
		if a[i] > b[i] {
			return false
		}
	}
	return true
}

// https://coderun.yandex.ru/selections/new-year-adventures/problems/closet_play
// ClosetPlay - problem 8
func ClosetPlay() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// K line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	k, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// str1 input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	s := strings.TrimRight(line, "\r\n")

	// str2 input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	t := strings.TrimRight(line, "\r\n")

	if len(s) < k || len(t) < k {
		writer.WriteString("NO")
		writer.WriteByte('\n')

		return
	}

	var cntT [26]int

	for _, c := range t {
		cntT[c-'a']++
	}

	var window [26]int

	for i := 0; i < k; i++ {
		window[s[i]-'a']++
	}

	if checkClosetPlay(window, cntT) {
		writer.WriteString("YES")
		writer.WriteByte('\n')

		return
	}

	for i := k; i < len(s); i++ {
		window[s[i-k]-'a']--
		window[s[i]-'a']++

		if checkClosetPlay(window, cntT) {
			writer.WriteString("YES")
			writer.WriteByte('\n')

			return
		}
	}

	writer.WriteString("NO")
	writer.WriteByte('\n')
}
