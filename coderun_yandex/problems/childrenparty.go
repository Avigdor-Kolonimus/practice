package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Worker struct {
	t, z, y int
}

func produced(w Worker, time int) int {
	cycle := w.t*w.z + w.y
	full := time / cycle
	rem := time % cycle

	res := full * w.z

	extra := rem / w.t
	if extra > w.z {
		extra = w.z
	}

	return res + extra
}

// https://coderun.yandex.ru/problem/children-party
// ChildrenParty - problem 238
func ChildrenParty() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	line, _ := reader.ReadString('\n')
	first := strings.Fields(line)

	m, _ := strconv.Atoi(first[0])
	n, _ := strconv.Atoi(first[1])

	// workers input
	workers := make([]Worker, n)
	for i := range n {
		line, _ = reader.ReadString('\n')
		fields := strings.Fields(line)

		t, _ := strconv.Atoi(fields[0])
		z, _ := strconv.Atoi(fields[1])
		y, _ := strconv.Atoi(fields[2])

		workers[i] = Worker{t, z, y}
	}

	if m == 0 {
		writer.WriteString("0\n")
		for i := 0; i < n; i++ {
			if i > 0 {
				writer.WriteByte(' ')
			}
			writer.WriteByte('0')
		}
		writer.WriteByte('\n')
		return
	}

	left, right := 0, 10000000

	for left < right {
		mid := (left + right) / 2

		total := 0
		for _, w := range workers {
			total += produced(w, mid)
			if total >= m {
				break
			}
		}

		if total >= m {
			right = mid
		} else {
			left = mid + 1
		}
	}

	answer := left

	writer.WriteString(strconv.Itoa(answer))
	writer.WriteByte('\n')

	remain := m

	for i, w := range workers {
		cnt := produced(w, answer)
		if cnt > remain {
			cnt = remain
		}

		if i > 0 {
			writer.WriteByte(' ')
		}

		writer.WriteString(strconv.Itoa(cnt))
		remain -= cnt
	}

	writer.WriteByte('\n')
}
