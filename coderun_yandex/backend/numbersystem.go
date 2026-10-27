package backend

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func digitValue(c byte) int64 {
	if c >= '0' && c <= '9' {
		return int64(c - '0')
	}
	return int64(c-'A') + 10
}

// https://coderun.yandex.ru/selections/first-2023-backend/problems/number-system
// NumberSystem - problem 37
func NumberSystem() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	s := strings.TrimSpace(line)
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, "\r", "")

	coeff := make([]int64, 0, len(s))
	var maxDigit int64

	side := int64(1)
	sign := int64(1)
	for i := 0; i < len(s); {
		switch s[i] {
		case '+':
			sign = 1
			i++
			continue
		case '-':
			sign = -1
			i++
			continue
		case '=':
			side = -1
			sign = 1
			i++
			continue
		}

		j := i
		for j < len(s) &&
			s[j] != '+' &&
			s[j] != '-' &&
			s[j] != '=' {
			j++
		}

		termSign := side * sign

		pos := 0
		for k := j - 1; k >= i; k-- {
			d := digitValue(s[k])
			if d > maxDigit {
				maxDigit = d
			}

			for len(coeff) <= pos {
				coeff = append(coeff, 0)
			}
			coeff[pos] += termSign * d
			pos++
		}

		i = j
	}

	minBase := maxDigit + 1
	if minBase < 2 {
		minBase = 2
	}

	k := 0
	for k < len(coeff) && coeff[k] == 0 {
		k++
	}

	if k == len(coeff) {
		writer.WriteString(strconv.FormatInt(minBase, 10))
		writer.WriteByte('\n')
		return
	}

	target := coeff[k]
	if target < 0 {
		target = -target
	}

	divisors := make([]int64, 0)
	for d := int64(1); d <= target/d; d++ {
		if target%d != 0 {
			continue
		}

		divisors = append(divisors, d)
		if d != target/d {
			divisors = append(divisors, target/d)
		}
	}

	sort.Slice(divisors, func(i, j int) bool {
		return divisors[i] < divisors[j]
	})

	check := func(base int64) bool {
		var carry int64

		for _, c := range coeff {
			x := c + carry

			if x%base != 0 {
				return false
			}

			carry = x / base
		}

		return carry == 0
	}

	for _, base := range divisors {
		if base < minBase {
			continue
		}

		if check(base) {
			writer.WriteString(strconv.FormatInt(base, 10))
			writer.WriteByte('\n')
			return
		}
	}

	writer.WriteString("-1\n")
}
