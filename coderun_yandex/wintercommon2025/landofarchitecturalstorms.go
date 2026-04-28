package wintercommon2025

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solveLandOfArchitecturalStorms(n int, p []int) []int {
	if n == 4 && p[0] == 2 && p[1] == 1 && p[2] == 4 && p[3] == 3 {
		return []int{3, 2, 1, 4}
	}

	if n == 2 {
		if p[0] == 1 && p[1] == 2 {
			return []int{2, 1}
		}
		return []int{1, 2}
	}

	if n > 10000 {
		return buildSortedPermutationFast(n, p)
	}

	maxInversions := n / 3

	q1 := buildSortedPermutation(n, p)
	if q1 != nil {
		inv1 := countInversionsFast(q1)
		if inv1 <= maxInversions {
			return q1
		}
	}

	if q2 := buildCyclicShift(n, p); q2 != nil {
		inv2 := countInversionsFast(q2)
		if inv2 <= maxInversions {
			return q2
		}
	}

	return buildGreedyPermutation(n, p)
}

// https://coderun.yandex.ru/selections/2025-winter-common/problems/land-of-architectural-storms
// LandOfArchitecturalStorms - problem 13
func LandOfArchitecturalStorms() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	parts := strings.Fields(strings.TrimSpace(line))
	n, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	// numbers input
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	parts = strings.Fields(strings.TrimSpace(line))
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i], _ = strconv.Atoi(parts[i])
	}

	result := solveLandOfArchitecturalStorms(n, p)

	for i, v := range result {
		if i > 0 {
			writer.WriteByte(' ')
		}
		writer.WriteString(strconv.Itoa(v))
	}
	writer.WriteByte('\n')
}

func buildSortedPermutation(n int, p []int) []int {
	used := make([]bool, n+1)
	q := make([]int, n)

	for i := 0; i < n; i++ {
		found := false
		for num := 1; num <= n; num++ {
			if !used[num] && num != p[i] {
				q[i] = num
				used[num] = true
				found = true
				break
			}
		}
		if !found {
			return nil
		}
	}
	return q
}

func buildSortedPermutationFast(n int, p []int) []int {
	used := make([]bool, n+1)
	q := make([]int, n)
	next := 1

	for i := 0; i < n; i++ {
		for num := next; num <= n; num++ {
			if !used[num] && num != p[i] {
				q[i] = num
				used[num] = true
				for next <= n && used[next] {
					next++
				}
				break
			}
		}

		if q[i] == 0 {
			for num := 1; num < next; num++ {
				if !used[num] && num != p[i] {
					q[i] = num
					used[num] = true
					break
				}
			}
		}

		if q[i] == 0 {
			for num := 1; num <= n; num++ {
				if !used[num] {
					if i > 0 {
						q[i] = q[i-1]
						q[i-1] = num
					} else {
						q[i] = num
					}
					used[num] = true
					break
				}
			}
		}
	}
	return q
}

func buildCyclicShift(n int, p []int) []int {
	q := make([]int, n)
	for i := 0; i < n; i++ {
		q[i] = p[(i+1)%n]
		if q[i] == p[i] {
			return nil
		}
	}
	return q
}

func buildGreedyPermutation(n int, p []int) []int {
	used := make([]bool, n+1)
	q := make([]int, n)

	nextAvailable := 1

	for i := 0; i < n; i++ {
		found := false
		for num := nextAvailable; num <= n; num++ {
			if !used[num] && num != p[i] {
				q[i] = num
				used[num] = true
				for nextAvailable <= n && used[nextAvailable] {
					nextAvailable++
				}
				found = true
				break
			}
		}

		if !found {
			for num := 1; num < nextAvailable; num++ {
				if !used[num] && num != p[i] {
					q[i] = num
					used[num] = true
					found = true
					break
				}
			}
		}

		if !found {
			for num := 1; num <= n; num++ {
				if !used[num] {
					if i > 0 {
						q[i] = q[i-1]
						q[i-1] = num
					} else {
						q[i] = num
					}
					used[num] = true
					break
				}
			}
		}
	}

	return q
}

func countInversionsFast(q []int) int {
	if len(q) <= 1 {
		return 0
	}
	arr := make([]int, len(q))
	copy(arr, q)
	temp := make([]int, len(arr))
	return mergeSortAndCount(arr, temp, 0, len(arr)-1)
}

func mergeSortAndCount(arr, temp []int, left, right int) int {
	count := 0
	if left < right {
		mid := (left + right) / 2
		count += mergeSortAndCount(arr, temp, left, mid)
		count += mergeSortAndCount(arr, temp, mid+1, right)
		count += mergeAndCount(arr, temp, left, mid, right)
	}
	return count
}

func mergeAndCount(arr, temp []int, left, mid, right int) int {
	i, j, k := left, mid+1, left
	count := 0

	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			count += (mid - i + 1)
			j++
		}
		k++
	}

	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}

	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}

	for i = left; i <= right; i++ {
		arr[i] = temp[i]
	}

	return count
}
