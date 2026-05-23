package summercommon2025

// https://coderun.yandex.ru/selections/2025-summer-common/problems/choose_primes
// ChoosePrimes - problem 6
func SolveChoosePrimes(n int) int {
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	var c1, c3 int
	for p := 3; p <= n; p++ {
		if isPrime[p] {
			if p%4 == 1 {
				c1++
			} else {
				c3++
			}
		}
	}

	return c1 * c3
}
