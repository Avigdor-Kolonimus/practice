package first2023mobiledev

import (
	"bufio"
	"os"
	"strings"
)

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/max-num-split
// MaxNumSplit - problem 17
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, _ := reader.ReadString('\n')
	n := strings.TrimSpace(line)

	length := len(n)
	total := 1 << (length - 1)

	best := n
	bestCount := 1
	for mask := 0; mask < total; mask++ {
		parts := make([]string, 0, length)
		seen := make(map[string]bool)
		start := 0
		valid := true

		for i := 0; i < length; i++ {
			if i == length-1 || mask&(1<<i) != 0 {
				part := n[start : i+1]

				if len(part) > 1 && part[0] == '0' {
					valid = false
					break
				}

				if seen[part] {
					valid = false
					break
				}

				seen[part] = true
				parts = append(parts, part)
				start = i + 1
			}
		}

		if valid && len(parts) > bestCount {
			bestCount = len(parts)
			best = strings.Join(parts, "-")
		}
	}

	writer.WriteString(best)
	writer.WriteByte('\n')
}
