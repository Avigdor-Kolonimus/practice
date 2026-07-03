package backend

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/backend/problems/random-swaps
// RandomSwaps - problem 9
func RandomSwaps() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	n := strings.TrimRight(line, "\r\n")

	// K input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	k, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	sumDigits := 0
	lastDigit := n[len(n)-1]
	L := len(n)
	C := L * (L - 1) / 2

	cnt5 := 0
	cntEven := 0
	for _, ch := range n {
		d := int(ch - '0')
		sumDigits += d
		if d == 5 {
			cnt5++
		}
		if d%2 == 0 {
			cntEven++
		}
	}

	probA := 0.0
	if cnt5 > 0 {
		s := 0.0
		if lastDigit == '5' {
			s = 1.0
		}

		if k == 0 {
			probA = s
		} else {
			p := float64(cnt5) / float64(C)
			q := 1.0 - float64(L-cnt5)/float64(C)

			for i := 0; i < k; i++ {
				s = s*q + (1.0-s)*p
			}
			probA = s
		}
	}

	probB := 0.0
	if sumDigits%3 == 0 && cntEven > 0 {
		e := 0.0
		if (lastDigit-'0')%2 == 0 {
			e = 1.0
		}

		if k == 0 {
			probB = e
		} else {
			p := float64(cntEven) / float64(C)
			q := 1.0 - float64(L-cntEven)/float64(C)

			for i := 0; i < k; i++ {
				e = e*q + (1.0-e)*p
			}
			probB = e
		}
	}

	result := probA + probB

	writer.WriteString(strconv.FormatFloat(result, 'f', 15, 64))
	writer.WriteByte('\n')
}
