package mobiledev

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/seventh-chord-2
// SeventhChord2 - problem 30
func SeventhChord2() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	maxA := 0
	a := make([]int, n)
	for i := 0; i < n; {
		line, _ = reader.ReadString('\n')
		fields := strings.Fields(line)

		for _, field := range fields {
			v, _ := strconv.Atoi(field)
			a[i] = v
			if v > maxA {
				maxA = v
			}
			i++
			if i == n {
				break
			}
		}
	}

	kLimit := 0
	for (kLimit+1)*(kLimit+1) <= maxA && kLimit < n {
		kLimit++
	}

	var answer int64
	pos := make([]int, 0, n)
	for k := 1; k <= kLimit; k++ {
		threshold := k * k
		pos = pos[:0]

		for i, v := range a {
			if v >= threshold {
				pos = append(pos, i+1)
			}
		}

		m := len(pos)
		if m < k {
			continue
		}

		for j := k; j <= m; j++ {
			leftCount := pos[j-k]
			if j > k {
				leftCount -= pos[j-k-1]
			}

			rightCount := n - pos[j-1] + 1
			answer += int64(leftCount) * int64(rightCount)
		}
	}

	writer.WriteString(strconv.FormatInt(answer, 10))
	writer.WriteByte('\n')
}
