package backend

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	n, m, k int

	cost     []float64
	allowed  []int
	discount []float64

	bestCost float64
	bestMask int
)

func calc(mask int) float64 {
	total := 0.0

	for i := 0; i < n; i++ {
		price := cost[i]

		for j := 0; j < m; j++ {
			if mask&(1<<j) != 0 && allowed[i]&(1<<j) != 0 {
				price *= (100 - discount[j]) / 100
			}
		}

		total += price
	}

	return total
}

func dfs(pos, count, mask int) {
	if count > k {
		return
	}

	if pos == m {
		cur := calc(mask)

		if cur < bestCost {
			bestCost = cur
			bestMask = mask
		}

		return
	}

	dfs(pos+1, count, mask)

	dfs(pos+1, count+1, mask|(1<<pos))
}

// https://coderun.yandex.ru/selections/backend/problems/discount-coupons
// DiscountCoupons - problem 46
func DiscountCoupons() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input n, m, k
	fmt.Fscan(reader, &n, &m, &k)

	//  cost of items
	cost = make([]float64, n)
	for i := range n {
		fmt.Fscan(reader, &cost[i])
	}

	allowed = make([]int, n)

	for i := range n {
		// C input
		var c int
		fmt.Fscan(reader, &c)

		// mask of allowed coupons
		mask := 0
		for j := 0; j < c; j++ {
			// X input
			var x int
			fmt.Fscan(reader, &x)

			mask |= 1 << (x - 1)
		}

		allowed[i] = mask
	}

	discount = make([]float64, m)

	for i := range m {
		fmt.Fscan(reader, &discount[i])
	}

	bestCost = math.MaxFloat64

	dfs(0, 0, 0)

	ans := make([]int, 0)

	for i := range m {
		if bestMask&(1<<i) != 0 {
			ans = append(ans, i+1)
		}
	}

	writer.WriteString(fmt.Sprintf("%d\n", len(ans)))

	if len(ans) > 0 {
		for i, v := range ans {
			if i > 0 {
				writer.WriteByte(' ')
			}

			writer.WriteString(fmt.Sprintf("%d", v))
		}

		writer.WriteByte('\n')
	}
}
