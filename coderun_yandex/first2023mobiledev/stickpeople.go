package first2023mobiledev

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/stick-people
// StickPeople - problem 8
func StickPeople() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	a := make([]int, n)
	b := make([]int, n)

	line, _ = reader.ReadString('\n')
	parts := strings.Fields(line)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(parts[i])
	}

	line, _ = reader.ReadString('\n')
	parts = strings.Fields(line)
	for i := 0; i < n; i++ {
		b[i], _ = strconv.Atoi(parts[i])
	}

	bestStart := 1
	bestCost := int64(1 << 60)

	for start := 0; start < n; start++ {
		var cost int64

		for i := 0; i < n; i++ {
			dog := (start + i) % n
			neck := a[dog]
			collar := b[i]

			if collar < neck {
				continue
			}

			diff := collar - neck

			if diff <= 100 {
				cost += int64(diff / 2)
			} else {
				cost += 30
			}
		}

		if cost < bestCost {
			bestCost = cost
			bestStart = start + 1
		}
	}

	writer.WriteString(strconv.Itoa(bestStart))
	writer.WriteByte(' ')
	writer.WriteString(strconv.FormatInt(bestCost, 10))
	writer.WriteByte('\n')
}
