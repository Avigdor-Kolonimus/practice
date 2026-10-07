package first2023mobiledev

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/first-2023-mobile-dev/problems/fan-ridesharing
// FanRidesharing - problem 15
func FanRidesharing() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	readInts := func() []int {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)

		result := make([]int, len(parts))

		for i, part := range parts {
			result[i], _ = strconv.Atoi(part)
		}

		return result
	}

	n := readInts()[0]

	teams := readInts()

	for len(teams) < n {
		teams = append(teams, readInts()...)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(teams)))

	k := readInts()[0]

	rooms := make([]int, 10001)

	for range k {
		room := readInts()

		capacity := room[0]
		count := room[1]

		rooms[capacity] += count
	}

	for i := range n {
		team := teams[i]
		found := false

		for capacity := team; capacity <= 10000; capacity++ {
			if rooms[capacity] > 0 {
				rooms[capacity]--
				found = true
				break
			}
		}

		if !found {
			writer.WriteString("No")
			writer.WriteByte('\n')
			return
		}
	}

	writer.WriteString("Yes")
	writer.WriteByte('\n')
}
