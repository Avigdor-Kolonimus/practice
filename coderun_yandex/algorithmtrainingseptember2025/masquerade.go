package algorithmtrainingseptember2025

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	INF = math.MaxInt
)

type Shop struct {
	P, R, Q, F int
}

func validateMasqueradeNLInput(n int) bool {
	return n >= 0 && n <= 100
}

func validateMasqueradeQPInput(n int) bool {
	return n >= 1 && n <= 1_000
}

func cost(s Shop, k int) int {
	if k == 0 {
		return 0
	}
	if k < s.R {
		return k * s.P
	}
	return k * s.Q
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/masquerade
// Masquerade - assignment 20
func Masquerade() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	parameters := strings.Fields(line)
	if len(parameters) != 2 {
		panic("input does not match 2")
	}

	// N
	n, err := strconv.Atoi(parameters[0])
	if err != nil {
		panic(err)
	}
	if n < 1 || !validateMasqueradeNLInput(n) {
		panic("number N out of range")
	}

	// L
	l, err := strconv.Atoi(parameters[1])
	if err != nil {
		panic(err)
	}
	if !validateMasqueradeNLInput(l) {
		panic("number L out of range")
	}
	maxLenght := l + 100

	shops := make([]Shop, n)
	for i := 0; i < n; i++ {
		// row input
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		strNum := strings.Fields(line)
		if len(strNum) != 4 {
			panic("numbers count does not match 4")
		}

		// Q
		q, err := strconv.Atoi(strNum[2])
		if err != nil {
			panic(err)
		}
		if !validateMasqueradeQPInput(q) {
			panic("number Q out of range")
		}

		// P
		p, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		if q > p || !validateMasqueradeQPInput(p) {
			panic("number P out of range")
		}

		// R
		r, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}
		if r < 1 || !validateMasqueradeNLInput(r) {
			panic("number R out of range")
		}

		// F
		f, err := strconv.Atoi(strNum[3])
		if err != nil {
			panic(err)
		}
		if !validateMasqueradeNLInput(f) {
			panic("number F out of range")
		}

		shops[i].P = p
		shops[i].R = r
		shops[i].Q = q
		shops[i].F = f
	}

	dp := make([]int, maxLenght+1)
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0

	prev := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		prev[i] = make([]int, maxLenght+1)
		for j := 0; j <= maxLenght; j++ {
			prev[i][j] = -1
		}
	}
	prev[0][0] = 0

	for i := 0; i < n; i++ {
		newdp := make([]int, maxLenght+1)
		for j := range newdp {
			newdp[j] = INF
		}

		for x := 0; x <= maxLenght; x++ {
			if dp[x] == INF {
				continue
			}
			for k := 0; k <= shops[i].F; k++ {
				nx := x + k
				if nx > maxLenght {
					break
				}
				val := dp[x] + cost(shops[i], k)
				if val < newdp[nx] {
					newdp[nx] = val
					prev[i+1][nx] = k
				}
			}
		}

		dp = newdp
	}

	// search result
	best := INF
	pos := -1
	for x := l; x <= maxLenght; x++ {
		if dp[x] < best {
			best = dp[x]
			pos = x
		}
	}

	if best != INF {
		writer.WriteString(strconv.Itoa(best))
		writer.WriteByte('\n')

		result := make([]int, n)
		for i := n; i > 0; i-- {
			k := prev[i][pos]
			result[i-1] = k
			pos -= k
		}

		results := ""
		for i := 0; i < n; i++ {
			results += strconv.Itoa(result[i]) + " "
		}
		writer.WriteString(results)
	} else {
		writer.WriteString("-1")
	}
	writer.WriteByte('\n')
}
