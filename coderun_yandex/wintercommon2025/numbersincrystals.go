package wintercommon2025

import (
	"bufio"
	"os"
	"strconv"
)

const (
	maxN = 700000
	maxK = 9

// 2*3*4*5*6*7*8*9 = 362880 <= 700000
// 2*3*4*5*6*7*8*9*10 = 3628800 > 700000
// maxK = 9
)

func solveNumbersInCrystals(prefixSums [][]int, k, l, r int) int {
	if k > maxK || l > maxN {
		return 0
	}

	if r > maxN {
		r = maxN
	}

	if l < 1 {
		l = 1
	}

	return prefixSums[k][r] - prefixSums[k][l-1]
}

func precompute() [][]int {
	prefixSums := make([][]int, maxK+1)
	for k := 1; k <= maxK; k++ {
		prefixSums[k] = make([]int, maxN+1)
	}

	generate(1, 2, 0, maxN, prefixSums)

	for n := 2; n <= maxN; n++ {
		prefixSums[1][n] = 1
	}

	for k := 1; k <= maxK; k++ {
		for n := 1; n <= maxN; n++ {
			prefixSums[k][n] += prefixSums[k][n-1]
		}
	}

	return prefixSums
}

func generate(product, minFactor, depth, limit int, prefixSums [][]int) {
	for factor := minFactor; product*factor <= limit; factor++ {
		newProduct := product * factor
		newDepth := depth + 1

		if newDepth >= 2 && newDepth <= maxK {
			prefixSums[newDepth][newProduct] = 1
		}

		generate(newProduct, factor+1, newDepth, limit, prefixSums)
	}
}

func validateNumbersInCrystalsQInput(n int) bool {
	return n >= 1 && n <= 50_000
}

// https://coderun.yandex.ru/selections/2025-winter-common/problems/numbers-in-crystals
// NumbersInCrystals - problem 8
func NumbersInCrystals() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	prefixSums := precompute()

	// Q input
	firstLine := mustReadIntArray(reader, 1)
	if !validateNumbersInCrystalsQInput(firstLine[0]) {
		panic("number Q out of range")
	}
	q := firstLine[0]

	for i := 0; i < q; i++ {
		// k,l and r input
		parts := mustReadIntArray(reader, 3)
		k, l, r := parts[0], parts[1], parts[2]

		result := solveNumbersInCrystals(prefixSums, k, l, r)
		writer.WriteString(strconv.Itoa(result))
		writer.WriteByte('\n')
	}
}
