package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func possible(s1, s2 string) bool {
	var to [26]int
	var from [26]int

	for i := range to {
		to[i] = -1
		from[i] = -1
	}

	for i := range s1 {
		a := int(s1[i] - 'a')
		b := int(s2[i] - 'a')

		if to[a] != -1 && to[a] != b {
			return false
		}

		if from[b] != -1 && from[b] != a {
			return false
		}

		to[a] = b
		from[b] = a
	}

	return true
}

// https://coderun.yandex.ru/problem/substitutions-playing
// SubstitutionsPlaying - problem 464
func SubstitutionsPlaying() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// t input
	line, err := in.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	t, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	for ; t > 0; t-- {
		// s1 and s2 lines
		s1, err := in.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		s1 = strings.TrimRight(s1, "\r\n")

		s2, err := in.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		s2 = strings.TrimRight(s2, "\r\n")

		if possible(s1, s2) {
			out.WriteString("YES")
		} else {
			out.WriteString("NO")
		}
		out.WriteByte('\n')
	}
}
