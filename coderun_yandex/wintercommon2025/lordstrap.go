package wintercommon2025

import (
	"bufio"
	"os"
	"strconv"
)

const (
	modLordsTrap int64 = 1000000007
)

func solveLordsTrap(n int) int64 {
	if n == 1 {
		return 1
	}

	nMod := int64(n) % modLordsTrap

	alpha := make([]int64, n+1)
	beta := make([]int64, n+1)

	alpha[0] = 0
	beta[0] = 0

	for k := 1; k < n; k++ {
		kMod := int64(k) % modLordsTrap
		nMinusK := int64(n-k) % modLordsTrap

		// denominator = n - k*alpha[k-1]
		denom := (nMod - kMod*alpha[k-1]%modLordsTrap + modLordsTrap) % modLordsTrap
		denomInv := modInverseLordsTrap(denom, modLordsTrap)

		// alpha[k] = (n-k) / denom
		alpha[k] = nMinusK * denomInv % modLordsTrap

		// beta[k] = (n + k*beta[k-1]) / denom
		numerator := (nMod + kMod*beta[k-1]%modLordsTrap) % modLordsTrap
		beta[k] = numerator * denomInv % modLordsTrap
	}

	oneMinusAlpha := (1 - alpha[n-1] + modLordsTrap) % modLordsTrap
	oneMinusAlphaInv := modInverseLordsTrap(oneMinusAlpha, modLordsTrap)
	En := (1 + beta[n-1]) % modLordsTrap * oneMinusAlphaInv % modLordsTrap

	return En
}

func modInverseLordsTrap(a, m int64) int64 {
	return modPowLordsTrap(a, m-2, m)
}

func modPowLordsTrap(base, exp, m int64) int64 {
	result := int64(1)
	base = base % m
	for exp > 0 {
		if exp%2 == 1 {
			result = result * base % m
		}
		exp = exp >> 1
		base = base * base % m
	}
	return result
}

func validateLordsTrapInput(n int) bool {
	return n >= 1 && n <= 2_000_000
}

// https://coderun.yandex.ru/selections/2025-winter-common/problems/lords-trap
// LordsTrap - problem 11
func LordsTrap() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	firstLine := mustReadIntArray(reader, 1)
	n := firstLine[0]
	if !validateLordsTrapInput(n) {
		panic("number N out of range")
	}

	result := solveLordsTrap(n)

	writer.WriteString(strconv.FormatInt(result, 10))
	writer.WriteByte('\n')
}
