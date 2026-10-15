package first2023mobiledev

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/perfect-word
// PerfectWord - problem 6
func PerfectWord() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// K input
	line, _ := reader.ReadString('\n')
	k, _ := strconv.Atoi(strings.TrimSpace(line))

	// Word input
	line, _ = reader.ReadString('\n')
	word := strings.TrimSpace(line)

	n := len(word)

	if k <= 0 {
		writer.WriteString(strconv.Itoa(n))
		writer.WriteByte('\n')
		return
	}

	if k > n {
		writer.WriteString("0\n")
		return
	}

	a := []byte(word)

	check := func(perfect int) bool {
		if perfect <= 1 {
			return true
		}

		remains := k
		i := 0

		for i <= n-perfect*remains {
			if remains <= 0 {
				break
			}

			left := i
			right := i + perfect - 1

			for left < right && a[left] == a[right] {
				left++
				right--
			}

			if left >= right {
				i += perfect
				remains--
			} else {
				i++
			}
		}

		return remains == 0
	}

	result := 1

	for perfect := n / k; perfect >= 2; perfect-- {
		if check(perfect) {
			result = perfect
			break
		}
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
