package algorithmtrainingseptember2025

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateTowerInput(n int) bool {
	return n >= 1 && n <= 1_000
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/tower
// Tower - assignment 15
func Tower() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	// N
	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	if !validateTowerInput(n) {
		panic("number N out of range")
	}

	// K
	k, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	if !validateTowerInput(n) {
		panic("number K out of range")
	}

	if k > n {
		panic("K > N")
	}

	// segments input
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	segments := strings.Fields(line)
	if len(segments) != n {
		panic("segments count does not match")
	}
	a := make([]int, 0, n)
	for _, rawInt := range segments {
		ai, err := strconv.Atoi(rawInt)
		if err != nil {
			panic(err)
		}
		if !validateTowerInput(ai) {
			panic("number Ai out of range")
		}

		a = append(a, ai)
	}

	// prefix sums
	pref := make([]int, n+1)
	for i := 1; i <= n; i++ {
		pref[i] = pref[i-1] + a[i-1]
	}

	// value[start]
	value := make([]int64, n+2)

	// deque (stores indices)
	deque := make([]int, 0)

	for i := 0; i < n; i++ {
		// remove elements > current (maintain increasing)
		for len(deque) > 0 && a[deque[len(deque)-1]] >= a[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)

		// remove out-of-window
		if deque[0] <= i-k {
			deque = deque[1:]
		}

		// when window is ready
		if i >= k-1 {
			start := i - k + 1 + 1 // 1-based
			minVal := a[deque[0]]
			sum := pref[i+1] - pref[i+1-k]
			value[start] = int64(sum) * int64(minVal)
		}
	}

	// DP
	dp := make([]int64, n+1)
	prev := make([]int, n+1)
	chosenStart := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1]
		prev[i] = i - 1

		if i >= k {
			start := i - k + 1
			cand := dp[i-k] + value[start]
			if cand > dp[i] {
				dp[i] = cand
				prev[i] = i - k
				chosenStart[i] = start
			}
		}
	}

	// reconstruct
	var starts []int
	for i := n; i > 0; {
		if chosenStart[i] != 0 {
			starts = append(starts, chosenStart[i])
			i = prev[i]
		} else {
			i--
		}
	}

	// reverse
	for i, j := 0, len(starts)-1; i < j; i, j = i+1, j-1 {
		starts[i], starts[j] = starts[j], starts[i]
	}

	writer.WriteString(strconv.Itoa(len(starts)))
	writer.WriteByte('\n')
	for i, v := range starts {
		if i > 0 {
			writer.WriteString(" ")
		}
		writer.WriteString(strconv.Itoa(v))
	}
	writer.WriteByte('\n')
}
