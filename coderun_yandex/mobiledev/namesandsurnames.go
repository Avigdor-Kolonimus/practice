package mobiledev

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hunter struct {
	level int
	index int
}

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/names-and-surnames
// NamesAndSurnames - problem 4
func NamesAndSurnames() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	line, _ = reader.ReadString('\n')
	fields := strings.Fields(line)

	hunters := make([]Hunter, n)
	for i := range n {
		level, _ := strconv.Atoi(fields[i])
		hunters[i] = Hunter{
			level: level,
			index: i + 1,
		}
	}

	sort.Slice(hunters, func(i, j int) bool {
		return hunters[i].level < hunters[j].level
	})

	level := 1
	penalties := 0
	order := make([]int, 0, n)
	for _, hunter := range hunters {
		if level < hunter.level {
			writer.WriteString("Impossible\n")
			return
		}

		level++
		order = append(order, hunter.index)

		if level < 2*hunter.level {
			penalties++

			if penalties == 3 {
				level--
				penalties = 0
			}
		}
	}

	writer.WriteString("Possible\n")

	for i, index := range order {
		if i > 0 {
			writer.WriteByte(' ')
		}
		writer.WriteString(strconv.Itoa(index))
	}
	writer.WriteByte('\n')
}
