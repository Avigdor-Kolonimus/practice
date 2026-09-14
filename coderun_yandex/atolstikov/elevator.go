package atolstikov

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const INF int64 = 1 << 62

// ------------------------------------------------------------
// Fast Scanner
// ------------------------------------------------------------

type FastScanner struct {
	r *bufio.Reader
}

func NewFastScanner() *FastScanner {
	return &FastScanner{
		r: bufio.NewReaderSize(os.Stdin, 1<<20),
	}
}

func (s *FastScanner) ReadInt64() int64 {
	var sign int64 = 1
	var num int64

	c, err := s.r.ReadByte()
	for err == nil && (c == ' ' || c == '\n' || c == '\r' || c == '\t') {
		c, err = s.r.ReadByte()
	}

	if c == '-' {
		sign = -1
		c, err = s.r.ReadByte()
	}

	for err == nil && c >= '0' && c <= '9' {
		num = num*10 + int64(c-'0')
		c, err = s.r.ReadByte()
	}

	return num * sign
}

// ------------------------------------------------------------
// Data
// ------------------------------------------------------------

type Entry struct {
	floor int64
	key   int64
	id    int
}

// ------------------------------------------------------------
// Direction Index
// ------------------------------------------------------------

type DirIndex struct {
	floors  []int64
	entries []Entry

	start []int
	end   []int
	ptr   []int

	m    int
	size int
	seg  []int64
}

