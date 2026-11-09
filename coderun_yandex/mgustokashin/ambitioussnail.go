package mgustokashin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Berry struct {
	a  int64
	d  int64
	id int
}

// https://coderun.yandex.ru/selections/mgustokashin/problems/ambitious-snail
// AmbitiousSnail - problem 6
func AmbitiousSnail() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	berries := make([]Berry, n)
	var sumPos int64
	for i := range n {
		var a, b int64
		fmt.Fscan(reader, &a, &b)

		d := a - b

		berries[i] = Berry{
			a:  a,
			d:  d,
			id: i + 1,
		}

		if d > 0 {
			sumPos += d
		}
	}

	bestIdx := 0
	var bestH int64
	for i, berry := range berries {
		candidate := berry.a + sumPos

		if berry.d > 0 {
			candidate -= berry.d
		}

		if candidate > bestH {
			bestH = candidate
			bestIdx = i
		}
	}

	answer := make([]int, 0, n)
	for i, berry := range berries {
		if berry.d > 0 && i != bestIdx {
			answer = append(answer, berry.id)
		}
	}

	answer = append(answer, berries[bestIdx].id)
	for i, berry := range berries {
		if berry.d <= 0 && i != bestIdx {
			answer = append(answer, berry.id)
		}
	}

	fmt.Fprintln(writer, bestH)

	for i, id := range answer {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, id)
	}
	fmt.Fprintln(writer)
}
