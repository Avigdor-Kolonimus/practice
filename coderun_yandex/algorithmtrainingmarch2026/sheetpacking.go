package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func needSheetPacking(x, limit int) int {
	cnt := 0
	for x > limit {
		x = (x + 1) / 2
		cnt++
	}

	return cnt
}

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/sheet-packing
// SheetPacking - assignment 4
func SheetPacking() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N, M, H and W input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 4 {
		panic("numbers count does not match 4")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	h, err := strconv.Atoi(strNum[2])
	if err != nil {
		panic(err)
	}
	w, err := strconv.Atoi(strNum[3])
	if err != nil {
		panic(err)
	}

	ans1 := needSheetPacking(n, h) + needSheetPacking(m, w)
	ans2 := needSheetPacking(n, w) + needSheetPacking(m, h)

	if ans2 < ans1 {
		ans1 = ans2
	}

	writer.WriteString(strconv.Itoa(ans1))
	writer.WriteByte('\n')
}
