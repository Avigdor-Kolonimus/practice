package eserajim

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func validateBoardWithCoinsInput(p int) bool {
	return p >= 1 && p <= 20
}

func getID(i, j, m int) int {
	return i*m + j
}

// https://coderun.yandex.ru/selections/eserajim/problems/board-with-coins
// BoardWithCoins - problem 4
func BoardWithCoins() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	line, err := reader.ReadString('\n')
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

	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	if !validateBoardWithCoinsInput(n * m) {
		panic("number N*M out of range")
	}

	// grid
	grid := make([]string, n)
	for i := range n {
		// points input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		grid[i] = line
	}

	// start
	start := 0
	pos := 0
	for i := range n {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				start |= (1 << pos)
			}
			pos++
		}
	}

	// targets
	target1, target2 := 0, 0
	pos = 0
	for i := range n {
		for j := 0; j < m; j++ {
			if (i+j)%2 == 0 {
				target1 |= (1 << pos)
			} else {
				target2 |= (1 << pos)
			}
			pos++
		}
	}

	// moves
	moves := []int{}
	for i := range n {
		for j := range m {
			if i+1 < n {
				a := getID(i, j, m)
				b := getID(i+1, j, m)
				moves = append(moves, (1<<a)|(1<<b))
			}
			if j+1 < m {
				a := getID(i, j, m)
				b := getID(i, j+1, m)
				moves = append(moves, (1<<a)|(1<<b))
			}
		}
	}

	maxState := 1 << (n * m)

	dist := make([]int, maxState)
	for i := range dist {
		dist[i] = -1
	}

	queue := make([]int, maxState)
	head, tail := 0, 0

	// start
	dist[start] = 0
	queue[tail] = start
	tail++

	for head < tail {
		cur := queue[head]
		head++

		if cur == target1 || cur == target2 {
			writer.WriteString(strconv.Itoa(dist[cur]))
			writer.WriteByte('\n')
			return
		}

		for _, mv := range moves {
			next := cur ^ mv
			if dist[next] == -1 {
				dist[next] = dist[cur] + 1
				queue[tail] = next
				tail++
			}
		}
	}

	writer.WriteString(strconv.Itoa(-1))
	writer.WriteByte('\n')
}
