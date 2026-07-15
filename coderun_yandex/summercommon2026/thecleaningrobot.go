package summercommon2026

import (
	"bufio"
	"fmt"
	"os"
)

type Interval struct {
	minK int
	maxK int
}

// Проверяет, инициализирован ли интервал (достижима ли клетка)
func (inv Interval) isValid() bool {
	return inv.minK <= inv.maxK && inv.minK != 1e9
}

// Слияние двух интервалов (объединение диапазонов)
func merge(a, b Interval) Interval {
	if !a.isValid() {
		return b
	}
	if !b.isValid() {
		return a
	}
	minK := a.minK
	if b.minK < minK {
		minK = b.minK
	}
	maxK := a.maxK
	if b.maxK > maxK {
		maxK = b.maxK
	}
	return Interval{minK: minK, maxK: maxK}
}

func solve(scanner *bufio.Scanner, writer *bufio.Writer) {
	if !scanner.Scan() {
		return
	}
	var n, m int
	fmt.Sscanf(scanner.Text(), "%d %d", &n, &m)

	grid := make([][]byte, n)
	for i := 0; i < n; i++ {
		if scanner.Scan() {
			grid[i] = append([]byte(nil), scanner.Bytes()...)
		}
	}

	// Предрасчет суффиксов грязи
	rowDirtySuff := make([][]int, n)
	for i := 0; i < n; i++ {
		rowDirtySuff[i] = make([]int, m)
		count := 0
		for j := m - 1; j >= 0; j-- {
			if grid[i][j] == '3' || grid[i][j] == '4' {
				count++
			}
			rowDirtySuff[i][j] = count
		}
	}

	colDirtySuff := make([][]int, n)
	for i := 0; i < n; i++ {
		colDirtySuff[i] = make([]int, m)
	}
	for j := 0; j < m; j++ {
		count := 0
		for i := n - 1; i >= 0; i-- {
			if grid[i][j] == '3' || grid[i][j] == '4' {
				count++
			}
			colDirtySuff[i][j] = count
		}
	}

	// Храним только текущую и следующую строку ДП для экономии памяти
	dpCurr := make([]Interval, m)
	dpNext := make([]Interval, m)

	// Инициализируем пустые интервалы
	clearRow := func(row []Interval) {
		for j := 0; j < m; j++ {
			row[j] = Interval{minK: 1e9, maxK: -1}
		}
	}
	clearRow(dpCurr)
	clearRow(dpNext)

	// Стартовая позиция (0,0)
	startK := 1
	if grid[0][0] == '1' || grid[0][0] == '3' {
		startK++
	}
	dpCurr[0] = Interval{minK: startK, maxK: startK}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			inv := dpCurr[j]
			if !inv.isValid() {
				continue
			}

			// 1. Движение ВПРАВО -> (i, j+1)
			if j+1 < m {
				// Фильтруем maxK по условию лазера для уходящего столбца j
				effMaxK := inv.maxK
				for effMaxK >= inv.minK {
					uncovered := 0
					if i+effMaxK < n {
						uncovered = colDirtySuff[i+effMaxK][j]
					}
					// Условие: uncovered <= curP  <=>  effMaxK + uncovered <= i + j + 2
					if effMaxK+uncovered <= i+j+2 {
						break
					}
					effMaxK--
				}

				// Если нашли валидный под-интервал
				if inv.minK <= effMaxK {
					nextKMin := inv.minK
					nextKMax := effMaxK
					if grid[i][j+1] == '1' || grid[i][j+1] == '3' {
						nextKMin++
						nextKMax++
					}
					dpCurr[j+1] = merge(dpCurr[j+1], Interval{minK: nextKMin, maxK: nextKMax})
				}
			}

			// 2. Движение ВНИЗ -> (i+1, j)
			if i+1 < n {
				// Фильтруем maxK по условию лазера для уходящей строки i
				effMaxK := inv.maxK
				for effMaxK >= inv.minK {
					uncovered := 0
					if j+effMaxK < m {
						uncovered = rowDirtySuff[i][j+effMaxK]
					}
					// Условие: uncovered <= curP  <=>  effMaxK + uncovered <= i + j + 2
					if effMaxK+uncovered <= i+j+2 {
						break
					}
					effMaxK--
				}

				if inv.minK <= effMaxK {
					nextKMin := inv.minK
					nextKMax := effMaxK
					if grid[i+1][j] == '1' || grid[i+1][j] == '3' {
						nextKMin++
						nextKMax++
					}
					dpNext[j] = merge(dpNext[j], Interval{minK: nextKMin, maxK: nextKMax})
				}
			}
		}

		// Переход к следующей строке
		if i+1 < n {
			copy(dpCurr, dpNext)
			clearRow(dpNext)
		}
	}

	// Финальная проверка в точке (n-1, m-1)
	possible := false
	finalInv := dpCurr[m-1]

	if finalInv.isValid() {
		// Проверяем, существует ли в финальном интервале хотя бы одно k,
		// при котором остаточная мощность лазера покроет хвосты последней строки и столбца
		for k := finalInv.maxK; k >= finalInv.minK; k-- {
			uncoveredRow := 0
			if m-1+k < m {
				uncoveredRow = rowDirtySuff[n-1][m-1+k]
			}
			uncoveredCol := 0
			if n-1+k < n {
				uncoveredCol = colDirtySuff[n-1+k][m-1]
			}

			// Проверка лазера для финальной клетки: k + uncovered <= n + m
			if k+uncoveredRow <= n+m && k+uncoveredCol <= n+m {
				possible = true
				break
			}
		}
	}

	if possible {
		writer.WriteString("Yes\n")
	} else {
		writer.WriteString("No\n")
	}
}

// https://coderun.yandex.ru/seasons/2026-summer/tracks/common/problem/the-cleaning-robot
// TheCleaningRobot - problem 15
func TheCleaningRobot() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 24 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	if !scanner.Scan() {
		return
	}
	var t int
	fmt.Sscanf(scanner.Text(), "%d", &t)

	for i := 0; i < t; i++ {
		solve(scanner, writer)
	}
}
