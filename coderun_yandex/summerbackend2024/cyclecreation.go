package summerbackend2024

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	maxN = 200001
)

var (
	g                = make([][]int, maxN)
	g2               = make([][]int, maxN)
	used             = make([]bool, maxN)
	used2            = make([]bool, maxN)
	depth            = make([]int, maxN)
	lowCycleCreation = make([]int, maxN)
	multi            = make(map[[2]int]bool)

	cnt int64
)

func edgeKey(a, b int) [2]int {
	if a > b {
		a, b = b, a
	}
	return [2]int{a, b}
}

func dfs(v, p int) {
	used[v] = true
	if v == p {
		depth[v] = 0
	} else {
		depth[v] = depth[p] + 1
	}
	lowCycleCreation[v] = depth[v]

	for _, u := range g[v] {
		if u == p {
			continue
		}
		if !used[u] {
			dfs(u, v)
			if lowCycleCreation[u] < lowCycleCreation[v] {
				lowCycleCreation[v] = lowCycleCreation[u]
			}
			if lowCycleCreation[u] > depth[v] && !multi[edgeKey(v, u)] {
				g2[v] = append(g2[v], u)
				g2[u] = append(g2[u], v)
			}
		} else if depth[u] < lowCycleCreation[v] {
			lowCycleCreation[v] = depth[u]
		}
	}
}

func dfs2(v int) {
	used2[v] = true
	cnt++

	for _, u := range g2[v] {
		if !used2[u] {
			dfs2(u)
		}
	}
}

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/cycle-creation
// CycleCreation - problem 16
func CycleCreation() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	line, _ := reader.ReadString('\n')
	fields := strings.Fields(line)

	n, _ := strconv.Atoi(fields[0])
	m, _ := strconv.Atoi(fields[1])

	// Edges input
	for i := 0; i < m; i++ {
		line, _ = reader.ReadString('\n')
		fields = strings.Fields(line)

		u, _ := strconv.Atoi(fields[0])
		v, _ := strconv.Atoi(fields[1])

		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	for v := 1; v <= n; v++ {
		seen := make(map[int]bool)
		for _, u := range g[v] {
			if seen[u] {
				multi[edgeKey(v, u)] = true
			}
			seen[u] = true
		}
	}

	for i := 1; i <= n; i++ {
		if !used[i] {
			dfs(i, i)
		}
	}

	var ans int64
	for i := 1; i <= n; i++ {
		if !used2[i] {
			cnt = 0
			dfs2(i)
			ans += cnt*(cnt-1)/2 - cnt + 1
		}
	}

	writer.WriteString(strconv.FormatInt(ans, 10))
	writer.WriteByte('\n')
}
