package mgustokashin

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type DequeGoblinsAndChess struct {
	data       []int
	head, tail int
}

func NewDeque(size int) *DequeGoblinsAndChess {
	m := 2*size + 5
	return &DequeGoblinsAndChess{
		data: make([]int, m),
		head: m / 2,
		tail: m / 2,
	}
}

func (d *DequeGoblinsAndChess) Len() int {
	return d.tail - d.head
}

func (d *DequeGoblinsAndChess) PushFront(x int) {
	d.head--
	d.data[d.head] = x
}

func (d *DequeGoblinsAndChess) PushBack(x int) {
	d.data[d.tail] = x
	d.tail++
}

func (d *DequeGoblinsAndChess) PopFront() int {
	x := d.data[d.head]
	d.head++
	return x
}

func (d *DequeGoblinsAndChess) PopBack() int {
	d.tail--
	return d.data[d.tail]
}

func balance(left, right *DequeGoblinsAndChess) {
	if left.Len() < right.Len() {
		left.PushBack(right.PopFront())
	}
	if left.Len() > right.Len()+1 {
		right.PushFront(left.PopBack())
	}
}

// https://coderun.yandex.ru/problem/goblins-and-chess
// GoblinsAndChess - problem 4
func GoblinsAndChess() {
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

	left := NewDeque(n)
	right := NewDeque(n)

	for i := 0; i < n; i++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum := strings.Fields(line)

		op := strNum[0]

		switch op {
		case "+":
			x, err := strconv.Atoi(strNum[1])
			if err != nil {
				panic(err)
			}

			right.PushBack(x)
			balance(left, right)

		case "*":
			x, err := strconv.Atoi(strNum[1])
			if err != nil {
				panic(err)
			}

			right.PushFront(x)
			balance(left, right)

		case "-":
			writer.WriteString(strconv.Itoa(left.PopFront()))
			writer.WriteByte('\n')

			balance(left, right)
		}
	}
}
