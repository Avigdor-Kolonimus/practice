package first2023mobiledev

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/woods
// Woods - problem 11
func Woods() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	data, _ := io.ReadAll(reader)
	pos := 0

	nextInt := func() int {
		for pos < len(data) && data[pos] <= ' ' {
			pos++
		}

		num := 0
		for pos < len(data) && data[pos] >= '0' && data[pos] <= '9' {
			num = num*10 + int(data[pos]-'0')
			pos++
		}
		return num
	}

	h := nextInt()
	n := nextInt()

	x := make([]int, n)
	total := 0

	for i := 0; i < n; i++ {
		x[i] = nextInt()
		total += x[i]
	}

	check := func(w int) bool {
		rows := 0
		length := 0

		for _, plank := range x {
			length += plank

			if length >= w {
				rows++
				length = 0

				if rows >= h {
					return true
				}
			}
		}

		return false
	}

	left, right := 1, total/h
	ans := 0

	for left <= right {
		mid := left + (right-left)/2

		if check(mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Fprintln(writer, ans)
}
