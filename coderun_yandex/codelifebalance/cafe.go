package codelifebalance

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	inf = 1 << 30
)

type State struct {
	prevJ int
	used  bool
}

func validateCafeInput(p int) bool {
	return p >= 0 && p <= 100
}

func validateCafeCostInput(p int) bool {
	return p >= 0 && p <= 300
}

// https://coderun.yandex.ru/selections/code-life-balance/problems/cafe
// Cafe - assignment 1
func Cafe() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateCafeInput(n) {
		panic("number N out of range")
	}

	// costs input
	costs := make([]int, n+1)
	for i := 1; i <= n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		cost, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		if !validateCafeCostInput(cost) {
			panic("number cost out of range")
		}

		costs[i] = cost

	}

	// dp[i][j] — cost minimal
	dp := make([][]int, n+1)
	parent := make([][]State, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		parent[i] = make([]State, n+1)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}

	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 0; j <= n; j++ {
			// 1. use coupon
			if j < n {
				if dp[i-1][j+1] < dp[i][j] {
					dp[i][j] = dp[i-1][j+1]
					parent[i][j] = State{j + 1, true}
				}
			}

			// 2. pay, no coupon
			if costs[i] <= 100 && j <= n {
				candidate := dp[i-1][j] + costs[i]
				if candidate < dp[i][j] {
					dp[i][j] = candidate
					parent[i][j] = State{j, false}
				}
			}

			// 3. pay and  coupon
			if costs[i] > 100 && j > 0 {
				candidate := dp[i-1][j-1] + costs[i]
				if candidate < dp[i][j] {
					dp[i][j] = candidate
					parent[i][j] = State{j - 1, false}
				}
			}
		}
	}

	// search best way
	minCost := inf
	k1 := 0
	for j := 0; j <= n; j++ {
		if dp[n][j] <= minCost {
			minCost = dp[n][j]
			k1 = j
		}
	}

	// repair
	usedDays := make([]int, 0)
	currentJ := k1
	for i := n; i >= 1; i-- {
		if parent[i][currentJ].used {
			usedDays = append(usedDays, i)
		}
		currentJ = parent[i][currentJ].prevJ
	}

	// Output
	writer.WriteString(strconv.Itoa(minCost))
	writer.WriteByte('\n')
	writer.WriteString(strconv.Itoa(k1) + " " + strconv.Itoa(len(usedDays)))
	writer.WriteByte('\n')
	for i := len(usedDays) - 1; i >= 0; i-- {
		writer.WriteString(strconv.Itoa(usedDays[i]))
		writer.WriteByte('\n')
	}
}
