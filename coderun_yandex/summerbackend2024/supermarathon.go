package summerbackend2024

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func parseTime(s string) int64 {
	var h, m int64
	fmt.Sscanf(s, "%d:%d", &h, &m)

	return h*60 + m
}

func gcdSuperMarathon(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}

	if a < 0 {
		return -a
	}

	return a
}

func extGCDSuperMarathon(a, b int64) (g, x, y int64) {
	if b == 0 {
		return a, 1, 0
	}
	g, x1, y1 := extGCDSuperMarathon(b, a%b)

	return g, y1, x1 - (a/b)*y1
}

func modSuperMarathon(a, m int64) int64 {
	a %= m
	if a < 0 {
		a += m
	}

	return a
}

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/super-marathon
// SuperMarathon - problem 3
func SuperMarathon() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// Andrei start input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	sA := parseTime(line)

	// Boris start input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	sB := parseTime(line)

	// Andrei pace input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	p := parseTime(line)

	// Boris pace input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	q := parseTime(line)

	g := gcdSuperMarathon(p, q)
	diff := sB - sA

	if diff%g != 0 {
		writer.WriteString("Never")
		writer.WriteByte('\n')

		return
	}

	// CRT:
	// sA + p*k ≡ sB (mod q)
	// p*k ≡ diff (mod q)
	p1 := p / g
	q1 := q / g
	diff1 := diff / g

	_, inv, _ := extGCDSuperMarathon(p1, q1)
	inv = modSuperMarathon(inv, q1)

	k0 := modSuperMarathon(diff1*inv, q1)

	x0 := sA + p*k0
	lcm := p / g * q

	x0 = modSuperMarathon(x0, lcm)

	threshold := sA + p
	if sB+q > threshold {
		threshold = sB + q
	}

	var t int64
	if x0 >= threshold {
		t = x0
	} else {
		steps := (threshold - x0 + lcm - 1) / lcm
		t = x0 + steps*lcm
	}

	weekdays := []string{
		"Saturday",
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
	}

	day := (t / 1440) % 7
	timeOfDay := t % 1440

	h := timeOfDay / 60
	m := timeOfDay % 60

	writer.WriteString(weekdays[day])
	writer.WriteByte('\n')

	ans := fmt.Sprintf("%02d:%02d\n", h, m)
	writer.WriteString(ans)
	writer.WriteByte('\n')
}
