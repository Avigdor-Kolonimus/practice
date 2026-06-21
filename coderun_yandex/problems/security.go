package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type GuardSecurity struct {
	a int
	b int
}

// https://coderun.yandex.ru/problem/security
// Security - problem 83
func Security() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// K input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	k, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	for ; k > 0; k-- {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum := strings.Fields(line)

		n, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}

		guards := make([]GuardSecurity, n)
		diff := make([]int, 10002)
		for g := 0; g < n; g++ {
			a, err := strconv.Atoi(strNum[1+2*g])
			if err != nil {
				panic(err)
			}
			b, err := strconv.Atoi(strNum[2+2*g])
			if err != nil {
				panic(err)
			}

			guards[g] = GuardSecurity{a, b}

			if a < b {
				diff[a+1]++
				diff[b+1]--
			}
		}

		cnt := make([]int, 10001)

		cur := 0
		ok := true
		for t := 1; t <= 10000; t++ {
			cur += diff[t]
			cnt[t] = cur

			if cnt[t] == 0 {
				ok = false
			}
		}

		if !ok {
			writer.WriteString("Wrong Answer")
			writer.WriteByte('\n')

			continue
		}

		alonePref := make([]int, 10001)

		for t := 1; t <= 10000; t++ {
			alonePref[t] = alonePref[t-1]
			if cnt[t] == 1 {
				alonePref[t]++
			}
		}

		for _, g := range guards {
			unique := alonePref[g.b] - alonePref[g.a]

			if unique == 0 {
				ok = false
				break
			}
		}

		if ok {
			writer.WriteString("Accepted")
		} else {
			writer.WriteString("Wrong Answer")
		}
		writer.WriteByte('\n')
	}
}
