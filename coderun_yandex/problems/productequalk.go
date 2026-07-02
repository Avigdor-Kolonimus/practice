package problems

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type PairProductEqualK struct {
	val int
	idx int
}

// https://coderun.yandex.ru/problem/product-equal-k
// ProductEqualK - problem 298
func ProductEqualK() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N, M and K input
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
	k, err := strconv.Atoi(strNum[2])
	if err != nil {
		panic(err)
	}

	// elements A input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum = strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	nums := make([]PairProductEqualK, n)
	for i := 0; i < n; i++ {
		ai, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		nums[i].val = ai
		nums[i].idx = i
	}

	var del []int

	for i := 2; i*i <= m; i++ {
		if m%i == 0 {
			del = append(del, i)
			if i*i != m {
				del = append(del, m/i)
			}
		}
	}

	del = append(del, m)
	sort.Ints(del)

	mp := make(map[int]int)

	for i, v := range del {
		mp[v] = i + 1
	}

	mp[1] = 0

	sort.Slice(nums, func(i, j int) bool {
		return nums[i].val < nums[j].val
	})

	var (
		a     []PairProductEqualK
		count []PairProductEqualK
	)

	for _, p := range nums {
		if p.val == 0 {
			continue
		}
		if p.val == 1 {
			count = append(count, p)
		} else {
			a = append(a, p)
		}
	}

	if m == 0 {
		ans := make([]int, 0, k)

		for _, p := range nums {
			if p.val == 0 && len(ans) < k {
				ans = append(ans, p.idx+1)
			}
		}

		for _, p := range nums {
			if p.val != 0 && len(ans) < k {
				ans = append(ans, p.idx+1)
			}
		}

		for _, x := range ans {
			writer.WriteString(strconv.Itoa(x))
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')

		return
	}

	ln := len(a)

	dp := make([][]struct {
		ok   bool
		path []int
	}, 32)

	for i := range dp {
		dp[i] = make([]struct {
			ok   bool
			path []int
		}, len(del)+1)
	}

	dp[0][0].ok = true

	// DP
	for i := ln - 1; i >= 0; i-- {
		val := a[i].val
		idx := a[i].idx

		for j := 31; j >= 1; j-- {
			for _, d := range del {
				if d%val != 0 {
					continue
				}

				from := mp[d/val]
				to := mp[d]

				if dp[j-1][from].ok && !dp[j][to].ok {
					prev := dp[j-1][from].path

					newPath := make([]int, len(prev), len(prev)+1)
					copy(newPath, prev)
					newPath = append(newPath, idx)

					dp[j][to].ok = true
					dp[j][to].path = newPath
				}
			}
		}
	}

	// find answer
	start := k - len(count)
	if start < 0 {
		start = 0
	}

	var res []int

	for kk := start; kk <= k; kk++ {
		if kk >= len(dp) {
			continue
		}
		if dp[kk][mp[m]].ok {
			res = append(res, dp[kk][mp[m]].path...)

			for i := 0; i < k-kk && i < len(count); i++ {
				res = append(res, count[i].idx)
			}
			break
		}
	}

	// output
	for i, v := range res {
		writer.WriteString(strconv.Itoa(v + 1))
		if i+1 < len(res) {
			writer.WriteByte(' ')
		}
	}
	writer.WriteByte('\n')
}
