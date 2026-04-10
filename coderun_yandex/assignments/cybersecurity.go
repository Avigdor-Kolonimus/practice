package assignments

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validatePasswordInput(password string) bool {
	passwordLength := len(password)

	return passwordLength >= 1 && passwordLength <= 100_000
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/cybersecurity
// Cybersecurity - assignment 3
func Cybersecurity() {
	uniquePassword := 0
	duplicateSwaps := 0
	totalSwaps := 0
	freq := make(map[rune]int)

	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	password := strings.TrimRight(line, "\r\n")
	if !validatePasswordInput(password) {
		panic("password length out of range")
	}

	for _, char := range password {
		freq[char]++
	}

	for _, count := range freq {
		if count > 1 {
			duplicateSwaps += count * (count - 1) / 2
		}
	}

	passwordLenght := len(password)
	totalSwaps = passwordLenght * (passwordLenght - 1) / 2
	uniquePassword = 1 + (totalSwaps - duplicateSwaps)

	// output
	writer.WriteString(strconv.Itoa(uniquePassword))
	writer.WriteByte('\n')
}
