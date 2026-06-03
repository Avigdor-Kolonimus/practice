package summerbackend2024

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	MODGCD int64 = 1000000000
)

func factorize(x int, cnt map[int]int) {
	d := 2

	for d*d <= x {
		for x%d == 0 {
			cnt[d]++
			x /= d
		}
		d++
	}

	if x > 1 {
		cnt[x]++
	}
}

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/gcd
// GCD - problem 9
func GCD() {
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

	// A number input input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	cntA := make(map[int]int)
	for i := 0; i < n; i++ {
		x, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		factorize(x, cntA)
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

	// B number input input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum = strings.Fields(line)
	if len(strNum) != m {
		panic("numbers count does not match m")
	}

	cntB := make(map[int]int)
	for i := 0; i < m; i++ {
		x, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		factorize(x, cntB)
	}

	var ans int64 = 1
	overflow := false
	for p, cA := range cntA {
		cB, ok := cntB[p]
		if !ok {
			continue
		}

		power := cA
		if cB < power {
			power = cB
		}

		for i := 0; i < power; i++ {
			ans *= int64(p)

			if ans >= MODGCD {
				overflow = true
				ans %= MODGCD
			}
		}
	}

	if overflow {
		writer.WriteString(fmt.Sprintf("%09d", ans))
	} else {
		writer.WriteString(strconv.FormatInt(ans, 10))
	}
	writer.WriteByte('\n')
}
