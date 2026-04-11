package assignments

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func validateFirstLineInput(n int) bool {
	return n >= 2 && n <= 100_000
}

func validateWeightInput(n int) bool {
	return n >= 1 && n <= 1_000
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/mushroom-sharing/description
// SharingMushrooms - assignment 1
// Description:
// Vasya and Masha are sharing mushrooms. Vasya is the first to take mushrooms, and Masha is the second.
func SharingMushrooms() {
	sumOdd := 0  // Vasya
	sumEven := 0 // Masha
	joy := 0
	maxEvenWeight := -1
	minOddWeight := math.MaxInt

	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first line
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	n, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}
	if !validateFirstLineInput(n) {
		panic("n out of range")
	}

	// second line
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	mushroomWeights := strings.Fields(line)
	weightsCount := len(mushroomWeights)
	if weightsCount != n {
		panic("mushroom weights count does not match n")
	}

	nums := make([]int, weightsCount)
	for i, v := range mushroomWeights {
		weight, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		if !validateWeightInput(weight) {
			panic("weight out of range")
		}

		nums[i] = weight
	}

	for index, a := range nums {
		if (index+1)%2 == 0 {
			sumEven += a
			if a > maxEvenWeight {
				maxEvenWeight = a
			}
		} else {
			sumOdd += a
			if a < minOddWeight {
				minOddWeight = a
			}
		}
	}

	if maxEvenWeight <= minOddWeight {
		joy = sumOdd - sumEven
	} else {
		joy = (sumOdd - minOddWeight + maxEvenWeight) - (sumEven - maxEvenWeight + minOddWeight)
	}

	writer.WriteString(strconv.Itoa(joy))
	writer.WriteByte('\n')
}
