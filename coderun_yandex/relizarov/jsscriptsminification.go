package relizarov

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Token struct {
	text string
	kind uint8
}

const (
	ReservedJsScriptsMinification uint8 = iota
	NumberJsScriptsMinification
	WordJsScriptsMinification
)

type TrieNode struct {
	children map[byte]*TrieNode
	terminal bool
}

type SolverJsScriptsMinification struct {
	reserved map[string]bool
	trie     *TrieNode

	rename map[string]string
	nextID int

	tokens []Token
}

func NewSolverJsScriptsMinification(keywords []string) *SolverJsScriptsMinification {
	s := &SolverJsScriptsMinification{
		reserved: make(map[string]bool, len(keywords)),
		trie: &TrieNode{
			children: make(map[byte]*TrieNode),
		},
		rename: make(map[string]string),
	}

	for _, kw := range keywords {
		s.reserved[kw] = true

		node := s.trie

		for i := 0; i < len(kw); i++ {
			ch := kw[i]

			if node.children[ch] == nil {
				node.children[ch] = &TrieNode{
					children: make(map[byte]*TrieNode),
				}
			}

			node = node.children[ch]
		}

		node.terminal = true
	}

	return s
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z')
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isWordStart(c byte) bool {
	return isLetter(c) || c == '_' || c == '$'
}

func isWordPart(c byte) bool {
	return isWordStart(c) || isDigit(c)
}

func (s *SolverJsScriptsMinification) ParseToken(src string, pos int) (Token, int) {
	bestEnd := pos
	var bestKind uint8

	node := s.trie
	for i := pos; i < len(src); i++ {
		node = node.children[src[i]]

		if node == nil {
			break
		}

		if node.terminal {
			bestEnd = i + 1
			bestKind = ReservedJsScriptsMinification
		}
	}

	if isWordStart(src[pos]) {
		end := pos + 1

		for end < len(src) && isWordPart(src[end]) {
			end++
		}

		if end > bestEnd {
			bestEnd = end
			bestKind = WordJsScriptsMinification
		}
	}

	if isDigit(src[pos]) {
		end := pos + 1

		for end < len(src) && isDigit(src[end]) {
			end++
		}

		if end > bestEnd {
			bestEnd = end
			bestKind = NumberJsScriptsMinification
		}
	}

	if bestEnd == pos {
		return Token{}, pos
	}

	return Token{
		text: src[pos:bestEnd],
		kind: bestKind,
	}, bestEnd
}

func (s *SolverJsScriptsMinification) ParseLine(line string) {
	if idx := strings.IndexByte(line, '#'); idx >= 0 {
		line = line[:idx]
	}

	pos := 0
	for pos < len(line) {
		if line[pos] == ' ' {
			pos++
			continue
		}

		token, next := s.ParseToken(line, pos)

		if next == pos {
			pos++
			continue
		}

		s.tokens = append(s.tokens, token)
		pos = next
	}
}

func encodeID(index int) string {
	index++

	var buf [16]byte
	pos := len(buf)

	for index > 0 {
		index--
		pos--

		buf[pos] = byte('a' + index%26)
		index /= 26
	}

	return string(buf[pos:])
}

func (s *SolverJsScriptsMinification) NextIdentifier() string {
	for {
		id := encodeID(s.nextID)
		s.nextID++

		if !s.reserved[id] {
			return id
		}
	}
}

func (s *SolverJsScriptsMinification) Rename() {
	for i := range s.tokens {
		token := &s.tokens[i]

		if token.kind != WordJsScriptsMinification {
			continue
		}

		id, exists := s.rename[token.text]

		if !exists {
			id = s.NextIdentifier()
			s.rename[token.text] = id
		}

		token.text = id
	}
}

func (s *SolverJsScriptsMinification) Matches(src string, expected []Token) bool {
	pos := 0
	for _, want := range expected {
		if pos >= len(src) {
			return false
		}

		got, next := s.ParseToken(src, pos)

		if next == pos {
			return false
		}

		if got.text != want.text || got.kind != want.kind {
			return false
		}

		pos = next
	}

	return pos == len(src)
}

func (s *SolverJsScriptsMinification) Minimize() string {
	var result strings.Builder

	var group strings.Builder
	groupTokens := make([]Token, 0)
	for _, token := range s.tokens {
		group.WriteString(token.text)
		groupTokens = append(groupTokens, token)

		candidate := group.String()

		if s.Matches(candidate, groupTokens) {
			continue
		}

		groupTokens = groupTokens[:len(groupTokens)-1]

		previous := group.String()
		previous = previous[:len(previous)-len(token.text)]

		if result.Len() > 0 {
			result.WriteByte(' ')
		}

		result.WriteString(previous)

		group.Reset()
		group.WriteString(token.text)

		groupTokens = groupTokens[:0]
		groupTokens = append(groupTokens, token)
	}

	if group.Len() > 0 {
		if result.Len() > 0 {
			result.WriteByte(' ')
		}

		result.WriteString(group.String())
	}

	return result.String()
}

// https://coderun.yandex.ru/selections/relizarov/problems/js-scripts-minification
// JsScriptsMinification - problem 5
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	line, _ = reader.ReadString('\n')
	keywords := strings.Fields(line)

	if n == 0 {
		keywords = nil
	}

	solver := NewSolverJsScriptsMinification(keywords)

	line, _ = reader.ReadString('\n')
	m, _ := strconv.Atoi(strings.TrimSpace(line))

	for range m {
		line, _ = reader.ReadString('\n')

		line = strings.TrimSuffix(line, "\n")
		line = strings.TrimSuffix(line, "\r")

		solver.ParseLine(line)
	}

	solver.Rename()

	result := solver.Minimize()

	writer.WriteString(result)
	writer.WriteByte('\n')
}
