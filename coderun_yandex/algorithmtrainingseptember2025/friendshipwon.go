package algorithmtrainingseptember2025

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func absFriendshipWon(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/friendship-won
// FriendshipWon - assignment 34
func FriendshipWon() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// tables input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	tables := make([]int, n)
	for i := range n {
		t, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		tables[i] = t
	}

	l, r := 0, n-1
	leftSum := tables[l]
	rightSum := tables[r]

	bestDiff := absFriendshipWon(leftSum - rightSum)
	bestL := l + 1
	bestR := r + 1

	for l+1 < r {
		if leftSum < rightSum {
			l++
			leftSum += tables[l]
		} else {
			r--
			rightSum += tables[r]
		}

		diff := absFriendshipWon(leftSum - rightSum)

		if diff < bestDiff {
			bestDiff = diff
			bestL = l + 1
			bestR = r + 1
		}
	}

	writer.WriteString(strconv.Itoa(bestDiff) + " " + strconv.Itoa(bestL) + " " + strconv.Itoa(bestR))
	writer.WriteByte('\n')
}
