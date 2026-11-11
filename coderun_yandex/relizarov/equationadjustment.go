package relizarov

import (
	"bufio"
	"os"
	"strings"
)

func evaluate(expr string) int64 {
	var result, current int64
	sign := int64(1)

	for i := 0; i < len(expr); i++ {
		c := expr[i]

		if c == '+' || c == '-' {
			result += sign * current
			current = 0

			if c == '+' {
				sign = 1
			} else {
				sign = -1
			}
		} else {
			current = current*10 + int64(c-'0')
		}
	}

	return result + sign*current
}

func checkPart(part string) bool {
	if len(part) == 0 {
		return false
	}

	if part[0] == '+' || part[0] == '-' {
		return false
	}

	last := part[len(part)-1]
	if last == '+' || last == '-' {
		return false
	}

	for i := 0; i < len(part); {
		if part[i] == '+' || part[i] == '-' {
			if i+1 < len(part) &&
				(part[i+1] == '+' || part[i+1] == '-') {
				return false
			}
			i++
			continue
		}

		if part[i] < '0' || part[i] > '9' {
			return false
		}

		start := i

		for i < len(part) &&
			part[i] >= '0' && part[i] <= '9' {
			i++
		}

		numLen := i - start

		if numLen > 10 {
			return false
		}

		if numLen > 1 && part[start] == '0' {
			return false
		}
	}

	return true
}

func check(s string) bool {
	eq := -1

	for i := 0; i < len(s); i++ {
		if s[i] == '=' {
			if eq != -1 {
				return false
			}
			eq = i
		}
	}

	if eq <= 0 || eq == len(s)-1 {
		return false
	}

	left := s[:eq]
	right := s[eq+1:]

	if !checkPart(left) || !checkPart(right) {
		return false
	}

	return evaluate(left) == evaluate(right)
}

func solveEquationAdjustment(s string) string {
	if check(s) {
		return "Correct"
	}

	n := len(s)

	for i := 0; i < n; i++ {
		if s[i] < '0' || s[i] > '9' {
			continue
		}

		digit := s[i]
		temp := s[:i] + s[i+1:]

		for j := 0; j <= len(temp); j++ {
			candidate := temp[:j] + string(digit) + temp[j:]

			if candidate == s {
				continue
			}

			if check(candidate) {
				return candidate
			}
		}
	}

	return "Impossible"
}

// https://coderun.yandex.ru/selections/relizarov/problems/equation-adjustment
// EquationAdjustment - problem 2
func EquationAdjustment() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	s := strings.TrimSpace(line)

	writer.WriteString(solveEquationAdjustment(s))
	writer.WriteByte('\n')
}
