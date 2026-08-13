package problems

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type NodeOS struct {
	children  map[string]int
	blacklist bool
	files     map[string]int
	total     map[string]int
}

var tree []NodeOS

func newNode() NodeOS {
	return NodeOS{
		children: make(map[string]int),
		files:    make(map[string]int),
		total:    make(map[string]int),
	}
}

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func getDir(path string) int {
	if path == "/" {
		return 0
	}

	parts := strings.Split(strings.Trim(path, "/"), "/")

	cur := 0

	for _, part := range parts {
		next, ok := tree[cur].children[part]

		if !ok {
			next = len(tree)
			tree = append(tree, newNode())
			tree[cur].children[part] = next
		}

		cur = next
	}

	return cur
}

func isBlacklisted(path string) (int, bool) {
	if path == "/" {
		return 0, tree[0].blacklist
	}

	parts := strings.Split(strings.Trim(path, "/"), "/")

	cur := 0
	removed := tree[0].blacklist

	for _, part := range parts {
		next, ok := tree[cur].children[part]

		if !ok {
			next = len(tree)
			tree = append(tree, newNode())
			tree[cur].children[part] = next
		}

		cur = next

		if tree[cur].blacklist {
			removed = true
		}
	}

	return cur, removed
}

func dfsOS(v int) {
	for ext, count := range tree[v].files {
		tree[v].total[ext] += count
	}

	for _, child := range tree[v].children {
		dfsOS(child)

		for ext, count := range tree[child].total {
			tree[v].total[ext] += count
		}
	}
}

// https://coderun.yandex.ru/problem/ivan-and-opensource
// IvanAndOpensource - problem 316
func IvanAndOpensource() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	tree = append(tree, newNode())

	// Blacklist input
	line := readLine(reader)
	n, _ := strconv.Atoi(line)

	blacklist := make([]string, n)
	for i := range n {
		blacklist[i] = readLine(reader)

		id := getDir(blacklist[i])
		tree[id].blacklist = true
	}

	// File's name input
	line = readLine(reader)
	m, _ := strconv.Atoi(line)

	files := make([]string, m)
	for i := range m {
		files[i] = readLine(reader)
	}

	for _, path := range files {
		pos := strings.LastIndex(path, "/")
		dir := path[:pos+1]

		getDir(dir)
	}

	for _, path := range files {
		pos := strings.LastIndex(path, "/")

		dir := path[:pos+1]
		name := path[pos+1:]

		id, removed := isBlacklisted(dir)

		if !removed {
			continue
		}

		dot := strings.LastIndex(name, ".")
		ext := name[dot:]

		tree[id].files[ext]++
	}

	line = readLine(reader)
	q, _ := strconv.Atoi(line)
	queries := make([]int, q)
	for i := range q {
		path := readLine(reader)
		queries[i] = getDir(path)
	}

	dfsOS(0)

	for _, id := range queries {
		stats := tree[id].total

		writer.WriteString(strconv.Itoa(len(stats)))
		writer.WriteByte('\n')

		for ext, count := range stats {
			writer.WriteString(ext)
			writer.WriteString(": ")
			writer.WriteString(strconv.Itoa(count))
			writer.WriteByte('\n')
		}
	}
}
