package wintercommon2025

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func checkLimits(maxTime time.Duration, maxMemoryMB int, fn func()) {
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

func solveProtectiveField(n int) (string, string) {
	// Find M such that 10^M > n-1
	m := 0
	limit := 1
	// Using int for limit is safe because n <= 10^4.
	// 10^4 fits in int.
	for limit <= n-1 {
		limit *= 10
		m++
	}
	if m == 0 {
		m = 1
	}

	s := strings.Repeat("9", m)
	return s, s
}

// https://coderun.yandex.ru/selections/2025-winter-common/problems/protective-field
// ProtectiveField - problem 17
func ProtectiveField() {
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

	checkLimits(1*time.Second, 256, func() {
		a, d := solveProtectiveField(n)

		writer.WriteString(a)
		writer.WriteByte('\n')
		writer.WriteString(d)
		writer.WriteByte('\n')
	})
}
