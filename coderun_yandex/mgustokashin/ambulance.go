package mgustokashin

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/mgustokashin/problems/ambulance
// Ambulance - problem 1
func Ambulance() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	fields := strings.Fields(line)

	k1, _ := strconv.Atoi(fields[0])
	m, _ := strconv.Atoi(fields[1])
	k2, _ := strconv.Atoi(fields[2])
	p2, _ := strconv.Atoi(fields[3])
	n2, _ := strconv.Atoi(fields[4])

	found := false
	p1Ans, n1Ans := -2, -2
	for x := 1; x <= 1_000_000; x++ {
		perEntrance := m * x

		p := (k2-1)/perEntrance + 1
		inside := (k2 - 1) % perEntrance
		n := inside/x + 1

		if p != p2 || n != n2 {
			continue
		}

		found = true

		p1 := (k1-1)/perEntrance + 1
		inside1 := (k1 - 1) % perEntrance
		n1 := inside1/x + 1

		if p1Ans == -2 {
			p1Ans = p1
		} else if p1Ans != p1 {
			p1Ans = 0
		}

		if n1Ans == -2 {
			n1Ans = n1
		} else if n1Ans != n1 {
			n1Ans = 0
		}
	}

	if !found {
		writer.WriteString("-1 -1")
		writer.WriteByte('\n')

		return
	}

	writer.WriteString(strconv.Itoa(p1Ans))
	writer.WriteByte(' ')
	writer.WriteString(strconv.Itoa(n1Ans))
	writer.WriteByte('\n')
}
