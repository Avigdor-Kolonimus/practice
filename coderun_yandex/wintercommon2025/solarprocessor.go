package wintercommon2025

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

func checkLimitsSolarProcessor(maxTime time.Duration, maxMemoryMB int, fn func()) {
	if os.Getenv("CHECK_LIMITS") == "" {
		fn()
		return
	}

	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	start := time.Now()
	fn()
	elapsed := time.Since(start)

	runtime.GC()
	runtime.ReadMemStats(&m2)

	allocated := m2.TotalAlloc - m1.TotalAlloc
	memoryMB := float64(allocated) / (1024 * 1024)
	maxMemoryBytes := uint64(maxMemoryMB) * 1024 * 1024

	timeOk := elapsed <= maxTime
	memoryOk := allocated <= maxMemoryBytes

	if !timeOk || !memoryOk {
		if elapsed > maxTime {
			fmt.Fprintf(os.Stderr, "time: %v (limit: %v)\n", elapsed, maxTime)
		}
		if memoryMB > float64(maxMemoryMB) {
			fmt.Fprintf(os.Stderr, "memory: %.2f МБ (limit: %d MB)\n", memoryMB, maxMemoryMB)
		}
	} else {
		fmt.Fprintf(os.Stderr, "✓ Time: %v, Memory: %.2f MB\n", elapsed, memoryMB)
	}
}

const MAX_NODES = 13000000
const MAX_BITS = 29

var (
	l    [MAX_NODES]int32
	r    [MAX_NODES]int32
	cnt  [MAX_NODES]int32
	memo [MAX_NODES]int32
	ptr  int32
)

func newNode() int32 {
	ptr++
	idx := ptr
	l[idx] = 0
	r[idx] = 0
	cnt[idx] = 0
	memo[idx] = 0
	return idx
}

func pushUp(u int32, bit int) {
	idx0 := l[u]
	idx1 := r[u]

	c0, c1 := int32(0), int32(0)
	if idx0 != 0 {
		c0 = cnt[idx0]
	}
	if idx1 != 0 {
		c1 = cnt[idx1]
	}

	cnt[u] = c0 + c1

	m0, m1 := int32(0), int32(0)
	if idx0 != 0 {
		m0 = memo[idx0]
	}
	if idx1 != 0 {
		m1 = memo[idx1]
	}

	full := int32(1) << bit

	if c0 == full {
		memo[u] = full + m1
	} else if c1 == full {
		memo[u] = full + m0
	} else {
		if m0 > m1 {
			memo[u] = m0
		} else {
			memo[u] = m1
		}
	}
}

func update(u int32, val int, bit int, add bool) {
	if bit < 0 {
		if add {
			cnt[u] = 1
			memo[u] = 1
		} else {
			cnt[u] = 0
			memo[u] = 0
		}
		return
	}

	dir := (val >> bit) & 1
	var childIdx int32

	if dir == 0 {
		if l[u] == 0 {
			l[u] = newNode()
		}
		childIdx = l[u]
	} else {
		if r[u] == 0 {
			r[u] = newNode()
		}
		childIdx = r[u]
	}

	update(childIdx, val, bit-1, add)
	pushUp(u, bit)
}

type FastReader struct {
	sc *bufio.Scanner
}

func NewFastReader(r *os.File) *FastReader {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	sc.Split(bufio.ScanWords)
	return &FastReader{sc: sc}
}

func (r *FastReader) ReadInt() int {
	r.sc.Scan()
	x, _ := strconv.Atoi(r.sc.Text())
	return x
}

func solveSolarProcessor() {
	reader := NewFastReader(os.Stdin)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	t := reader.ReadInt()

	ptr = 0

	for i := 0; i < t; i++ {
		n := reader.ReadInt()
		q := reader.ReadInt()

		a := make([]int, n)
		freq := make(map[int]int, n)

		root := newNode()

		for j := 0; j < n; j++ {
			a[j] = reader.ReadInt()
			freq[a[j]]++

			if freq[a[j]] == 1 {
				update(root, a[j], MAX_BITS, true)
			}
		}

		writer.WriteString(strconv.Itoa(int(memo[root])))
		writer.WriteByte('\n')

		for k := 0; k < q; k++ {
			j := reader.ReadInt()
			v := reader.ReadInt()
			j--

			oldVal := a[j]
			if oldVal != v {
				freq[oldVal]--
				if freq[oldVal] == 0 {
					update(root, oldVal, MAX_BITS, false)
				}

				a[j] = v

				freq[v]++
				if freq[v] == 1 {
					update(root, v, MAX_BITS, true)
				}
			}

			writer.WriteString(strconv.Itoa(int(memo[root])))
			writer.WriteByte('\n')
		}
	}
}

// https://coderun.yandex.ru/selections/2025-winter-common/problems/solar-processor
// SolarProcessor - problem 20
func SolarProcessor() {
	checkLimitsSolarProcessor(4*time.Second, 256, solveSolarProcessor)
}
