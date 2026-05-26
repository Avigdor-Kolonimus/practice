package algorithmtrainingseptember2025

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type TaxeRate struct {
	Power int
	Rate  int
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/transport-taxes
// TransportTaxes - assignment 32
func TransportTaxes() {
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

	// taxes input
	taxeRates := make([]TaxeRate, n)
	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")
		strNum := strings.Fields(line)
		if len(strNum) != 2 {
			panic("numbers count does not match 2")
		}

		p, err := strconv.Atoi(strNum[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		taxeRates[i] = TaxeRate{p, r}
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

	// car input
	for i := 0; i < m; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		q, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		// first index with taxeRates[idx].Power >= q
		idx := sort.Search(n, func(j int) bool {
			return taxeRates[j].Power >= q
		})

		var rate int
		if idx == n {
			rate = taxeRates[n-1].Rate
		} else if taxeRates[idx].Power == q {
			if idx == 0 {
				rate = taxeRates[0].Rate
			} else {
				rate = taxeRates[idx-1].Rate
			}
		} else if idx == 0 {
			rate = taxeRates[0].Rate
		} else {
			rate = taxeRates[idx-1].Rate
		}

		ans := q * rate
		writer.WriteString(strconv.Itoa(ans))
		writer.WriteByte('\n')
	}
}
