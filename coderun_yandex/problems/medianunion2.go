package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	INFMEDIAN2 = int(1 << 60)
)

func leftMedian2(a, b []int) int {
	n := len(a)

	lo, hi := 0, n

	for lo <= hi {
		i := (lo + hi) / 2
		j := n - i

		aleft := -INFMEDIAN2
		if i > 0 {
			aleft = a[i-1]
		}

		aright := INFMEDIAN2
		if i < n {
			aright = a[i]
		}

		bleft := -INFMEDIAN2
		if j > 0 {
			bleft = b[j-1]
		}

		bright := INFMEDIAN2
		if j < n {
			bright = b[j]
		}

		if aleft <= bright && bleft <= aright {
			if aleft > bleft {
				return aleft
			}
			return bleft
		}

		if aleft > bright {
			hi = i - 1
		} else {
			lo = i + 1
		}
	}

	return -1
}

// https://coderun.yandex.ru/problem/median-union-2
// MedianUnion2 - problem 81
func MedianUnion2() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and L input
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
	l, err := strconv.Atoi(strKeys[1])
	if err != nil {
		panic(err)
	}

	// list input
	seqs := make([][]int, n)
	for k := 0; k < n; k++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		strNum := strings.Fields(line)
		if len(strNum) != 5 {
			panic("numbers count does not match 5")
		}

		x1, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		d1, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}
		a, err := strconv.Atoi(strNum[2])
		if err != nil {
			panic(err)
		}
		c, err := strconv.Atoi(strNum[3])
		if err != nil {
			panic(err)
		}
		m, err := strconv.Atoi(strNum[4])
		if err != nil {
			panic(err)
		}

		arr := make([]int, l)
		arr[0] = x1

		d := d1

		for i := 1; i < l; i++ {
			arr[i] = arr[i-1] + d
			d = (a*d + c) % m
		}

		seqs[k] = arr
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			median := leftMedian2(seqs[i], seqs[j])

			writer.WriteString(strconv.Itoa(median))
			writer.WriteByte('\n')
		}
	}
}
