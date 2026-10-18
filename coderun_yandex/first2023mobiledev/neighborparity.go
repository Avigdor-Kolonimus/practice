package first2023mobiledev

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Odd  = 1
	Even = 0
)

func expectedParity(index, positions int) int {
	if positions == Even {
		return index % 2
	}
	if index%2 == 0 {
		return 1
	}
	return 0
}

func buildResultNeighborParity(a []int, positions int) (int, []int) {
	result := append([]int(nil), a...)

	wrongEven := make([]int, 0)
	wrongOdd := make([]int, 0)

	for i, x := range a {
		expected := expectedParity(i, positions)
		if x%2 != expected {
			if x%2 == 0 {
				wrongEven = append(wrongEven, x)
			} else {
				wrongOdd = append(wrongOdd, x)
			}
		}
	}

	evenIndex, oddIndex := 0, 0
	changed := 0

	for i, x := range a {
		expected := expectedParity(i, positions)
		if x%2 == expected {
			continue
		}

		changed++
		if expected == 0 {
			result[i] = wrongEven[evenIndex]
			evenIndex++
		} else {
			result[i] = wrongOdd[oddIndex]
			oddIndex++
		}
	}

	return changed, result
}

func solve(a []int) (int, []int) {
	odd, even := 0, 0

	for _, x := range a {
		if x%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	if abs(odd-even) > 1 {
		return -1, nil
	}

	if even > odd {
		return buildResultNeighborParity(a, Even)
	}
	if odd > even {
		return buildResultNeighborParity(a, Odd)
	}

	evenK, evenResult := buildResultNeighborParity(a, Even)
	oddK, oddResult := buildResultNeighborParity(a, Odd)

	if evenK <= oddK {
		return evenK, evenResult
	}
	return oddK, oddResult
}

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/neighbor-parity
// NeighborParity - problem 20
func NeighborParity() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}

	k, result := solve(a)

	if k == -1 {
		fmt.Fprintln(writer, -1)
		return
	}

	fmt.Fprintln(writer, k)
	for i, x := range result {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, x)
	}
	fmt.Fprintln(writer)
}
