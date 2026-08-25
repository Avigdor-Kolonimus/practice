package ababin

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type NodeSandwiches struct {
	parent int
	value  int64
	depth  int
}

type KeySandwiches struct {
	parent int
	value  int64
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// https://coderun.yandex.ru/selections/ababin/problems/sandwiches
// Sandwiches - problem 5
func Sandwiches() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// K and N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	K, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	N, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// Node 0 = empty sandwich.
	nodes := make([]NodeSandwiches, 1, N+1)

	// For every (parent, ingredient) we store the unique node.
	next := make(map[KeySandwiches]int, N)

	// top[i] is the current top node of sandwich i.
	top := make([]int, K)

	// Every non-empty sandwich that appeared.
	seen := make(map[int]struct{}, N)

	// Number of different sandwiches.
	answer := 0

	// Creates or returns a sandwich obtained by putting value
	// on top of parent.
	push := func(parent int, value int64) int {
		key := KeySandwiches{
			parent: parent,
			value:  value,
		}

		if id, ok := next[key]; ok {
			return id
		}

		id := len(nodes)

		nodes = append(nodes, NodeSandwiches{
			parent: parent,
			value:  value,
			depth:  nodes[parent].depth + 1,
		})

		next[key] = id

		return id
	}

	// Adds sandwich to the set of all sandwiches that appeared.
	addSeen := func(id int) {
		if id == 0 {
			return
		}

		if _, ok := seen[id]; ok {
			return
		}

		seen[id] = struct{}{}
		answer++
	}

	for q := 0; q < N; q++ {
		// typ input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)

		typ, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}

		switch typ {
		case 1:
			i, err := strconv.Atoi(strNum[1])
			if err != nil {
				panic(err)
			}
			x, err := strconv.Atoi(strNum[2])
			if err != nil {
				panic(err)
			}

			i--

			top[i] = push(top[i], int64(x))
			addSeen(top[i])

		case 2:
			i, err := strconv.Atoi(strNum[1])
			if err != nil {
				panic(err)
			}

			i--

			top[i] = nodes[top[i]].parent

		case 3:
			i, err := strconv.Atoi(strNum[1])
			if err != nil {
				panic(err)
			}
			j, err := strconv.Atoi(strNum[2])
			if err != nil {
				panic(err)
			}

			i--
			j--

			di := nodes[top[i]].depth
			dj := nodes[top[j]].depth

			// Already balanced.
			if abs(di-dj) <= 1 {
				continue
			}

			if di > dj {
				// Move from i to j.
				diff := (di - dj) / 2

				for step := 0; step < diff; step++ {
					x := nodes[top[i]].value

					top[i] = nodes[top[i]].parent
					top[j] = push(top[j], x)

					// Both intermediate sandwiches count.
					addSeen(top[i])
					addSeen(top[j])
				}
			} else {
				// Move from j to i.
				diff := (dj - di) / 2

				for step := 0; step < diff; step++ {
					x := nodes[top[j]].value

					top[j] = nodes[top[j]].parent
					top[i] = push(top[i], x)

					addSeen(top[i])
					addSeen(top[j])
				}
			}
		}
	}

	writer.WriteString(strconv.Itoa(answer))
	writer.WriteByte('\n')
}
