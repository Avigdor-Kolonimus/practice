package backend

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/backend/problems/diversity-scoring
// DiversityScoring - problem 48
func DiversityScoring() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// product and category inputs
	productToCategory := make(map[int]int, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum := strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		p, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		c, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		productToCategory[p] = c
	}

	// products input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)

	order := make([]int, len(strNum))
	for i := 0; i < len(strNum); i++ {
		p, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		order[i] = p
	}

	ans := n
	lastPos := make(map[int]int)
	for pos, productID := range order {
		category := productToCategory[productID]

		if prev, ok := lastPos[category]; ok {
			dist := pos - prev
			if dist < ans {
				ans = dist
			}
		}

		lastPos[category] = pos
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
