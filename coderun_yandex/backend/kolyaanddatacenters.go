package backend

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/backend/problems/kolya-and-data-centers
// KolyaAndDataCenters - problem 35
func KolyaAndDataCenters() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N, M and Q input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 3 {
		panic("numbers count does not match 3")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	q, err := strconv.Atoi(strNum[2])
	if err != nil {
		panic(err)
	}

	R := make([]int64, n+1)
	A := make([]int64, n+1)
	version := make([]int, n+1)
	for i := 1; i <= n; i++ {
		A[i] = int64(m)
	}

	// key = (dc-1)*m + server
	disabled := make(map[int]int)

	for ; q > 0; q-- {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		cmd := strings.Fields(line)

		switch cmd[0] {

		case "RESET":
			dc, err := strconv.Atoi(cmd[1])
			if err != nil {
				panic(err)
			}

			R[dc]++
			A[dc] = int64(m)
			version[dc]++

		case "DISABLE":
			dc, err := strconv.Atoi(cmd[1])
			if err != nil {
				panic(err)
			}
			srv, err := strconv.Atoi(cmd[2])
			if err != nil {
				panic(err)
			}

			key := (dc-1)*m + srv

			if disabled[key] != version[dc] {
				disabled[key] = version[dc]
				A[dc]--
			}

		case "GETMAX":
			ans := 1
			best := R[1] * A[1]

			for i := 2; i <= n; i++ {
				cur := R[i] * A[i]

				if cur > best {
					best = cur
					ans = i
				}
			}
			writer.WriteString(strconv.Itoa(ans))
			writer.WriteByte('\n')

		case "GETMIN":
			ans := 1
			best := R[1] * A[1]

			for i := 2; i <= n; i++ {
				cur := R[i] * A[i]

				if cur < best {
					best = cur
					ans = i
				}
			}

			writer.WriteString(strconv.Itoa(ans))
			writer.WriteByte('\n')
		}
	}
}
