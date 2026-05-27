package devgointerview

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// https://coderun.yandex.ru/selections/dev-go-interview/problems/leader-tariff-history
// LeaderTariffHistory - problem 2
func LeaderTariffHistory() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// line input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	cnt := make(map[byte]int)
	maxFreq := 0
	leaders := 0
	var leaderChar byte

	n := len(line)
	ans := make([]byte, n)
	for i := 0; i < n; i++ {
		c := line[i]

		cnt[c]++
		f := cnt[c]

		if f > maxFreq {
			maxFreq = f
			leaders = 1
			leaderChar = c
		} else if f == maxFreq {
			if c != leaderChar || leaders > 0 {
				leaders++
			}
		}

		if f == maxFreq {
			leaders = 0

			for ch, v := range cnt {
				if v == maxFreq {
					leaders++
					leaderChar = ch
				}
			}
		}

		if leaders == 1 {
			ans[i] = leaderChar
		} else {
			ans[i] = '-'
		}
	}

	writer.WriteString(string(ans))
	writer.WriteByte('\n')
}
