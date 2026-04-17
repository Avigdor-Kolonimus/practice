package algorithmtrainingseptember2025

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateNewRegulationsInput(p int) bool {
	return p >= 1 && p <= 1_000_000_000
}

func calculateMoves(fromX, fromY, toF, toG int) (actionsCount int) {
	if fromX == toF && fromY == toG {
		return 0
	}

	north := fromY - toG
	west := fromX - toF

	south := toG - fromY
	east := toF - fromX

	for _, requiredMoves := range [...]int{north, west, south, east} {
		if requiredMoves < 1 {
			continue
		}

		// first move in direction we did in one step
		actionsCount++
		requiredMoves--

		// move left, return, move left
		actionsCount += 3 * requiredMoves
	}

	// first move is free
	return actionsCount - 1
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/new-regulations
// NewRegulations - assignment 9
func NewRegulations() {
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

	// X
	x, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	if !validateNewRegulationsInput(x) {
		panic("X out of range")
	}

	// Y
	y, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	if !validateNewRegulationsInput(y) {
		panic("y out of range")
	}

	// second line
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	strNum = strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	// F
	f, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	if !validateNewRegulationsInput(f) {
		panic("F out of range")
	}

	// G
	g, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	if !validateNewRegulationsInput(g) {
		panic("G out of range")
	}

	result := calculateMoves(x, y, f, g)

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
