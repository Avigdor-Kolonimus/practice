package eserajim

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func validateTriangleSimilarityInput(p int) bool {
	return p >= 1 && p <= 1_000_000
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func gcd3(a, b, c int) int {
	return gcd(gcd(a, b), c)
}

// https://coderun.yandex.ru/selections/eserajim/problems/triangle-similarity
// TriangleSimilarity - problem 2
func TriangleSimilarity() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateTriangleSimilarityInput(n) {
		panic("number N out of range")
	}

	// tasks
	set := make(map[[3]int]struct{})
	for range n {
		// triangle input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		strNum := strings.Fields(line)
		if len(strNum) != 3 {
			panic("numbers count does not match 3")
		}

		// A
		a, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}

		// B
		b, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		// C
		c, err := strconv.Atoi(strNum[2])
		if err != nil {
			panic(err)
		}

		// sorting
		if a > b {
			a, b = b, a
		}
		if b > c {
			b, c = c, b
		}
		if a > b {
			a, b = b, a
		}

		g := gcd3(a, b, c)

		a /= g
		b /= g
		c /= g

		set[[3]int{a, b, c}] = struct{}{}
	}

	writer.WriteString(strconv.Itoa(len(set)))
	writer.WriteByte('\n')
}
