package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/numbers
// Numbers - problem 44
func Numbers() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	lCount := len(line)
	if lCount != 4 {
		panic("digits does not match 4")
	}

	startNum, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// second input
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	line = strings.TrimRight(line, "\r\n")
	lCount = len(line)
	if lCount != 4 {
		panic("digits does not match 4")
	}

	endNum, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	parent := make(map[int]int)
	queue := []int{startNum}
	parent[startNum] = -1

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == endNum {
			break
		}

		digit1 := current / 1000
		digit2 := current / 100 % 10
		digit3 := current / 10 % 10
		digit4 := current % 10

		neighbors := make([]int, 0, 4)

		if digit1 < 9 {
			neighbors = append(neighbors, (digit1+1)*1000+digit2*100+digit3*10+digit4)
		}

		if digit4 > 1 {
			neighbors = append(neighbors, digit1*1000+digit2*100+digit3*10+(digit4-1))
		}

		neighbors = append(neighbors, digit4*1000+digit1*100+digit2*10+digit3)
		neighbors = append(neighbors, digit2*1000+digit3*100+digit4*10+digit1)

		for _, nb := range neighbors {
			if _, exists := parent[nb]; !exists {
				parent[nb] = current
				queue = append(queue, nb)
			}
		}
	}

	path := []int{}
	for curr := endNum; curr != -1; curr = parent[curr] {
		path = append(path, curr)
	}

	for i := len(path) - 1; i >= 0; i-- {
		writer.WriteString(strconv.Itoa(path[i]))
		writer.WriteByte('\n')
	}
}
