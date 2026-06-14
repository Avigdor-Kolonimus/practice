package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/test-training-sessions
// TestTrainingSessions - assignment 22
func TestTrainingSessions() {
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

	// right input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	right := strings.TrimRight(line, "\r\n")

	// m input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	m, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// student answer input
	ans := make([]string, m)
	correctCnt := make([]int, m)
	for i := 0; i < m; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}

		ans[i] = strings.TrimRight(line, "\r\n")
		for j := 0; j < n; j++ {
			if ans[i][j] == right[j] {
				correctCnt[i]++
			}
		}
	}

	var pairs [][2]int
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {

			correct := 0
			wrong := 0

			for k := 0; k < n; k++ {
				if ans[i][k] == ans[j][k] {
					if ans[i][k] == right[k] {
						correct++
					} else {
						wrong++
					}
				}
			}

			c1 := correctCnt[i]
			c2 := correctCnt[j]

			if 2*correct > c1 && 2*correct > c2 && 2*wrong > (n-c1) && 2*wrong > (n-c2) {
				pairs = append(pairs, [2]int{i + 1, j + 1})
			}
		}
	}

	writer.WriteString(strconv.Itoa(len(pairs)))
	writer.WriteByte('\n')
	for _, p := range pairs {
		writer.WriteString(strconv.Itoa(p[0]))
		writer.WriteByte(' ')
		writer.WriteString(strconv.Itoa(p[1]))
		writer.WriteByte('\n')
	}
}
