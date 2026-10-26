package summerml2024

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	i int
	j int
}

// https://coderun.yandex.ru/selections/2024-summer-ml/problems/concatenation
// Concatenation - problem 23
func Concatenation() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	words := make([]string, n)
	dictionary := make(map[string]int, n)
	for i := range n {
		line, _ = reader.ReadString('\n')
		words[i] = strings.TrimSpace(line)
		dictionary[words[i]] = i + 1
	}

	res := make([]Pair, 0)
	for _, word := range words {
		index := dictionary[word]
		if index == 0 {
			continue
		}

		l := len(word)
		for i := 1; i < l; i++ {
			if 2*i > l {
				suffix := word[i:]

				if strings.HasPrefix(word, suffix) {
					partner := word[l-i : i]

					if j, ok := dictionary[partner]; ok && j != index {
						res = append(res, Pair{index, j})
					}
				}
			} else if 2*i < l {
				prefix := word[:i]

				if strings.HasSuffix(word[i:], prefix) {
					partner := word[i : l-i]

					if j, ok := dictionary[partner]; ok && j != index {
						res = append(res, Pair{j, index})
					}
				}
			}
		}

		delete(dictionary, word)
		dictionary[word] = index
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i].i != res[j].i {
			return res[i].i < res[j].i
		}
		return res[i].j < res[j].j
	})

	for _, pair := range res {
		writer.WriteString(strconv.Itoa(pair.i))
		writer.WriteByte(' ')
		writer.WriteString(strconv.Itoa(pair.j))
		writer.WriteByte('\n')
	}
}
