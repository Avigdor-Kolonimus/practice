package winterintern2024

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ItemRound struct {
	rem int
	idx int
}

// https://coderun.yandex.ru/selections/winter-intern-2024/problems/round
// Round - problem 4
func Round() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and X input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("input does not match 2")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	x, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// a input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum = strings.Fields(line)
	if len(strNum) != n {
		panic("input does not match n")
	}

	sum := 0
	a := make([]int, n)
	for i := range n {
		ai, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		a[i] = ai
		sum += ai
	}

	curSum := 0
	c := make([]int, n)
	items := make([]ItemRound, n)
	for i := 0; i < n; i++ {
		val := x * a[i]

		c[i] = val / sum
		rem := val % sum

		curSum += c[i]

		items[i] = ItemRound{
			rem: rem,
			idx: i,
		}
	}

	need := x - curSum

	sort.Slice(items, func(i, j int) bool {
		return items[i].rem > items[j].rem
	})

	for i := 0; i < need; i++ {
		c[items[i].idx]++
	}

	for i := 0; i < n; i++ {
		writer.WriteString(strconv.Itoa(c[i]) + " ")
	}
	writer.WriteByte('\n')
}
