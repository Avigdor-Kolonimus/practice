package summercommon2025

import (
	"bufio"
	"os"
	"strconv"
)

func solveThreeLetterBallad(ballad string, _ int) int64 {
	total := make([]int64, 26)

	for i := 0; i < len(ballad); i++ {
		c := ballad[i]
		if c == ' ' {
			continue
		}
		total[c-'a']++
	}

	left := make([]int64, 26)
	right := make([]int64, 26)
	copy(right, total)

	var ans int64 = 0

	for i := 0; i < len(ballad); i++ {
		c := ballad[i]
		if c == ' ' {
			continue
		}

		idx := c - 'a'
		right[idx]--

		for j := 0; j < 26; j++ {
			ans += left[j] * right[j]
		}

		left[idx]++
	}

	return ans
}

// https://coderun.yandex.ru/selections/2025-summer-common/problems/three-letter-ballad
// ThreeLetterBallad - problem 15
func ThreeLetterBallad() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// ballad input
	ballad := mustReadLine(reader)

	answer := solveThreeLetterBallad(ballad, len(ballad))

	writer.WriteString(strconv.FormatInt(answer, 10))
	writer.WriteByte('\n')
}
