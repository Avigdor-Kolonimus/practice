package summercommon2026

import (
	"bufio"
	"os"
)

// ─── Fast I/O ─────────────────────────────────────────────────────────────────
var reader *bufio.Reader
var writer *bufio.Writer

func initIO() {
	reader = bufio.NewReaderSize(os.Stdin, 1<<20)
	writer = bufio.NewWriterSize(os.Stdout, 1<<20)
}

func readInt32() int32 {
	b, _ := reader.ReadByte()
	for b <= ' ' {
		b, _ = reader.ReadByte()
	}
	var x int32
	for b >= '0' && b <= '9' {
		x = x*10 + int32(b-'0')
		b, _ = reader.ReadByte()
	}
	return x
}

func readInt64() int64 {
	b, _ := reader.ReadByte()
	for b <= ' ' {
		b, _ = reader.ReadByte()
	}
	var x int64
	for b >= '0' && b <= '9' {
		x = x*10 + int64(b-'0')
		b, _ = reader.ReadByte()
	}
	return x
}

func readOp() byte {
	b, _ := reader.ReadByte()
	for b <= ' ' {
		b, _ = reader.ReadByte()
	}
	first := b
	for {
		b, _ = reader.ReadByte()
		if b <= ' ' || b == 0 {
			break
		}
	}
	return first
}

func writeInt64VacationPhotos(x int64) {
	if x == 0 {
		writer.WriteByte('0')
		return
	}
	var buf [20]byte
	i := len(buf)
	for x > 0 {
		i--
		buf[i] = byte('0' + x%10)
		x /= 10
	}
	writer.Write(buf[i:])
}

// ─── Генератор приоритетов (XORShift) ─────────────────────────────────────────
var rngState uint64 = 88172645463325252

func nextRand() uint32 {
	rngState ^= rngState << 13
	rngState ^= rngState >> 7
	rngState ^= rngState << 17
	return uint32(rngState >> 32)
}

// ─── Глобальные структуры данных ──────────────────────────────────────────────
var (
	tLeft   []int32
	tRight  []int32
	tParent []int32

	tWidth  []int64 // Размеры блоков МОГУТ быть больше 2 млрд!
	tMaxGap []int64 // Максимальный пробел в поддереве

	tPrio  []uint32
	tIsGap []bool

	nextB []int32
	prevB []int32

	actionNode []int32

	nodeCnt   int32
	root      int32
	totalFree int64
)

func allocNodes(maxNodes, N int32) {
	tLeft = make([]int32, maxNodes)
	tRight = make([]int32, maxNodes)
	tParent = make([]int32, maxNodes)
	tWidth = make([]int64, maxNodes)
	tMaxGap = make([]int64, maxNodes)
	tPrio = make([]uint32, maxNodes)
	tIsGap = make([]bool, maxNodes)
	nextB = make([]int32, maxNodes)
	prevB = make([]int32, maxNodes)
	actionNode = make([]int32, N+5)
}

func newNode(w int64, gap bool) int32 {
	nodeCnt++
	idx := nodeCnt
	tWidth[idx] = w
	tIsGap[idx] = gap
	if gap {
		tMaxGap[idx] = w
	} else {
		tMaxGap[idx] = 0
	}
	tPrio[idx] = nextRand()
	return idx
}

func updateNode(n int32) {
	if n == 0 {
		return
	}
	l, r := tLeft[n], tRight[n]

	if tIsGap[n] {
		tMaxGap[n] = tWidth[n]
	} else {
		tMaxGap[n] = 0
	}

	if l != 0 {
		if tMaxGap[l] > tMaxGap[n] {
			tMaxGap[n] = tMaxGap[l]
		}
		tParent[l] = n
	}
	if r != 0 {
		if tMaxGap[r] > tMaxGap[n] {
			tMaxGap[n] = tMaxGap[r]
		}
		tParent[r] = n
	}
}

func updateChainUp(n int32) {
	for n != 0 {
		updateNode(n)
		n = tParent[n]
	}
}

