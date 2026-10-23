package mobiledev

import (
	"bufio"
	"fmt"
	"os"
)

// https://coderun.yandex.ru/selections/2024-summer-mobile-dev/problems/dead-battery-2
// DeadBattery2 - problem 28
func DeadBattery2() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20)
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	sum := int64(0)

	for i := range a {
		fmt.Fscan(in, &a[i])
		sum += int64(a[i])
	}
	if sum >= 40000 {
		fmt.Fprintln(out, 0)

		return
	}

	total := int(sum)
	limit := total / 2

	dp := make([]uint64, limit/64+1)
	dp[0] = 1
	for _, x := range a {
		if x > limit {
			continue
		}

		wordShift := x / 64
		bitShift := uint(x % 64)

		for i := len(dp) - 1; i >= wordShift; i-- {
			v := dp[i-wordShift] << bitShift

			if bitShift != 0 && i-wordShift > 0 {
				v |= dp[i-wordShift-1] >> (64 - bitShift)
			}

			dp[i] |= v
		}
	}

	ans := int64(1 << 62)
	for s := 1; s <= limit; s++ {
		if dp[s/64]&(uint64(1)<<uint(s%64)) == 0 {
			continue
		}

		t1 := int64(10000 / s)
		t2 := int64(10000 / (total - s))

		if t1+t2 < ans {
			ans = t1 + t2
		}
	}

	fmt.Fprintln(out, ans)
}
