package summercommon2025

import (
	"bufio"
	"os"
	"strconv"
)

const (
	mod int64 = 1e9 - 7538
)

var res = map[int64]int64{0: 1}

func solveFirstSeedRandomGarden(n int64) int64 {
	if n == 0 {
		return 1
	}

	r, ok := res[n]
	if ok {
		return r
	}

	a := solveFirstSeedRandomGarden(n / 2)
	b := solveFirstSeedRandomGarden(n / 3)
	result := powMod(a, b)

	c := solveFirstSeedRandomGarden(n / 4)
	result = (result + (5*c)%mod + n) % mod

	res[n] = result

	return result
}

func powMod(a, b int64) int64 {
	base := a % mod
	if base < 0 {
		base += mod
	}

	result := int64(1)
	exp := b
	for exp > 0 {
		if exp&1 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}

	return result
}

// https://coderun.yandex.ru/selections/2025-summer-common/problems/first-seed-random-garden
// FirstSeedRandomGarden - problem 7
func FirstSeedRandomGarden() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	line := mustReadIntArray(reader, 1)
	n := line[0]

	answer := solveFirstSeedRandomGarden(int64(n))
	writer.WriteString(strconv.FormatInt(answer, 10))
	writer.WriteByte('\n')
}
