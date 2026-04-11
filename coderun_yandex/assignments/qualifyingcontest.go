package assignments

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func validateNandKInput(n, k int) bool {
	if n < k {
		return false
	}

	if k < 1 {
		return false
	}

	if n > 100_000 {
		return false
	}

	return true
}

func validateTaskInput(task int) bool {
	return task >= 1 && task <= 100_000
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/qualifying-contest
// QualifyingContest - assignment 4
func QualifyingContest() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// first line
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	tmpNandK := strings.Fields(line)
	if len(tmpNandK) != 2 {
		panic("first line is not correct")
	}

	n, err := strconv.Atoi(tmpNandK[0])
	if err != nil {
		panic(err)
	}

	k, err := strconv.Atoi(tmpNandK[1])
	if err != nil {
		panic(err)
	}

	if !validateNandKInput(n, k) {
		panic("n or k out of range")
	}

	// second line
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	tasks := strings.Fields(line)
	if len(tasks) < k {
		panic("few tasks")
	}

	uniqueThemesCount := make(map[int]int)
	for _, t := range tasks {
		task, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}

		if !validateTaskInput(task) {
			panic("tasks out of range")
		}

		uniqueThemesCount[task]++
	}

	// sort tasks
	keys := make([]int, 0, len(uniqueThemesCount))
	for k := range uniqueThemesCount {
		keys = append(keys, k)
	}

	// sort keys
	slices.Sort(keys)

	results := make([]int, 0, k)
	for _, key := range keys {
		if uniqueThemesCount[key] > 0 {
			uniqueThemesCount[key]--
			results = append(results, key)
			k--
		}

		if k < 1 {
			break
		}
	}

	for _, key := range keys {
		for uniqueThemesCount[key] > 0 {
			if uniqueThemesCount[key] > 0 {
				uniqueThemesCount[key]--
				results = append(results, key)
				k--
			}
			if k < 1 {
				break
			}
		}

		if k < 1 {
			break
		}
	}

	// output
	for i, result := range results {
		_, err = fmt.Fprintf(writer, "%d", result)
		if err != nil {
			panic(err)
		}

		if i < len(results)-1 {
			err = writer.WriteByte(' ')
			if err != nil {
				panic(err)
			}
		}
	}
}

func QualifyingContest2() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var tasksCount, perContestCount int

	_, err := fmt.Fscanf(reader, "%d %d\n", &tasksCount, &perContestCount)
	if err != nil {
		panic(err)
	}

	if perContestCount > tasksCount {
		panic("invalid input")
	}

	tasks := make([]int, tasksCount)

	for i := range tasks {
		_, err = fmt.Fscanf(reader, "%d", &tasks[i])
		if err != nil {
			panic(err)
		}

		_, err = reader.ReadByte()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			panic(err)
		}
	}

	slices.Sort(tasks)

	uniqueThemesCount := make([]int, 0, len(tasks))
	themesValues := make([]int, 0, len(tasks))

	themesValues = append(themesValues, tasks[0])
	uniqueThemesCount = append(uniqueThemesCount, 1)

	for i := range tasks[1:] {
		i++

		if tasks[i] == themesValues[len(themesValues)-1] {
			uniqueThemesCount[len(uniqueThemesCount)-1]++

			continue
		}

		themesValues = append(themesValues, tasks[i])
		uniqueThemesCount = append(uniqueThemesCount, 1)
	}

	results := make([]int, 0, perContestCount)

	for i := range uniqueThemesCount {
		if len(results) == perContestCount {
			break
		}

		uniqueThemesCount[i]--
		results = append(results, themesValues[i])
	}

Loop:
	for i := range uniqueThemesCount {
		for range uniqueThemesCount[i] {
			if len(results) == perContestCount {
				break Loop
			}

			results = append(results, themesValues[i])
		}
	}

	for i, result := range results {
		_, err = fmt.Fprintf(writer, "%d", result)
		if err != nil {
			panic(err)
		}

		if i < len(results)-1 {
			err = writer.WriteByte(' ')
			if err != nil {
				panic(err)
			}
		}
	}
}
