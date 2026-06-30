package mobiledev

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/mobile-dev/problems/game
// Game - problem 28
func Game() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N, M and K line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 3 {
		panic("numbers count does not match 3")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	k, err := strconv.Atoi(strNum[2])
	if err != nil {
		panic(err)
	}

	// current input
	cur := make([][]int, n)
	for i := range cur {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
		if len(strNum) != m {
			panic("numbers count does not match m")
		}

		cur[i] = make([]int, m)
		for j := 0; j < m; j++ {
			cur[i][j], err = strconv.Atoi(strNum[j])
			if err != nil {
				panic(err)
			}
		}
	}

	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}

	dx := [4]int{-1, 1, 0, 0}
	dy := [4]int{0, 0, -1, 1}

	for step := 0; step < k; step++ {
		next := make([][]int, n)
		for i := range next {
			next[i] = make([]int, m)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {

				stable := 0
				active := 0

				for d := 0; d < 4; d++ {
					ni := i + dx[d]
					nj := j + dy[d]

					if ni < 0 || ni >= n || nj < 0 || nj >= m {
						continue
					}

					if cur[ni][nj] == 2 {
						stable++
					}

					if cur[ni][nj] == 2 || cur[ni][nj] == 3 {
						active++
					}
				}

				if stable > 1 {
					next[i][j] = 2
				} else if active > 0 {
					next[i][j] = 3
				} else {
					next[i][j] = 1
				}

				if next[i][j] != cur[i][j] {
					ans[i][j]++
				}
			}
		}

		cur = next
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			writer.WriteString(strconv.Itoa(ans[i][j]))
			if j+1 < m {
				writer.WriteByte(' ')
			}
		}
		writer.WriteByte('\n')
	}
}
