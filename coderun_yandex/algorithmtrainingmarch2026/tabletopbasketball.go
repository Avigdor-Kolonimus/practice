package algorithmtrainingmarch2026

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/tabletop-basketball
// TabletopBasketball - assignment 7
func TabletopBasketball() {
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

	// names input
	players := make(map[string]int, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		players[line] = 0
	}

	//  M input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	m, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// score input
	prevA, prevB := 0, 0
	for i := 0; i < m; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		var a, b int
		var name string

		fmt.Sscanf(line, "%d:%d %s", &a, &b, &name)

		delta := (a - prevA) + (b - prevB)

		players[name] += delta

		prevA = a
		prevB = b
	}

	bestName := ""
	bestScore := -1
	for name, points := range players {
		if points > bestScore {
			bestScore = points
			bestName = name
		}
	}

	writer.WriteString(bestName)
	writer.WriteByte(' ')
	writer.WriteString(strconv.Itoa(bestScore))
	writer.WriteByte('\n')
}
