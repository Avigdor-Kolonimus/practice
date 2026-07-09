package winterintern2024

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type FenwickTree struct {
	tree []int64
	size int
}

func NewFenwickTree(size int) *FenwickTree {
	return &FenwickTree{
		tree: make([]int64, size+2),
		size: size,
	}
}

func (f *FenwickTree) Update(idx int, val int64) {
	for idx <= f.size {
		f.tree[idx] += val
		idx += idx & -idx
	}
}

func (f *FenwickTree) Query(idx int) int64 {
	var res int64
	for idx > 0 {
		res += f.tree[idx]
		idx -= idx & -idx
	}

	return res
}

// https://coderun.yandex.ru/selections/winter-intern-2024/problems/lap-2
// Lap2 - problem 3
func Lap2() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	fields := strings.Fields(line)

	// N, T, S input
	n, _ := strconv.Atoi(fields[0])
	t, _ := strconv.ParseInt(fields[1], 10, 64)
	s, _ := strconv.ParseInt(fields[2], 10, 64)

	// speeds input
	line, _ = reader.ReadString('\n')
	fields = strings.Fields(line)

	cars := make([]int64, n)
	for i := range n {
		cars[i], _ = strconv.ParseInt(fields[i], 10, 64)
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i] < cars[j]
	})

	ftCounts := NewFenwickTree(int(s))
	ftDivs := NewFenwickTree(int(s))

	var totalOvertakes int64
	for _, v := range cars {
		currentVal := v * t
		currentLaps := currentVal / s
		currentRem := currentVal % s

		countSmallerV := ftCounts.Query(int(s))
		totalLapsContribution := countSmallerV * currentLaps

		sumPrevLaps := ftDivs.Query(int(s))

		countRemGeCurRem := ftCounts.Query(int(s)) - ftCounts.Query(int(currentRem))

		currentCarOvertakes := totalLapsContribution - sumPrevLaps - countRemGeCurRem
		totalOvertakes += currentCarOvertakes

		ftCounts.Update(int(currentRem)+1, 1)
		ftDivs.Update(int(currentRem)+1, currentLaps)
	}

	writer.WriteString(strconv.FormatInt(totalOvertakes, 10))
	writer.WriteByte('\n')
}
