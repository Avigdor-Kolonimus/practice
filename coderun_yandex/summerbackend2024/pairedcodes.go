package summerbackend2024

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/paired-codes
// PairedCodes - problem 25
func PairedCodes() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	const bits = 10
	const size = 1 << bits
	const fullMask = size - 1

	count := make([]int64, size)
	for range n {
		x, _ := reader.ReadString(' ')
		x = strings.TrimSpace(x)

		value, _ := strconv.Atoi(x)

		mask := 0

		if value == 0 {
			mask = 1
		} else {
			for value > 0 {
				digit := value % 10
				mask |= 1 << digit
				value /= 10
			}
		}

		count[mask]++
	}

	subsetSum := make([]int64, size)
	copy(subsetSum, count)

	for bit := 0; bit < bits; bit++ {
		for mask := 0; mask < size; mask++ {
			if mask&(1<<bit) != 0 {
				subsetSum[mask] += subsetSum[mask^(1<<bit)]
			}
		}
	}

	var bad int64
	for mask := 0; mask < size; mask++ {
		complement := fullMask ^ mask
		bad += count[mask] * subsetSum[complement]
	}

	bad /= 2

	total := int64(n) * int64(n-1) / 2

	writer.WriteString(strconv.FormatInt(total-bad, 10))
	writer.WriteByte('\n')
}
