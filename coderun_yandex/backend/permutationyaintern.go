package backend

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/backend/problems/permutation-ya-intern
// PermutationYaIntern - problem 39
func PermutationYaIntern() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// sum inputs
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 3 {
		panic("numbers count does not match 3")
	}

	sm, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	sm2, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	sm3, err := strconv.Atoi(strNum[2])
	if err != nil {
		panic(err)
	}

	// totals
	total1 := n * (n + 1) / 2
	total2 := n * (n + 1) * (2*n + 1) / 6
	total3 := total1 * total1

	p1 := total1 - sm
	p2 := total2 - sm2
	p3 := total3 - sm3

	e1 := p1
	e2 := (p1*p1 - p2) / 2
	e3 := (p3 - e1*p2 + e2*e1) / 3

	x := 0
	y := 0
	z := 0
	for t := 1; t <= n; t++ {
		val := t*t*t - e1*t*t + e2*t - e3
		if val != 0 {
			continue
		}

		// y + z
		sumYZ := e1 - t

		// yz
		mulYZ := e3 / t

		// t^2 - sumYZ*t + mulYZ = 0
		D := sumYZ*sumYZ - 4*mulYZ
		if D < 0 {
			continue
		}

		sqrtD := int(math.Sqrt(float64(D)))
		if sqrtD*sqrtD != D {
			continue
		}

		x = t
		y = (sumYZ + sqrtD) / 2
		z = (sumYZ - sqrtD) / 2

		if y > 0 && z > 0 && y <= n && z <= n && x != y && x != z && y != z {
			break
		}
	}

	writer.WriteString(strconv.Itoa(x) + " " + strconv.Itoa(y) + " " + strconv.Itoa(z))
	writer.WriteByte('\n')
}
