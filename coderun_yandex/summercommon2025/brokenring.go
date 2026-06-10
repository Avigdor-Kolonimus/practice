package summercommon2025

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

func union(parent []int, x, y int) {
	rootX := find(parent, x)
	rootY := find(parent, y)
	if rootX != rootY {
		parent[rootX] = rootY
	}
}

func calculateBrokenRingAnswer(n int, a []int, b []int) int {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	for i := 0; i < n; i++ {
		union(parent, i, a[i]-1)
	}

	root := find(parent, 0)
	allConnected := true
	for i := 1; i < n; i++ {
		if find(parent, i) != root {
			allConnected = false
			break
		}
	}
	if allConnected {
		return 0
	}

	d := make(map[int]int)
	for i := 0; i < n; i++ {
		r := find(parent, i)
		if _, exists := d[r]; !exists {
			d[r] = math.MaxInt
		}
		if b[i] < d[r] {
			d[r] = b[i]
		}
	}

	answer := 0
	for _, val := range d {
		answer += val
	}

	return answer
}

// ввод/вывод
// не изменяйте сигнатуру метода
// https://coderun.yandex.ru/selections/2025-summer-common/problems/broken-ring
// BrokenRing - problem 19
func BrokenRing() {
	input := NewFastScanner(os.Stdin)

	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	t := input.readInt()
	for test := 0; test < t; test++ {
		n := input.readInt()
		a := input.readIntArray(n)
		b := input.readIntArray(n)

		answer := calculateBrokenRingAnswer(n, a, b)
		output.WriteString(strconv.Itoa(answer))
		output.WriteByte('\n')
	}
}

type FastScanner struct {
	reader *bufio.Reader
	tokens []string
	pos    int
}

func NewFastScanner(r io.Reader) *FastScanner {
	return &FastScanner{
		reader: bufio.NewReader(r),
		tokens: make([]string, 0),
		pos:    0,
	}
}

func (fs *FastScanner) loadTokens() {
	for {
		line, err := fs.reader.ReadString('\n')
		if err != nil {
			fs.tokens = nil
			return
		}
		fs.tokens = strings.Fields(line)
		if len(fs.tokens) > 0 {
			fs.pos = 0
			return
		}
	}
}

func (fs *FastScanner) readToken() string {
	if fs.pos >= len(fs.tokens) {
		fs.loadTokens()
	}
	if fs.tokens == nil || fs.pos >= len(fs.tokens) {
		return ""
	}
	token := fs.tokens[fs.pos]
	fs.pos++
	return token
}

func (fs *FastScanner) readInt() int {
	val, _ := strconv.Atoi(fs.readToken())
	return val
}

func (fs *FastScanner) readIntArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = fs.readInt()
	}
	return arr
}
