package first2023mobiledev

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/upgrade-your-hero
// UpgradeYourHero - problem 10
func UpgradeYourHero() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var input []string
	for {
		line, err := reader.ReadString('\n')
		if len(line) > 0 {
			input = append(input, strings.Fields(line)...)
		}
		if err != nil {
			break
		}
	}

	n, _ := strconv.Atoi(input[0])

	var zeroCount int
	var negativeCount int

	var minPositive int64
	var minNegativeAbs int64
	var maxNegativeAbs int64

	var anyElement int64

	for i := range n {
		x, _ := strconv.ParseInt(input[i+1], 10, 64)

		if i == 0 {
			anyElement = x
		}

		switch {
		case x == 0:
			zeroCount++

		case x < 0:
			negativeCount++
			abs := -x

			if minNegativeAbs == 0 || abs < minNegativeAbs {
				minNegativeAbs = abs
			}

			if abs > maxNegativeAbs {
				maxNegativeAbs = abs
			}

		default:
			if minPositive == 0 || x < minPositive {
				minPositive = x
			}
		}
	}

	var answer int64

	switch {
	case zeroCount >= 2:
		answer = anyElement

	case zeroCount == 1:
		if negativeCount%2 == 0 {
			answer = 0
		} else {
			for i := 1; i <= n; i++ {
				x, _ := strconv.ParseInt(input[i], 10, 64)
				if x != 0 {
					answer = x
					break
				}
			}
		}

	default:
		if negativeCount%2 == 1 {
			answer = -minNegativeAbs
		} else if minPositive != 0 {
			answer = minPositive
		} else {
			answer = -maxNegativeAbs
		}
	}

	writer.WriteString(strconv.FormatInt(answer, 10))
	writer.WriteByte('\n')
}
