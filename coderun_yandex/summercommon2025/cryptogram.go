package summercommon2025

import (
	"bufio"
	"os"
	"strconv"
)

func calculateCryptogramAnswer(n int, k int, a []int) int {
	freq := make([]int, n+1)

	mx := 0
	for _, v := range a {
		if v%k == 0 {
			x := v / k
			freq[x]++
			if x > mx {
				mx = x
			}
		}
	}

	if mx == 0 {
		return 0
	}

	cnt := make([]int, mx+1)
	for d := 1; d <= mx; d++ {
		for multiple := d; multiple <= mx; multiple += d {
			cnt[d] += freq[multiple]
		}
	}

	g := make([]int, mx+1)

	for d := mx; d >= 1; d-- {
		c := cnt[d]
		g[d] = c * (c - 1) / 2

		for multiple := 2 * d; multiple <= mx; multiple += d {
			g[d] -= g[multiple]
		}
	}

	return g[1]
}

// ввод/вывод
// не изменяйте сигнатуру метода
// https://coderun.yandex.ru/selections/2025-summer-common/problems/cryptogram
// Cryptogram - problem 20
func Cryptogram() {
	input := NewFastScanner(os.Stdin)

	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	t := input.readInt()
	for test := 0; test < t; test++ {
		n := input.readInt()
		k := input.readInt()
		a := input.readIntArray(n)

		answer := calculateCryptogramAnswer(n, k, a)
		output.WriteString(strconv.Itoa(answer))
		output.WriteByte('\n')
	}
}
