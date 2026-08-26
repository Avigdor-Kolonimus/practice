package ababin

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	MAXX = 1_000_000
	MOD  = int64(1_000_000_007)
)

type ExpMap struct {
	keys []uint64
	vals []int32
	mask uint64
}

func NewExpMap(capacity int) *ExpMap {
	size := 1
	for size < capacity {
		size <<= 1
	}

	return &ExpMap{
		keys: make([]uint64, size),
		vals: make([]int32, size),
		mask: uint64(size - 1),
	}
}

func hash64(x uint64) uint64 {
	x ^= x >> 33
	x *= 0xff51afd7ed558ccd
	x ^= x >> 33
	x *= 0xc4ceb9fe1a85ec53
	x ^= x >> 33

	return x
}

func (m *ExpMap) Get(key uint64) int32 {
	pos := hash64(key) & m.mask

	for {
		if m.keys[pos] == 0 {
			return 0
		}

		if m.keys[pos] == key {
			return m.vals[pos]
		}

		pos = (pos + 1) & m.mask
	}
}

func (m *ExpMap) Set(key uint64, value int32) {
	pos := hash64(key) & m.mask

	for {
		if m.keys[pos] == 0 {
			m.keys[pos] = key
			m.vals[pos] = value
			return
		}

		if m.keys[pos] == key {
			m.vals[pos] = value
			return
		}

		pos = (pos + 1) & m.mask
	}
}

func makeKey(v, p int) uint64 {
	return (uint64(v) << 20) | uint64(p)
}

func buildSPF() []int {
	spf := make([]int, MAXX+1)

	for i := 2; i <= MAXX; i++ {
		if spf[i] != 0 {
			continue
		}

		spf[i] = i

		if i <= MAXX/i {
			for j := i * i; j <= MAXX; j += i {
				if spf[j] == 0 {
					spf[j] = i
				}
			}
		}
	}

	return spf
}

func powMod(a int64, e int32) int64 {
	result := int64(1)

	for e > 0 {
		if e&1 != 0 {
			result = result * a % MOD
		}

		a = a * a % MOD
		e >>= 1
	}

	return result
}

func nextInt(reader *bufio.Reader) int {
	line, _ := reader.ReadString('\n')
	value, _ := strconv.Atoi(strings.TrimSpace(line))

	return value
}

// https://coderun.yandex.ru/selections/ababin/problems/bonsai-tree
// BonsaiTree - problem 1
func BonsaiTree() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and Q input
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	n, _ := strconv.Atoi(parts[0])
	q, _ := strconv.Atoi(parts[1])

	// Node input
	line, _ = reader.ReadString('\n')
	parts = strings.Fields(line)
	parent := make([]int, n+1)
	left := make([]int, n+1)
	right := make([]int, n+1)
	for v := 2; v <= n; v++ {
		p, _ := strconv.Atoi(parts[v-2])

		parent[v] = p

		if left[p] == 0 {
			left[p] = v
		} else {
			right[p] = v
		}
	}

	ans := make([]int64, n+1)
	for v := 1; v <= n; v++ {
		ans[v] = 1
	}

	expMap := NewExpMap(1 << 23)
	spf := buildSPF()
	for range q {
		line, _ = reader.ReadString('\n')
		parts = strings.Fields(line)

		t, _ := strconv.Atoi(parts[0])

		if t == 1 {
			v, _ := strconv.Atoi(parts[1])
			x, _ := strconv.Atoi(parts[2])

			for x > 1 {
				p := spf[x]
				cnt := 0

				for x%p == 0 {
					x /= p
					cnt++
				}

				key := makeKey(v, p)

				oldChild := expMap.Get(key)
				newChild := oldChild + int32(cnt)

				expMap.Set(key, newChild)

				ans[v] = ans[v] * powMod(int64(p), int32(cnt)) % MOD

				child := v

				for parent[child] != 0 {
					pv := parent[child]

					sibling := left[pv]

					if sibling == child {
						sibling = right[pv]
					}

					siblingExp := expMap.Get(
						makeKey(sibling, p),
					)

					oldParent := min(
						oldChild,
						siblingExp,
					)

					newParent := min(
						newChild,
						siblingExp,
					)

					if oldParent == newParent {
						break
					}

					expMap.Set(
						makeKey(pv, p),
						newParent,
					)

					diff := newParent - oldParent

					ans[pv] = ans[pv] *
						powMod(int64(p), diff) % MOD

					child = pv
					oldChild = oldParent
					newChild = newParent
				}
			}
		} else {
			v, _ := strconv.Atoi(parts[1])

			writer.WriteString(strconv.FormatInt(ans[v], 10))
			writer.WriteByte('\n')
		}
	}
}
