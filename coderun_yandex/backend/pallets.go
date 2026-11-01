package backend

import (
	"bufio"
	"fmt"
	"os"
)

// https://coderun.yandex.ru/selections/first-2023-backend/problems/pallets
// Pallets - problem 38
func Pallets() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20)
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	if n == 0 {
		return
	}

	deliveryIDs := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &deliveryIDs[i])
	}

	parentIDs := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &parentIDs[i])
	}

	var k int
	fmt.Fscan(in, &k)

	missing := make(map[int]struct{}, k)
	for i := 0; i < k; i++ {
		var id int
		fmt.Fscan(in, &id)
		missing[id] = struct{}{}
	}

	isBad := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		if _, exists := missing[deliveryIDs[i]]; exists {
			isBad[i] = true
		}
	}

	for i := 1; i <= n; i++ {
		if isBad[i] {
			curr := parentIDs[i]

			for curr != 0 && !isBad[curr] {
				isBad[curr] = true
				curr = parentIDs[curr]
			}
		}
	}

	validPallets := make([]int, 0)
	for i := 1; i <= n; i++ {
		if parentIDs[i] == 0 && !isBad[i] {
			validPallets = append(validPallets, i)
		}
	}

	fmt.Fprintln(out, len(validPallets))

	for i, id := range validPallets {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, id)
	}

	if len(validPallets) > 0 {
		fmt.Fprintln(out)
	}
}
