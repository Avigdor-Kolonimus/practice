package problems

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Solver struct {
	n        int
	words    int
	lastMask uint64
	tree     [][]int
	answer   []uint64
}

func NewSolver(n int) *Solver {
	words := (n + 1 + 63) / 64

	lastMask := ^uint64(0)
	if rem := (n + 1) % 64; rem != 0 {
		lastMask = (uint64(1) << rem) - 1
	}

	return &Solver{
		n:        n,
		words:    words,
		lastMask: lastMask,
		tree:     make([][]int, 4*n+5),
		answer:   make([]uint64, words),
	}
}

func (s *Solver) add(v, l, r, ql, qr, x int) {
	if ql <= l && r <= qr {
		s.tree[v] = append(s.tree[v], x)
		return
	}

	mid := (l + r) / 2

	if ql <= mid {
		s.add(v*2, l, mid, ql, qr, x)
	}

	if qr > mid {
		s.add(v*2+1, mid+1, r, ql, qr, x)
	}
}

func (s *Solver) shiftOr(mask []uint64, shift int) {
	wordShift := shift / 64
	bitShift := uint(shift % 64)

	for i := s.words - 1; i >= wordShift; i-- {
		src := i - wordShift

		value := mask[src] << bitShift

		if bitShift > 0 && src > 0 {
			value |= mask[src-1] >> (64 - bitShift)
		}

		mask[i] |= value
	}

	mask[s.words-1] &= s.lastMask
}

func (s *Solver) dfs(v, l, r int, mask []uint64) {
	if len(s.tree[v]) > 0 {
		next := make([]uint64, s.words)
		copy(next, mask)
		mask = next

		for _, x := range s.tree[v] {
			s.shiftOr(mask, x)
		}
	}

	if l == r {
		for i := 0; i < s.words; i++ {
			s.answer[i] |= mask[i]
		}
		return
	}

	mid := (l + r) / 2

	s.dfs(v*2, l, mid, mask)
	s.dfs(v*2+1, mid+1, r, mask)
}

// https://coderun.yandex.ru/problem/questions-about-maximum
// QuestionsAboutMaximum - problem 417
func QuestionsAboutMaximum() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	if len(parts) < 2 {
		return
	}

	n, _ := strconv.Atoi(parts[0])
	q, _ := strconv.Atoi(parts[1])

	s := NewSolver(n)

	for range q {
		var l, r, x int
		fmt.Fscan(reader, &l, &r, &x)

		s.add(1, 1, n, l, r, x)
	}

	initial := make([]uint64, s.words)
	initial[0] = 1

	s.dfs(1, 1, n, initial)

	s.answer[0] &^= 1
	result := make([]int, 0)
	for i, word := range s.answer {
		for word != 0 {
			bit := 0

			for (word & (uint64(1) << uint(bit))) == 0 {
				bit++
			}

			value := i*64 + bit

			if value >= 1 && value <= n {
				result = append(result, value)
			}

			word &= word - 1
		}
	}

	fmt.Fprintln(writer, len(result))

	for i, value := range result {
		if i > 0 {
			writer.WriteByte(' ')
		}

		writer.WriteString(strconv.Itoa(value))
	}

	writer.WriteByte('\n')
}
