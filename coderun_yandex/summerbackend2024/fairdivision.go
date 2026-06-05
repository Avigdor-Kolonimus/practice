package summerbackend2024

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/fair-division
// FairDivision - problem 32
func FairDivision() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// str input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n := len(line)

	pref := make([][]int, 26)
	for c := 0; c < 26; c++ {
		pref[c] = make([]int, n+1)
	}

	for i := 0; i < n; i++ {
		for c := 0; c < 26; c++ {
			pref[c][i+1] = pref[c][i]
		}

		pref[int(line[i]-'a')][i+1]++
	}

	ans := 1
	for blockLen := 1; blockLen <= n; blockLen++ {
		if n%blockLen != 0 {
			continue
		}

		pattern := make([]int, 26)

		for c := 0; c < 26; c++ {
			pattern[c] = pref[c][blockLen] - pref[c][0]
		}

		ok := true

		for l := blockLen; l < n && ok; l += blockLen {
			r := l + blockLen

			for c := 0; c < 26; c++ {
				cnt := pref[c][r] - pref[c][l]

				if cnt != pattern[c] {
					ok = false

					break
				}
			}
		}

		if ok {
			ans = n / blockLen

			break
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
