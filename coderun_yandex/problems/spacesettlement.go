package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func correctPlacement(n, a, b, w, h, s int64) bool {
	eW := a + 2*s
	eH := b + 2*s

	if eW > w || eH > h {
		return false
	}

	columns := w / eW
	rows := h / eH

	requiredRows := n / columns
	if n%columns != 0 {
		requiredRows++
	}

	return rows >= requiredRows
}

func findShield(n, a, b, w, h int64) int64 {
	left := int64(0)
	right := w + 1 - a

	if !correctPlacement(n, a, b, w, h, left) {
		return 0
	}

	for left+1 < right {
		mid := (left + right) / 2

		if correctPlacement(n, a, b, w, h, mid) {
			left = mid
		} else {
			right = mid
		}
	}

	return left
}

// https://coderun.yandex.ru/problem/space-settlement
// SpaceSettlement - problem 77
func SpaceSettlement() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// N, A, B, W and H input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	fields := strings.Fields(line)

	n, err := strconv.ParseInt(fields[0], 10, 64)
	if err != nil {
		panic(err)
	}
	a, err := strconv.ParseInt(fields[1], 10, 64)
	if err != nil {
		panic(err)
	}
	b, err := strconv.ParseInt(fields[2], 10, 64)
	if err != nil {
		panic(err)
	}
	w, err := strconv.ParseInt(fields[3], 10, 64)
	if err != nil {
		panic(err)
	}
	h, err := strconv.ParseInt(fields[4], 10, 64)
	if err != nil {
		panic(err)
	}

	shield1 := findShield(n, a, b, w, h)
	shield2 := findShield(n, b, a, w, h)

	if shield2 > shield1 {
		shield1 = shield2
	}

	writer.WriteString(strconv.FormatInt(shield1, 10))
	writer.WriteByte('\n')
}
