package atolstikov

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strings"
)

const (
	alphabetSize = 26
	textLength   = (1 << 26) - 1
)

func gcdRestoreAlphabet(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// https://coderun.yandex.ru/selections/atolstikov/problems/restore-alphabet
// RestoreAlphabet - problem 6
func RestoreAlphabet() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	n := len(s)

	var first [alphabetSize]int
	var last [alphabetSize]int
	var gcdDiff [alphabetSize]int
	var count [alphabetSize]int
	var feasible [alphabetSize]uint32

	for i := range alphabetSize {
		first[i] = -1
		last[i] = -1
	}

	used := make([]int, 0, alphabetSize)
	for i := range n {
		u := int(s[i] - 'A')

		if first[u] == -1 {
			first[u] = i
			used = append(used, u)
		} else {
			gcdDiff[u] = gcdRestoreAlphabet(gcdDiff[u], i-last[u])
		}

		last[u] = i
		count[u]++
	}

	for _, u := range used {
		if count[u] == 1 {
			feasible[u] = (1 << alphabetSize) - 1
			continue
		}

		maxK := bits.TrailingZeros(uint(gcdDiff[u])) - 1

		if maxK > 25 {
			maxK = 25
		}

		if maxK >= 0 {
			feasible[u] = (uint32(1) << (maxK + 1)) - 1
		}
	}

	m := len(used)
	maxStart := textLength - n + 1

	getResidue := func(u, k int) int {
		mod := 1 << (k + 1)
		return ((1 << k) - first[u]) & (mod - 1)
	}

	canMatchOthers := func(fixedU, kMax, p int) bool {
		rightN := kMax

		if rightN == 0 {
			return m == 1
		}

		var masks [alphabetSize]uint32
		for _, u := range used {
			if u == fixedU {
				continue
			}

			for k := 0; k < kMax; k++ {
				if feasible[u]&(uint32(1)<<k) == 0 {
					continue
				}

				modMask := (1 << (k + 1)) - 1

				if getResidue(u, k) == p&modMask {
					masks[u] |= uint32(1) << k
				}
			}

			if masks[u] == 0 {
				return false
			}
		}

		matchR := make([]int, rightN)
		for i := range matchR {
			matchR[i] = -1
		}

		var dfs func(u int, seen []bool) bool

		dfs = func(u int, seen []bool) bool {
			for k := range rightN {
				if masks[u]&(uint32(1)<<k) == 0 || seen[k] {
					continue
				}

				seen[k] = true

				if matchR[k] == -1 || dfs(matchR[k], seen) {
					matchR[k] = u
					return true
				}
			}

			return false
		}

		for _, u := range used {
			if u == fixedU {
				continue
			}

			seen := make([]bool, rightN)

			if !dfs(u, seen) {
				return false
			}
		}

		return true
	}

	answer := -1
	for kMax := range alphabetSize {
		mod := 1 << (kMax + 1)

		for _, fixedU := range used {
			if feasible[fixedU]&(uint32(1)<<kMax) == 0 {
				continue
			}

			p := getResidue(fixedU, kMax)

			if p == 0 {
				p = mod
			}

			if p > maxStart {
				continue
			}

			if answer != -1 && p >= answer {
				continue
			}

			if canMatchOthers(fixedU, kMax, p) {
				answer = p
			}
		}
	}

	if answer == -1 {
		fmt.Fprintln(writer, "No solution")
		return
	}

	var lowerToUpper [alphabetSize]byte
	var upperUsed [alphabetSize]bool
	for i := range n {
		u := int(s[i] - 'A')
		k := bits.TrailingZeros(uint(answer + i))

		lowerToUpper[k] = s[i]
		upperUsed[u] = true
	}

	nextUpper := 0
	for k := range alphabetSize {
		if lowerToUpper[k] != 0 {
			continue
		}

		for nextUpper < alphabetSize && upperUsed[nextUpper] {
			nextUpper++
		}

		lowerToUpper[k] = byte('A' + nextUpper)
		upperUsed[nextUpper] = true
	}

	writer.Write(lowerToUpper[:])
	writer.WriteByte('\n')
	fmt.Fprintln(writer, answer)
}
