package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/symbols-set-min-susbstr
// SymbolsSetMinSusbstr - problem 146
func SymbolsSetMinSusbstr() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// s group input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	s := strings.TrimRight(line, "\r\n")

	// c group input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	c := strings.TrimRight(line, "\r\n")

	need := make(map[byte]bool)
	for i := range c {
		need[c[i]] = true
	}

	left := 0
	have := 0
	ans := len(s) + 1
	cnt := make(map[byte]int)
	for right := 0; right < len(s); right++ {
		ch := s[right]

		if !need[ch] {
			cnt = make(map[byte]int)
			have = 0
			left = right + 1

			continue
		}

		cnt[ch]++
		if cnt[ch] == 1 {
			have++
		}

		for have == len(need) {
			if cnt[s[left]] > 1 {
				cnt[s[left]]--
				left++
			} else {
				break
			}
		}

		if have == len(need) {
			if right-left+1 < ans {
				ans = right - left + 1
			}
		}
	}

	if ans == len(s)+1 {
		writer.WriteByte('0')
	} else {
		writer.WriteString(strconv.Itoa(ans))
	}
	writer.WriteByte('\n')
}
