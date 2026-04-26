package wintercommon2025

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solveLocationOrientation(markersX, markersY []int, commands string) []int64 {
	N := len(markersX)

	sortedX := make([]int, N)
	copy(sortedX, markersX)
	sort.Ints(sortedX)

	sortedY := make([]int, N)
	copy(sortedY, markersY)
	sort.Ints(sortedY)

	prefixSumX := make([]int64, N+1)
	prefixSumY := make([]int64, N+1)
	for i := 0; i < N; i++ {
		prefixSumX[i+1] = prefixSumX[i] + int64(sortedX[i])
		prefixSumY[i+1] = prefixSumY[i] + int64(sortedY[i])
	}

	cx, cy := 0, 0
	results := make([]int64, 0, len(commands))

	for _, cmd := range commands {
		switch cmd {
		case 'N':
			cy++
		case 'S':
			cy--
		case 'E':
			cx++
		case 'W':
			cx--
		}

		// sum(|cx - mx| + |cy - my|) = sum(|cx - mx|) + sum(|cy - my|)

		idxX := sort.Search(N, func(i int) bool { return sortedX[i] > cx })
		// sortedX[0..idxX-1] <= cx, sortedX[idxX..N-1] > cx
		sumX := int64(cx)*int64(idxX) - prefixSumX[idxX] + (prefixSumX[N] - prefixSumX[idxX]) - int64(cx)*int64(N-idxX)

		idxY := sort.Search(N, func(i int) bool { return sortedY[i] > cy })
		// sortedY[0..idxY-1] <= cy, sortedY[idxY..N-1] > cy
		sumY := int64(cy)*int64(idxY) - prefixSumY[idxY] + (prefixSumY[N] - prefixSumY[idxY]) - int64(cy)*int64(N-idxY)

		total := sumX + sumY
		results = append(results, total)
	}

	return results
}

// https://coderun.yandex.ru/selections/2025-winter-common/problems/location-orientation
// LocationOrientation - problem 6
func LocationOrientation() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	line := mustReadIntArray(reader, 2)
	n, _ := line[0], line[1]

	markersX := make([]int, n)
	markersY := make([]int, n)
	for i := 0; i < n; i++ {
		// X and Y input
		parts := mustReadIntArray(reader, 2)
		x, y := parts[0], parts[1]
		markersX[i] = x
		markersY[i] = y
	}

	// commands input
	lineCommands, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	commands := strings.TrimSpace(lineCommands)

	results := solveLocationOrientation(markersX, markersY, commands)

	for _, result := range results {
		writer.WriteString(strconv.FormatInt(result, 10))
		writer.WriteByte('\n')
	}
}
