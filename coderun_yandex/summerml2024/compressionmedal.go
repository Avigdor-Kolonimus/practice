package summerml2024

import (
	"bufio"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/2024-summer-ml/problems/compression-medal
// CompressionMedal - problem 25
func CompressionMedal() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	fields := strings.Fields(line)

	n, _ := strconv.Atoi(fields[0])
	m, _ := strconv.Atoi(fields[1])
	_, _ = strconv.ParseFloat(fields[2], 64)

	x := make([][]float64, n)
	mins := make([]float64, m)
	maxs := make([]float64, m)

	for j := range m {
		mins[j] = math.Inf(1)
		maxs[j] = math.Inf(-1)
	}

	for i := range n {
		line, _ = reader.ReadString('\n')
		fields = strings.Fields(line)

		x[i] = make([]float64, m)

		for j := range m {
			value, _ := strconv.ParseFloat(fields[j], 64)
			x[i][j] = value

			if value < mins[j] {
				mins[j] = value
			}
			if value > maxs[j] {
				maxs[j] = value
			}
		}
	}

	v := 2 * m / 3
	v = max(1, v)

	for i := range n {
		for j := range m {
			if maxs[j] == mins[j] {
				x[i][j] = 0
			} else {
				x[i][j] = (x[i][j] - mins[j]) / (maxs[j] - mins[j])
			}
		}
	}

	rng := rand.New(rand.NewSource(42))
	projection := make([][]float64, m)
	scale := 1.0 / math.Sqrt(float64(v))

	for j := range m {
		projection[j] = make([]float64, v)
		for k := range v {
			if rng.Intn(2) == 0 {
				projection[j][k] = -scale
			} else {
				projection[j][k] = scale
			}
		}
	}

	writer.WriteString(strconv.Itoa(v))
	writer.WriteByte('\n')

	for i := range n {
		result := make([]float64, v)
		for j := range m {
			for k := range v {
				result[k] += x[i][j] * projection[j][k]
			}
		}

		for k := range v {
			if k > 0 {
				writer.WriteByte(' ')
			}
			writer.WriteString(strconv.FormatFloat(result[k], 'g', -1, 64))
		}

		writer.WriteByte('\n')
	}
}
