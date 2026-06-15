package algorithmtrainingmarch2026

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type BestVasyasUnluckyNumber struct {
	val int
	mod int
}

// https://coderun.yandex.ru/selections/algorithm-training-march-2026/problems/vasyas-unlucky-number
// VasyasUnluckyNumber - assignment 18
func VasyasUnluckyNumber() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and K input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 2 {
		panic("numbers count does not match 2")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	k, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// values input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum = strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	values := make([]int, n)
	for i := 0; i < n; i++ {
		values[i], err = strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}
	}

	minPref := make([]int, k)
	for i := range minPref {
		minPref[i] = math.MaxInt
	}

	// pref[0] = 0
	minPref[0] = 0

	best1 := BestVasyasUnluckyNumber{0, 0}
	best2 := BestVasyasUnluckyNumber{math.MaxInt, -1}

	ans, pref := 0, 0
	for i := 0; i < n; i++ {
		pref += values[i]

		mod := ((pref % k) + k) % k

		minOther := math.MaxInt
		if best1.mod != mod {
			minOther = best1.val
		} else {
			minOther = best2.val
		}

		if minOther != math.MaxInt {
			cur := pref - minOther
			if cur > ans {
				ans = cur
			}
		}

		if pref < minPref[mod] {
			minPref[mod] = pref

			switch mod {
			case best1.mod:
				best1.val = pref
			case best2.mod:
				best2.val = pref
				if best2.val < best1.val {
					best1, best2 = best2, best1
				}
			default:
				if pref < best1.val {
					best2 = best1
					best1 = BestVasyasUnluckyNumber{pref, mod}
				} else if pref < best2.val {
					best2 = BestVasyasUnluckyNumber{pref, mod}
				}
			}
		}
	}

	if ans < 0 {
		ans = 0
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
