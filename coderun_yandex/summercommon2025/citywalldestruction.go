package summercommon2025

import (
	"bufio"
	"os"
	"strconv"
)

func validateCityWallDestructionInput(n int) bool {
	return n >= 1 && n <= 300_000
}

func solveCityWallDestruction(n int, a []int) int {
	if n == 0 {
		return 0
	}

	leftGreater := make([]int, n)
	rightGreater := make([]int, n)
	stack := []int{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && a[stack[len(stack)-1]] <= a[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			leftGreater[i] = stack[len(stack)-1]
		} else {
			leftGreater[i] = -1
		}
		stack = append(stack, i)
	}

	stack = []int{}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && a[stack[len(stack)-1]] < a[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			rightGreater[i] = stack[len(stack)-1]
		} else {
			rightGreater[i] = n
		}
		stack = append(stack, i)
	}

	var sumMax int64
	for i := 0; i < n; i++ {
		left := int64(i - leftGreater[i])
		right := int64(rightGreater[i] - i)
		sumMax += int64(a[i]) * left * right
	}

	leftSmaller := make([]int, n)
	rightSmaller := make([]int, n)
	stack = []int{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && a[stack[len(stack)-1]] >= a[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			leftSmaller[i] = stack[len(stack)-1]
		} else {
			leftSmaller[i] = -1
		}
		stack = append(stack, i)
	}

	stack = []int{}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && a[stack[len(stack)-1]] > a[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			rightSmaller[i] = stack[len(stack)-1]
		} else {
			rightSmaller[i] = n
		}
		stack = append(stack, i)
	}

	var sumMin int64
	for i := 0; i < n; i++ {
		left := int64(i - leftSmaller[i])
		right := int64(rightSmaller[i] - i)
		sumMin += int64(a[i]) * left * right
	}

	return int(sumMax - sumMin)
}

// https://coderun.yandex.ru/selections/2025-summer-common/problems/city-wall-destruction
// CityWallDestruction - problem 17
func CityWallDestruction() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line := mustReadIntArray(reader, 1)
	n := line[0]
	if !validateStairsInput(n) {
		panic("number N out of range")
	}

	// height input
	height := make([]int, n)
	line = mustReadIntArray(reader, n)
	for i := range n {
		height[i] = line[i]
	}

	answer := solveCityWallDestruction(n, height)

	writer.WriteString(strconv.Itoa(answer))
	writer.WriteByte('\n')
}
