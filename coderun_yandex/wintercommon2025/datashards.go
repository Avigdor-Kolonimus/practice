package wintercommon2025

import (
	"bufio"
	"os"
	"strconv"
)

const (
	mod = 998244353
)

func solveDataShards(n, s int, fact, invFact []int) int {
	if n > s {
		return 0
	}

	// C(s, n) = s! / (n! * (s-n)!)
	c := comb(s, n, fact, invFact)

	// (n+1)! × C(s, n)
	return fact[n+1] * c % mod
}

func validateDataShardsTInput(n int) bool {
	return n >= 1 && n <= 50_000
}

func validateDataShardsSNInput(n int) bool {
	return n >= 1 && n <= 200_000
}

// https://coderun.yandex.ru/selections/2025-winter-common/problems/data-shards
// DataShards - problem 7
func DataShards() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	maxN := 400001
	fact := precomputeFactorials(maxN)
	invFact := precomputeInvFactorials(fact, maxN)

	// T input
	firstLine := mustReadIntArray(reader, 1)
	if !validateDataShardsTInput(firstLine[0]) {
		panic("number T out of range")
	}
	t := firstLine[0]

	for i := 0; i < t; i++ {
		// n and s input
		parts := mustReadIntArray(reader, 2)
		n, s := parts[0], parts[1]
		if !validateDataShardsSNInput(s) {
			panic("number s out of range")
		}
		if !validateDataShardsSNInput(n) {
			panic("number N out of range")
		}

		result := solveDataShards(n, s, fact, invFact)
		writer.WriteString(strconv.Itoa(result))
		writer.WriteByte('\n')
	}
}

func comb(n, k int, fact, invFact []int) int {
	if k < 0 || k > n {
		return 0
	}

	return fact[n] * invFact[k] % mod * invFact[n-k] % mod
}

func precomputeFactorials(n int) []int {
	fact := make([]int, n+1)
	fact[0] = 1

	for i := 1; i <= n; i++ {
		fact[i] = fact[i-1] * i % mod
	}

	return fact
}

func precomputeInvFactorials(fact []int, n int) []int {
	invFact := make([]int, n+1)
	invFact[n] = modPow(fact[n], mod-2)

	for i := n; i > 0; i-- {
		invFact[i-1] = invFact[i] * i % mod
	}

	return invFact
}

// modPow вычисляет a^b mod mod
func modPow(a, b int) int {
	result := 1
	a %= mod

	for b > 0 {
		if b&1 == 1 {
			result = result * a % mod
		}
		a = a * a % mod
		b >>= 1
	}

	return result
}
