package algorithmtrainingseptember2025

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type GroupInterviewSignUp struct {
	end int
	cnt int
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/interview-sign-up
// InterviewSignUp - assignment 30
func InterviewSignUp() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// interview input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	sumA := 0
	a := make([]int, n+1)
	for i := range n {
		ai, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		sumA += ai
		a[i+1] = ai
	}

	// limit interview input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum = strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	sumB := 0
	b := make([]int, n+1)
	for i := range n {
		bi, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		sumB += bi
		b[i+1] = bi
	}

	if sumA > sumB {
		writer.WriteString(strconv.Itoa(-1))
		writer.WriteByte('\n')

		return
	}

	check := func(k int) bool {
		starts := make([][]GroupInterviewSignUp, n+2)

		for i := 1; i <= n; i++ {
			s := i - k
			if s < 1 {
				s = 1
			}

			e := i + k
			if e > n {
				e = n
			}

			starts[s] = append(starts[s], GroupInterviewSignUp{
				end: e,
				cnt: a[i],
			})
		}

		queue := make([]GroupInterviewSignUp, 0)
		head := 0

		for day := 1; day <= n; day++ {
			for _, g := range starts[day] {
				queue = append(queue, g)
			}

			capacity := b[day]

			for capacity > 0 && head < len(queue) {
				take := queue[head].cnt
				if take > capacity {
					take = capacity
				}

				queue[head].cnt -= take
				capacity -= take

				if queue[head].cnt == 0 {
					head++
				}
			}

			if head < len(queue) &&
				queue[head].cnt > 0 &&
				queue[head].end <= day {
				return false
			}
		}

		for head < len(queue) {
			if queue[head].cnt > 0 {
				return false
			}
			head++
		}

		return true
	}

	l, r := 0, n-1
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	writer.WriteString(strconv.Itoa(l))
	writer.WriteByte('\n')
}
