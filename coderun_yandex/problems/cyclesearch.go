package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	nCycleSearch       int
	gCycleSearch       [][]int
	visitedCycleSearch []bool
	parentCycleSearch  []int
	cycleStart         = -1
	cycleEnd           = -1
)

func dfsCycleSearch(v, p int) bool {
	visitedCycleSearch[v] = true
	for to := 0; to < nCycleSearch; to++ {
		if gCycleSearch[v][to] == 0 {
			continue
		}

		if to == p {
			continue
		}

		if !visitedCycleSearch[to] {
			parentCycleSearch[to] = v
			if dfsCycleSearch(to, v) {
				return true
			}
		} else {
			cycleStart = to
			cycleEnd = v

			return true
		}
	}

	return false
}

// https://coderun.yandex.ru/problem/cycle-search
// CycleSearch - problem 11
func CycleSearch() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	nCycleSearch, err = strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// graph input
	gCycleSearch = make([][]int, nCycleSearch)
	for i := 0; i < nCycleSearch; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		tokens := strings.Fields(line)
		if len(tokens) != nCycleSearch {
			panic("invalid input")
		}

		gCycleSearch[i] = make([]int, nCycleSearch)
		for j := 0; j < nCycleSearch; j++ {
			x, err := strconv.Atoi(tokens[j])
			if err != nil {
				panic(err)
			}

			gCycleSearch[i][j] = x
		}
	}

	visitedCycleSearch = make([]bool, nCycleSearch)
	parentCycleSearch = make([]int, nCycleSearch)
	found := false
	for v := 0; v < nCycleSearch; v++ {
		if !visitedCycleSearch[v] {
			if dfsCycleSearch(v, -1) {
				found = true
				break
			}
		}
	}

	if !found {
		writer.WriteString("NO")
		writer.WriteByte('\n')

		return
	}

	writer.WriteString("YES")
	writer.WriteByte('\n')

	cycle := []int{cycleStart}
	v := cycleEnd

	for v != cycleStart {
		cycle = append(cycle, v)
		v = parentCycleSearch[v]
	}

	// reverse
	for i, j := 0, len(cycle)-1; i < j; i, j = i+1, j-1 {
		cycle[i], cycle[j] = cycle[j], cycle[i]
	}

	writer.WriteString(strconv.Itoa(len(cycle)))
	writer.WriteByte('\n')

	for _, x := range cycle {
		writer.WriteString(strconv.Itoa(x + 1))
		writer.WriteByte(' ')
	}
	writer.WriteByte('\n')
}
