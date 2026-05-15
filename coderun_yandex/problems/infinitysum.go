package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func coefficients(rows int) []int64 {
	row := []int64{1}

	for n := 1; n < rows; n++ {
		prev := row
		cur := []int64{1}

		for k := 1; k < n; k++ {
			val := int64(n+1-k)*prev[k-1] + int64(k+1)*prev[k]
			cur = append(cur, val)
		}

		cur = append(cur, 1)
		row = cur
	}

	return row
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func pow(base, exp int64) int64 {
	res := int64(1)

	for exp > 0 {
		if exp&1 == 1 {
			res *= base
		}
		base *= base
		exp >>= 1
	}

	return res
}

// https://coderun.yandex.ru/problem/infinity-sum
// InfinitySum - problem 583
func InfinitySum() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// a and b line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	a, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	if b == 1 {
		writer.WriteString("infinity")
		writer.WriteByte('\n')

		return
	}

	coefs := coefficients(int(a))

	var numerator int64
	for k := int64(0); k < int64(a); k++ {
		numerator += coefs[k] * pow(int64(b), k)
	}

	numerator *= int64(b)

	denominator := pow(int64(b)-1, int64(a)+1)

	g := gcd(numerator, denominator)

	writer.WriteString(strconv.FormatInt(numerator/g, 10) + "/" + strconv.FormatInt(denominator/g, 10))
	writer.WriteByte('\n')
}
