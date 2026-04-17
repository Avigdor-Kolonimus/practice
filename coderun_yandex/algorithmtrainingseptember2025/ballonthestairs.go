package algorithmtrainingseptember2025

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateBallOnTheStairsInput(n int) bool {
	return n >= 1 && n <= 30
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/ball-on-the-stairs
// BallOnTheStairs - assignment 11
func BallOnTheStairs() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateBallOnTheStairsInput(n) {
		panic("number N out of range")
	}

	a, b, c := 1, 1, 2
	switch n {
	case 0:
		c = 1
	case 1:
		c = 1
	case 2:
		c = 2
	default:
		for i := 3; i <= n; i++ {
			a, b, c = b, c, a+b+c
		}
	}

	writer.WriteString(strconv.Itoa(c))
	writer.WriteByte('\n')
}
