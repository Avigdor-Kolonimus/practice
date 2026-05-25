package summercommon2025

type Answer struct {
	Sum   int
	Field []string
}

// https://coderun.yandex.ru/selections/2025-summer-common/problems/playing-field
// PlayingField - problem 14
func SolutionPlayingField(n int) Answer {
	totalSum := (n - 1) * (3*n - 2)

	row := make([]byte, n)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			row[i] = 'x'
		} else {
			row[i] = '-'
		}
	}

	field := make([]string, n)

	s := string(row)

	for i := 0; i < n; i++ {
		field[i] = s
	}

	return Answer{
		Sum:   totalSum,
		Field: field,
	}
}
