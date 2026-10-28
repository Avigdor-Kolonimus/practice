package backend

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pallet struct {
	w int
	h int
}

// https://coderun.yandex.ru/selections/first-2023-backend/problems/warehouse-pallets
// WarehousePallets - problem 28
func WarehousePallets() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	pallets := make([]Pallet, n)

	for i := range n {
		line, _ = reader.ReadString('\n')
		fields := strings.Fields(line)

		w, _ := strconv.Atoi(fields[0])
		h, _ := strconv.Atoi(fields[1])

		if w > h {
			w, h = h, w
		}

		pallets[i] = Pallet{w: w, h: h}
	}

	sort.Slice(pallets, func(i, j int) bool {
		if pallets[i].w != pallets[j].w {
			return pallets[i].w > pallets[j].w
		}
		return pallets[i].h > pallets[j].h
	})

	ans := 0
	maxH := 0
	for i := 0; i < n; {
		j := i

		for j < n && pallets[j].w == pallets[i].w {
			j++
		}

		for k := i; k < j; k++ {
			if pallets[k].h >= maxH {
				ans++
			}
		}

		for k := i; k < j; k++ {
			if pallets[k].h > maxH {
				maxH = pallets[k].h
			}
		}

		i = j
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
