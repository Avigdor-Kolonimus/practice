package backend

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Block struct {
	start int
	end   int
}

func digits(x int) int {
	switch {
	case x < 10:
		return 1
	case x < 100:
		return 2
	case x < 1_000:
		return 3
	case x < 10_000:
		return 4
	case x < 100_000:
		return 5
	case x < 1_000_000:
		return 6
	case x < 10_000_000:
		return 7
	case x < 100_000_000:
		return 8
	case x < 1_000_000_000:
		return 9
	default:
		return 10
	}
}

func encodedLength(length int) int {
	if length == 1 {
		return 1
	}

	return 1 + digits(length)
}

// https://coderun.yandex.ru/selections/backend/problems/find-rle-string-length
// FindRLEStringLength - problem 16
func FindRLEStringLength() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// Read compressed string
	line, _ := reader.ReadString('\n')
	s := strings.TrimSpace(line)

	// Parse RLE string
	pos := 1
	blocks := make([]Block, 0, len(s))
	starts := make([]int, 0, len(s))
	for i := 0; i < len(s); {
		count := 1

		// Read number if it exists
		if s[i] >= '0' && s[i] <= '9' {
			j := i

			for j < len(s) && s[j] >= '0' && s[j] <= '9' {
				j++
			}

			count, _ = strconv.Atoi(s[i:j])
			i = j
		}

		// s[i] is the character of the block
		i++

		start := pos
		end := pos + count - 1

		blocks = append(blocks, Block{
			start: start,
			end:   end,
		})

		starts = append(starts, start)

		pos = end + 1
	}

	// prefix[i] = compressed length of first i complete blocks
	prefix := make([]int, len(blocks)+1)

	for i := range blocks {
		length := blocks[i].end - blocks[i].start + 1
		prefix[i+1] = prefix[i] + encodedLength(length)
	}

	// Read q
	line, _ = reader.ReadString('\n')
	q, _ := strconv.Atoi(strings.TrimSpace(line))

	for ; q > 0; q-- {
		line, _ = reader.ReadString('\n')
		fields := strings.Fields(line)

		l, _ := strconv.Atoi(fields[0])
		r, _ := strconv.Atoi(fields[1])

		// Find the block containing l
		leftBlock := sort.Search(len(starts), func(i int) bool {
			return starts[i] > l
		}) - 1

		// Find the block containing r
		rightBlock := sort.Search(len(starts), func(i int) bool {
			return starts[i] > r
		}) - 1

		var answer int

		// Both positions are inside the same block
		if leftBlock == rightBlock {
			answer = encodedLength(r - l + 1)
		} else {
			// Part of the left block
			leftLength := blocks[leftBlock].end - l + 1
			answer = encodedLength(leftLength)

			// Complete blocks between left and right
			answer += prefix[rightBlock] - prefix[leftBlock+1]

			// Part of the right block
			rightLength := r - blocks[rightBlock].start + 1
			answer += encodedLength(rightLength)
		}

		writer.WriteString(strconv.Itoa(answer))
		writer.WriteByte('\n')
	}
}
