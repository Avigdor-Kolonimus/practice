package wintercommon2025

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func mustReadLine(reader *bufio.Reader) string {
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(line)
}

func mustReadIntArray(reader *bufio.Reader, size int) []int {
	rawNumbers := strings.Split(mustReadLine(reader), " ")
	if len(rawNumbers) != size {
		panic("len must be eq size")
	}

	result := make([]int, 0, size)
	for _, rawNumber := range rawNumbers {
		number, err := strconv.Atoi(rawNumber)
		if err != nil {
			panic(err)
		}
		result = append(result, number)
	}

	return result
}

func solveThermalPanels(R, B int) (int, int) {
	sum := (R + 4) / 2

	for d := 1; d*d <= B; d++ {
		if B%d != 0 {
			continue
		}

		// Проверяем оба варианта: (W-2, H-2) = (d, B/d) и (B/d, d)
		for _, w := range []int{B/d + 2, d + 2} {
			h := sum - w
			if (w-2)*(h-2) == B && w >= h {
				return w, h
			}
		}
	}

	return 0, 0
}

func validateThermalPanelsInput(n int) bool {
	return n >= 1 && n <= 1_000_000_000
}

// https://coderun.yandex.ru/selections/2025-winter-common/problems/thermal-panels
// ThermalPanels - problem 1
func ThermalPanels() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// R and B input
	firstLine := mustReadIntArray(reader, 2)
	if !validateThermalPanelsInput(firstLine[0]) {
		panic("number R out of range")
	}
	if !validateThermalPanelsInput(firstLine[1]) {
		panic("number B out of range")
	}
	r, b := firstLine[0], firstLine[1]

	w, h := solveThermalPanels(r, b)

	writer.WriteString(strconv.Itoa(w) + " " + strconv.Itoa(h))
	writer.WriteByte('\n')
}
