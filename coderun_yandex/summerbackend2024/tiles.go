package summerbackend2024

import (
	"bufio"
	"os"
	"strconv"
)

func validateTilesBInput(n int) bool {
	return n >= 8 && n <= 5_000
}

func validateTilesWInput(n int) bool {
	return n >= 1 && n <= 2_000_000
}

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/tiles
// Tiles - problem 31
func Tiles() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// B and W input
	line := mustReadIntArray(reader, 2)
	if !validateTilesBInput(line[0]) {
		panic("number B out of range")
	}
	b := line[0]
	if !validateTilesWInput(line[1]) {
		panic("number W out of range")
	}
	w := line[1]

	/*
		W = (n-2) * (m-2)
		B = n*m - (n*m - 2n - 2m + 4) = 2n + 2m - 4
	*/
	n, m := 0, 0
	for a := 1; a*a <= w; a++ {
		if w%a != 0 {
			continue
		}

		n = a + 2
		m = w/a + 2

		if 2*n+2*m-4 == b {
			if n < m {
				n, m = m, n
			}
			break
		}
	}

	writer.WriteString(strconv.Itoa(n))
	writer.WriteByte(' ')
	writer.WriteString(strconv.Itoa(m))
	writer.WriteByte('\n')
}
