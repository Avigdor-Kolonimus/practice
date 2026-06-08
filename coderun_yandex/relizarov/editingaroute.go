package relizarov

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/relizarov/problems/editing-a-route
// EditingARoute - problem 1
func EditingARoute() {
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

	// route input
	for i := 0; i < n; i++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum := strings.Fields(line)

		s := strNum[0]
		t := strNum[1]

		// need[c] = how many occurrences of character c
		// must remain in the final string
		var need [26]int
		for _, ch := range t {
			need[ch-'A']++
		}

		// Copy of need that will be used while scanning from right to left
		remain := need

		// keep[i] indicates whether s[i] survives all deletions
		keep := make([]bool, len(s))

		// For each character, keep only the last need[c] occurrences.
		// Scanning from right to left makes this easy.
		for i := len(s) - 1; i >= 0; i-- {
			c := s[i] - 'A'
			if remain[c] > 0 {
				keep[i] = true
				remain[c]--
			}
		}

		// Build the uniquely determined final string
		// after all possible deletions.
		res := make([]byte, 0, len(s))
		for i := 0; i < len(s); i++ {
			if keep[i] {
				res = append(res, s[i])
			}
		}

		if string(res) == t {
			writer.WriteString("YES")
		} else {
			writer.WriteString("NO")
		}
		writer.WriteByte('\n')
	}
}
