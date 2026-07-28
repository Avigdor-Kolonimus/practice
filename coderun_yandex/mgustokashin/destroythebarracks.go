package mgustokashin

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/mgustokashin/problems/destroy-the-barracks
// DestroyTheBarracks - problem 9
func DestroyTheBarracks() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// X input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	x1, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// Y input
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	y1, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// P input
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	p, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	ans := math.MaxInt
	for i := 0; i <= y1; i++ {
		round := 0
		x := x1
		y := y1
		solderCount := 0

		valid := true
		for y > i {
			round++
			if solderCount >= x {
				valid = false
				break
			}

			y -= (x - solderCount)
			solderCount = 0
			if y > 0 {
				solderCount += p
			}
		}

		if !valid {
			continue
		}

		success := false
		for x > 0 {
			if solderCount <= 0 && y <= 0 {
				success = true
				break
			}
			round++
			temp := y
			y = max(0, y-x)
			solderCount = max(0, solderCount-max(0, x-temp))

			x = x - solderCount

			if y > 0 {
				solderCount += p
			}
		}

		if success {
			ans = min(ans, round)
		}
	}

	if ans == math.MaxInt {
		ans = -1
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
