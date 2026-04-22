package algorithmtrainingseptember2025

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Trie struct {
	next   [][]int
	wordID []int
}

func NewTrie() *Trie {
	root := make([]int, 26)
	for i := 0; i < 26; i++ {
		root[i] = -1
	}

	return &Trie{
		next:   [][]int{root},
		wordID: []int{-1},
	}
}

func (t *Trie) Insert(bytes []byte, id int) {
	v := 0
	for _, b := range bytes {
		c := int(b - 'a')
		if t.next[v][c] == -1 {
			t.next[v][c] = len(t.next)

			newNode := make([]int, 26)
			for i := 0; i < 26; i++ {
				newNode[i] = -1
			}

			t.next = append(t.next, newNode)
			t.wordID = append(t.wordID, -1)
		}
		v = t.next[v][c]
	}
	t.wordID[v] = id
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/dictionary
// Dictionary - assignment 14
func Dictionary() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// Text input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	text := strings.TrimRight(line, "\r\n")
	if len(text) > 100 {
		panic("text length exceeds 100")
	}

	s := []byte(text)
	m := len(s)

	// N input
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if n > 2000 {
		panic("N is out of range")
	}

	words := make([]string, 0, n)
	trie := NewTrie()

	// Read words
	for i := range n {
		line, _ = reader.ReadString('\n')
		w := strings.TrimSpace(line)
		words = append(words, w)
		trie.Insert([]byte(w), i)
	}

	// DP arrays
	dp := make([]bool, m+1)
	nextPos := make([]int, m+1)
	nextWord := make([]int, m+1)

	for i := 0; i <= m; i++ {
		nextPos[i] = -1
		nextWord[i] = -1
	}

	dp[m] = true

	// Backward DP
	for i := m - 1; i >= 0; i-- {
		v := 0
		limit := i + 20
		if limit > m {
			limit = m
		}

		for j := i; j < limit; j++ {
			c := int(s[j] - 'a')
			nx := trie.next[v][c]
			if nx == -1 {
				break
			}
			v = nx

			wid := trie.wordID[v]
			if wid != -1 && dp[j+1] {
				dp[i] = true
				nextPos[i] = j + 1
				nextWord[i] = wid
				break
			}
		}
	}

	// Reconstruct answer
	result := []string{}
	pos := 0

	for pos < m {
		if nextWord[pos] == -1 {
			break // safety guard
		}
		result = append(result, words[nextWord[pos]])
		pos = nextPos[pos]
	}

	writer.WriteString(strings.Join(result, " "))
	writer.WriteByte('\n')
}
