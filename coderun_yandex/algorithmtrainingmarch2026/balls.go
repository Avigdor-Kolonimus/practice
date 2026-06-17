package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type GroupBalls struct {
	color int
	cnt   int
}

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/balls
// Balls - assignment 23
func Balls() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// numbers input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)

	n := len(strNum)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		ai, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		a[i] = ai
	}

	var groups []GroupBalls
	for _, x := range a {
		if len(groups) == 0 || groups[len(groups)-1].color != x {
			groups = append(groups, GroupBalls{x, 1})
		} else {
			groups[len(groups)-1].cnt++
		}
	}

	removed := 0

	for {
		pos := -1
		for i := range groups {
			if groups[i].cnt >= 3 {
				pos = i
				break
			}
		}

		if pos == -1 {
			break
		}

		removed += groups[pos].cnt
		groups = append(groups[:pos], groups[pos+1:]...)

		for pos > 0 && pos < len(groups) &&
			groups[pos-1].color == groups[pos].color {

			groups[pos-1].cnt += groups[pos].cnt
			groups = append(groups[:pos], groups[pos+1:]...)
			pos--
		}
	}

	writer.WriteString(strconv.Itoa(removed))
	writer.WriteByte('\n')
}
