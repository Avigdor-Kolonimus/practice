package backend

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Fenwick struct {
	tree []int
	n    int
}

func NewFenwick(n int) *Fenwick {
	return &Fenwick{
		tree: make([]int, n+1),
		n:    n,
	}
}

func (f *Fenwick) Add(i, delta int) {
	for i <= f.n {
		f.tree[i] += delta
		i += i & -i
	}
}

func (f *Fenwick) Sum(i int) int {
	res := 0

	for i > 0 {
		res += f.tree[i]
		i -= i & -i
	}

	return res
}

func (f *Fenwick) Kth(k int) int {
	idx := 0

	step := 1
	for step*2 <= f.n {
		step *= 2
	}

	for step > 0 {
		next := idx + step

		if next <= f.n && f.tree[next] < k {
			idx = next
			k -= f.tree[next]
		}

		step /= 2
	}

	return idx + 1
}

// https://coderun.yandex.ru/selections/first-2023-backend/problems/mtf-encoding-decoding
// MtfEncodingDecoding - problem 22
func MtfEncodingDecoding() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	params := strings.Fields(line)

	n, _ := strconv.Atoi(params[0])
	m, _ := strconv.Atoi(params[1])
	typ, _ := strconv.Atoi(params[2])

	size := n + m

	fw := NewFenwick(size)

	pos := make([]int, m+1)
	symbolAt := make([]int, size+1)
	for x := 1; x <= m; x++ {
		idx := n + x

		pos[x] = idx
		symbolAt[idx] = x

		fw.Add(idx, 1)
	}

	front := n
	for i := range n {
		line, _ = reader.ReadString('\n')
		fields := strings.Fields(line)

		for _, field := range fields {
			a, _ := strconv.Atoi(field)

			var x int
			if typ == 1 {
				x = a
				writer.WriteString(strconv.Itoa(fw.Sum(pos[x])))
			} else {
				idx := fw.Kth(a)
				x = symbolAt[idx]

				writer.WriteString(strconv.Itoa(x))
			}

			fw.Add(pos[x], -1)

			pos[x] = front
			symbolAt[front] = x

			fw.Add(front, 1)
			front--

			if i+1 < n {
				writer.WriteByte(' ')
			}

			i++
		}
		i--
	}

	writer.WriteByte('\n')
}
