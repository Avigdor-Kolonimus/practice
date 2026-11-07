package problems

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Customer struct {
	l, r  int64
	valid bool
}

// https://coderun.yandex.ru/problem/advertising
// Advertising - problem 236
func Advertising() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	customers := make([]Customer, 0, n)
	candidateSet := make(map[int64]struct{}, 2*n)

	validCount := 0

	for i := 0; i < n; i++ {
		var a, b int64
		fmt.Fscan(reader, &a, &b)

		l := a
		r := b - 5

		valid := l <= r

		customers = append(customers, Customer{
			l:     l,
			r:     r,
			valid: valid,
		})

		if valid {
			validCount++
			candidateSet[l] = struct{}{}
			candidateSet[r] = struct{}{}
		}
	}

	if validCount == 0 {
		fmt.Fprintln(writer, "0 1 6")
		return
	}

	candidates := make([]int64, 0, len(candidateSet))
	for t := range candidateSet {
		candidates = append(candidates, t)
	}

	slices.Sort(candidates)

	covered := make([]bool, n)
	starts := make([]int64, 0, n)
	endsPlus := make([]int64, 0, n)

	bestUnion := -1
	var bestT1, bestT2 int64 = 1, 6

	for _, t1 := range candidates {
		count1 := 0

		for i, c := range customers {
			covered[i] = c.valid && c.l <= t1 && t1 <= c.r

			if covered[i] {
				count1++
			}
		}

		lb := t1 + 5

		starts = starts[:0]
		endsPlus = endsPlus[:0]

		for i, c := range customers {
			if covered[i] || !c.valid {
				continue
			}

			s := c.l
			if s < lb {
				s = lb
			}

			if s <= c.r {
				starts = append(starts, s)
				endsPlus = append(endsPlus, c.r+1)
			}
		}

		count2 := 0
		t2 := lb

		if len(starts) > 0 {
			slices.Sort(starts)
			slices.Sort(endsPlus)

			i, j := 0, 0
			active := 0
			best := 0
			bestTime := starts[0]
			for i < len(starts) {
				if j >= len(endsPlus) || starts[i] < endsPlus[j] {
					active++

					if active > best {
						best = active
						bestTime = starts[i]
					}

					i++
				} else {
					active--
					j++
				}
			}

			count2 = best
			t2 = bestTime
		}

		unionCount := count1 + count2

		if unionCount > bestUnion {
			bestUnion = unionCount
			bestT1 = t1
			bestT2 = t2

			if bestUnion == validCount {
				break
			}
		}
	}

	fmt.Fprintf(writer, "%d %d %d\n", bestUnion, bestT1, bestT2)
}
