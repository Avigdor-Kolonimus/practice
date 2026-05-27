package wintercommon2025

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

const MOD int = 1_000_000_007

// https://coderun.yandex.ru/selections/2025-winter-common/problems/equipment
// Equipment - problem 3
func Equipment() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// M and N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	m, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	n, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// input equipment
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum = strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	w := make([]int, n)
	for i := 0; i < n; i++ {
		wi, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		w[i] = wi
	}

	var sum int
	for _, x := range w {
		sum += x
	}

	// Total deficit
	D := sum - m

	sort.Slice(w, func(i, j int) bool {
		return w[i] < w[j]
	})

	var ans int

	left := n

	for i := 0; i < n; i++ {

		// Target equal deficit
		k := D / left

		if w[i] <= k {
			// Fully remove this group
			ans = (ans + w[i]%MOD*w[i]%MOD) % MOD
			D -= w[i]
			left--
		} else {
			break
		}
	}

	if left > 0 {
		base := D / left
		extra := D % left

		cntBig := extra
		cntSmall := left - extra

		ans = (ans + cntSmall%MOD*(base%MOD)*(base%MOD)) % MOD

		big := base + 1

		ans = (ans + cntBig%MOD*(big%MOD)*(big%MOD)) % MOD
	}

	result := ans % MOD

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
