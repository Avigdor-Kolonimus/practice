package ababin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	to int
	id int
}

// https://coderun.yandex.ru/selections/ababin/problems/dark-starts-and-win
// DarkStartsAndWin - problem 3
func DarkStartsAndWin() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	board := make([]byte, n*n)
	whites := make([]int, 0)
	blacks := make([]int, 0)
	isBlack := make([]bool, n*n)
	for r := range n {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)

		for c := range n {
			cell := r*n + c
			board[cell] = line[c]

			switch line[c] {
			case 'W':
				whites = append(whites, cell)
			case 'B':
				blacks = append(blacks, cell)
				isBlack[cell] = true
			}
		}
	}

	if len(whites) == 0 {
		if len(blacks) == 0 {
			fmt.Fprintln(writer, -1)
			return
		}

		cell := blacks[0]
		fmt.Fprintln(writer, 1)
		fmt.Fprintln(writer, cell/n+1, cell%n+1)
		return
	}

	firstWhiteType := (whites[0] / n) & 1

	for _, w := range whites {
		if (w/n)&1 != firstWhiteType {
			fmt.Fprintln(writer, -1)
			return
		}
	}

	inBounds := func(r, c int) bool {
		return r >= 0 && r < n && c >= 0 && c < n
	}

	compParity := func(r, c int) int {
		return ((r >> 1) + (c >> 1)) & 1
	}

	tryBuild := func(moverType, parityP int) ([]int, bool) {
		if firstWhiteType != 1-moverType {
			return nil, false
		}

		idOfCell := make([]int, n*n)
		for i := range idOfCell {
			idOfCell[i] = -1
		}

		vertexCells := make([]int, 0, len(whites)*2)
		deg := make([]int, 0, len(whites)*2)

		getID := func(cell int) int {
			if idOfCell[cell] >= 0 {
				return idOfCell[cell]
			}

			id := len(vertexCells)
			idOfCell[cell] = id

			vertexCells = append(vertexCells, cell)
			deg = append(deg, 0)

			return id
		}

		e := len(whites)
		u := make([]int, 0, e)
		v := make([]int, 0, e)
		for _, w := range whites {
			r, c := w/n, w%n

			uCell, vCell := -1, -1

			r1, c1 := r-1, c-1
			r2, c2 := r+1, c+1

			if inBounds(r1, c1) && inBounds(r2, c2) {
				if (r1&1) == moverType &&
					compParity(r1, c1) == parityP {

					uCell = r1*n + c1
					vCell = r2*n + c2
				}
			}

			if uCell == -1 {
				r3, c3 := r-1, c+1
				r4, c4 := r+1, c-1

				if inBounds(r3, c3) && inBounds(r4, c4) {
					if (r3&1) == moverType &&
						compParity(r3, c3) == parityP {

						uCell = r3*n + c3
						vCell = r4*n + c4
					}
				}
			}

			if uCell == -1 {
				return nil, false
			}

			uid := getID(uCell)
			vid := getID(vCell)

			u = append(u, uid)
			v = append(v, vid)

			deg[uid]++
			deg[vid]++
		}

		start := -1
		for i, cell := range vertexCells {
			if deg[i] == 0 {
				continue
			}

			if isBlack[cell] {
				if start == -1 {
					start = i
				} else {
					return nil, false
				}
			}
		}

		if start == -1 {
			return nil, false
		}

		odd := 0
		for _, d := range deg {
			if d > 0 && d%2 == 1 {
				odd++
			}
		}

		if odd != 0 && odd != 2 {
			return nil, false
		}

		if odd == 2 && deg[start]%2 == 0 {
			return nil, false
		}

		adj := make([][]Edge, len(vertexCells))
		for i := 0; i < e; i++ {
			a, b := u[i], v[i]

			adj[a] = append(adj[a], Edge{to: b, id: i})
			adj[b] = append(adj[b], Edge{to: a, id: i})
		}

		used := make([]bool, e)
		usedCnt := 0
		stack := make([]int, 0, e+1)
		result := make([]int, 0, e+1)
		it := make([]int, len(vertexCells))
		stack = append(stack, start)
		for len(stack) > 0 {
			v := stack[len(stack)-1]

			for it[v] < len(adj[v]) &&
				used[adj[v][it[v]].id] {
				it[v]++
			}

			if it[v] == len(adj[v]) {
				result = append(result, v)
				stack = stack[:len(stack)-1]
				continue
			}

			edge := adj[v][it[v]]
			it[v]++

			if used[edge.id] {
				continue
			}

			used[edge.id] = true
			usedCnt++

			stack = append(stack, edge.to)
		}

		if usedCnt != e || len(result) != e+1 {
			return nil, false
		}

		path := make([]int, len(result))
		for i := range result {
			path[i] = vertexCells[result[len(result)-1-i]]
		}

		return path, true
	}

	for moverType := 0; moverType <= 1; moverType++ {
		for p := 0; p <= 1; p++ {
			path, ok := tryBuild(moverType, p)

			if !ok {
				continue
			}

			fmt.Fprintln(writer, len(path))

			for _, cell := range path {
				fmt.Fprintln(writer, cell/n+1, cell%n+1)
			}

			return
		}
	}

	fmt.Fprintln(writer, -1)
}
