package summercommon2025

import (
	"bufio"
	"os"
	"strconv"
)

// https://coderun.yandex.ru/selections/2025-summer-common/problems/ancient-tunnels
// AncientTunnels - problem 22
func AncientTunnels(n int, a []int) {
	result := make([]int, n)
	for i := range result {
		result[i] = 1
	}

	inbound := make([]int, n)
	for _, to := range a {
		if to != -1 {
			inbound[to-1]++
		}
	}

	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if inbound[i] == 0 {
			queue = append(queue, i)
		}
	}

	head := 0
	for head < len(queue) {
		room := queue[head]
		head++

		if a[room] == -1 {
			continue
		}

		next := a[room] - 1
		result[next] += result[room]
		inbound[next]--

		if inbound[next] == 0 {
			queue = append(queue, next)
		}
	}

	for room := 0; room < n; room++ {
		if inbound[room] == 0 {
			continue
		}

		count := 0
		cur := room
		for inbound[cur] != 0 {
			count += result[cur]
			inbound[cur] = 0
			cur = a[cur] - 1
		}

		for result[cur] != count {
			result[cur] = count
			cur = a[cur] - 1
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i, v := range result {
		if i > 0 {
			writer.WriteByte(' ')
		}
		writer.WriteString(strconv.Itoa(v))
	}
	writer.WriteByte('\n')
}

// ввод/вывод
// не изменяйте сигнатуру метода
func AncientTunnelSolve() {
	input := NewFastScanner(os.Stdin)

	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	t := input.readInt()
	for test := 0; test < t; test++ {
		n := input.readInt()
		a := input.readIntArray(n)

		AncientTunnels(n, a)
	}
}
