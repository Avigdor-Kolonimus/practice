package summerbackend2024

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type Group struct {
	ch    byte
	count int
}

func encode(s string) []Group {
	if len(s) == 0 {
		return nil
	}

	res := make([]Group, 0)

	cur := s[0]
	cnt := 1

	for i := 1; i < len(s); i++ {
		if s[i] == cur {
			cnt++
		} else {
			res = append(res, Group{cur, cnt})
			cur = s[i]
			cnt = 1
		}
	}

	res = append(res, Group{cur, cnt})
	return res
}

func median(a, b, c int) int {
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}

	return b
}

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/crystals
// Crystals - problem 17
func Crystals() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// str1 input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	str1 := strings.TrimRight(line, "\r\n")

	// str2 input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	str2 := strings.TrimRight(line, "\r\n")

	// str3 input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	str3 := strings.TrimRight(line, "\r\n")

	g1 := encode(str1)
	g2 := encode(str2)
	g3 := encode(str3)

	if len(g1) != len(g2) || len(g1) != len(g3) {
		writer.WriteString("IMPOSSIBLE\n")

		return
	}

	var ans strings.Builder

	for i := range g1 {
		if g1[i].ch != g2[i].ch || g1[i].ch != g3[i].ch {
			writer.WriteString("IMPOSSIBLE\n")

			return
		}

		k := median(g1[i].count, g2[i].count, g3[i].count)

		for j := 0; j < k; j++ {
			ans.WriteByte(g1[i].ch)
		}
	}

	writer.WriteString(ans.String())
	writer.WriteByte('\n')
}
