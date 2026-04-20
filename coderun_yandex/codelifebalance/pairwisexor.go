package codelifebalance

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func validatePairwiseXorTInput(p int) bool {
	return p >= 1 && p <= 1_000
}

func validatePairwiseXorNInput(p int) bool {
	return p >= 2 && p <= 1_000_000
}

// https://coderun.yandex.ru/selections/code-life-balance/problems/pairwise-xor
// PairwiseXor - assignment 6
func PairwiseXor() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// T input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	t, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validatePairwiseXorTInput(t) {
		panic("number T out of range")
	}

	// tests
	results := make([]int, 0, t)
	for range t {
		// N input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		if !validatePairwiseXorNInput(n) {
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
		sliceA := make([]int, n)
		for j := range n {
			curr, err := strconv.Atoi(strNum[j])
			if err != nil {
				panic(err)
			}

			sliceA[j] = curr
		}

		sort.Ints(sliceA)

		minXor := int(1<<31 - 1)
		for i := 0; i < n-1; i++ {
			x := sliceA[i] ^ sliceA[i+1]
			if x < minXor {
				minXor = x
			}
		}

		results = append(results, minXor)
	}

	for _, v := range results {
		writer.WriteString(strconv.Itoa(v))
		writer.WriteByte('\n')
	}
}
