package devgointerview

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	task int
	time int
}

// https://coderun.yandex.ru/selections/dev-go-interview/problems/ctf-team-score
// CTFTeamScore - problem 12
func CTFTeamScore() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	// tasks info input
	a := make([]Task, n)
	for i := range n {
		line, _ = reader.ReadString('\n')
		p := strings.Fields(line)

		task, _ := strconv.Atoi(p[0])
		time, _ := strconv.Atoi(p[1])

		a[i] = Task{task, time}
	}

	// M input
	line, _ = reader.ReadString('\n')
	m, _ := strconv.Atoi(strings.TrimSpace(line))

	// tasks info input
	b := make([]Task, m)
	for i := range m {
		line, _ = reader.ReadString('\n')
		p := strings.Fields(line)

		task, _ := strconv.Atoi(p[0])
		time, _ := strconv.Atoi(p[1])

		b[i] = Task{task, time}
	}

	solvedA := make(map[int]bool)
	solvedB := make(map[int]bool)

	pa, pb := 0, 0
	scoreA, scoreB := 0, 0
	for pa < n || pb < m {
		// same time and same task
		if pa < n && pb < m &&
			a[pa].time == b[pb].time &&
			a[pa].task == b[pb].task {

			solvedA[a[pa].task] = true
			solvedB[b[pb].task] = true

			pa++
			pb++
			continue
		}

		// Eda's turn
		if pb == m || (pa < n && a[pa].time < b[pb].time) {
			if !solvedB[a[pa].task] {
				scoreA++
			}

			solvedA[a[pa].task] = true
			pa++

		} else {
			// Lavka's turn
			if !solvedA[b[pb].task] {
				scoreB++
			}

			solvedB[b[pb].task] = true
			pb++
		}
	}

	writer.WriteString(strconv.Itoa(scoreA))
	writer.WriteByte(' ')
	writer.WriteString(strconv.Itoa(scoreB))
	writer.WriteByte('\n')
}
