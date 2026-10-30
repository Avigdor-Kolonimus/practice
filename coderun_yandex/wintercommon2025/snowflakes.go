package wintercommon2025

import (
	"bufio"
	"fmt"
	"os"
)

const (
	modSnowflakes int64 = 1_000_000_007
	inv2          int64 = 500_000_004
	inv3          int64 = 333_333_336
)

func sum1(x int64) int64 {
	x %= mod
	return x * (x + 1) % mod * inv2 % mod
}

func sum2(x int64) int64 {
	x %= modSnowflakes
	return x * (x + 1) % modSnowflakes * ((2*x + 1) % modSnowflakes) % modSnowflakes * inv2 % modSnowflakes * inv3 % modSnowflakes
}

func solveSnowflakes(a, q, L, R int64) int64 {
	if a == 0 || q == 1 {
		return 0
	}

	N := R - L + 1
	if N <= 0 {
		return 0
	}

	n := N % modSnowflakes

	if q == -1 {
		evens := N / 2
		if N%2 == 1 && L%2 == 0 {
			evens++
		}

		odds := N - evens

		return n * n % modSnowflakes *
			(2 * (evens % modSnowflakes) % modSnowflakes) % modSnowflakes *
			(odds % modSnowflakes) % modSnowflakes
	}

	ans := n * n % modSnowflakes * ((n - 1 + modSnowflakes) % modSnowflakes) % modSnowflakes

	if N <= 1 {
		return ans
	}

	limit := N - 1

	termN := n * ((n + 1) % modSnowflakes) % modSnowflakes
	term2N1 := (2*n + 1) % modSnowflakes

	for l := int64(1); l <= limit; {
		kMax := limit / l
		r := limit / kMax

		k := kMax % modSnowflakes

		s1 := sum1(kMax)
		s2 := sum2(kMax)

		ss1 := (sum1(r) - sum1(l-1) + modSnowflakes) % modSnowflakes
		ss2 := (sum2(r) - sum2(l-1) + modSnowflakes) % modSnowflakes
		ss0 := (r - l + 1) % modSnowflakes

		p1 := s2 * ss2 % modSnowflakes
		p2 := term2N1 * s1 % modSnowflakes * ss1 % modSnowflakes
		p3 := termN * k % modSnowflakes * ss0 % modSnowflakes

		block := (p1 - p2 + modSnowflakes) % modSnowflakes
		block = (block + p3) % modSnowflakes
		block = block * 2 % modSnowflakes

		ans = (ans + block) % modSnowflakes
		l = r + 1
	}

	if q == -2 && N >= 3 {
		maxOdd := limit
		if maxOdd%2 == 0 {
			maxOdd--
		}

		cnt := ((maxOdd + 1) / 2) % modSnowflakes

		sumOdd1 := cnt * cnt % modSnowflakes

		sq := 4 * sumOdd1 % modSnowflakes
		sq = (sq - 1 + modSnowflakes) % modSnowflakes

		sumOdd2 := cnt * sq % modSnowflakes * inv3 % modSnowflakes

		val := sumOdd2
		val = (val - term2N1*sumOdd1%modSnowflakes + modSnowflakes) % modSnowflakes
		val = (val + termN*cnt%modSnowflakes) % modSnowflakes

		correction := val * 2 % modSnowflakes
		correction = (correction - 4 + modSnowflakes) % modSnowflakes

		ans = (ans + correction) % modSnowflakes
	}

	return ans
}

// https://coderun.yandex.ru/selections/2025-winter-common/problems/snowflakes
// Snowflakes - problem 14
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var a, q, L, R int64
	fmt.Fscan(reader, &a, &q, &L, &R)

	fmt.Fprintln(writer, solveSnowflakes(a, q, L, R))
}
