package summerbackend2024

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateTableauNInput(n int) bool {
	return n >= 1 && n <= 50
}

func validateTableauMInput(n int) bool {
	return n >= 1 && n <= 500
}

// https://coderun.yandex.ru/selections/2024-summer-backend/problems/tableau
// Tableau - problem 30
func Tableau() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line := mustReadIntArray(reader, 1)
	if !validateTableauNInput(line[0]) {
		panic("number N out of range")
	}
	n := line[0]

	// player inputs
	player := make(map[string]int, n)
	for range n {
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		name := strings.TrimSpace(line)

		player[name] = 0
	}

	// M input
	line = mustReadIntArray(reader, 1)
	if !validateTableauMInput(line[0]) {
		panic("number M out of range")
	}
	m := line[0]

	bestPlayer, bestPoints := "", 0
	team1, team2 := 0, 0
	for range m {
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		line = strings.TrimSpace(line)

		ss1 := strings.Split(line, ":")
		ss2 := strings.Split(ss1[1], " ")

		p1, err := strconv.Atoi(ss1[0])
		if err != nil {
			panic(err)
		}
		p2, err := strconv.Atoi(ss2[0])
		if err != nil {
			panic(err)
		}
		name := ss2[1]

		if p1-team1 > 0 {
			player[name] += p1 - team1
			team1 = p1
		} else if p2-team2 > 0 {
			player[name] += p2 - team2
			team2 = p2
		}

		if player[name] > bestPoints || (player[name] == bestPoints && name > bestPlayer) {
			bestPlayer, bestPoints = name, player[name]
		}
	}

	writer.WriteString(bestPlayer + " " + strconv.Itoa(bestPoints))
	writer.WriteByte('\n')
}