func NewDirIndex(entries []Entry, served []bool, floorPos []int) *DirIndex {
	if len(entries) == 0 {
		return nil
	}

	e := make([]Entry, len(entries))
	copy(e, entries)

	sort.Slice(e, func(i, j int) bool {
		if e[i].floor != e[j].floor {
			return e[i].floor < e[j].floor
		}
		return e[i].key < e[j].key
	})

	floors := make([]int64, 0, len(e))
	start := make([]int, 0, len(e))
	end := make([]int, 0, len(e))
	ptr := make([]int, 0, len(e))

	i := 0

	for i < len(e) {
		floor := e[i].floor
		floorIndex := len(floors)

		floors = append(floors, floor)
		start = append(start, i)

		j := i

		for j < len(e) && e[j].floor == floor {
			floorPos[e[j].id] = floorIndex
			j++
		}

		end = append(end, j)
		ptr = append(ptr, i)

		i = j
	}

	m := len(floors)

	size := 1
	for size < m {
		size <<= 1
	}

	seg := make([]int64, 2*size)

	for i := range seg {
		seg[i] = INF
	}

	for i := 0; i < m; i++ {
		seg[size+i] = e[start[i]].key
	}

	for i := size - 1; i >= 1; i-- {
		seg[i] = min(seg[i<<1], seg[i<<1|1])
	}

	return &DirIndex{
		floors:  floors,
		entries: e,
		start:   start,
		end:     end,
		ptr:     ptr,
		m:       m,
		size:    size,
		seg:     seg,
	}
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func (d *DirIndex) updateLeaf(i int, value int64) {
	p := d.size + i
	d.seg[p] = value

	p >>= 1

	for p >= 1 {
		newValue := min(
			d.seg[p<<1],
			d.seg[p<<1|1],
		)

		if d.seg[p] == newValue {
			break
		}

		d.seg[p] = newValue

		if p == 1 {
			break
		}

		p >>= 1
	}
}

// Refreshes one floor after passengers have been served.
func (d *DirIndex) refreshFloor(floorIndex int, served []bool) {
	p := d.ptr[floorIndex]
	e := d.end[floorIndex]

	for p < e && served[d.entries[p].id] {
		p++
	}

	d.ptr[floorIndex] = p

	value := INF

	if p < e {
		value = d.entries[p].key
	}

	d.updateLeaf(floorIndex, value)
}

// ------------------------------------------------------------
// Segment tree range minimum
// ------------------------------------------------------------

func (d *DirIndex) rangeMin(l, r int) int64 {
	if l > r {
		return INF
	}

	left := l + d.size
	right := r + d.size

	res := INF

	for left <= right {
		if left&1 == 1 {
			res = min(res, d.seg[left])
			left++
		}

		if right&1 == 0 {
			res = min(res, d.seg[right])
			right--
		}

		left >>= 1
		right >>= 1
	}

	return res
}

// Finds a leaf containing value <= limit.
func (d *DirIndex) findLeaf(
	node, nl, nr int,
	ql, qr int,
	limit int64,
) int {

	if nr < ql || qr < nl {
		return -1
	}

	if d.seg[node] > limit {
		return -1
	}

	if nl == nr {
		return nl
	}

	mid := (nl + nr) >> 1

	left := d.findLeaf(
		node<<1,
		nl,
		mid,
		ql,
		qr,
		limit,
	)

	if left != -1 {
		return left
	}

	return d.findLeaf(
		node<<1|1,
		mid+1,
		nr,
		ql,
		qr,
		limit,
	)
}

// Extracts one passenger whose key <= limit
// from floors [l, r].
func (d *DirIndex) extract(
	l, r int,
	limit int64,
	served []bool,
) (int, bool) {

	if d.m == 0 || l > r {
		return -1, false
	}

	for {
		if d.rangeMin(l, r) > limit {
			return -1, false
		}

		pos := d.findLeaf(
			1,
			0,
			d.size-1,
			l,
			r,
			limit,
		)

		if pos < 0 || pos >= d.m {
			return -1, false
		}

		d.refreshFloor(pos, served)

		if d.seg[d.size+pos] > limit {
			continue
		}

		p := d.ptr[pos]

		if p >= d.end[pos] {
			continue
		}

		id := d.entries[p].id

		d.ptr[pos] = p + 1

		d.refreshFloor(pos, served)

		return id, true
	}
}

// ------------------------------------------------------------
// Binary search
// ------------------------------------------------------------

func (d *DirIndex) lowerBound(x int64) int {
	return sort.Search(len(d.floors), func(i int) bool {
		return d.floors[i] >= x
	})
}

func (d *DirIndex) upperBound(x int64) int {
	return sort.Search(len(d.floors), func(i int) bool {
		return d.floors[i] > x
	})
}

// https://coderun.yandex.ru/selections/atolstikov/problems/elevator
// Elevator - problem 10
func Elevator() {
	in := NewFastScanner()

	n := int(in.ReadInt64())
	tMove := in.ReadInt64()

	ti := make([]int64, n)
	si := make([]int64, n)
	di := make([]int64, n)
	dir := make([]int8, n)

	upEntries := make([]Entry, 0, n)
	downEntries := make([]Entry, 0, n)

	for i := 0; i < n; i++ {
		t := in.ReadInt64()
		s := in.ReadInt64()
		d := in.ReadInt64()

		ti[i] = t
		si[i] = s
		di[i] = d

		if d > s {
			dir[i] = 1

			// key = ti - si * tMove
			key := t - s*tMove

			upEntries = append(upEntries, Entry{
				floor: s,
				key:   key,
				id:    i,
			})
		} else {
			dir[i] = -1

			// key = ti + si * tMove
			key := t + s*tMove

			downEntries = append(downEntries, Entry{
				floor: s,
				key:   key,
				id:    i,
			})
		}
	}

	served := make([]bool, n)
	wait := make([]int64, n)

	upFloorPos := make([]int, n)
	downFloorPos := make([]int, n)

	for i := range upFloorPos {
		upFloorPos[i] = -1
		downFloorPos[i] = -1
	}

	var upIndex *DirIndex
	var downIndex *DirIndex

	if len(upEntries) > 0 {
		upIndex = NewDirIndex(
			upEntries,
			served,
			upFloorPos,
		)
	}

	if len(downEntries) > 0 {
		downIndex = NewDirIndex(
			downEntries,
			served,
			downFloorPos,
		)
	}

	// --------------------------------------------------------
	// Current state
	// --------------------------------------------------------

	var curTime int64 = 0
	var curFloor int64 = 1

	ptrTime := 0

	// --------------------------------------------------------
	// Process passengers
	// --------------------------------------------------------

	for {
		// Find the first passenger that hasn't been served.
		for ptrTime < n && served[ptrTime] {
			ptrTime++
		}

		if ptrTime >= n {
			break
		}

		callTime := ti[ptrTime]

		if curTime < callTime {
			curTime = callTime
		}

		startFloor := si[ptrTime]
		destFloor := di[ptrTime]

		// Empty elevator goes directly to this passenger.
		travel := abs(startFloor-curFloor) * tMove

		pickupTime := curTime + travel

		wait[ptrTime] = pickupTime - callTime
		served[ptrTime] = true

		// Remove this passenger from the corresponding index.
		if dir[ptrTime] == 1 && upIndex != nil {
			pos := upFloorPos[ptrTime]

			if pos >= 0 {
				upIndex.refreshFloor(pos, served)
			}
		} else if dir[ptrTime] == -1 && downIndex != nil {
			pos := downFloorPos[ptrTime]

			if pos >= 0 {
				downIndex.refreshFloor(pos, served)
			}
		}

		curTime = pickupTime
		curFloor = startFloor

		// ----------------------------------------------------
		// Elevator moves UP
		// ----------------------------------------------------

		if destFloor > startFloor {
			end := destFloor

			T0 := curTime
			F0 := curFloor

			// For a passenger at floor f:
			//
			// key = callTime - f*tMove
			//
			// Passenger can be picked up if:
			//
			// callTime <= T0 + (f-F0)*tMove
			//
			// => callTime - f*tMove <= T0 - F0*tMove
			//
			C := T0 - F0*tMove

			if upIndex != nil {
				l := upIndex.lowerBound(F0)

				for {
					r := upIndex.upperBound(end) - 1

					if l > r {
						break
					}

					id, ok := upIndex.extract(
						l,
						r,
						C,
						served,
					)

					if !ok {
						break
					}

					// Time when elevator reaches this passenger.
					pickup := T0 +
						(si[id]-F0)*tMove

					wait[id] = pickup - ti[id]
					served[id] = true

					pos := upFloorPos[id]

					if pos >= 0 {
						upIndex.refreshFloor(
							pos,
							served,
						)
					}

					// This passenger may extend the trip.
					if di[id] > end {
						end = di[id]
					}
				}
			}

			curTime = T0 + (end-F0)*tMove
			curFloor = end

		} else {

			// ------------------------------------------------
			// Elevator moves DOWN
			// ------------------------------------------------

			end := destFloor

			T0 := curTime
			F0 := curFloor

			// key = callTime + floor*tMove
			//
			// Passenger can be picked up if:
			//
			// callTime <= T0 + (F0-floor)*tMove
			//
			// => callTime + floor*tMove <=
			//    T0 + F0*tMove
			//
			C := T0 + F0*tMove

			if downIndex != nil {
				r := downIndex.upperBound(F0) - 1

				for {
					l := downIndex.lowerBound(end)

					if l > r {
						break
					}

					id, ok := downIndex.extract(
						l,
						r,
						C,
						served,
					)

					if !ok {
						break
					}

					pickup := T0 +
						(F0-si[id])*tMove

					wait[id] = pickup - ti[id]
					served[id] = true

					pos := downFloorPos[id]

					if pos >= 0 {
						downIndex.refreshFloor(
							pos,
							served,
						)
					}

					// This passenger may extend the trip downward.
					if di[id] < end {
						end = di[id]
					}
				}
			}

			curTime = T0 + (F0-end)*tMove
			curFloor = end
		}
	}

	// ------------------------------------------------------------
	// Output
	// ------------------------------------------------------------

	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	for i := 0; i < n; i++ {
		fmt.Fprintln(out, wait[i])
	}
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
