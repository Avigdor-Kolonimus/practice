package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/problem/robot
// Robot - problem 217
func Robot() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	k, _ := strconv.Atoi(strings.TrimSpace(line))

	line, _ = reader.ReadString('\n')
	s := strings.TrimSpace(line)

	n := len(s)
	if k == 0 {
		writer.WriteString("0\n")
		return
	}

	var ans int64
	cnt := int64(0)
	for i := n - 1; i >= k; i-- {
		if s[i] == s[i-k] {
			cnt++
		} else {
			cnt = 0
		}

		ans += cnt
	}

	writer.WriteString(strconv.FormatInt(ans, 10))
	writer.WriteByte('\n')
}
