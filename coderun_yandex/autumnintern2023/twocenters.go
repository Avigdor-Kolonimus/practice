package autumnintern2023

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	to int
	w  int
}

type Pair struct {
	u, v int
}

var (
	g    [][]Edge
	dist [][]int
)

func dfs(v, p, root int, d int) {
	dist[root][v] = d
	for _, e := range g[v] {
		if e.to != p {
			dfs(e.to, v, root, d+e.w)
		}
	}
}

// https://coderun.yandex.ru/selections/autumn-intern-2023/problems/two-centers
// TwoCenters - problem 5
func TwoCenters() {
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

	// edge input
	g = make([][]Edge, n)
	edges := make([]Pair, 0, n-1)

	for i := 0; i < n-1; i++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		strNum := strings.Fields(line)
		if len(strNum) != 3 {
			panic("numbers count does not match 3")
		}

		a, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}
		c, err := strconv.Atoi(strNum[2])
		if err != nil {
			panic(err)
		}

		a--
		b--

		g[a] = append(g[a], Edge{b, c})
		g[b] = append(g[b], Edge{a, c})
		edges = append(edges, Pair{a, b})
	}

	dist = make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		dfs(i, -1, i, 0)
	}

	ans := math.MaxInt
	for _, e := range edges {
		var cur int
		for w := 0; w < n; w++ {
			x := dist[e.u][w]
			if dist[e.v][w] < x {
				x = dist[e.v][w]
			}
			if x > cur {
				cur = x
			}
		}
		if cur < ans {
			ans = cur
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
