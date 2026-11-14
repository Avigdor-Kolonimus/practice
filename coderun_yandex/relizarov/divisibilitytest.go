package relizarov

import (
	"bufio"
	"fmt"
	"os"
)

func solveDivisibilityTest(b, n int64) (int, int64) {
	pow := int64(1)
	for k := int64(1); k <= n; k++ {
		pow = (pow * b) % n

		if pow == 0 {
			return 1, k
		}

		if pow == 1 {
			return 2, k
		}

		if pow == n-1 {
			return 3, k
		}
	}

	return 0, 0
}

// https://coderun.yandex.ru/selections/relizarov/problems/divisibility-test
// DivisibilityTest - problem 3
func DivisibilityTest() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var b, n int64
		fmt.Fscan(reader, &b, &n)

		a, k := solveDivisibilityTest(b, n)

		if a == 0 {
			fmt.Fprintln(writer, 0)
		} else {
			fmt.Fprintln(writer, a, k)
		}
	}
}
