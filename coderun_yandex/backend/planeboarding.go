package backend

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Passenger struct {
	a      int
	row    int
	seat   byte
	pos    int
	seated bool
	doneAt int
}

func needToMove(seat byte) string {
	switch seat {
	case 'A':
		return "BC"
	case 'B':
		return "C"
	case 'C':
		return ""
	case 'D':
		return ""
	case 'E':
		return "D"
	case 'F':
		return "DE"
	}

	return ""
}

// https://coderun.yandex.ru/selections/backend/problems/plane-boarding
// PlaneBoarding - problem 51
func PlaneBoarding() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// Passengers input
	p := make([]Passenger, n)
	for i := range n {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum := strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		a, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		place := strNum[1]

		seat := place[len(place)-1]
		row, _ := strconv.Atoi(place[:len(place)-1])

		p[i] = Passenger{
			a:    a,
			row:  row,
			seat: seat,
			pos:  0,
		}
	}

	finish := make([][]int, 100_000)
	seatedCount, answer := 0, 0
	for t := 0; seatedCount < n; t++ {
		for _, id := range finish[t] {
			p[id].seated = true
			p[id].doneAt = t
			seatedCount++

			if t > answer {
				answer = t
			}
		}

		occupied := make([]int, 31)
		for i := range occupied {
			occupied[i] = -1
		}

		for i := range n {
			if !p[i].seated && p[i].pos > 0 {
				occupied[p[i].pos] = i
			}
		}

		target := make([]int, 31)
		for i := range target {
			target[i] = -1
		}

		for i := range n {
			if p[i].seated {
				continue
			}

			next := p[i].pos + 1

			if next > p[i].row {
				continue
			}

			if target[next] == -1 {
				target[next] = i
			}
		}

		moving := make([]bool, n)
		for i := range n {
			if p[i].seated {
				continue
			}

			next := p[i].pos + 1

			if next > p[i].row || target[next] != i {
				continue
			}

			if occupied[next] == -1 ||
				moving[occupied[next]] {

				moving[i] = true
			}
		}

		for i := range n {
			if !moving[i] {
				continue
			}

			next := p[i].pos + 1
			p[i].pos = next

			if next == p[i].row {
				need := needToMove(p[i].seat)
				cnt := 0
				for j := range n {
					if !p[j].seated {
						continue
					}

					if p[j].row != p[i].row {
						continue
					}

					for k := range len(need) {
						if p[j].seat == need[k] {
							cnt++
						}
					}
				}

				extra := 0
				if cnt == 1 {
					extra = 5
				} else if cnt >= 2 {
					extra = 15
				}

				doneAt := t + 1 + p[i].a + extra

				finish[doneAt] = append(finish[doneAt], i)
			}
		}
	}

	writer.WriteString(strconv.Itoa(answer))
	writer.WriteByte('\n')
}
