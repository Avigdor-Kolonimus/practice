package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var dirs = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func validateSapperMandNInput(p int) bool {
	return p >= 1 && p <= 100
}

func validateSapperPandQInput(p, nOrm int) bool {
	return p >= 1 && p <= nOrm
}

func validateSapperMineInput(k, nm int) bool {
	return k >= 0 && k <= nm
}

// https://coderun.yandex.ru/problem/sapper
// Sapper - problem 63
func Sapper() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first line
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	strNum := strings.Fields(line)
	if len(strNum) != 3 {
		panic("numbers count does not match 3")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	if !validateSapperMandNInput(n) {
		panic("number N out of range")
	}

	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	if !validateSapperMandNInput(m) {
		panic("number M out of range")
	}

	k, err := strconv.Atoi(strNum[2])
	if err != nil {
		panic(err)
	}
	if !validateSapperMineInput(k, n*m) {
		panic("number K out of range")
	}

	// mine inputs
	sapperBoard := make([][]int, n)
	for i := range sapperBoard {
		sapperBoard[i] = make([]int, m)
	}

	for range k {
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		strNum := strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		p, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		if !validateSapperPandQInput(p, n) {
			panic("number P out of range")
		}

		q, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}
		if !validateSapperPandQInput(q, m) {
			panic("number Q out of range")
		}

		p--
		q--
		sapperBoard[p][q] = -1
		for _, d := range dirs {
			nr, nc := p+d[0], q+d[1]
			if nr >= 0 && nr < n && nc >= 0 && nc < m && sapperBoard[nr][nc] != -1 {
				sapperBoard[nr][nc]++
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j > 0 {
				writer.WriteString(" ")
			}
			if sapperBoard[i][j] == -1 {
				writer.WriteString("*")
			} else {
				writer.WriteString(strconv.Itoa(sapperBoard[i][j]))
			}
		}
		writer.WriteByte('\n')
	}
}
