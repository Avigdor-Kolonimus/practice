package backend

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type TemplatePhoneFormatting struct {
	countryCode  string
	operatorCode string
	mask         string
	country      string
	operator     string
}

func digitsOnly(s string) string {
	var b strings.Builder
	for _, r := range s {
		if r >= '0' && r <= '9' {
			b.WriteRune(r)
		}
	}

	return b.String()
}

func matchMask(mask, number string) bool {
	if len(mask) != len(number) {
		return false
	}

	for i := 0; i < len(mask); i++ {
		if mask[i] == 'X' {
			continue
		}
		if mask[i] != number[i] {
			return false
		}
	}

	return true
}

// https://coderun.yandex.ru/selections/backend/problems/phone-formatting
// PhoneFormatting - problem 23
func PhoneFormatting() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// phones input
	phones := make([]string, n)
	for i := 0; i < n; i++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}

		phones[i] = strings.TrimSpace(line)
	}

	// M input
	line, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	m, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// templates input
	templates := make(map[string][]TemplatePhoneFormatting)
	for i := 0; i < m; i++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimSpace(line)

		parts := strings.Split(line, " - ")

		numberPart := parts[0]

		rest := strings.Fields(parts[1])
		country := rest[0]
		operator := rest[1]

		pos1 := strings.Index(numberPart, "+")
		pos2 := strings.Index(numberPart, "(")
		pos3 := strings.Index(numberPart, ")")

		countryCode := numberPart[pos1+1 : pos2-1]
		operatorCode := numberPart[pos2+1 : pos3]

		mask := strings.TrimSpace(numberPart[pos3+1:])

		t := TemplatePhoneFormatting{
			countryCode:  countryCode,
			operatorCode: operatorCode,
			mask:         mask,
			country:      country,
			operator:     operator,
		}

		key := countryCode + "#" + operatorCode
		templates[key] = append(templates[key], t)
	}

	for _, raw := range phones {
		number := digitsOnly(raw)

		found := false
		for countryLen := 1; countryLen <= 3 && !found; countryLen++ {
			if len(number) <= countryLen {
				continue
			}

			countryCode := number[:countryLen]

			for operatorLen := 2; operatorLen <= 4 && !found; operatorLen++ {

				if len(number) <= countryLen+operatorLen {
					continue
				}

				operatorCode := number[countryLen : countryLen+operatorLen]
				personal := number[countryLen+operatorLen:]

				key := countryCode + "#" + operatorCode

				candidates, ok := templates[key]
				if !ok {
					continue
				}

				for _, t := range candidates {
					if matchMask(t.mask, personal) {
						writer.WriteByte('+')
						writer.WriteString(t.countryCode)
						writer.WriteString(" (")
						writer.WriteString(t.operatorCode)
						writer.WriteString(") ")
						writer.WriteString(personal)
						writer.WriteString(" - ")
						writer.WriteString(t.country)
						writer.WriteByte(' ')
						writer.WriteString(t.operator)
						writer.WriteByte('\n')
						found = true

						break
					}
				}
			}
		}
	}
}
