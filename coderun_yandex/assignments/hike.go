package assignments

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateHikeInput(n int) bool {
	return n <= 200
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/hike
// Hike - assignment 12
func Hike() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	line = strings.TrimRight(line, "\r\n")
	if !validateHikeInput(len(line)) {
		panic("lenght of string out of range")
	}

	left, right := 0, 1
	for _, label := range line {
		switch label {
		case 'L':
			left = min(left+1, right+1)
		case 'R':
			right = min(left+1, right+1)
		default:
			left++
			right++
		}
	}
	result := min(left+1, right)

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
