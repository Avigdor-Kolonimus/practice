package summerbackend2024

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type EventGCD2 struct {
	a int
	b int
}

type Factor struct {
	p int
	e int
}

type Entry struct {
	idx int
	cnt int
}

type FastScanner struct {
	r *bufio.Reader
}

func NewFastScanner() *FastScanner {
	return &FastScanner{
		r: bufio.NewReaderSize(os.Stdin, 1<<20),
	}
}

func (fs *FastScanner) NextInt() int {
	x := 0
	c, err := fs.r.ReadByte()

	for err == nil && (c < '0' || c > '9') {
		c, err = fs.r.ReadByte()
	}

	for err == nil && c >= '0' && c <= '9' {
		x = x*10 + int(c-'0')
		c, err = fs.r.ReadByte()
	}

	return x
}

func sieve(limit int) []int {
	isPrime := make([]bool, limit+1)

	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}

	primes := make([]int, 0)

	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

func factorizeGCD2(x int, primes []int) []Factor {
	res := make([]Factor, 0)

	for _, p := range primes {
		if p*p > x {
			break
		}

		if x%p == 0 {
			cnt := 0

			for x%p == 0 {
				x /= p
				cnt++
			}

			res = append(res, Factor{p, cnt})
		}
	}

	if x > 1 {
		res = append(res, Factor{x, 1})
	}

	return res
}

func buildIndex(arr []int, primes []int) map[int][]Entry {
	index := make(map[int][]Entry)
	counts := make(map[int]int)

	for i, x := range arr {
		factors := factorizeGCD2(x, primes)

		for _, f := range factors {
			counts[f.p] += f.e

			index[f.p] = append(
				index[f.p],
				Entry{
					idx: i,
					cnt: counts[f.p],
				},
			)
		}
	}

	return index
}

func firstIndex(index map[int][]Entry, p int, need int, length int) int {
	list, ok := index[p]

	if !ok {
		return length
	}

	pos := sort.Search(len(list), func(i int) bool {
		return list[i].cnt >= need
	})

	if pos == len(list) {
		return length
	}

	return list[pos].idx
}

func findThreshold(x int, index map[int][]Entry, primes []int, length int) int {
	factors := factorizeGCD2(x, primes)

	result := 0

	for _, f := range factors {
		pos := firstIndex(index, f.p, f.e, length)

		if pos == length {
			return length
		}

		if pos > result {
			result = pos
		}
	}

	return result
}

type SparseTable struct {
	st  [][]int64
	log []int
}

func NewSparseTable(arr []int64) *SparseTable {
	n := len(arr)

	log := make([]int, n+1)

	for i := 2; i <= n; i++ {
		log[i] = log[i/2] + 1
	}

	k := log[n] + 1

	st := make([][]int64, k)

	st[0] = make([]int64, n)
	copy(st[0], arr)

	for j := 1; j < k; j++ {
		st[j] = make([]int64, n)

		length := 1 << j
		half := length >> 1

		for i := 0; i+length <= n; i++ {
			a := st[j-1][i]
			b := st[j-1][i+half]

			if a < b {
				st[j][i] = a
			} else {
				st[j][i] = b
			}
		}
	}

	return &SparseTable{
		st:  st,
		log: log,
	}
}

func (s *SparseTable) Query(l, r int) int64 {
	length := r - l + 1
	j := s.log[length]

	a := s.st[j][l]
	b := s.st[j][r-(1<<j)+1]

	if a < b {
		return a
	}

	return b
}

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/gcd-2
// GCD2 - problem 35
func GCD2() {
	in := NewFastScanner()
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	n := in.NextInt()
	m := in.NextInt()
	q := in.NextInt()

	A := make([]int, n+1)
	B := make([]int, m+1)

	for i := range A {
		A[i] = in.NextInt()
	}

	for i := range B {
		B[i] = in.NextInt()
	}

	CA := make([]int64, n+1)
	CB := make([]int64, m+1)

	for i := range CA {
		CA[i] = int64(in.NextInt())
	}

	for i := range CB {
		CB[i] = int64(in.NextInt())
	}

	primes := sieve(31623)

	indexA := buildIndex(A, primes)
	indexB := buildIndex(B, primes)

	rmq := NewSparseTable(CA)

	prefixMinB := make([]int64, m+1)
	prefixMinB[0] = CB[0]
	for i := 1; i <= m; i++ {
		prefixMinB[i] = prefixMinB[i-1]

		if CB[i] < prefixMinB[i] {
			prefixMinB[i] = CB[i]
		}
	}

	for range q {
		k := in.NextInt()

		EventGCD2s := make([]EventGCD2, 0, k)

		for t := 0; t < k; t++ {
			x := in.NextInt()

			a := findThreshold(x, indexA, primes, n+1)
			b := findThreshold(x, indexB, primes, m+1)

			if a <= n {
				EventGCD2s = append(EventGCD2s, EventGCD2{a, b})
			}
		}

		sort.Slice(EventGCD2s, func(i, j int) bool {
			return EventGCD2s[i].a < EventGCD2s[j].a
		})

		limit := m + 1

		left := 0

		const INF int64 = 1 << 62
		ans := INF

		process := func(l, r int) {
			if l > r {
				return
			}

			minCA := rmq.Query(l, r)

			minCB := prefixMinB[limit-1]

			cost := minCA + minCB

			if cost < ans {
				ans = cost
			}
		}

		pos := 0
		for pos < len(EventGCD2s) {
			currentA := EventGCD2s[pos].a

			process(left, currentA-1)

			for pos < len(EventGCD2s) &&
				EventGCD2s[pos].a == currentA {

				if EventGCD2s[pos].b < limit {
					limit = EventGCD2s[pos].b
				}

				pos++
			}

			left = currentA
		}

		process(left, n)

		fmt.Fprintln(out, ans)
	}
}
