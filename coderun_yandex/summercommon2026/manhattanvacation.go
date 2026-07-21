package summercommon2026

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
)

// ---------- fast IO ----------

type fastReader struct {
	r *bufio.Reader
}

func newFastReader() *fastReader {
	return &fastReader{r: bufio.NewReaderSize(os.Stdin, 1<<20)}
}

func (fr *fastReader) readInt64() int64 {
	c, _ := fr.r.ReadByte()
	for c == ' ' || c == '\n' || c == '\r' || c == '\t' {
		c, _ = fr.r.ReadByte()
	}
	neg := false
	if c == '-' {
		neg = true
		c, _ = fr.r.ReadByte()
	}
	var n int64
	for c >= '0' && c <= '9' {
		n = n*10 + int64(c-'0')
		c, _ = fr.r.ReadByte()
	}
	if neg {
		n = -n
	}
	return n
}

func (fr *fastReader) readInt() int {
	return int(fr.readInt64())
}

// ---------- DSU ----------

type dsu struct {
	parent []int32
	rank   []int8
}

func newDSU(n int) *dsu {
	p := make([]int32, n)
	for i := range p {
		p[i] = int32(i)
	}
	return &dsu{parent: p, rank: make([]int8, n)}
}

func (d *dsu) find(x int32) int32 {
	for d.parent[x] != x {
		d.parent[x] = d.parent[d.parent[x]]
		x = d.parent[x]
	}
	return x
}

