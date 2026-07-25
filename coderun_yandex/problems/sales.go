package problems

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// https://coderun.yandex.ru/problem/sales
// Sales - problem 71
func Sales() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	data := make(map[string]map[string]int)

	var buyer, product string
	var count int

	for {
		_, err := fmt.Fscan(reader, &buyer, &product, &count)
		if err != nil {
			break
		}

		if _, ok := data[buyer]; !ok {
			data[buyer] = make(map[string]int)
		}
		data[buyer][product] += count
	}

	buyers := make([]string, 0, len(data))
	for buyer := range data {
		buyers = append(buyers, buyer)
	}
	sort.Strings(buyers)

	for _, buyer := range buyers {
		fmt.Fprintf(writer, "%s:\n", buyer)

		products := make([]string, 0, len(data[buyer]))
		for product := range data[buyer] {
			products = append(products, product)
		}
		sort.Strings(products)

		for _, product := range products {
			fmt.Fprintf(writer, "%s %d\n", product, data[buyer][product])
		}
	}
}
