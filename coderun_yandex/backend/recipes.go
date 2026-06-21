package backend

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var unitToBase = map[string]float64{
	"g":    1,
	"kg":   1000,
	"ml":   1,
	"l":    1000,
	"cnt":  1,
	"tens": 10,
}

func toBase(amount float64, unit string) float64 {
	return amount * unitToBase[unit]
}

func convert(amount float64, from, to string) float64 {
	base := toBase(amount, from)
	return base / unitToBase[to]
}

type Quantity struct {
	amount float64
	unit   string
}

type Cost struct {
	price    int
	quantity Quantity
}

type Nutritional struct {
	quantity Quantity
	protein  float64
	fats     float64
	carbs    float64
	cal      float64
}

type Dish struct {
	name        string
	friends     int
	ingredients map[string]Quantity
}

func validateRecipesInput(n int) bool {
	return n >= 1 && n <= 1_000
}

// https://coderun.yandex.ru/selections/backend/problems/recipes
// Recipes - problem 10
func Recipes() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// N input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	n, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateRecipesInput(n) {
		panic("number N out of range")
	}

	dishes := make([]Dish, n)
	for i := 0; i < n; i++ {
		// input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		strNum := strings.Fields(line)
		if len(strNum) != 3 {
			panic("numbers count does not match 3")
		}

		name := strNum[0]

		c, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		z, err := strconv.Atoi(strNum[2])
		if err != nil {
			panic(err)
		}

		d := Dish{
			name:        name,
			friends:     c,
			ingredients: make(map[string]Quantity),
		}

		for j := 0; j < z; j++ {
			// ingredients input
			line, err = reader.ReadString('\n')
			if err != nil && err != io.EOF {
				panic(err)
			}
			line = strings.TrimRight(line, "\r\n")

			strNum = strings.Fields(line)
			if len(strNum) != 3 {
				panic("numbers count does not match 3")
			}

			ing := strNum[0]

			a, err := strconv.ParseFloat(strNum[1], 64)
			if err != nil {
				panic(err)
			}

			u := strNum[2]

			d.ingredients[ing] = Quantity{a, u}
		}

		dishes[i] = d
	}

	// ---- COST CATALOG ----
	// K input
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	k, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateRecipesInput(k) {
		panic("number K out of range")
	}

	costs := make(map[string]Cost)

	for i := 0; i < k; i++ {
		// input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		strNum := strings.Fields(line)
		if len(strNum) != 4 {
			panic("numbers count does not match 4")
		}

		name := strNum[0]

		price, err := strconv.Atoi(strNum[1])
		if err != nil {
			panic(err)
		}

		amount, err := strconv.ParseFloat(strNum[2], 64)
		if err != nil {
			panic(err)
		}

		unit := strNum[3]

		costs[name] = Cost{price, Quantity{amount, unit}}
	}

	// ---- NUTRITION ----
	// M input
	line, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimRight(line, "\r\n")

	m, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	if !validateRecipesInput(m) {
		panic("number M out of range")
	}

	nutrition := make(map[string]Nutritional)

	for i := 0; i < m; i++ {
		// input
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimRight(line, "\r\n")

		strNum := strings.Fields(line)
		if len(strNum) != 7 {
			panic("numbers count does not match 7")
		}

		name := strNum[0]

		amount, err := strconv.ParseFloat(strNum[1], 64)
		if err != nil {
			panic(err)
		}

		unit := strNum[2]

		pr, err := strconv.ParseFloat(strNum[3], 64)
		if err != nil {
			panic(err)
		}

		f, err := strconv.ParseFloat(strNum[4], 64)
		if err != nil {
			panic(err)
		}

		ch, err := strconv.ParseFloat(strNum[5], 64)
		if err != nil {
			panic(err)
		}

		cal, err := strconv.ParseFloat(strNum[6], 64)
		if err != nil {
			panic(err)
		}

		nutrition[name] = Nutritional{
			Quantity{amount, unit},
			pr, f, ch, cal,
		}
	}

	// ---- TOTAL INGREDIENTS ----
	total := make(map[string]float64)

	for _, d := range dishes {
		for ing, q := range d.ingredients {
			total[ing] += toBase(q.amount*float64(d.friends), q.unit)
		}
	}

	// ---- SHOPPING ----
	totalCost := 0
	packages := make(map[string]int)

	for name, cost := range costs {
		need := total[name]

		packSize := toBase(cost.quantity.amount, cost.quantity.unit)

		cnt := int(math.Ceil(need / packSize))
		packages[name] = cnt

		totalCost += cnt * cost.price
	}

	writer.WriteString(strconv.Itoa(totalCost))
	writer.WriteByte('\n')

	for name := range costs {
		writer.WriteString(name)
		writer.WriteByte(' ')
		writer.WriteString(strconv.Itoa(packages[name]))
		writer.WriteByte('\n')
	}

	// ---- NUTRITION PER DISH ----
	for _, d := range dishes {
		var pr, f, ch, cal float64

		for ing, q := range d.ingredients {
			info, ok := nutrition[ing]
			if !ok {
				continue
			}

			converted := convert(q.amount, q.unit, info.quantity.unit)
			k := converted / info.quantity.amount

			pr += info.protein * k
			f += info.fats * k
			ch += info.carbs * k
			cal += info.cal * k
		}

		line = fmt.Sprintf("%s %.6f %.6f %.6f %.6f\n", d.name, pr, f, ch, cal)
		writer.WriteString(line)
		writer.WriteByte('\n')
	}
}