// ─── Вращения Treap ──────────────────────────────────────────────────────────
func rotateLeft(p int32) {
	g := tParent[p]
	r := tRight[p]
	tRight[p] = tLeft[r]
	if tLeft[r] != 0 {
		tParent[tLeft[r]] = p
	}
	tLeft[r] = p
	tParent[p] = r
	tParent[r] = g
	if g != 0 {
		if tLeft[g] == p {
			tLeft[g] = r
		} else {
			tRight[g] = r
		}
	} else {
		root = r
	}
	updateNode(p)
	updateNode(r)
}

func rotateRight(p int32) {
	g := tParent[p]
	l := tLeft[p]
	tLeft[p] = tRight[l]
	if tRight[l] != 0 {
		tParent[tRight[l]] = p
	}
	tRight[l] = p
	tParent[p] = l
	tParent[l] = g
	if g != 0 {
		if tLeft[g] == p {
			tLeft[g] = l
		} else {
			tRight[g] = l
		}
	} else {
		root = l
	}
	updateNode(p)
	updateNode(l)
}

func bubbleUp(x int32) {
	for tParent[x] != 0 && tPrio[tParent[x]] < tPrio[x] {
		p := tParent[x]
		if tLeft[p] == x {
			rotateRight(p)
		} else {
			rotateLeft(p)
		}
	}
	updateChainUp(x)
}

// Быстрая вставка после узла
func insertAfter(node, x int32) {
	if tRight[node] == 0 {
		tRight[node] = x
		tParent[x] = node
	} else {
		curr := tRight[node]
		for tLeft[curr] != 0 {
			curr = tLeft[curr]
		}
		tLeft[curr] = x
		tParent[x] = curr
	}
	bubbleUp(x)
}

func mergeVacationPhotos(a, b int32) int32 {
	if a == 0 || b == 0 {
		if a != 0 {
			return a
		}
		return b
	}
	if tPrio[a] > tPrio[b] {
		tRight[a] = mergeVacationPhotos(tRight[a], b)
		updateNode(a)
		return a
	} else {
		tLeft[b] = mergeVacationPhotos(a, tLeft[b])
		updateNode(b)
		return b
	}
}

func eraseByPointer(n int32) {
	p := tParent[n]
	merged := mergeVacationPhotos(tLeft[n], tRight[n])
	if merged != 0 {
		tParent[merged] = p
	}
	if p == 0 {
		root = merged
	} else if tLeft[p] == n {
		tLeft[p] = merged
		updateChainUp(p)
	} else {
		tRight[p] = merged
		updateChainUp(p)
	}
	tLeft[n] = 0
	tRight[n] = 0
	tParent[n] = 0
}

func findLeftmostGap(t int32, s int64) int32 {
	if t == 0 || tMaxGap[t] < s {
		return 0
	}
	for {
		l := tLeft[t]
		if l != 0 && tMaxGap[l] >= s {
			t = l
			continue
		}
		if tIsGap[t] && tWidth[t] >= s {
			return t
		}
		t = tRight[t]
	}
}

func getRightmost(t int32) int32 {
	if t == 0 {
		return 0
	}
	for tRight[t] != 0 {
		t = tRight[t]
	}
	return t
}

func printPhotos() {
	head := root
	if head == 0 {
		return
	}
	for tLeft[head] != 0 {
		head = tLeft[head]
	}
	currCoord := int64(0)
	cur := head
	for cur != 0 {
		if !tIsGap[cur] {
			writeInt64VacationPhotos(currCoord)
			writer.WriteByte(' ')
			writeInt64VacationPhotos(tWidth[cur])
			writer.WriteByte('\n')
		}
		currCoord += tWidth[cur]
		cur = nextB[cur]
	}
}

