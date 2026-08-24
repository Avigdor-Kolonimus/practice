package backend

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	firstChild int32
	next       int32
	child      byte
	hasWord    bool
}

var trie []Node

func findChild(v int32, ch byte) int32 {
	for u := trie[v].firstChild; u != -1; u = trie[u].next {
		if trie[u].child == ch {
			return u
		}
	}

	return -1
}

func addChild(v int32, ch byte) int32 {
	u := int32(len(trie))

	trie = append(trie, Node{
		firstChild: -1,
		next:       trie[v].firstChild,
		child:      ch,
	})

	trie[v].firstChild = u

	return u
}

// https://coderun.yandex.ru/selections/backend/problems/autocomplete
// Autocomplete - problem 13
func Autocomplete() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	//  N input
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	// Words input
	line, _ = reader.ReadString('\n')
	words := strings.Fields(line)

	trie = make([]Node, 1)
	trie[0].firstChild = -1
	totalPresses := 0
	for i := range n {
		s := words[i]

		cur := int32(0)
		prefixIndex := -1

		for j := 0; j < len(s); j++ {
			ch := s[j]
			child := findChild(cur, ch)

			if child == -1 {
				child = addChild(cur, ch)
				prefixIndex = j
			}

			cur = child

			if j == len(s)-1 {
				if trie[cur].firstChild != -1 {
					prefixIndex = j
				}
				break
			}

			if trie[cur].hasWord {
				prefixIndex = -1
			} else {
				first := trie[cur].firstChild

				if first != -1 {
					second := trie[first].next

					if second == -1 && prefixIndex == -1 {
						prefixIndex = j
					} else if second != -1 {
						prefixIndex = -1
					}
				}
			}
		}

		trie[cur].hasWord = true

		if prefixIndex == -1 {
			totalPresses += len(s)
		} else {
			totalPresses += prefixIndex + 1
		}
	}

	writer.WriteString(strconv.Itoa(totalPresses))
	writer.WriteByte('\n')
}
