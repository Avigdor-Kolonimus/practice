package problems

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func validateCubesInput(n int) bool {
	return n >= 0 && n <= 100_000
}

// https://coderun.yandex.ru/problem/cubes
// Cubes - problem 199
func Cubes() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strKeys := strings.Fields(line)
	if len(strKeys) != 2 {
		panic("invalid input")
	}

	n, err := strconv.Atoi(strKeys[0])
	if err != nil {
		panic(err)
	}
	if !validateCubesInput(n) {
		panic("number N out of range")
	}

	m, err := strconv.Atoi(strKeys[1])
	if err != nil {
		panic(err)
	}
	if !validateCubesInput(m) {
		panic("number M out of range")
	}

	// Anya input
	anyaCubes := make(map[int]struct{})
	for i := 0; i < n; i++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		key, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		anyaCubes[key] = struct{}{}
	}

	// Borya input
	boryaCubes := make(map[int]struct{})
	for i := 0; i < m; i++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		key, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		boryaCubes[key] = struct{}{}
	}

	var common []int
	var onlyA []int
	var onlyB []int

	for x := range anyaCubes {
		if _, ok := boryaCubes[x]; ok {
			common = append(common, x)
		} else {
			onlyA = append(onlyA, x)
		}
	}

	for x := range boryaCubes {
		if _, ok := anyaCubes[x]; !ok {
			onlyB = append(onlyB, x)
		}
	}

	sort.Ints(common)
	sort.Ints(onlyA)
	sort.Ints(onlyB)

	writer.WriteString(strconv.Itoa(len(common)))
	writer.WriteByte('\n')
	for i, x := range common {
		if i > 0 {
			writer.WriteByte(' ')
		}
		writer.WriteString(strconv.Itoa(x))
	}
	writer.WriteByte('\n')

	writer.WriteString(strconv.Itoa(len(onlyA)))
	writer.WriteByte('\n')
	for i, x := range onlyA {
		if i > 0 {
			writer.WriteByte(' ')
		}
		writer.WriteString(strconv.Itoa(x))
	}
	writer.WriteByte('\n')

	writer.WriteString(strconv.Itoa(len(onlyB)))
	writer.WriteByte('\n')
	for i, x := range onlyB {
		if i > 0 {
			writer.WriteByte(' ')
		}
		writer.WriteString(strconv.Itoa(x))
	}
	writer.WriteByte('\n')
}
