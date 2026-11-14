package summerbackend2024

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/nested-cycle
// NestedCycle - problem 22
func NestedCycle() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	typ := make([]int, n+1)
	L := make([]int, n+1)
	R := make([]int, n+1)
	parent := make([]int, n+1)
	children := make([][]int, n+1)

	maxConst := 0

	for i := 1; i <= n; i++ {
		line, _ = reader.ReadString('\n')
		fields := strings.Fields(line)

		left := fields[0]
		right := fields[1]

		lNum := left[0] >= '0' && left[0] <= '9'
		rNum := right[0] >= '0' && right[0] <= '9'

		if lNum {
			L[i], _ = strconv.Atoi(left)
			if L[i] > maxConst {
				maxConst = L[i]
			}
		} else {
			parent[i] = int(left[0]-'a') + 1
		}

		if rNum {
			R[i], _ = strconv.Atoi(right)
			if R[i] > maxConst {
				maxConst = R[i]
			}
		} else {
			parent[i] = int(right[0]-'a') + 1
		}

		if lNum && rNum {
			typ[i] = 0
		} else if lNum {
			typ[i] = 1
		} else {
			typ[i] = 2
		}

		if parent[i] != 0 {
			children[parent[i]] = append(children[parent[i]], i)
		}
	}

	answer := int64(1)

	var dfs func(v int) []int64

	dfs = func(v int) []int64 {
		G := make([]int64, maxConst+1)

		for x := 1; x <= maxConst; x++ {
			G[x] = 1
		}

		for _, ch := range children[v] {
			arr := dfs(ch)

			for x := 1; x <= maxConst; x++ {
				G[x] = G[x] * arr[x] % MOD
			}
		}

		if parent[v] == 0 {
			var total int64

			for x := L[v]; x <= R[v]; x++ {
				total = (total + G[x]) % MOD
			}

			answer = answer * total % MOD
			return nil
		}

		f := make([]int64, maxConst+1)

		if typ[v] == 1 {
			var pref int64
			lo := L[v]

			for x := 1; x <= maxConst; x++ {
				if x >= lo {
					pref = (pref + G[x]) % MOD
				}
				f[x] = pref
			}
		} else {
			var suff int64
			hi := R[v]

			for x := maxConst; x >= 1; x-- {
				if x <= hi {
					suff = (suff + G[x]) % MOD
				}
				f[x] = suff
			}
		}

		return f
	}

	for i := 1; i <= n; i++ {
		if parent[i] == 0 {
			dfs(i)
		}
	}

	writer.WriteString(strconv.FormatInt(answer, 10))
	writer.WriteByte('\n')
}
