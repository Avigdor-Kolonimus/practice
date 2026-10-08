package first2023mobiledev

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func compatible(first []int, firstPos int, second []int, secondPos int) bool {
	left := max(firstPos, secondPos)
	right := min(
		firstPos+len(first),
		secondPos+len(second),
	)

	if left >= right {
		return true
	}

	for pos := left; pos < right; pos++ {
		firstIndex := pos - firstPos
		secondIndex := pos - secondPos

		if first[firstIndex] != second[secondIndex] {
			return false
		}
	}

	return true
}

func placeSequence(result []int, used []bool, sequence []int, position int, left int) bool {
	for i, value := range sequence {
		index := position + i - left

		if used[index] {
			if result[index] != value {
				return false
			}
		} else {
			result[index] = value
			used[index] = true
		}
	}

	return true
}

func buildResult(seq1 []int, seq2 []int, seq3 []int, pos2 int, pos3 int) ([]int, bool) {
	left := min(0, min(pos2, pos3))

	right := max(
		len(seq1),
		max(
			pos2+len(seq2),
			pos3+len(seq3),
		),
	)

	result := make([]int, right-left)
	used := make([]bool, right-left)

	if !placeSequence(result, used, seq1, 0, left) {
		return nil, false
	}

	if !placeSequence(result, used, seq2, pos2, left) {
		return nil, false
	}

	if !placeSequence(result, used, seq3, pos3, left) {
		return nil, false
	}

	return result, true
}

func solveForOrder(seq1 []int, seq2 []int, seq3 []int) []int {
	var best []int
	for pos2 := -len(seq2); pos2 <= len(seq1); pos2++ {
		if !compatible(seq1, 0, seq2, pos2) {
			continue
		}

		for pos3 := -len(seq3); pos3 <= len(seq1); pos3++ {
			if !compatible(seq1, 0, seq3, pos3) {
				continue
			}

			result, ok := buildResult(seq1, seq2, seq3, pos2, pos3)

			if !ok {
				continue
			}

			if best == nil || len(result) < len(best) {
				best = result
			}
		}
	}

	return best
}

func parseSequence(line string) []int {
	parts := strings.Fields(line)

	n, _ := strconv.Atoi(parts[0])

	result := make([]int, n)
	for i := range n {
		result[i], _ = strconv.Atoi(parts[i+1])
	}

	return result
}

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/sequences
// Sequences - problem 24
func Sequences() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	seq1 := parseSequence(line)

	line, _ = reader.ReadString('\n')
	seq2 := parseSequence(line)

	line, _ = reader.ReadString('\n')
	seq3 := parseSequence(line)

	permutations := [][3][]int{
		{seq1, seq2, seq3},
		{seq1, seq3, seq2},
		{seq2, seq1, seq3},
		{seq2, seq3, seq1},
		{seq3, seq1, seq2},
		{seq3, seq2, seq1},
	}

	var best []int
	for _, permutation := range permutations {
		result := solveForOrder(permutation[0], permutation[1], permutation[2])

		if best == nil || len(result) < len(best) {
			best = result
		}
	}

	writer.WriteString(strconv.Itoa(len(best)))
	writer.WriteByte('\n')

	for i, value := range best {
		if i > 0 {
			writer.WriteByte(' ')
		}

		writer.WriteString(strconv.Itoa(value))
	}

	writer.WriteByte('\n')
}
