package problems

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func dfsPedigreeCountingLevels(height map[string]int, children map[string][]string, node string, h int) {
	height[node] = h
	for _, child := range children[node] {
		dfsPedigreeCountingLevels(height, children, child, h+1)
	}
}

// https://coderun.yandex.ru/problem/pedigree-counting-levels
// PedigreeCountingLevels - problem 243
func PedigreeCountingLevels() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	children := make(map[string][]string) // parent -> children
	hasParent := make(map[string]bool)    // child -> parent
	all := make(map[string]struct{})      // sort output by A-B
	child, parent := "", ""
	for i := 0; i < n-1; i++ {
		// child and parent input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		strNames := strings.Fields(line)
		if len(strNames) != 2 {
			panic("numbers count does not match 2")
		}

		child = strNames[0]
		parent = strNames[1]

		children[parent] = append(children[parent], child)
		hasParent[child] = true

		all[child] = struct{}{}
		all[parent] = struct{}{}
	}

	root := ""
	for name := range all {
		if !hasParent[name] {
			root = name
			break
		}
	}

	if root == "" {
		panic("root not found")
	}

	height := make(map[string]int)
	dfsPedigreeCountingLevels(height, children, root, 0)

	names := make([]string, 0, len(all))
	for name := range all {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		writer.WriteString(name + " " + strconv.Itoa(height[name]))
		writer.WriteByte('\n')
	}
}
