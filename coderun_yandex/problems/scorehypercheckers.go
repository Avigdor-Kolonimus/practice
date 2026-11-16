package problems

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// https://coderun.yandex.ru/problem/score-hypercheckers
// ScoreHypercheckers - problem 215
func ScoreHypercheckers() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var k int64
	fmt.Fscan(in, &n, &k)

	a := make([]int64, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	v := make([]int64, 0, n)
	cnt := make([]int64, 0, n)
	for _, x := range a {
		if len(v) == 0 || v[len(v)-1] != x {
			v = append(v, x)
			cnt = append(cnt, 1)
		} else {
			cnt[len(cnt)-1]++
		}
	}

	m := len(v)
	pref := make([]int64, m+1)
	for i := range m {
		pref[i+1] = pref[i]
		if cnt[i] >= 2 {
			pref[i+1]++
		}
	}

	var ans int64
	r := 0

	for i := range m {
		if r < i {
			r = i
		}

		for r+1 < m && v[r+1] <= v[i]*k {
			r++
		}

		d := int64(r - i)

		if cnt[i] >= 3 {
			ans++
		}

		ans += 3 * (pref[r+1] - pref[i+1])

		if cnt[i] >= 2 {
			ans += 3 * d
		}

		ans += 6 * d * (d - 1) / 2
	}

	fmt.Fprintln(out, ans)
}
