package problems

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/collector-diego
// CollectorDiego - problem 245
func CollectorDiego() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// stickers input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	stickersMap := make(map[int]struct{})
	for i := 0; i < n; i++ {
		x, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		stickersMap[x] = struct{}{}
	}

	var stickers []int
	for x := range stickersMap {
		stickers = append(stickers, x)
	}

	sort.Ints(stickers)

	// K line
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	k, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// p input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum = strings.Fields(line)
	if len(strNum) != k {
		panic("numbers count does not match k")
	}
	for i := 0; i < k; i++ {
		p, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		ans := sort.SearchInts(stickers, p)

		writer.WriteString(strconv.Itoa(ans))
		writer.WriteByte('\n')
	}
}
