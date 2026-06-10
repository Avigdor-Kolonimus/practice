package summerbackend2024

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const MOD int64 = 1_000_000_007

type Update struct {
	l, r int
}

type Query struct {
	l, r int
	id   int
}

func modPow(a, e int64) int64 {
	res := int64(1)

	for e > 0 {
		if e&1 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		e >>= 1
	}

	return res
}

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/balls-and-baskets
// BallsAndBaskets - problem 1
func BallsAndBaskets() {
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

	// balls input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	a := make([]int64, n)
	for i := 0; i < n; i++ {
		ai, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		a[i] = int64(ai)
	}

	// Q input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	q, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// requests input
	updates := make([]Update, 0, 400)

	// queriesByStage[s] = queries after s updates
	queriesByStage := make([][]Query, 405)

	answerCnt := 0
	stage := 0
	for i := 0; i < q; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
		if len(strNum) != 3 {
			panic("numbers count does not match 3")
		}

		typ, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		l, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(strNum[2])
		if err != nil {
			panic(err)
		}

		l--
		r--

		if typ == 0 {
			updates = append(updates, Update{l, r})
			stage++
		} else {
			queriesByStage[stage] = append(
				queriesByStage[stage],
				Query{l, r, answerCnt},
			)
			answerCnt++
		}
	}

	answers := make([]int64, answerCnt)

	cur := make([]int64, n)
	copy(cur, a)

	pref := make([]int64, n+1)
	invPref := make([]int64, n+1)
	zeroPref := make([]int, n+1)
	factor := make([]int64, n+1)

	for s := 0; s <= len(updates); s++ {

		pref[0] = 1
		zeroPref[0] = 0

		for i := 1; i <= n; i++ {
			x := cur[i-1] % MOD

			if x == 0 {
				zeroPref[i] = zeroPref[i-1] + 1
				factor[i] = 1
				pref[i] = pref[i-1]
			} else {
				zeroPref[i] = zeroPref[i-1]
				factor[i] = x
				pref[i] = pref[i-1] * x % MOD
			}
		}

		invPref[n] = modPow(pref[n], MOD-2)

		for i := n; i >= 1; i-- {
			invPref[i-1] = invPref[i] * factor[i] % MOD
		}

		for _, qu := range queriesByStage[s] {

			if zeroPref[qu.r+1]-zeroPref[qu.l] > 0 {
				answers[qu.id] = 0
				continue
			}

			res := pref[qu.r+1] * invPref[qu.l] % MOD
			answers[qu.id] = res
		}

		if s == len(updates) {
			break
		}

		up := updates[s]

		for i := up.l; i <= up.r; i++ {
			cur[i]++
		}
	}

	for _, x := range answers {
		writer.WriteString(strconv.FormatInt(x, 10))
		writer.WriteByte('\n')
	}
}
