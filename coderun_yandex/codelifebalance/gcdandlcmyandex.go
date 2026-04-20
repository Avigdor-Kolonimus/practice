package codelifebalance

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func validateGcdAndLcmYandexInput(p float64) bool {
	return p >= 1 && p <= 1_000_000_000_000_000_000_000_000_000_000_000_000
}

// https://coderun.yandex.ru/selections/code-life-balance/problems/gcd-and-lcm-yandex
// GcdAndLcmYandex - assignment 4
func GcdAndLcmYandex() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	parameters := strings.Fields(line)
	if len(parameters) != 2 {
		panic("input does not match 2")
	}

	x, err := strconv.ParseFloat(parameters[0], 64)
	if err != nil {
		panic(err)
	}
	if !validateGcdAndLcmYandexInput(x) {
		panic("number X out of range")
	}

	y, err := strconv.ParseFloat(parameters[1], 64)
	if err != nil {
		panic(err)
	}
	if !validateGcdAndLcmYandexInput(y) {
		panic("number Y out of range")
	}

	if x > y {
		panic("X > Y")
	}

	eps := 1e-9
	if math.Abs(math.Mod(y, x)) > eps {
		writer.WriteByte('0')
		writer.WriteByte('\n')
		return
	}

	n := int64(y / x)

	// считаем количество различных простых делителей
	cnt := 0
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			cnt++
			for n%i == 0 {
				n /= i
			}
		}
	}

	if n > 1 {
		cnt++
	}

	// result = 2^cnt
	result := int(1)
	for i := 0; i < cnt; i++ {
		result *= 2
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
