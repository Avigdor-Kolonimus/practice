package algorithmtrainingseptember2025

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/ancestor
// Ancestor - assignment 25
func Ancestor() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// tree input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	tree := make([][]int, n+1)
	roots := []int{}
	for i := 1; i <= n; i++ {
		p, err := strconv.Atoi(strNum[i-1])
		if err != nil {
			panic(err)
		}

		if p == 0 {
			roots = append(roots, i)
		} else {
			tree[p] = append(tree[p], i)
		}
	}

	tin := make([]int, n+1)
	tout := make([]int, n+1)
	timer := 0

	var dfsLocal func(int)
	dfsLocal = func(v int) {
		timer++
		tin[v] = timer

		for _, to := range tree[v] {
			dfsLocal(to)
		}

		timer++
		tout[v] = timer
	}

	for _, root := range roots {
		dfsLocal(root)
	}

	isAncestor := func(a, b int) bool {
		return tin[a] <= tin[b] && tout[a] >= tout[b]
	}

	// M input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	m, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	for i := 0; i < m; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		a, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		if isAncestor(a, b) {
			writer.WriteByte('1')
		} else {
			writer.WriteByte('0')
		}
		writer.WriteByte('\n')
	}
}
