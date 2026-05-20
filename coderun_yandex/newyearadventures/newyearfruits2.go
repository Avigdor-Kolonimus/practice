package newyearadventures

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type BoxNewYearFruits2 struct {
	mandarin int
	orange   int
	idx      int
}

// https://coderun.yandex.ru/selections/new-year-adventures/problems/new-year-fruits-2
// NewYearFruits2 - problem 9
func NewYearFruits2() {
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

	// box input
	mo := make([]BoxNewYearFruits2, 2*n-1)
	for i := 0; i < 2*n-1; i++ {
		// fruits input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum := strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		m, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}

		o, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		mo[i].mandarin = m
		mo[i].orange = o
		mo[i].idx = i + 1
	}

	sort.Slice(mo, func(i, j int) bool {
		return mo[i].mandarin > mo[j].mandarin
	})

	ans := []int{mo[0].idx}

	for i := 1; i < 2*n-1; i += 2 {
		if mo[i].orange > mo[i+1].orange {
			ans = append(ans, mo[i].idx)
		} else {
			ans = append(ans, mo[i+1].idx)
		}
	}

	for _, x := range ans {
		writer.WriteString(strconv.Itoa(x) + " ")
	}
	writer.WriteByte('\n')
}
