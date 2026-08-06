package devgointerview

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Order struct {
	time int
	id   int
}

// https://coderun.yandex.ru/selections/dev-go-interview/problems/busiest-restaurants
// BusiestRestaurants - problem 8
func BusiestRestaurants() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// K and limit input
	line, _ := reader.ReadString('\n')
	fields := strings.Fields(line)
	k, _ := strconv.Atoi(fields[0])
	limit, _ := strconv.Atoi(fields[1])

	// Q input
	line, _ = reader.ReadString('\n')
	q, _ := strconv.Atoi(strings.TrimSpace(line))

	cnt := make(map[int]int)
	queue := make([]Order, 0, q)
	head, busy, active := 0, 0, 0
	for i := 0; i < q; i++ {
		line, _ = reader.ReadString('\n')
		fields = strings.Fields(line)

		now, _ := strconv.Atoi(fields[0])
		id, _ := strconv.Atoi(fields[1])

		left := now - k + 1

		for head < len(queue) && queue[head].time < left {
			old := queue[head]
			head++

			before := cnt[old.id]

			if before == limit {
				busy--
			}

			if before == 1 {
				active--
				delete(cnt, old.id)
			} else {
				cnt[old.id] = before - 1
			}
		}

		before := cnt[id]

		if before == 0 {
			active++
		}

		cnt[id] = before + 1

		if before+1 == limit {
			busy++
		}

		queue = append(queue, Order{time: now, id: id})

		writer.WriteString(strconv.Itoa(busy))
		writer.WriteByte(' ')
		writer.WriteString(strconv.Itoa(active))
		writer.WriteByte('\n')
	}
}
