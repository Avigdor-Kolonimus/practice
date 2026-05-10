package algorithmtrainingseptember2025

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/shelf-of-discord
// ShelfOfDiscord - assignment 21
func ShelfOfDiscord() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// a,b and S input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	tokens := strings.Fields(line)
	if len(tokens) != 3 {
		panic("invalid input")
	}
	a, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}
	S, err := strconv.Atoi(tokens[2])
	if err != nil {
		panic(err)
	}

	// D = (a+b)^2 - 4(ab-S)
	D := (a-b)*(a-b) + 4*S
	sqrtD := int(math.Sqrt(float64(D)))

	result := -1
	if sqrtD*sqrtD != D {
		writer.WriteString(strconv.Itoa(result))
		writer.WriteByte('\n')

		return
	}

	// L = ((a+b) + sqrt(D)) / 2
	num := a + b + sqrtD

	if num%2 != 0 {
		writer.WriteString(strconv.Itoa(result))
		writer.WriteByte('\n')

		return
	}

	L := num / 2

	// validation
	if (L-a)*(L-b) != S || L < a || L < b {
		writer.WriteString(strconv.Itoa(result))
		writer.WriteByte('\n')

		return
	}

	writer.WriteString(strconv.Itoa(L))
	writer.WriteByte('\n')
}
