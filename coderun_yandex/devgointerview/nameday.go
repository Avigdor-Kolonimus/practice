package devgointerview

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func removeChar(s string, pos int) string {
	return s[:pos] + s[pos+1:]
}

// https://coderun.yandex.ru/selections/dev-go-interview/problems/name-day
// NameDay - problem 7
func NameDay() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N and M input
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
	m, err := strconv.Atoi(strNum[1])
	if err != nil {
		panic(err)
	}

	// person celebrating their birthday input
	birthdays := make([]string, n)
	birthdaySet := make(map[string]struct{}, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}

		birthdays[i] = strings.TrimRight(line, "\r\n")
		birthdaySet[birthdays[i]] = struct{}{}
	}

	// taxi users input
	taxiUsers := make([]string, m)
	taxiSet := make(map[string]struct{}, m)
	for i := 0; i < m; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}

		taxiUsers[i] = strings.TrimRight(line, "\r\n")
		taxiSet[taxiUsers[i]] = struct{}{}
	}

	answer := make(map[string]struct{})

	// Exact match + removal of character from user
	for _, user := range taxiUsers {
		if _, ok := birthdaySet[user]; ok {
			answer[user] = struct{}{}
			continue
		}

		for i := 0; i < len(user); i++ {
			candidate := removeChar(user, i)
			if _, ok := birthdaySet[candidate]; ok {
				answer[user] = struct{}{}

				break
			}
		}
	}

	// Adding a symbol to the user
	// Equivalent to removing the symbol from the birthday person
	for _, birthday := range birthdays {
		for i := 0; i < len(birthday); i++ {
			candidate := removeChar(birthday, i)
			if _, ok := taxiSet[candidate]; ok {
				answer[candidate] = struct{}{}
			}
		}
	}

	result := make([]string, 0, len(answer))
	for surname := range answer {
		result = append(result, surname)
	}
	sort.Strings(result)

	writer.WriteString(strconv.Itoa(len(result)))
	writer.WriteByte('\n')
	for _, surname := range result {
		writer.WriteString(surname)
		writer.WriteByte('\n')
	}
}
