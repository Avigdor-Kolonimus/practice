package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Deque struct {
	data   []int
	writer *bufio.Writer
}

func (d *Deque) PushFront(n int) {
	d.data = append([]int{n}, d.data...)

	d.writer.WriteString("ok")
	d.writer.WriteByte('\n')
}

func (d *Deque) PushBack(n int) {
	d.data = append(d.data, n)

	d.writer.WriteString("ok")
	d.writer.WriteByte('\n')
}

func (d *Deque) PopFront() {
	if len(d.data) == 0 {
		d.writer.WriteString("error")
		d.writer.WriteByte('\n')

		return
	}
	val := d.data[0]
	d.data = d.data[1:]

	d.writer.WriteString(strconv.Itoa(val))
	d.writer.WriteByte('\n')
}

func (d *Deque) PopBack() {
	if len(d.data) == 0 {
		d.writer.WriteString("error")
		d.writer.WriteByte('\n')

		return
	}
	val := d.data[len(d.data)-1]
	d.data = d.data[:len(d.data)-1]

	d.writer.WriteString(strconv.Itoa(val))
	d.writer.WriteByte('\n')
}

func (d *Deque) Front() {
	if len(d.data) == 0 {
		d.writer.WriteString("error")
		d.writer.WriteByte('\n')

		return
	}

	d.writer.WriteString(strconv.Itoa(d.data[0]))
	d.writer.WriteByte('\n')
}

// back
func (d *Deque) Back() {
	if len(d.data) == 0 {
		d.writer.WriteString("error")
		d.writer.WriteByte('\n')

		return
	}

	d.writer.WriteString(strconv.Itoa(d.data[len(d.data)-1]))
	d.writer.WriteByte('\n')
}

func (d *Deque) Size() {
	d.writer.WriteString(strconv.Itoa(len(d.data)))
	d.writer.WriteByte('\n')
}

func (d *Deque) Clear() {
	d.data = []int{}

	d.writer.WriteString("ok")
	d.writer.WriteByte('\n')
}

// https://coderun.yandex.ru/problem/dec-with-error-protection
// DecWithErrorProtection - problem 101
func DecWithErrorProtection() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// deque
	deque := &Deque{
		writer: writer,
	}
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		parts := strings.Fields(line)

		switch parts[0] {

		case "push_front":
			n, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			deque.PushFront(n)

		case "push_back":
			n, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			deque.PushBack(n)

		case "pop_front":
			deque.PopFront()

		case "pop_back":
			deque.PopBack()

		case "front":
			deque.Front()

		case "back":
			deque.Back()

		case "size":
			deque.Size()

		case "clear":
			deque.Clear()

		case "exit":
			writer.WriteString("bye")
			writer.WriteByte('\n')
			return
		}
	}
}
