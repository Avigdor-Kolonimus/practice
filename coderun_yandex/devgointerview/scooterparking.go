package devgointerview

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type EventScooterParking struct {
	time   int
	factor int // 1 - return, -1 - take
}

// https://coderun.yandex.ru/selections/dev-go-interview/problems/scooter-parking
// ScooterParking - problem 9
func ScooterParking() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	// Event input
	events := make([]EventScooterParking, 0, 2*n)
	for i := 0; i < n; i++ {
		line, _ = reader.ReadString('\n')
		parts := strings.Fields(line)

		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		events = append(events, EventScooterParking{start, 1})
		events = append(events, EventScooterParking{end, -1})
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].time == events[j].time {
			return events[i].factor < events[j].factor
		}
		return events[i].time < events[j].time
	})

	minPlaces := math.MinInt
	minScooters := 0

	current, currentTime, currentFactor, parked := 0, -1, 0, 0
	for _, event := range events {
		if currentTime == -1 || event.time > currentTime {
			parked += currentFactor
			currentTime = event.time
			currentFactor = 0

			if parked > minPlaces {
				minPlaces = parked
			}
		}

		currentFactor -= event.factor
		current -= event.factor

		if current < minScooters {
			minScooters = current
		}
	}

	writer.WriteString(strconv.Itoa(minPlaces - minScooters))
	writer.WriteByte(' ')
	writer.WriteString(strconv.Itoa(-minScooters))
	writer.WriteByte('\n')
}
