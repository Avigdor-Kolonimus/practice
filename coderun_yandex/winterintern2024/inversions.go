package winterintern2024

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// https://coderun.yandex.ru/selections/winter-intern-2024/problems/inversions
// Inversions - problem 1
func Inversions() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	// P input
	line, _ = reader.ReadString('\n')
	fields := strings.Fields(line)
	p := make([]int, n)
	for i := range n {
		p[i], _ = strconv.Atoi(fields[i])
	}

	var initialInversions int64 = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if p[i] > p[j] {
				initialInversions++
			}
		}
	}

	var totalPairs int64 = int64(n) * int64(n-1) / 2
	var totalInversionsSum int64 = initialInversions * totalPairs
	for i := range n {
		for j := i + 1; j < n; j++ {
			var delta int64 = 0

			if p[i] < p[j] {
				delta += 1
			} else {
				delta -= 1
			}
		}
	}

	var totalDelta int64 = 0
	for i := range n {
		for j := i + 1; j < n; j++ {
			if p[i] < p[j] {
				totalDelta += 1
			} else {
				totalDelta -= 1
			}
		}
	}

	for k := range n {
		lessLeft := 0
		greaterLeft := 0
		for i := 0; i < k; i++ {
			if p[i] < p[k] {
				lessLeft++
			} else {
				greaterLeft++
			}
		}

		lessRight := 0
		greaterRight := 0
		for j := k + 1; j < n; j++ {
			if p[j] < p[k] {
				lessRight++
			} else {
				greaterRight++
			}
		}

		totalDelta += int64(lessLeft*greaterRight) * 2
		totalDelta -= int64(greaterLeft*lessRight) * 2
	}

	totalInversionsSum += totalDelta

	g := gcd(totalInversionsSum, totalPairs)
	numerator := totalInversionsSum / g
	denominator := totalPairs / g

	writer.WriteString(strconv.FormatInt(numerator, 10))
	writer.WriteByte('/')
	writer.WriteString(strconv.FormatInt(denominator, 10))
	writer.WriteByte('\n')
}
