package codelifebalance

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func validateStylishClothesInput(p int) bool {
	return p >= 1 && p <= 100_000
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// https://coderun.yandex.ru/selections/code-life-balance/problems/stylish-clothes
// StylishClothes - assignment 8
func StylishClothes() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateStylishClothesInput(n) {
		panic("number N out of range")
	}

	// slice input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("input does not match n")
	}
	sliceShirt := make([]int, n)
	for j := range n {
		curr, err := strconv.Atoi(strNum[j])
		if err != nil {
			panic(err)
		}

		sliceShirt[j] = curr
	}

	// M input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	m, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateStylishClothesInput(m) {
		panic("number M out of range")
	}

	// slice input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	strNum = strings.Fields(line)
	if len(strNum) != m {
		panic("input does not match m")
	}
	slicePants := make([]int, m)
	for j := range m {
		curr, err := strconv.Atoi(strNum[j])
		if err != nil {
			panic(err)
		}

		slicePants[j] = curr
	}

	i, j := 0, 0
	bestDiff := int(1e18)
	bestShirt, bestPants := 0, 0
	for i < n && j < m {
		diff := abs(sliceShirt[i] - slicePants[j])

		if diff < bestDiff {
			bestDiff = diff
			bestShirt = sliceShirt[i]
			bestPants = slicePants[j]
		}

		if sliceShirt[i] < slicePants[j] {
			i++
		} else {
			j++
		}
	}

	writer.WriteString(strconv.Itoa(bestShirt) + " " + strconv.Itoa(bestPants))
	writer.WriteByte('\n')
}
