package summercommon2025

import (
	"bufio"
	"os"
	"slices"
	"strconv"
)

func validateFavoriteMelodiesInput(n int) bool {
	return n >= 1 && n <= 100_000
}

func solveFavoriteMelodies(n int, _ int, a []int64, queries [][2]int) int64 {
	events := make([]int, n+1)

	for _, qr := range queries {
		l, r := qr[0], qr[1]
		events[l-1] += 1
		events[r] -= 1
	}

	nowCovered := 0
	covered := make([]int, n)

	for i := 0; i < n; i++ {
		nowCovered += events[i]
		covered[i] = nowCovered
	}

	slices.Sort(covered)
	slices.Sort(a)

	var res int64 = 0
	for i := 0; i < n; i++ {
		res += a[i] * int64(covered[i])
	}

	return res
}

// https://coderun.yandex.ru/selections/2025-summer-common/problems/favorite-melodies
// FavoriteMelodies - problem 8
func FavoriteMelodies() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and Q input
	line := mustReadIntArray(reader, 2)
	n, q := line[0], line[1]
	if !validateFavoriteMelodiesInput(n) {
		panic("number N out of range")
	}
	if !validateFavoriteMelodiesInput(q) {
		panic("number Q out of range")
	}

	// friendliness input
	friendliness := make([]int64, n)
	line = mustReadIntArray(reader, n)
	for i := range friendliness {
		friendliness[i] = int64(line[i])
	}

	// pairs input
	pairs := make([][2]int, q)
	for i := range pairs {
		line = mustReadIntArray(reader, 2)
		pairs[i][0] = line[0]
		pairs[i][1] = line[1]
	}

	answer := solveFavoriteMelodies(n, q, friendliness, pairs)

	writer.WriteString(strconv.FormatInt(answer, 10))
	writer.WriteByte('\n')
}
