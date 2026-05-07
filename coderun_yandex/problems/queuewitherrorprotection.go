package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	data   []int
	head   int
	writer *bufio.Writer
}

func (q *Queue) Push(x int) {
	q.data = append(q.data, x)

	q.writer.WriteString("ok")
	q.writer.WriteByte('\n')
}

func (q *Queue) Pop() (int, bool) {
	if (len(q.data) - q.head) == 0 {
		q.writer.WriteString("error")
		q.writer.WriteByte('\n')

		return 0, false
	}

	value := q.data[q.head]
	q.head++

	q.writer.WriteString(strconv.Itoa(value))
	q.writer.WriteByte('\n')

	return value, true
}

func (q *Queue) Front() (int, bool) {
	if (len(q.data) - q.head) == 0 {
		q.writer.WriteString("error")
		q.writer.WriteByte('\n')

		return 0, false
	}

	q.writer.WriteString(strconv.Itoa(q.data[q.head]))
	q.writer.WriteByte('\n')

	return q.data[q.head], true
}

func (q *Queue) Size() int {
	size := len(q.data) - q.head

	q.writer.WriteString(strconv.Itoa(size))
	q.writer.WriteByte('\n')

	return size
}

func (q *Queue) Clear() {
	q.data = make([]int, 0)
	q.head = 0

	q.writer.WriteString("ok")
	q.writer.WriteByte('\n')
}

// https://coderun.yandex.ru/problem/queue-with-error-protection
// QueueWithErrorProtection - problem 99
func QueueWithErrorProtection() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// queue
	queue := &Queue{
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

		case "push":
			n, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			queue.Push(n)

		case "pop":
			queue.Pop()

		case "front":
			queue.Front()

		case "size":
			queue.Size()

		case "clear":
			queue.Clear()

		case "exit":
			writer.WriteString("bye")
			writer.WriteByte('\n')
			return
		}
	}
}
