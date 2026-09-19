package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/krosh-and-game
// KroshAndGame - problem 591
func KroshAndGame() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	q, _ := strconv.Atoi(strings.TrimSpace(line))

	maxN := 0
	numbers := make([]int, q)
	for i := 0; i < q; i++ {
		line, _ = reader.ReadString('\n')
		numbers[i], _ = strconv.Atoi(strings.TrimSpace(line))

		if numbers[i] > maxN {
			maxN = numbers[i]
		}
	}

	win := make([]bool, maxN+1)
	for n := 1; n <= maxN; n++ {
		for x := 1; x*x <= n; x++ {
			if !win[n-x*x] {
				win[n] = true
				break
			}
		}
	}

	for _, n := range numbers {
		if win[n] {
			writer.WriteString("1\n")
		} else {
			writer.WriteString("0\n")
		}
	}
}