package wintercommon2025

import (
	"bufio"
	"os"
	"strconv"
)

func solveGlitchChallenge(n, _ int, a []int) int {
	threshold := min(n, 1000000)

	factorialPrimes := factorizeFactorialSmall(n, threshold)

	aPrimes := factorizeProduct(a)

	sPrimes := make(map[int]int)

	for prime, exp := range factorialPrimes {
		if aExp, ok := aPrimes[prime]; ok {
			exp -= aExp
		}

		if exp > 0 {
			sPrimes[prime] = exp
		}
	}

	result := 1
	for _, exp := range sPrimes {
		result = (result * (exp + 1)) % mod
	}

	return result
}

// https://coderun.yandex.ru/selections/2025-winter-common/problems/glitch-challenge
// GlitchChallenge - problem 9
func GlitchChallenge() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and K input
	firstLine := mustReadIntArray(reader, 2)
	n, k := firstLine[0], firstLine[1]

	parts := mustReadIntArray(reader, k)
	a := make([]int, k)
	for i := 0; i < k; i++ {
		a[i] = parts[i]
	}

	result := solveGlitchChallenge(n, k, a)

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}

func factorizeFactorialSmall(n, threshold int) map[int]int {
	primes := sieve(threshold)
	result := make(map[int]int)

	for _, p := range primes {
		result[p] = legendre(n, p)
	}

	return result
}

func processLargePrimes(low, high int, callback func(int)) {
	if low > high {
		return
	}

	sqrtHigh := intSqrt(high)
	basePrimes := sieve(sqrtHigh)

	segmentSize := 100000
	for segmentLow := low; segmentLow <= high; segmentLow += segmentSize {
		segmentHigh := min(segmentLow+segmentSize-1, high)

		segment := make([]bool, segmentHigh-segmentLow+1)
		for i := range segment {
			segment[i] = true
		}

		for _, p := range basePrimes {
			start := max(((segmentLow+p-1)/p)*p, p*p)
			for j := start; j <= segmentHigh; j += p {
				segment[j-segmentLow] = false
			}
		}

		for i, isPrime := range segment {
			if isPrime {
				num := segmentLow + i
				if num >= 2 {
					callback(num)
				}
			}
		}
	}
}

func intSqrt(n int) int {
	if n < 2 {
		return n
	}
	left, right := 1, n
	for left < right {
		mid := (left + right + 1) / 2
		if mid*mid <= n {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func legendre(n, p int) int {
	result := 0
	power := p
	for power <= n {
		result += n / power
		power *= p
	}
	return result
}

func factorizeProduct(arr []int) map[int]int {
	result := make(map[int]int)
	for _, num := range arr {
		factors := factorize(num)
		for prime, exp := range factors {
			result[prime] += exp
		}
	}
	return result
}

func factorize(n int) map[int]int {
	result := make(map[int]int)
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			result[i]++
			n /= i
		}
	}
	if n > 1 {
		result[n]++
	}
	return result
}

func sieve(n int) []int {
	if n < 2 {
		return []int{}
	}

	isPrime := make([]bool, n+1)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0], isPrime[1] = false, false

	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	primes := []int{}
	for i := range isPrime {
		if i >= 2 && isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
