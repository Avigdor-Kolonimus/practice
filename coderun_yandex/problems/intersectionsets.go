package problems

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/intersection-sets
// IntersectionSets - problem 65
func IntersectionSets() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first line input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	tokens := strings.Fields(line)

	set1 := make(map[int]struct{}, len(tokens))
	for i := range tokens {
		n, err := strconv.Atoi(tokens[i])
		if err != nil {
			panic(err)
		}

		set1[n] = struct{}{}
	}

	// second line input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	tokens = strings.Fields(line)
	set2 := make(map[int]struct{}, len(tokens))
	result := make([]int, 0, len(set1))
	for i := range tokens {
		n, err := strconv.Atoi(tokens[i])
		if err != nil {
			panic(err)
		}

		set2[n] = struct{}{}
		if _, ok := set1[n]; ok {
			result = append(result, n)
		}
	}

	sort.Ints(result)

	// output
	for i, x := range result {
		if i > 0 {
			writer.WriteByte(' ')
		}
		writer.WriteString(strconv.Itoa(x))
	}
	writer.WriteByte('\n')
}
