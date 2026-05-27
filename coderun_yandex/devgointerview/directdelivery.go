package devgointerview

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://coderun.yandex.ru/selections/dev-go-interview/problems/direct-delivery
// DirectDelivery - problem 1
func DirectDelivery() {
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

	// coordinates input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	strNum = strings.Fields(line)
	if len(strNum) != n {
		panic("numbers count does not match n")
	}

	posCoordinates := make([]int, 0, n)
	negCoordinates := make([]int, 0, n)
	for i := 0; i < n; i++ {
		c, err := strconv.Atoi(strNum[i])
		if err != nil {
			panic(err)
		}

		if c > 0 {
			posCoordinates = append(posCoordinates, c)
		} else if c < 0 {
			negCoordinates = append(negCoordinates, -c)
		}
	}

	sort.Ints(posCoordinates)
	sort.Ints(negCoordinates)

	maxDist := 0

	if len(posCoordinates) > 0 {
		maxDist = max(maxDist, posCoordinates[len(posCoordinates)-1])
	}
	if len(negCoordinates) > 0 {
		maxDist = max(maxDist, negCoordinates[len(negCoordinates)-1])
	}

	ans := 0
	for i := len(posCoordinates) - 1; i >= 0; i -= k {
		ans += 2 * posCoordinates[i]
	}

	for i := len(negCoordinates) - 1; i >= 0; i -= k {
		ans += 2 * negCoordinates[i]
	}

	ans -= maxDist

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
