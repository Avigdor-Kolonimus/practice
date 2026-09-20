package wintercommon2025

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const MOD_SnowstormIntensifying int64 = 998244353

func solveSnowstormIntensifying(s, t string) int64 {
	n := len(s)
	m := len(t)

	// two[i] = 2^i mod MOD
	two := make([]int64, n+1)
	two[0] = 1

	for i := 1; i <= n; i++ {
		two[i] = two[i-1] * 2 % MOD_SnowstormIntensifying
	}

	// lcp[i][j] = LCP(s[i:], t[j:])
	lcp := make([][]int, n)
	for i := 0; i < n; i++ {
		lcp[i] = make([]int, m)
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if s[i] == t[j] {
				lcp[i][j] = 1

				if i+1 < n && j+1 < m {
					lcp[i][j] += lcp[i+1][j+1]
				}
			}
		}
	}

	dp := make([][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int64, m+1)
	}

	for i := 0; i < n; i++ {
		for j := m - 1; j >= 0; j-- {
			var cur int64
			for k := 0; k <= i; k++ {
				l := lcp[k][j]

				if k+l-1 >= i {
					length := i - k + 1

					if k > 0 && j+length < m {
						cur += dp[k-1][j+length]
						if cur >= MOD_SnowstormIntensifying {
							cur -= MOD_SnowstormIntensifying
						}
					} else if k == 0 {
						add := int64(m - (j + length))

						cur += add
						if cur >= MOD_SnowstormIntensifying {
							cur -= MOD_SnowstormIntensifying
						}
					}

					continue
				}

				posS := k + l
				posT := j + l
				if posT < m && s[posS] < t[posT] {
					mult := int64(m - posT)

					if k > 0 {
						mult = mult * two[k-1] % MOD_SnowstormIntensifying
					}

					cur += mult
					if cur >= MOD_SnowstormIntensifying {
						cur -= MOD_SnowstormIntensifying
					}
				}
			}

			dp[i][j] = cur
		}
	}

	var ans int64
	for j := 0; j < m; j++ {
		ans += dp[n-1][j]
		if ans >= MOD_SnowstormIntensifying {
			ans -= MOD_SnowstormIntensifying
		}
	}

	return ans
}

// https://coderun.yandex.ru/selections/2025-winter-common/problems/snowstorm-intensifying
// SnowstormIntensifying - problem 15
func SnowstormIntensifying() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	T, _ := strconv.Atoi(strings.TrimSpace(line))

	for ; T > 0; T-- {
		line, _ = reader.ReadString('\n')
		s := strings.TrimSpace(line)

		line, _ = reader.ReadString('\n')
		t := strings.TrimSpace(line)

		ans := solveSnowstormIntensifying(s, t)

		writer.WriteString(strconv.FormatInt(ans, 10))
		writer.WriteByte('\n')
	}
}
