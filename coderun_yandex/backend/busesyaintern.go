package backend

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/backend/problems/buses-ya-intern
// BusesYaIntern - problem 30
func BusesYaIntern() {
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

	// time input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	m := int64(1 << 60)
	D := big.NewInt(1)
	t := make([]int, n)
	for i := 0; i < n; i++ {
		ti, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		if int64(ti) < m {
			m = int64(ti)
		}
		D.Mul(D, big.NewInt(int64(ti)))

		t[i] = ti
	}

	// P(x) = product (t_i - x)
	coeff := []*big.Int{big.NewInt(1)}
	for _, tk := range t {
		next := make([]*big.Int, len(coeff)+1)
		for i := range next {
			next[i] = big.NewInt(0)
		}

		for i := 0; i < len(coeff); i++ {
			// * tk
			tmp := new(big.Int).Mul(coeff[i], big.NewInt(int64(tk)))
			next[i].Add(next[i], tmp)

			// * (-x)
			next[i+1].Sub(next[i+1], coeff[i])
		}

		coeff = next
	}

	// sum a_i * m^(i+1)/(i+1)
	res := new(big.Rat)
	mpow := big.NewInt(m)
	for i := 0; i < len(coeff); i++ {
		num := new(big.Int).Mul(coeff[i], mpow)

		term := new(big.Rat).SetFrac(num, big.NewInt(int64(i+1)))
		res.Add(res, term)

		mpow = new(big.Int).Mul(mpow, big.NewInt(m))
	}

	res.Quo(res, new(big.Rat).SetInt(D))

	ans := fmt.Sprintf("%s/%s\n",
		res.Num().String(),
		res.Denom().String(),
	)
	writer.WriteString(ans)
	writer.WriteByte('\n')
}
