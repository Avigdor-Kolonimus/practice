package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type IntervalLOS struct {
	l int
	r int
}

func intersectLOS(a, b IntervalLOS) bool {
	return !(a.r < b.l || b.r < a.l)
}

// https://coderun.yandex.ru/problem/lite-operating-systems
// LiteOperatingSystems - problem 246
func LiteOperatingSystems() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// M input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	_, err = strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

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

	// part input
	parts := make([]IntervalLOS, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		strNum := strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		l, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		parts[i].l = l
		parts[i].r = r
	}

	ans := 0
	for i := n - 1; i >= 0; i-- {
		erased := false

		for j := i + 1; j < n; j++ {
			if intersectLOS(parts[i], parts[j]) {
				erased = true

				break
			}
		}

		if !erased {
			ans++
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
