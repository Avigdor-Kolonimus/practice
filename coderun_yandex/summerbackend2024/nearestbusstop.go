package summerbackend2024

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/nearest-bus-stop
// NearestBusStop - problem 5
func NearestBusStop() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	first := strings.Fields(line)

	n, _ := strconv.Atoi(first[0])
	k, _ := strconv.Atoi(first[1])

	line, _ = reader.ReadString('\n')
	fields := strings.Fields(line)

	a := make([]int, n)
	for i := range n {
		a[i], _ = strconv.Atoi(fields[i])
	}

	for range k {
		line, err := reader.ReadString('\n')
		if len(strings.TrimSpace(line)) == 0 && err != nil {
			break
		}

		fields := strings.Fields(line)
		for _, field := range fields {
			x, _ := strconv.Atoi(field)

			l, r := -1, n

			for l+1 < r {
				m := (l + r) / 2

				if a[m] < x {
					l = m
				} else {
					r = m
				}
			}

			var ans int

			switch {
			case r == 0:
				ans = 1
			case r == n:
				ans = n
			case x-a[l] <= a[r]-x:
				ans = l + 1
			default:
				ans = r + 1
			}

			writer.WriteString(strconv.Itoa(ans))
			writer.WriteByte('\n')
		}
	}
}
