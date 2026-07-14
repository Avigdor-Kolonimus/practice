package summerbackend2024

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INF int32 = 1_000_000_000
)

func computeD(S string) []int32 {
	n := len(S)
	D := make([]int32, n)
	var lastPos [256]int32
	for i := range lastPos {
		lastPos[i] = -1
	}

	for i := range n {
		c := S[i]
		if lastPos[c] == -1 {
			D[i] = INF
		} else {
			D[i] = int32(i) - lastPos[c]
		}
		lastPos[c] = int32(i)
	}

	return D
}

func match(Dt, Ds []int32, textIdx, patternIdx int32) bool {
	valT := Dt[textIdx]
	valS := Ds[patternIdx]
	if valS == INF {
		return valT > patternIdx
	}

	return valT == valS
}

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/substitution-code-v2
// SubstitutionCodeV2 - problem 8
func SubstitutionCodeV2() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t, s string
	if _, err := fmt.Fscan(reader, &t); err != nil {
		return
	}
	if _, err := fmt.Fscan(reader, &s); err != nil {
		return
	}

	n := int32(len(t))
	m := int32(len(s))

	if m > n {
		fmt.Fprintln(writer, 0)
		return
	}

	Dt := computeD(t)
	Ds := computeD(s)

	pi := make([]int32, m)
	for i := int32(1); i < m; i++ {
		j := pi[i-1]
		for j > 0 && !match(Ds, Ds, i, j) {
			j = pi[j-1]
		}
		if match(Ds, Ds, i, j) {
			j++
		}
		pi[i] = j
	}

	var ans []int32
	j := int32(0)
	for i := int32(0); i < n; i++ {
		for j > 0 && !match(Dt, Ds, i, j) {
			j = pi[j-1]
		}
		if match(Dt, Ds, i, j) {
			j++
		}
		if j == m {
			ans = append(ans, i-m+2)
			j = pi[j-1]
		}
	}

	fmt.Fprintln(writer, len(ans))
	for i, pos := range ans {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, pos)
	}
	fmt.Fprintln(writer)
}
