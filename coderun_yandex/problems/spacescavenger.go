package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	rules map[byte]string
	dp    map[byte][101]int
)

func solveSpaceScavenger(c byte, m int) int {
	if m == 1 {
		return 1
	}

	// calculated yet
	if dp[c][m] != 0 {
		return dp[c][m]
	}

	ans := 1 // first step

	for i := 0; i < len(rules[c]); i++ {
		ans += solveSpaceScavenger(rules[c][i], m-1)
	}

	tmp := dp[c]
	tmp[m] = ans
	dp[c] = tmp

	return ans
}

// https://coderun.yandex.ru/problem/space-scavenger
// SpaceScavenger - problem 34
func SpaceScavenger() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	dp = make(map[byte][101]int)
	order := []byte{'N', 'S', 'W', 'E', 'U', 'D'}
	rules = make(map[byte]string)
	for i := 0; i < 6; i++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		s := strings.TrimRight(line, "\r\n")

		rules[order[i]] = s
	}

	// traps input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	c := strNum[0][0]
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	ans := solveSpaceScavenger(c, m)

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
