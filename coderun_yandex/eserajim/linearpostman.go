package eserajim

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	cur  uint32
	a, b uint32
)

func nextRand24() uint32 {
	cur = cur*a + b
	return cur >> 8
}

func nextRand32() uint32 {
	A := nextRand24()
	B := nextRand24()

	return (A << 8) ^ B
}

// https://coderun.yandex.ru/selections/eserajim/problems/linear-postman
// LinearPostman - problem 6
func LinearPostman() {
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

	// A and B input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	t, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	a = uint32(t)

	t, err = strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	b = uint32(t)

	// generation
	x := make([]uint32, n)
	for i := 0; i < n; i++ {
		x[i] = nextRand32()
	}

	// radix sort by 4 byte
	tmp := make([]uint32, n)
	for shift := 0; shift < 32; shift += 8 {
		cnt := make([]int, 256)

		// counting sort
		for _, v := range x {
			byteVal := (v >> shift) & 255
			cnt[byteVal]++
		}

		pos := make([]int, 256)
		for i := 1; i < 256; i++ {
			pos[i] = pos[i-1] + cnt[i-1]
		}

		for _, v := range x {
			byteVal := (v >> shift) & 255
			tmp[pos[byteVal]] = v
			pos[byteVal]++
		}

		x, tmp = tmp, x
	}

	// mediana
	med := x[n/2]

	var ans uint64
	for _, v := range x {
		if v > med {
			ans += uint64(v - med)
		} else {
			ans += uint64(med - v)
		}
	}

	writer.WriteString(strconv.Itoa(int(ans)))
	writer.WriteByte('\n')
}
