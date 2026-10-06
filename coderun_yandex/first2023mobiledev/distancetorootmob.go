package first2023mobiledev

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/distance-to-root-mob
// DistanceToRootMob - problem 14
func DistanceToRootMob() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var input []string
	for {
		line, err := reader.ReadString('\n')

		if len(line) > 0 {
			input = append(input, strings.Fields(line)...)
		}

		if err != nil {
			break
		}
	}

	n, _ := strconv.Atoi(input[0])

	p := make([]int, n+1)
	for i := 1; i <= n; i++ {
		p[i], _ = strconv.Atoi(input[i])
	}

	d := make([]int, n+1)

	for i := 1; i <= n; i++ {
		d[i] = -1
	}

	for i := 1; i <= n; i++ {
		if p[i] == 0 {
			d[i] = 0
		}
	}

	for i := 1; i <= n; i++ {
		if d[i] != -1 {
			continue
		}

		cur := i
		path := make([]int, 0)
		for d[cur] == -1 {
			path = append(path, cur)
			cur = p[cur]
		}

		depth := d[cur]

		for j := len(path) - 1; j >= 0; j-- {
			depth++
			d[path[j]] = depth
		}
	}

	for i := 1; i <= n; i++ {
		if i > 1 {
			writer.WriteByte(' ')
		}

		writer.WriteString(strconv.Itoa(d[i]))
	}

	writer.WriteByte('\n')
}
