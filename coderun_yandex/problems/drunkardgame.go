package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type QueueDrunkardGame struct {
	data []int
}

func (q *QueueDrunkardGame) Push(v int) {
	q.data = append(q.data, v)
}

func (q *QueueDrunkardGame) Pop() int {
	v := q.data[0]
	q.data = q.data[1:]

	return v
}

func (q *QueueDrunkardGame) Size() int {
	return len(q.data)
}

// https://coderun.yandex.ru/problem/drunkard-game
// DrunkardGame - problem 100
func DrunkardGame() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// playerA input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	tokens := strings.Fields(line)
	if len(tokens) != 5 {
		panic("invalid input")
	}
	cardsA := make([]int, 5)
	for i := range cardsA {
		n, err := strconv.Atoi(tokens[i])
		if err != nil {
			panic(err)
		}

		cardsA[i] = n
	}

	// playerB input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	tokens = strings.Fields(line)
	if len(tokens) != 5 {
		panic("invalid input")
	}
	cardsB := make([]int, 5)
	for i := range cardsB {
		n, err := strconv.Atoi(tokens[i])
		if err != nil {
			panic(err)
		}

		cardsB[i] = n
	}

	// game
	player1 := QueueDrunkardGame{data: cardsA}
	player2 := QueueDrunkardGame{data: cardsB}
	steps := 0
	for ; player1.Size() > 0 && player2.Size() > 0 && steps < 1_000_000; steps++ {
		c1 := player1.Pop()
		c2 := player2.Pop()

		if c1 == 0 && c2 == 9 {
			player1.Push(c1)
			player1.Push(c2)
		} else if c1 == 9 && c2 == 0 {
			player2.Push(c1)
			player2.Push(c2)
		} else if c1 > c2 {
			player1.Push(c1)
			player1.Push(c2)
		} else {
			player2.Push(c1)
			player2.Push(c2)
		}
	}

	if player1.Size() > 0 && player2.Size() > 0 {
		writer.WriteString("botva")
	} else if player1.Size() == 0 {
		writer.WriteString("second " + strconv.Itoa(steps))
	} else {
		writer.WriteString("first " + strconv.Itoa(steps))
	}
	writer.WriteByte('\n')
}
