package algorithmtrainingseptember2025

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateIncrementalScoreboardNumAndSecInput(p int) bool {
	return p >= 0 && p <= 1_000_000_000
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/incremental-scoreboard
// Incremental Scoreboard - assignment 5
func IncrementalScoreboard() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first line
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	// num
	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	if !validateIncrementalScoreboardNumAndSecInput(n) {
		panic("N out of range")
	}

	// sec
	k, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	if !validateIncrementalScoreboardNumAndSecInput(k) {
		panic("K out of range")
	}

	for ; k > 0; k-- {
		lastDigit := n % 10

		if lastDigit == 0 {
			break
		} else if lastDigit == 2 && k > 4 {
			cycles := k >> 2
			n += cycles * 20 // 20 = (2 + 4 + 8 + 6)
			k = k%4 + 1
			continue
		}

		n += lastDigit
	}

	writer.WriteString(strconv.Itoa(n))
	writer.WriteByte('\n')
}
