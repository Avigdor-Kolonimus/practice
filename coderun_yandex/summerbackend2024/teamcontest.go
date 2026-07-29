package summerbackend2024

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func solve(first, second, third []int64, n int) int64 {
	best := first[1] - second[1]
	ans := int64(math.MaxInt64)

	for j := 2; j <= n-1; j++ {
		cur := first[j-1] - second[j-1]
		if cur < best {
			best = cur
		}

		cost := best + (second[j] - third[j]) + third[n]
		if cost < ans {
			ans = cost
		}
	}

	return ans
}

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/team-contest
// TeamContest - problem 7
func TeamContest() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	pref := make([][]int64, 3)

	for k := 0; k < 3; k++ {
		pref[k] = make([]int64, n+1)

		line, _ = reader.ReadString('\n')
		fields := strings.Fields(line)

		for i := 1; i <= n; i++ {
			x, _ := strconv.ParseInt(fields[i-1], 10, 64)
			pref[k][i] = pref[k][i-1] + x
		}
	}

	perms := [][3]int{
		{0, 1, 2},
		{0, 2, 1},
		{1, 0, 2},
		{1, 2, 0},
		{2, 0, 1},
		{2, 1, 0},
	}

	ans := int64(math.MaxInt64)
	for _, p := range perms {
		cur := solve(pref[p[0]], pref[p[1]], pref[p[2]], n)
		if cur < ans {
			ans = cur
		}
	}

	writer.WriteString(strconv.FormatInt(ans, 10))
	writer.WriteByte('\n')
}
