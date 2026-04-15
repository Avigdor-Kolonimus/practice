package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateMinimumOfTheSegmentNInput(n int) bool {
	return n >= 0 && n <= 150_000
}

func validateMinimumOfTheSegmentKInput(k int) bool {
	return k >= 0 && k <= 10_000
}

// https://coderun.yandex.ru/selections/hr-tech-interview/problems/minimum-of-the-segment
// MinimumOfTheSegment - problem 11
func MinimumOfTheSegment() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first line
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	nk := strings.Fields(line)
	if len(nk) != 2 {
		panic("nk count does not match 2")
	}

	n, err := strconv.Atoi(strings.TrimSpace(nk[0]))
	if err != nil {
		panic(err)
	}
	if !validateMinimumOfTheSegmentNInput(n) {
		panic("n out of range")
	}

	k, err := strconv.Atoi(strings.TrimSpace(nk[1]))
	if err != nil {
		panic(err)
	}
	if !validateMinimumOfTheSegmentKInput(k) {
		panic("k out of range")
	}
	if k > n {
		panic("k > n")
	}

	// segments input
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	segments := strings.Fields(line)
	if len(segments) != n {
		panic("segments count does not match")
	}
	numsSlice := make([]int, 0, n)
	for _, rawInt := range segments {
		num, err := strconv.Atoi(rawInt)
		if err != nil {
			panic(err)
		}

		numsSlice = append(numsSlice, num)
	}

	indexesSlice := make([]int, 0, k)
	for i := range segments {
		if len(indexesSlice) > 0 && indexesSlice[0] == i-k {
			indexesSlice = indexesSlice[1:]
		}

		for len(indexesSlice) > 0 && numsSlice[indexesSlice[len(indexesSlice)-1]] >= numsSlice[i] {
			indexesSlice = indexesSlice[:len(indexesSlice)-1]
		}

		indexesSlice = append(indexesSlice, i)

		if i >= k-1 {
			writer.WriteString(strconv.Itoa(numsSlice[indexesSlice[0]]))
			writer.WriteByte('\n')
		}
	}
}
