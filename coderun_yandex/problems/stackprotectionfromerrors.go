package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	data   []int
	writer *bufio.Writer
}

func (s *Stack) Push(n int) {
	s.data = append(s.data, n)

	s.writer.WriteString("ok")
	s.writer.WriteByte('\n')
}

func (s *Stack) Pop() {
	if len(s.data) == 0 {
		s.writer.WriteString("error")
		s.writer.WriteByte('\n')

		return
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	s.writer.WriteString(strconv.Itoa(val))
	s.writer.WriteByte('\n')
}

func (s *Stack) Back() {
	if len(s.data) == 0 {
		s.writer.WriteString("error")
		s.writer.WriteByte('\n')

		return
	}

	s.writer.WriteString(strconv.Itoa(s.data[len(s.data)-1]))
	s.writer.WriteByte('\n')
}

func (s *Stack) Size() {
	s.writer.WriteString(strconv.Itoa(len(s.data)))
	s.writer.WriteByte('\n')
}

func (s *Stack) Clear() {
	s.data = []int{}

	s.writer.WriteString("ok")
	s.writer.WriteByte('\n')
}

// https://coderun.yandex.ru/problem/stack-protection-from-errors
// StackProtectionFromErrors - problem 248
func StackProtectionFromErrors() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// stack
	stack := &Stack{
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
			stack.Push(n)

		case "pop":
			stack.Pop()

		case "back":
			stack.Back()

		case "size":
			stack.Size()

		case "clear":
			stack.Clear()

		case "exit":
			writer.WriteString("bye")
			writer.WriteByte('\n')
			return
		}
	}
}
