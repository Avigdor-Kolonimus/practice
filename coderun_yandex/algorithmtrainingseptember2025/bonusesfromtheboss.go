package algorithmtrainingseptember2025

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/bonuses-from-the-boss
// BonusesFromTheBoss - assignment 38
func BonusesFromTheBoss() {
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

	// overtime input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	a := make([]int, n+1)
	for i := range n {
		ai, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		a[i+1] = ai
	}

	diff := make([]int, n+3)
	for j := 1; j <= n; j++ {
		if a[j] == 0 {
			continue
		}

		l := j + 1
		r := min(n, j+int(a[j])-1)

		if l <= r {
			diff[l]++
			diff[r+1]--
		}
	}

	ans, cur := 0, 0
	for i := 1; i <= n; i++ {
		cur += diff[i]
		ans += cur * a[i]
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
