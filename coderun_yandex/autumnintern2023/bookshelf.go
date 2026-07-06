package autumnintern2023

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	SIZE = 102
	NEG  = -1
)

func id(l, r int) int {
	return l*SIZE + r
}

// https://coderun.yandex.ru/selections/autumn-intern-2023/problems/book-shelf
// BookShelf - problem 1
func BookShelf() {
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

	// height of the book input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	height := make([]int, n)
	for i := 0; i < n; i++ {
		height[i], err = strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}
	}

	dp := make([]int, SIZE*SIZE)
	for i := range dp {
		dp[i] = NEG
	}

	dp[id(0, 101)] = 0
	for _, x := range height {
		next := make([]int, len(dp))
		copy(next, dp)

		for l := 0; l < SIZE; l++ {
			base := l * SIZE
			for r := 0; r < SIZE; r++ {
				cur := dp[base+r]
				if cur < 0 {
					continue
				}

				if l <= x && x <= r {
					// left side
					idx := id(x, r)
					if next[idx] < cur+1 {
						next[idx] = cur + 1
					}

					// right side
					idx = id(l, x)
					if next[idx] < cur+1 {
						next[idx] = cur + 1
					}
				}
			}
		}

		dp = next
	}

	ans := 0
	for _, v := range dp {
		if v > ans {
			ans = v
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
