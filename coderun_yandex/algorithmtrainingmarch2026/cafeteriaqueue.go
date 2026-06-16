package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/cafeteria-queue
// CafeteriaQueue - assignment 16
func CafeteriaQueue() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// casses input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	cases, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	for ; cases > 0; cases-- {
		// N and D input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum := strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		n, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		d, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		pref := 0
		slack := make([]int, n)
		for i := 0; i < n; i++ {
			// T and K input
			line, err = reader.ReadString('\n')
			if err != nil && err != io.EOF {
				panic(err)
			}
			line = strings.TrimRight(line, "\r\n")
			strNum = strings.Fields(line)
			if len(strNum) != 2 {
				panic("numbers count does not match 2")
			}

			t, err := strconv.Atoi(strNum[0])
			if err != nil {
				panic(err)
			}
			k, err := strconv.Atoi(strNum[1])
			if err != nil {
				panic(err)
			}

			slack[i] = t - pref
			pref += k
		}

		suf := make([]int, n+1)
		suf[n] = math.MaxInt
		for i := n - 1; i >= 0; i-- {
			suf[i] = min(slack[i], suf[i+1])
		}

		ans := n + 1
		for i := 0; i < n; i++ {
			if suf[i] >= d {
				ans = i + 1

				break
			}
		}

		writer.WriteString(strconv.Itoa(ans))
		writer.WriteByte('\n')

	}
}
