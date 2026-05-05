package problems

import (
	"bufio"
	"os"
	"sort"
)

// https://coderun.yandex.ru/problem/Histogram
// Histogram - problem 91
func Histogram() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	freq := make(map[byte]int, 60)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}

		if b == ' ' || b == '\n' {
			continue
		}

		freq[b]++
	}

	// sorting
	var chars []byte
	maxFreq := 0
	for c := range freq {
		if freq[c] > maxFreq {
			maxFreq = freq[c]
		}

		chars = append(chars, c)
	}
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	// histogram
	for level := maxFreq; level > 0; level-- {
		for _, c := range chars {
			if freq[c] >= level {
				writer.WriteByte('#')
			} else {
				writer.WriteByte(' ')
			}
		}
		writer.WriteByte('\n')
	}

	for _, c := range chars {
		writer.WriteByte(c)
	}
	writer.WriteByte('\n')
}
