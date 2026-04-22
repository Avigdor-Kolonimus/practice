package backend

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func mustReadLine(reader *bufio.Reader) string {
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(line)
}

func mustReadIntArray(reader *bufio.Reader, size int) []int {
	rawNumbers := strings.Split(mustReadLine(reader), " ")
	if len(rawNumbers) != size {
		panic("len must be eq size")
	}

	result := make([]int, 0, size)
	for _, rawNumber := range rawNumbers {
		number, err := strconv.Atoi(rawNumber)
		if err != nil {
			panic(err)
		}
		result = append(result, number)
	}

	return result
}

func validateTradingYaInternInput(n int) bool {
	return n >= 1 && n <= 100_000
}

// https://coderun.yandex.ru/selections/backend/problems/trading-ya-intern
// TradingYaIntern - problem 40
func TradingYaIntern() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	firstLine := mustReadIntArray(reader, 2)
	if !validateTradingYaInternInput(firstLine[0]) {
		panic("number N out of range")
	}
	if !validateTradingYaInternInput(firstLine[1]) {
		panic("number M out of range")
	}
	n, m := firstLine[0], firstLine[1]

	sellers := mustReadIntArray(reader, n)
	buyers := mustReadIntArray(reader, m)

	sort.Ints(sellers)
	sort.Slice(buyers, func(i, j int) bool {
		return buyers[i] > buyers[j]
	})

	k := min(n, m)
	profit := 0
	for i := range k {
		if buyers[i] > sellers[i] {
			profit += buyers[i] - sellers[i]
		} else {
			break
		}
	}

	writer.WriteString(strconv.Itoa(profit))
	writer.WriteByte('\n')
}
