package backend

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	MOD = 1_000_000_007
)

func power(base, exp int64) int64 {
	res := int64(1)
	base %= MOD
	for exp > 0 {
		if exp%2 == 1 {
			res = (res * base) % MOD
		}
		base = (base * base) % MOD
		exp /= 2
	}
	return res
}

func modInverse(n int64) int64 {
	return power(n, MOD-2)
}

// https://coderun.yandex.ru/selections/backend/problems/card-game
// CardGame - problem 44
func CardGame() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line1, _ := reader.ReadString('\n')
	N, _ := strconv.Atoi(strings.TrimSpace(line1))

	// A input
	line2, _ := reader.ReadString('\n')
	parts := strings.Fields(strings.TrimSpace(line2))

	a := make([]int, N)
	S := 0
	for i := 0; i < N; i++ {
		a[i], _ = strconv.Atoi(parts[i])
		S += a[i]
	}

	halfS := S / 2

	C := make([][]int64, 11)
	for n := 0; n <= 10; n++ {
		C[n] = make([]int64, n+1)
		C[n][0] = 1
		for k := 1; k <= n; k++ {
			if k == n {
				C[n][k] = 1
			} else {
				C[n][k] = C[n-1][k-1] + C[n-1][k]
			}
		}
	}

	dp := make([][][]int64, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([][]int64, halfS+1)
		for j := 0; j <= halfS; j++ {
			dp[i][j] = make([]int64, 2*N+1)
		}
	}

	dp[0][0][N] = 1

	for i := range N {
		A := a[i]
		for sum_x := 0; sum_x <= halfS; sum_x++ {
			for sum_delta_idx := 0; sum_delta_idx <= 2*N; sum_delta_idx++ {
				val := dp[i][sum_x][sum_delta_idx]
				if val == 0 {
					continue
				}
				for x := 0; x <= A; x++ {
					if sum_x+x > halfS {
						break
					}
					var delta int
					switch x {
					case 0:
						delta = -1
					case A:
						delta = 1
					default:
						delta = 0
					}

					new_delta_idx := sum_delta_idx + delta
					if new_delta_idx >= 0 && new_delta_idx <= 2*N {
						ways := (val * C[A][x]) % MOD
						dp[i+1][sum_x+x][new_delta_idx] = (dp[i+1][sum_x+x][new_delta_idx] + ways) % MOD
					}
				}
			}
		}
	}

	P := dp[N][halfS][N]

	fact := make([]int64, S+1)
	fact[0] = 1
	for i := 1; i <= S; i++ {
		fact[i] = (fact[i-1] * int64(i)) % MOD
	}

	num := fact[S]
	den := (fact[halfS] * fact[S-halfS]) % MOD
	Q := (num * modInverse(den)) % MOD

	ans := (P * modInverse(Q)) % MOD

	writer.WriteString(strconv.Itoa(int(ans)))
	writer.WriteByte('\n')
}
