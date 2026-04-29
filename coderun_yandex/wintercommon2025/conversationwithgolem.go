package wintercommon2025

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

const (
	modConversationWithGolem = 1000000007
)

func checkLimitsConversationWithGolem(maxTime time.Duration, maxMemoryMB int, fn func()) {
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

type FastScanner struct {
	r    io.Reader
	buf  []byte
	pos  int
	size int
}

func NewFastScanner(r io.Reader) *FastScanner {
	return &FastScanner{
		r:   r,
		buf: make([]byte, 1<<16), // 64KB buffer
	}
}

func (s *FastScanner) NextInt() int {
	var n int
	// Skip whitespace
	for {
		if s.pos >= s.size {
			var err error
			s.size, err = s.r.Read(s.buf)
			s.pos = 0
			if err != nil || s.size == 0 {
				return 0
			}
		}
		c := s.buf[s.pos]
		s.pos++
		if c >= '0' && c <= '9' {
			n = int(c - '0')
			break
		}
	}
	// Read number
	for {
		if s.pos >= s.size {
			var err error
			s.size, err = s.r.Read(s.buf)
			s.pos = 0
			if err != nil || s.size == 0 {
				break
			}
		}
		c := s.buf[s.pos]
		if c < '0' || c > '9' {
			break
		}
		s.pos++
		n = n*10 + int(c-'0')
	}
	return n
}

// https://coderun.yandex.ru/selections/2025-winter-common/problems/conversation-with-golem
// ConversationWithGolem - problem 19
func ConversationWithGolem() {
	checkLimitsConversationWithGolem(1*time.Second, 256, solveConversationWithGolem)
}

func solveConversationWithGolem() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	scanner := NewFastScanner(reader)

	m := scanner.NextInt()
	if m == 0 {
		return
	}

	var N int64 = 0
	var lastCount int64 = 0

	for i := 0; i < m; i++ {
		val := int64(scanner.NextInt())
		N += val
		if val > 0 {
			lastCount = val
		}
	}

	if N == 0 {
		fmt.Fprintln(writer, 0)
		return
	}

	S := N - lastCount
	if S == 0 {
		fmt.Fprintln(writer, 0)
		return
	}

	inv := make([]int64, N+1)
	inv[1] = 1
	for i := int64(2); i <= N; i++ {
		inv[i] = (modConversationWithGolem - (modConversationWithGolem/i)*inv[modConversationWithGolem%i]%modConversationWithGolem) % modConversationWithGolem
	}

	var sumInv int64 = 0

	// sum 1/(k(N-k)) для k от 1 до S
	for k := int64(1); k <= S; k++ {
		// term = 1/(k*(N-k)) = inv[k] * inv[N-k]
		term := (inv[k] * inv[N-k]) % modConversationWithGolem
		sumInv = (sumInv + term) % modConversationWithGolem
	}

	// Coeff = N*(N-1)/2
	coeff := (N * (N - 1)) % modConversationWithGolem
	coeff = (coeff * inv[2]) % modConversationWithGolem

	ans := (coeff * sumInv) % modConversationWithGolem
	fmt.Fprintln(writer, ans)
}
