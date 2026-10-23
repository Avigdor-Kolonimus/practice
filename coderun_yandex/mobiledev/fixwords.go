package mobiledev

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	negInf = -1e300
	eps    = 1e-12
)

var (
	logSub       = math.Log(0.5)
	logDel       = math.Log(0.1)
	logIns       = math.Log(0.4)
	logThreshold = math.Log(1e-10)
)

func bestLogProbability(word, scanned string) float64 {
	w := []byte(word)
	s := []byte(scanned)

	lw, ls := len(w), len(s)
	cols := ls + 1

	dp := make([]float64, (lw+1)*cols)
	for i := range dp {
		dp[i] = negInf
	}

	idx := func(i, j int) int {
		return i*cols + j
	}

	dp[idx(0, 0)] = 0

	for i := 0; i <= lw; i++ {
		for j := 0; j <= ls; j++ {
			cur := dp[idx(i, j)]
			if cur <= negInf/2 {
				continue
			}

			if i < lw && j < ls {
				cand := cur
				if w[i] != s[j] {
					cand += logSub
				}

				ni := idx(i+1, j+1)
				if cand > dp[ni] {
					dp[ni] = cand
				}
			}

			if i < lw {
				ni := idx(i+1, j)
				cand := cur + logDel
				if cand > dp[ni] {
					dp[ni] = cand
				}
			}

			if j < ls {
				ni := idx(i, j+1)
				cand := cur + logIns
				if cand > dp[ni] {
					dp[ni] = cand
				}
			}
		}
	}

	return dp[idx(lw, ls)]
}

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/fix-words
// FixWords - problem 13
func FixWords() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var scanned string
	fmt.Fscan(reader, &scanned)

	var m int
	fmt.Fscan(reader, &m)

	bestWord := ""
	bestLogP := negInf

	for i := 0; i < m; i++ {
		var length int
		var word string

		fmt.Fscan(reader, &length)
		fmt.Fscan(reader, &word)

		lp := bestLogProbability(word, scanned)

		if lp > bestLogP+eps {
			bestLogP = lp
			bestWord = word
		} else if math.Abs(lp-bestLogP) <= eps {
			if bestWord == "" || word < bestWord {
				bestWord = word
			}
		}
	}

	if bestWord == "" || bestLogP <= logThreshold+eps {
		fmt.Fprintln(writer, "NO MATCH")
	} else {
		fmt.Fprintln(writer, bestWord)
	}
}
