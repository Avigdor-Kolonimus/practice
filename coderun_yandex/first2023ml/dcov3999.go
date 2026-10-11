package first2023ml

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/first-2023-ml/problems/dcov3999
// Dcov3999 - problem 24
func Dcov3999() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	readInts := func() []int {
		line, _ := reader.ReadString('\n')
		fields := strings.Fields(line)
		values := make([]int, len(fields))

		for i, field := range fields {
			values[i], _ = strconv.Atoi(field)
		}

		return values
	}

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	initial := readInts()
	infected := make([]bool, n)

	for i, v := range initial {
		infected[i] = v == 1
	}

	meetings := make(map[int][]int)

	for i := 0; i < n; i++ {
		values := readInts()
		k := values[0]

		for j := 1; j <= k; j++ {
			id := values[j]
			meetings[id] = append(meetings[id], i)
		}
	}

	ids := make([]int, 0, len(meetings))
	for id := range meetings {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	for _, id := range ids {
		participants := meetings[id]

		hasInfected := false
		for _, person := range participants {
			if infected[person] {
				hasInfected = true
				break
			}
		}

		if hasInfected {
			for _, person := range participants {
				infected[person] = true
			}
		}
	}

	for i, v := range infected {
		if i > 0 {
			writer.WriteByte(' ')
		}

		if v {
			writer.WriteString("1")
		} else {
			writer.WriteString("0")
		}
	}
	writer.WriteByte('\n')
}
