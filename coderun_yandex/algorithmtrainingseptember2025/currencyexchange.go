package algorithmtrainingseptember2025

import (
	"bufio"
	"io"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

type PairCurrencyExchange struct {
	val int64
	idx int
}

func absCurrencyExchange(x int64) int64 {
	if x < 0 {
		return -x
	}

	return x
}

func lowerBoundCurrencyExchange(a []PairCurrencyExchange, x int64) int {
	return sort.Search(len(a), func(i int) bool {
		return a[i].val >= x
	})
}

// true if num1/den1 < num2/den2
func lessFraction(num1, den1, num2, den2 int64) bool {
	left := new(big.Int).Mul(
		big.NewInt(num1),
		big.NewInt(den2),
	)

	right := new(big.Int).Mul(
		big.NewInt(num2),
		big.NewInt(den1),
	)

	return left.Cmp(right) < 0
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/currency-exchange
// CurrencyExchange - problem 24
func CurrencyExchange() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and P input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	c, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	p := int64(c)

	// Ci input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum = strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	cArray := make([]PairCurrencyExchange, n)
	for i := range n {
		c, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		cArray[i].val = int64(c)
		cArray[i].idx = i + 1
	}

	sort.Slice(cArray, func(i, j int) bool {
		return cArray[i].val < cArray[j].val
	})

	hasBest := false

	var bestNum int64
	var bestDen int64

	bestI, bestJ := 1, 2
	for _, den := range cArray {
		target := den.val * p

		pos := lowerBoundCurrencyExchange(cArray, target)

		for delta := -3; delta <= 3; delta++ {
			k := pos + delta

			if k < 0 || k >= n {
				continue
			}

			numEntry := cArray[k]

			if numEntry.idx == den.idx {
				continue
			}

			num := absCurrencyExchange(numEntry.val - target)
			denominator := den.val

			if !hasBest ||
				lessFraction(num, denominator, bestNum, bestDen) {

				hasBest = true

				bestNum = num
				bestDen = denominator

				bestI = numEntry.idx
				bestJ = den.idx
			}
		}
	}

	writer.WriteString(strconv.Itoa(bestI) + " " + strconv.Itoa(bestJ))
	writer.WriteByte('\n')
}
