package wintercommon2025

import (
	"bufio"
	"os"
	"strconv"
)

func solveCalibratingMagicAltimeter(nums []int) int {
	const target = 100
	bestSum := 0
	bestDist := target

	for mask := 0; mask < (1 << 10); mask++ {
		sum := 0
		for i := 0; i < 10; i++ {
			if mask&(1<<i) != 0 {
				sum += nums[i]
			}
		}

		dist := absCalibratingMagicAltimeter(sum - target)

		if dist < bestDist || (dist == bestDist && sum > bestSum) {
			bestSum = sum
			bestDist = dist
		}
	}

	return bestSum
}

func absCalibratingMagicAltimeter(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// https://coderun.yandex.ru/selections/2025-winter-common/problems/calibrating-magic-altimeter
// CalibratingMagicAltimeter - problem 2
func CalibratingMagicAltimeter() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		line := mustReadIntArray(reader, 1)
		nums[i] = line[0]
	}

	answer := solveCalibratingMagicAltimeter(nums)

	writer.WriteString(strconv.Itoa(answer))
	writer.WriteByte('\n')
}
