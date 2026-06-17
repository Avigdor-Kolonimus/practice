package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/candy-subsegment
// CandySubsegment - assignment 14
func CandySubsegment() {
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

	// candy input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	candy := make([]int, n)
	for i := 0; i < n; i++ {
		candy[i], err = strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}
	}

	l := 0
	ans := 0
	cnt := make(map[int]int)
	for r := 0; r < n; r++ {
		cnt[candy[r]]++

		for len(cnt) > 2 {
			cnt[candy[l]]--
			if cnt[candy[l]] == 0 {
				delete(cnt, candy[l])
			}

			l++
		}

		if len(cnt) == 2 {
			if r-l+1 > ans {
				ans = r - l + 1
			}
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
