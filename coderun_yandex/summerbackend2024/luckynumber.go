package summerbackend2024

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func minNS(n, s int) []int {
	if n*9 < s {
		return nil
	}

	a := make([]int, n)
	i := n - 1

	for i >= 0 && s > 0 {
		if s <= 9 {
			a[i] = s
		} else {
			a[i] = 9
		}
		s -= a[i]
		i--
	}

	return a
}

func nextS(a []int, s int) []int {
	n := len(a)
	s1 := 0
	for _, v := range a {
		s1 += v
	}

	if s > s1 && s-s1+a[n-1] <= 9 {
		a[n-1] += s - s1
		return a
	}

	for i := n - 1; i >= 0; i-- {
		s1 -= a[i]

		if a[i] == 9 {
			continue
		}

		a[i]++

		s2 := s - s1 - a[i]

		if s2 < 0 || 9*(n-i) < s2 {
			continue
		}

		j := n - 1
		for j > i {
			if s2 <= 9 {
				a[j] = s2
			} else {
				a[j] = 9
			}
			s2 -= a[j]
			j--
		}

		a[i] += s2
		if a[i] <= 9 {
			return a
		}
	}

	return nil
}

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/lucky-number
// LuckyLumber - problem 24
func LuckyLumber() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// X input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	x := strings.TrimRight(line, "\r\n")

	n := len(x) / 2

	s := make([]int, len(x))
	for i := 0; i < len(x); i++ {
		s[i] = int(x[i] - '0')
	}

	s1 := 0
	for i := 0; i < n; i++ {
		s1 += s[i]
	}

	ans := nextS(append([]int(nil), s[n:]...), s1)
	if ans != nil {
		for i := 0; i < n; i++ {
			s[n+i] = ans[i]
		}
		goto out
	}

	if n*9 != s1 {
		for i := n - 1; i >= 0; i-- {
			if s[i] == 9 {
				s[i] = 0
			} else {
				s[i]++
				break
			}
		}

		s1 = 0
		for i := 0; i < n; i++ {
			s1 += s[i]
		}

		ans = minNS(n, s1)
		if ans != nil {
			for i := 0; i < n; i++ {
				s[n+i] = ans[i]
			}
			goto out
		}
	}

	for i := 0; i < len(s); i++ {
		s[i] = 0
	}
	s[n-1] = 1
	s[len(s)-1] = 1

out:
	for _, v := range s {
		writer.WriteString(strconv.Itoa(v))
	}
	writer.WriteByte('\n')
}
