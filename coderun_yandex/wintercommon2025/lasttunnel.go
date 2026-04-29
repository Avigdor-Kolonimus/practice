package wintercommon2025

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	modLastTunnel = 998244353
)

func matrixPower(M [][]int, power int64) [][]int {
	n := len(M)

	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		result[i][i] = 1
	}

	base := make([][]int, n)
	for i := 0; i < n; i++ {
		base[i] = make([]int, n)
		copy(base[i], M[i])
	}

	for power > 0 {
		if power&1 == 1 {
			result = matrixMultiply(result, base)
		}
		base = matrixMultiply(base, base)
		power >>= 1
	}

	return result
}

func matrixMultiply(A, B [][]int) [][]int {
	n := len(A)
	m := len(B[0])
	k := len(B)

	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
		for j := 0; j < m; j++ {
			sum := 0
			for t := 0; t < k; t++ {
				sum = (sum + A[i][t]*B[t][j]%modLastTunnel) % modLastTunnel
			}
			result[i][j] = sum
		}
	}

	return result
}

func solveLastTunnel(n int64, good map[int]bool) int {
	if n < 3 {
		return 0
	}

	M := make([][]int, 100)
	for i := 0; i < 100; i++ {
		M[i] = make([]int, 100)
	}

	for d1 := 0; d1 < 10; d1++ {
		for d2 := 0; d2 < 10; d2++ {
			from := d1*10 + d2
			for d3 := 0; d3 < 10; d3++ {
				sum := d1 + d2 + d3
				if good[sum] {
					to := d2*10 + d3
					M[from][to] = 1
				}
			}
		}
	}

	start := make([]int, 100)
	for d1 := 1; d1 < 10; d1++ {
		for d2 := 0; d2 < 10; d2++ {
			idx := d1*10 + d2
			start[idx] = 1
		}
	}

	if n == 2 {
		result := 0
		for i := 0; i < 100; i++ {
			result = (result + start[i]) % modLastTunnel
		}
		return result
	}

	power := n - 2
	M_power := matrixPower(M, power)

	result := 0
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			result = (result + start[i]*M_power[i][j]%modLastTunnel) % modLastTunnel
		}
	}

	return result
}

// https://coderun.yandex.ru/selections/2025-winter-common/problems/last-tunnel
// LastTunnel - problem 12
func LastTunnel() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	parts := strings.Fields(strings.TrimSpace(line))
	n, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	// Lucky numbers input
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	parts = strings.Fields(strings.TrimSpace(line))
	good := make(map[int]bool, m)
	for i := 0; i < m; i++ {
		val, err := strconv.Atoi(parts[i])
		if err != nil {
			panic(err)
		}

		good[val] = true
	}

	result := solveLastTunnel(int64(n), good)

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
