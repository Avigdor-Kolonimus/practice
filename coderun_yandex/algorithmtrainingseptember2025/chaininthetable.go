package algorithmtrainingseptember2025

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	n, m int
	a    [][]int
	dp   [][]int
	dir  = [][2]int{
		{1, 0}, {-1, 0},
		{0, 1}, {0, -1},
	}
)

func validateChainInTheTableNMInput(n int) bool {
	return n >= 1 && n <= 1_000
}

func dfs(i, j int) int {
	if dp[i][j] != 0 {
		return dp[i][j]
	}

	best := 1

	for _, d := range dir {
		ni, nj := i+d[0], j+d[1]

		if ni >= 0 && ni < n && nj >= 0 && nj < m {
			if a[ni][nj] == a[i][j]+1 {
				length := 1 + dfs(ni, nj)
				if length > best {
					best = length
				}
			}
		}
	}

	dp[i][j] = best
	return best
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/chain-in-the-table
// ChainInTheTable - assignment 19
func ChainInTheTable() {
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
	n, err = strconv.Atoi(parameters[0])
	if err != nil {
		panic(err)
	}

	if !validateChainInTheTableNMInput(n) {
		panic("number N out of range")
	}

	// M
	m, err = strconv.Atoi(parameters[1])
	if err != nil {
		panic(err)
	}

	if !validateChainInTheTableNMInput(m) {
		panic("number M out of range")
	}

	a = make([][]int, n)
	dp = make([][]int, n)

	for i := 0; i < n; i++ {
		// row input
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		strNum := strings.Fields(line)
		if len(strNum) != m {
			panic("numbers count does not match M")
		}

		a[i] = make([]int, m)
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			x, err := strconv.Atoi(strNum[j])
			if err != nil {
				panic(err)
			}
			a[i][j] = x
		}
	}

	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			val := dfs(i, j)
			if val > result {
				result = val
			}
		}
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
