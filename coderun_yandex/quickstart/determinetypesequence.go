package quickstart

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	constSequence   = "CONSTANT"
	ascSequence     = "ASCENDING"
	weakAscSequence = "WEAKLY ASCENDING"
	desSequence     = "DESCENDING"
	weakDesSequence = "WEAKLY DESCENDING"
	rndSequence     = "RANDOM"
)

func validateDetermineTypeSequenceStopInput(p int) bool {
	return p == -2_000_000_000
}

// https://coderun.yandex.ru/selections/quickstart/problems/determine-type-sequence
// DetermineTypeSequence - assignment 15
func DetermineTypeSequence() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")
	prev, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	isConst := true
	isAsc := true
	isWeakAsc := true
	isDesc := true
	isWeakDesc := true
	for {
		// input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		cur, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		if validateDetermineTypeSequenceStopInput(cur) {
			break
		}

		if cur > prev {
			isConst = false
			isDesc = false
			isWeakDesc = false
		} else if cur < prev {
			isConst = false
			isAsc = false
			isWeakAsc = false
		} else { // cur == prev
			isAsc = false
			isDesc = false
		}

		prev = cur
	}

	result := rndSequence
	switch {
	case isConst:
		result = constSequence
	case isAsc:
		result = ascSequence
	case isWeakAsc:
		result = weakAscSequence
	case isDesc:
		result = desSequence
	case isWeakDesc:
		result = weakDesSequence
	}

	writer.WriteString(result)
	writer.WriteByte('\n')
}
