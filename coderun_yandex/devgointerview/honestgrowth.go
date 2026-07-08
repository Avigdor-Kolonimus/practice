package devgointerview

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/dev-go-interview/problems/honest-growth
// HonestGrowth - problem 13
func HonestGrowth() {
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

	// honest input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	price, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}

	bestL, bestR := 0, 0
	prevStart := 0
	gapStart := 0
	last := price
	preLast := price - 1
	for i := 1; i < n; i++ {
		preLast = last
		last = price
		price, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		if price <= last && price <= preLast {
			gapStart = i - 1
		} else if price > preLast && (price <= last || prevStart < gapStart) {
			gapStart = prevStart
		}

		if last <= preLast {
			prevStart = i - 1
		}

		bestStart := prevStart
		if gapStart < bestStart {
			bestStart = gapStart
		}

		if i-bestStart > bestR-bestL {
			bestL = bestStart
			bestR = i
		}
	}

	writer.WriteString(strconv.Itoa(bestL))
	writer.WriteByte(' ')
	writer.WriteString(strconv.Itoa(bestR))
	writer.WriteByte('\n')
}