func (d *dsu) union(a, b int32) bool {
	ra, rb := d.find(a), d.find(b)
	if ra == rb {
		return false
	}
	if d.rank[ra] < d.rank[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	if d.rank[ra] == d.rank[rb] {
		d.rank[ra]++
	}
	return true
}

// ---------- edges ----------

type edge struct {
	weight int64
	u, v   int32
}

func abs64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// ---------- class construction (flood fill over rotated coords, L_inf) ----------
//
// activeByV is emulated with:
//   - order[]: indices sorted ascending by V
//   - sortedV[]: corresponding V values (parallel to order)
//   - a "next free slot" DSU (nextParent) that lets us jump from any
//     position to the next still-active position in O(alpha(n)) amortized,
//     exactly like std::set<pair<V,idx>> lower_bound + erase does.

func buildClasses(n int, U, V, R []int64) [][]int32 {
	d := newDSU(n)

	maxR := int64(0)
	for i := 0; i < n; i++ {
		if R[i] > maxR {
			maxR = R[i]
		}
	}

	order := make([]int32, n)
	for i := range order {
		order[i] = int32(i)
	}
	sort.Slice(order, func(a, b int) bool { return V[order[a]] < V[order[b]] })

	sortedV := make([]int64, n)
	for i, idx := range order {
		sortedV[i] = V[idx]
	}

	// next-alive "union find" over positions [0..n]; position n is a sentinel.
	nextParent := make([]int32, n+1)
	for i := 0; i <= n; i++ {
		nextParent[i] = int32(i)
	}
	var findNext func(x int32) int32
	findNext = func(x int32) int32 {
		for nextParent[x] != x {
			nextParent[x] = nextParent[nextParent[x]]
			x = nextParent[x]
		}
		return x
	}
	remove := func(pos int32) {
		nextParent[pos] = pos + 1
	}

	nn := int32(n)
	stack := make([]int32, 0, n)

	outerPos := int32(0)
	for {
		start := findNext(outerPos)
		if start >= nn {
			break
		}
		remove(start)
		stack = stack[:0]
		stack = append(stack, order[start])

		for len(stack) > 0 {
			c := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			window := R[c]
			if maxR > window {
				window = maxR
			}
			vLow := V[c] - window
			vHigh := V[c] + window

			posStart := int32(sort.Search(n, func(i int) bool { return sortedV[i] >= vLow }))
			cur := findNext(posStart)
			for cur < nn && sortedV[cur] <= vHigh {
				j := order[cur]
				du := abs64(U[c] - U[j])
				dv := abs64(V[c] - V[j])
				dist := du
				if dv > dist {
					dist = dv
				}
				if dist <= R[c] || dist <= R[j] {
					d.union(c, j)
					remove(cur)
					stack = append(stack, j)
					cur = findNext(cur)
				} else {
					cur = findNext(cur + 1)
				}
			}
		}
		outerPos = start
	}

	classes := make([][]int32, n)
	for v := 0; v < n; v++ {
		r := d.find(int32(v))
		classes[r] = append(classes[r], int32(v))
	}
	return classes
}

// ---------- octant candidate edges for one class (Hwang 1979) ----------

func appendOctantEdges(members []int32, globalX, globalY []int64, candidates *[]edge) {
	m := len(members)
	baseX := make([]int64, m)
	baseY := make([]int64, m)
	for k := 0; k < m; k++ {
		baseX[k] = globalX[members[k]]
		baseY[k] = globalY[members[k]]
	}

	const inf = int64(math.MaxInt64)
	x := make([]int64, m)
	y := make([]int64, m)
	order := make([]int32, m)

	for orientation := 0; orientation < 8; orientation++ {
		for k := 0; k < m; k++ {
			px, py := baseX[k], baseY[k]
			if orientation&1 != 0 {
				px, py = py, px
			}
			if orientation&2 != 0 {
				px = -px
			}
			if orientation&4 != 0 {
				py = -py
			}
			x[k] = px
			y[k] = py
		}

		for k := range order {
			order[k] = int32(k)
		}
		sort.Slice(order, func(a, b int) bool {
			ia, ib := order[a], order[b]
			diffA := x[ia] - y[ia]
			diffB := x[ib] - y[ib]
			if diffA != diffB {
				return diffA > diffB
			}
			return y[ia] > y[ib]
		})

		sortedY := append([]int64(nil), y...)
		sort.Slice(sortedY, func(a, b int) bool { return sortedY[a] < sortedY[b] })
		uniqueY := sortedY[:0]
		for i, v := range sortedY {
			if i == 0 || v != uniqueY[len(uniqueY)-1] {
				uniqueY = append(uniqueY, v)
			}
		}
		K := len(uniqueY)

		compress := func(value int64) int {
			return sort.Search(K, func(i int) bool { return uniqueY[i] >= value })
		}

		bestValue := make([]int64, K+1)
		bestOwner := make([]int32, K+1)
		for i := range bestValue {
			bestValue[i] = inf
			bestOwner[i] = -1
		}

		update := func(pos int, value int64, owner int32) {
			for i := K - pos; i <= K; i += i & (-i) {
				if value < bestValue[i] {
					bestValue[i] = value
					bestOwner[i] = owner
				}
			}
		}
		query := func(pos int) int32 {
			best := inf
			owner := int32(-1)
			for i := K - pos; i > 0; i -= i & (-i) {
				if bestValue[i] < best {
					best = bestValue[i]
					owner = bestOwner[i]
				}
			}
			return owner
		}

		for _, k := range order {
			pos := compress(y[k])
			owner := query(pos)
			if owner != -1 {
				gu := members[k]
				gv := members[owner]
				weight := abs64(globalX[gu]-globalX[gv]) + abs64(globalY[gu]-globalY[gv])
				*candidates = append(*candidates, edge{weight, gu, gv})
			}
			update(pos, x[k]+y[k], int32(k))
		}
	}
}

// ---------- main ----------
// https://coderun.yandex.ru/selections/2026-summer-common/problems/manhattan-vacation
// ManhattanVacation - problem 11
func ManhattanVacation() {
	fr := newFastReader()
	w := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer w.Flush()

	t := fr.readInt()

	for ; t > 0; t-- {
		n := fr.readInt()

		X := make([]int64, n)
		Y := make([]int64, n)
		R := make([]int64, n)
		U := make([]int64, n)
		V := make([]int64, n)
		for i := 0; i < n; i++ {
			X[i] = fr.readInt64()
			Y[i] = fr.readInt64()
			R[i] = fr.readInt64()
			U[i] = X[i] + Y[i]
			V[i] = X[i] - Y[i]
		}

		classes := buildClasses(n, U, V, R)

		const smallClass = 32
		candidates := make([]edge, 0, 8*n)
		classCount := 0

		for root := 0; root < n; root++ {
			members := classes[root]
			m := len(members)
			if m == 0 {
				continue
			}
			classCount++
			if m == 1 {
				continue
			}
			if m <= smallClass {
				for a := 0; a < m; a++ {
					ia := members[a]
					for b := a + 1; b < m; b++ {
						ib := members[b]
						weight := abs64(X[ia]-X[ib]) + abs64(Y[ia]-Y[ib])
						candidates = append(candidates, edge{weight, ia, ib})
					}
				}
			} else {
				appendOctantEdges(members, X, Y, &candidates)
			}
		}

		sort.Slice(candidates, func(a, b int) bool { return candidates[a].weight < candidates[b].weight })

		forest := newDSU(n)
		needed := n - classCount
		taken := 0
		var total int64
		chosen := make([][2]int32, 0, needed)

		for _, e := range candidates {
			if taken == needed {
				break
			}
			if forest.union(e.u, e.v) {
				total += e.weight
				chosen = append(chosen, [2]int32{e.u + 1, e.v + 1})
				taken++
			}
		}

		w.WriteString(strconv.FormatInt(total, 10))
		w.WriteByte('\n')
		w.WriteString(strconv.Itoa(len(chosen)))
		w.WriteByte('\n')
		for _, p := range chosen {
			w.WriteString(strconv.Itoa(int(p[0])))
			w.WriteByte(' ')
			w.WriteString(strconv.Itoa(int(p[1])))
			w.WriteByte('\n')
		}
	}
}
