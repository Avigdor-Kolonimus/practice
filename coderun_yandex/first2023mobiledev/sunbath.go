package first2023mobiledev

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const maxTime = 1000005

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/sunbath
// Sunbath - problem 26
func Sunbath() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N, S and T input
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	n, _ := strconv.Atoi(parts[0])
	s, _ := strconv.Atoi(parts[1])
	t, _ := strconv.Atoi(parts[2])

	var input strings.Builder
	for {
		line, err := reader.ReadString('\n')
		input.WriteString(line)
		if err != nil {
			break
		}
	}

	times := strings.Fields(input.String())

	tree := make([]int, 4*maxTime)
	lazy := make([]int, 4*maxTime)

	var push func(int)
	push = func(node int) {
		if lazy[node] != 0 {
			v := lazy[node]
			left, right := node*2, node*2+1
			tree[left] += v
			lazy[left] += v
			tree[right] += v
			lazy[right] += v
			lazy[node] = 0
		}
	}

	var update func(int, int, int, int, int, int)
	update = func(node, l, r, ql, qr, val int) {
		if ql <= l && r <= qr {
			tree[node] += val
			lazy[node] += val
			return
		}

		push(node)
		mid := (l + r) / 2

		if ql <= mid {
			update(node*2, l, mid, ql, qr, val)
		}
		if qr > mid {
			update(node*2+1, mid+1, r, ql, qr, val)
		}

		if tree[node*2] > tree[node*2+1] {
			tree[node] = tree[node*2]
		} else {
			tree[node] = tree[node*2+1]
		}
	}

	var query func(int, int, int, int, int) int
	query = func(node, l, r, ql, qr int) int {
		if ql <= l && r <= qr {
			return tree[node]
		}

		push(node)
		mid := (l + r) / 2
		result := 0

		if ql <= mid {
			result = query(node*2, l, mid, ql, qr)
		}
		if qr > mid {
			v := query(node*2+1, mid+1, r, ql, qr)
			if v > result {
				result = v
			}
		}

		return result
	}

	for i := range n {
		ti, _ := strconv.Atoi(times[i])
		end := ti + t - 1

		if query(1, 0, maxTime-1, ti, end) < s {
			writer.WriteString("YES\n")
			update(1, 0, maxTime-1, ti, end, 1)
		} else {
			writer.WriteString("NO\n")
		}
	}
}
