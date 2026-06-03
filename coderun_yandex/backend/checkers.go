package backend

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type PointCheckers struct {
	x int
	y int
}

// https://coderun.yandex.ru/selections/backend/problems/checkers
// Checkers - problem 31
func Checkers() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M line
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

	board := make(map[PointCheckers]int)

	// white pieces input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	w, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	white := make([]PointCheckers, w)
	for i := 0; i < w; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		white[i].x, err = strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		white[i].y, err = strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		board[white[i]] = 1
	}

	// black pieces input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	b, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	black := make([]PointCheckers, b)
	for i := 0; i < b; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		black[i].x, err = strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		black[i].y, err = strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		board[black[i]] = 2
	}

	// turn input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	turn := strings.TrimRight(line, "\r\n")

	var myPieces []PointCheckers
	enemyColor := 2
	if turn == "black" {
		myPieces = black
		enemyColor = 1
	} else {
		myPieces = white
	}

	dirs := [][2]int{
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}

	for _, p := range myPieces {
		for _, d := range dirs {
			mx := p.x + d[0]
			my := p.y + d[1]

			if board[PointCheckers{mx, my}] != enemyColor {
				continue
			}

			lx := p.x + 2*d[0]
			ly := p.y + 2*d[1]

			if lx < 1 || lx > n || ly < 1 || ly > m {
				continue
			}

			if _, occupied := board[PointCheckers{lx, ly}]; !occupied {
				writer.WriteString("Yes")
				writer.WriteByte('\n')

				return
			}
		}
	}

	writer.WriteString("No")
	writer.WriteByte('\n')
}
