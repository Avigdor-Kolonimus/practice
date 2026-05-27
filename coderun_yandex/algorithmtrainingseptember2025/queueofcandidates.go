package algorithmtrainingseptember2025

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/queue-of-candidates
// QueueOfCandidates - assignment 33
func QueueOfCandidates() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and X input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	parameters := strings.Fields(line)
	if len(parameters) != 2 {
		panic("input does not match 2")
	}

	n, err := strconv.Atoi(parameters[0])
	if err != nil {
		panic(err)
	}
	x, err := strconv.Atoi(parameters[1])
	if err != nil {
		panic(err)
	}

	// order input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	parameters = strings.Fields(line)
	if len(parameters) != n {
		panic("input does not match n")
	}

	good := make([]int, n)
	pref := []int{0}
	for i := 0; i < n; i++ {
		a, err := strconv.Atoi(parameters[i])
		if err != nil {
			panic(err)
		}

		val := 0
		if a >= x {
			val = 1
		}

		good[i] = val
		pref = append(pref, pref[len(pref)-1]+val)
	}

	// M input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	m, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// event input
	head := 0
	for i := 0; i < m; i++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}

		line = strings.TrimSpace(line)
		parameters := strings.Fields(line)

		t, err := strconv.Atoi(parameters[0])
		if err != nil {
			panic(err)
		}

		switch t {

		case 1:
			a, err := strconv.Atoi(parameters[1])
			if err != nil {
				panic(err)
			}

			val := 0
			if a >= x {
				val = 1
			}

			good = append(good, val)
			pref = append(pref, pref[len(pref)-1]+val)

		case 2:
			head++

		case 3:
			k, err := strconv.Atoi(parameters[1])
			if err != nil {
				panic(err)
			}

			l := head
			r := head + k

			ans := pref[r] - pref[l]

			writer.WriteString(strconv.Itoa(ans))
			writer.WriteByte('\n')
		}
	}
}
