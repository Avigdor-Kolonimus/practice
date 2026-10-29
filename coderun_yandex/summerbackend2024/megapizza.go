package summerbackend2024

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/mega-pizza
// MegaPizza - problem 18
func MegaPizza() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	tokens := []string{}
	tokenPos := 0

	nextInt := func() int {
		for tokenPos >= len(tokens) {
			line, err := reader.ReadString('\n')
			tokens = strings.Fields(strings.TrimSpace(line))
			tokenPos = 0

			if len(tokens) == 0 && err != nil {
				return 0
			}
		}

		value, _ := strconv.Atoi(strings.TrimSpace(tokens[tokenPos]))
		tokenPos++
		return value
	}

	n := nextInt()
	m := nextInt()
	k := nextInt()

	friends := make([][]int, n)
	for range k {
		a := nextInt() - 1
		b := nextInt()

		var lit int
		if b > 0 {
			lit = 2 * (b - 1)
		} else {
			lit = 2*(-b-1) + 1
		}

		friends[a] = append(friends[a], lit)
	}

	auxStart := make([]int, n)
	totalVars := m
	for i, wishes := range friends {
		if len(wishes) >= 2 {
			auxStart[i] = totalVars
			totalVars += len(wishes) - 1
		}
	}

	V := 2 * totalVars
	graph := make([][]int, V)
	reverseGraph := make([][]int, V)

	addEdge := func(from, to int) {
		graph[from] = append(graph[from], to)
		reverseGraph[to] = append(reverseGraph[to], from)
	}

	addClause := func(a, b int) {
		addEdge(a^1, b)
		addEdge(b^1, a)
	}

	for i, wishes := range friends {
		t := len(wishes)
		if t <= 1 {
			continue
		}

		base := auxStart[i]

		addClause(wishes[0], 2*base)
		for j := 1; j < t-1; j++ {
			prev := 2 * (base + j - 1)
			curr := 2 * (base + j)

			addClause(wishes[j], curr)

			addClause(prev^1, curr)

			addClause(wishes[j], prev^1)
		}

		last := 2 * (base + t - 2)
		addClause(wishes[t-1], last^1)
	}

	visited := make([]bool, V)
	nextEdge := make([]int, V)
	order := make([]int, 0, V)
	stack := make([]int, 0, V)
	for start := 0; start < V; start++ {
		if visited[start] {
			continue
		}

		visited[start] = true
		stack = append(stack, start)

		for len(stack) > 0 {
			v := stack[len(stack)-1]

			if nextEdge[v] < len(graph[v]) {
				to := graph[v][nextEdge[v]]
				nextEdge[v]++

				if !visited[to] {
					visited[to] = true
					stack = append(stack, to)
				}
			} else {
				order = append(order, v)
				stack = stack[:len(stack)-1]
			}
		}
	}

	comp := make([]int, V)
	for i := range comp {
		comp[i] = -1
	}

	compID := 0
	for i := len(order) - 1; i >= 0; i-- {
		start := order[i]
		if comp[start] != -1 {
			continue
		}

		stack = append(stack[:0], start)
		comp[start] = compID

		for len(stack) > 0 {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			for _, to := range reverseGraph[v] {
				if comp[to] == -1 {
					comp[to] = compID
					stack = append(stack, to)
				}
			}
		}

		compID++
	}

	for i := range m {
		pos := 2 * i
		neg := pos + 1

		if comp[pos] == comp[neg] {
			writer.WriteString("-1\n")
			return
		}
	}

	result := make([]int, 0, m)
	for i := range m {
		pos := 2 * i
		neg := pos + 1

		if comp[pos] > comp[neg] {
			result = append(result, i+1)
		}
	}

	writer.WriteString(strconv.Itoa(len(result)))
	writer.WriteByte('\n')

	for i, topping := range result {
		if i > 0 {
			writer.WriteByte(' ')
		}
		writer.WriteString(strconv.Itoa(topping))
	}
	writer.WriteByte('\n')
}
