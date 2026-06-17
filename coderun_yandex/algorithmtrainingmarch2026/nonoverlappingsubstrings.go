package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	BASE uint64 = 911_382_323
)

var (
	s    string
	n    int
	pref []uint64
	pow  []uint64
)

func getHash(l, r int) uint64 {
	return pref[r] - pref[l]*pow[r-l]
}

func check(pos []int, length int) bool {
	h := getHash(pos[0], pos[0]+length)

	for i := 1; i < len(pos); i++ {
		if getHash(pos[i], pos[i]+length) != h {
			return false
		}
	}

	return true
}

func commonPrefix(pos []int, limit int) int {
	left, right := 0, limit
	ans := 0

	for left <= right {
		mid := (left + right) / 2

		if check(pos, mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/non-overlapping-substrings
// NonOverlappingSubstrings - assignment 10
func NonOverlappingSubstrings() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// string input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	s = strings.TrimRight(line, "\r\n")

	n = len(s)

	pos := make([][]int, 26)
	for i := 0; i < n; i++ {
		pos[s[i]-'a'] = append(pos[s[i]-'a'], i)
	}

	maxFreq := 0
	for c := 0; c < 26; c++ {
		if len(pos[c]) > maxFreq {
			maxFreq = len(pos[c])
		}
	}

	pref = make([]uint64, n+1)
	pow = make([]uint64, n+1)
	pow[0] = 1
	for i := 0; i < n; i++ {
		pref[i+1] = pref[i]*BASE + uint64(s[i])
		pow[i+1] = pow[i] * BASE
	}

	answer := 0
	for c := 0; c < 26; c++ {
		if len(pos[c]) != maxFreq {
			continue
		}

		p := pos[c]

		limit := n - p[len(p)-1]

		for i := 0; i+1 < len(p); i++ {
			limit = min(limit, p[i+1]-p[i])
		}

		cur := commonPrefix(p, limit)

		if cur > answer {
			answer = cur
		}
	}

	writer.WriteString(strconv.Itoa(answer))
	writer.WriteByte('\n')
}
