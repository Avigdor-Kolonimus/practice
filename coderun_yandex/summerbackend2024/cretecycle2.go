package summerbackend2024

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	MODCreteCycle2 = 998244353
)

type EdgeCreteCycle2 struct {
	to int
	id int
}

var (
	n, m int

	graph [][]EdgeCreteCycle2
	edges [][2]int

	tin      []int
	low      []int
	timer    int
	isBridge []bool
)

func dfsCreteCycle2(v, parentEdge int) {
	tin[v] = timer
	low[v] = timer
	timer++

	for _, e := range graph[v] {
		to := e.to
		id := e.id

		if id == parentEdge {
			continue
		}

		if tin[to] == -1 {
			dfsCreteCycle2(to, id)

			if low[to] < low[v] {
				low[v] = low[to]
			}

			if low[to] > tin[v] {
				isBridge[id] = true
			}
		} else {
			if tin[to] < low[v] {
				low[v] = tin[to]
			}
		}
	}
}

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/crete-cycle-2
// CreteCycle2 - problem 34
func CreteCycle2() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// stations input
	graph = make([][]EdgeCreteCycle2, n)
	edges = make([][2]int, m)
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

		u, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		v, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		u--
		v--

		edges[i] = [2]int{u, v}

		graph[u] = append(graph[u], EdgeCreteCycle2{v, i})
		graph[v] = append(graph[v], EdgeCreteCycle2{u, i})
	}

	tin = make([]int, n)
	low = make([]int, n)

	for i := 0; i < n; i++ {
		tin[i] = -1
	}

	isBridge = make([]bool, m)

	// Graph is connected according to the statement.
	dfsCreteCycle2(0, -1)

	// Degree in the bridge forest.
	deg := make([]int, n)

	for i, e := range edges {
		if isBridge[i] {
			u := e[0]
			v := e[1]

			deg[u]++
			deg[v]++
		}
	}

	maxDeg := 0
	for _, d := range deg {
		if d > maxDeg {
			maxDeg = d
		}
	}

	// I[k] = number of involutions:
	// I[k] = I[k-1] + (k-1)*I[k-2]
	I := make([]int64, maxDeg+1)

	I[0] = 1
	if maxDeg >= 1 {
		I[1] = 1
	}

	for k := 2; k <= maxDeg; k++ {
		I[k] = (I[k-1] + int64(k-1)*I[k-2]) % MODCreteCycle2
	}

	var ans int64 = 1

	for _, d := range deg {
		ans = ans * I[d] % MODCreteCycle2
	}

	writer.WriteString(strconv.FormatInt(ans, 10))
	writer.WriteByte('\n')
}