// Дефрагментация
func compact() {
	for {
		g := findLeftmostGap(root, 1)
		if g == 0 {
			break
		}
		p := prevB[g]
		nxt := nextB[g]
		if p != 0 {
			nextB[p] = nxt
		}
		if nxt != 0 {
			prevB[nxt] = p
		}
		eraseByPointer(g)
	}
	if totalFree > 0 {
		lastPhoto := getRightmost(root)
		idx := newNode(totalFree, true)

		if lastPhoto == 0 {
			root = idx
			tParent[idx] = 0
		} else {
			tRight[lastPhoto] = idx
			tParent[idx] = lastPhoto
			bubbleUp(idx)
		}

		if lastPhoto != 0 {
			nextB[lastPhoto] = idx
			prevB[idx] = lastPhoto
			nextB[idx] = 0
		} else {
			prevB[idx] = 0
			nextB[idx] = 0
		}
	}
}

func addPhoto(s int64, actionID int32) bool {
	if s > totalFree {
		return false
	}
	found := findLeftmostGap(root, s)
	if found == 0 {
		compact()
		found = findLeftmostGap(root, s)
	}
	gapW := tWidth[found]

	tIsGap[found] = false
	tWidth[found] = s
	tMaxGap[found] = 0
	updateChainUp(found)
	actionNode[actionID] = found

	if gapW > s {
		leftoverIdx := newNode(gapW-s, true)
		insertAfter(found, leftoverIdx)

		nxt := nextB[found]
		nextB[found] = leftoverIdx
		prevB[leftoverIdx] = found
		nextB[leftoverIdx] = nxt
		if nxt != 0 {
			prevB[nxt] = leftoverIdx
		}
	}
	totalFree -= s
	return true
}

// In-place удаление и схлопывание соседей
func deletePhoto(k int32) {
	n := actionNode[k]
	if n == 0 || tIsGap[n] {
		return // Защита
	}
	w := tWidth[n]
	totalFree += w

	p := prevB[n]
	nxt := nextB[n]

	mergeP := p != 0 && tIsGap[p]
	mergeNxt := nxt != 0 && tIsGap[nxt]

	if mergeP && mergeNxt {
		tWidth[n] += tWidth[p] + tWidth[nxt]
		eraseByPointer(p)
		eraseByPointer(nxt)

		pp := prevB[p]
		nnxt := nextB[nxt]

		prevB[n] = pp
		if pp != 0 {
			nextB[pp] = n
		}
		nextB[n] = nnxt
		if nnxt != 0 {
			prevB[nnxt] = n
		}

		tIsGap[n] = true
		updateChainUp(n)
	} else if mergeP {
		tWidth[n] += tWidth[p]
		eraseByPointer(p)

		pp := prevB[p]

		prevB[n] = pp
		if pp != 0 {
			nextB[pp] = n
		}
		nextB[n] = nxt
		if nxt != 0 {
			prevB[nxt] = n
		}

		tIsGap[n] = true
		updateChainUp(n)
	} else if mergeNxt {
		tWidth[n] += tWidth[nxt]
		eraseByPointer(nxt)

		nnxt := nextB[nxt]

		nextB[n] = nnxt
		if nnxt != 0 {
			prevB[nnxt] = n
		}

		tIsGap[n] = true
		updateChainUp(n)
	} else {
		tIsGap[n] = true
		updateChainUp(n)
	}
}

// https://coderun.yandex.ru/seasons/2026-summer/tracks/common/problem/vacation-photos
// VacationPhotos - problem 12
func VacationPhotos() {
	initIO()
	defer writer.Flush()

	M := readInt64()
	N := readInt32()

	maxNodes := int32(2*N + 5)
	allocNodes(maxNodes, N)

	totalFree = M
	if M > 0 {
		root = newNode(M, true)
		prevB[root] = 0
		nextB[root] = 0
	}

	for i := int32(1); i <= N; i++ {
		op := readOp()
		if op == 'A' {
			s := readInt64()
			if !addPhoto(s, i) {
				writer.WriteString("OVERFLOW\n")
				printPhotos()
				return
			}
		} else {
			k := readInt32()
			deletePhoto(k)
		}
	}

	printPhotos()
}
