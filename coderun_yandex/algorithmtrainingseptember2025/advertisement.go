package algorithmtrainingseptember2025

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type WordAdvertisement struct {
	a float64
	b float64
}

func canAdvertisement(words []WordAdvertisement, n int, W, H, k float64) bool {
	totalH := 0.0

	curB := -1.0
	curW := 0.0

	for i := 0; i < n; i++ {
		wordW := k * words[i].a
		wordH := k * words[i].b

		if wordW > W || wordH > H {
			return false
		}

		if curB == words[i].b && curW+wordW <= W {
			curW += wordW
		} else {
			if curB != -1 {
				totalH += k * curB
			}

			curB = words[i].b
			curW = wordW
		}
	}

	if curB != -1 {
		totalH += k * curB
	}

	return totalH <= H
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/advertisement
// Advertisement - assignment 23
func Advertisement() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N, W and H input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum := strings.Fields(line)
	if len(strNum) != 3 {
		panic("numbers count does not match 3")
	}

	n, err := strconv.Atoi(strNum[0])
	if err != nil {
		panic(err)
	}
	w, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}
	h, err := strconv.Atoi(strNum[2])
	if err != nil {
		panic(err)
	}
	W, H := float64(w), float64(h)

	// word input
	words := make([]WordAdvertisement, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum = strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		a, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		words[i] = WordAdvertisement{float64(a), float64(b)}
	}

	l, r := 0.0, 1e9
	for it := 0; it < 80; it++ {
		mid := (l + r) / 2

		if canAdvertisement(words, n, W, H, mid) {
			l = mid
		} else {
			r = mid
		}
	}

	writer.WriteString(strconv.FormatFloat(l, 'f', 10, 64))
	writer.WriteByte('\n')
}
