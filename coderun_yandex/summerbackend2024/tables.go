package summerbackend2024

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/tables
// Tables - problem 2
func Tables() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	// First table input
	first := make([][]int, n)
	for i := range n {
		line, _ = reader.ReadString('\n')
		parts = strings.Fields(line)
		first[i] = make([]int, m)
		for j := 0; j < m; j++ {
			first[i][j], _ = strconv.Atoi(parts[j])
		}
	}

	// Second table input
	second := make([][]int, n)
	for i := range n {
		line, _ = reader.ReadString('\n')
		parts = strings.Fields(line)
		second[i] = make([]int, m)
		for j := 0; j < m; j++ {
			second[i][j], _ = strconv.Atoi(parts[j])
		}
	}
	answer := 0
	writer.WriteString(strconv.Itoa(answer))
	writer.WriteByte('\n')
}
