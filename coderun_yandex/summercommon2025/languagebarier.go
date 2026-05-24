package summercommon2025

import (
	"sort"
	"strings"
)

type WordLanguageBarier struct {
	Index int
	Word  string
}

// https://coderun.yandex.ru/selections/2025-summer-common/problems/language-barier
// LanguageBarier - problem 9
func SolveLanguageBarier(n int, words []string) [][2]int {
	arr := make([]WordLanguageBarier, len(words))
	for i, w := range words {
		arr[i] = WordLanguageBarier{
			Index: i + 1,
			Word:  w,
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i].Word == arr[j].Word {
			return arr[i].Index < arr[j].Index
		}
		return arr[i].Word > arr[j].Word
	})

	stack := make([]WordLanguageBarier, 0, n)
	ans := make([][2]int, 0, n)

	for _, cur := range arr {
		if len(stack) > 0 &&
			strings.HasPrefix(stack[len(stack)-1].Word, cur.Word) {

			full := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ans = append(ans, [2]int{
				cur.Index,
				full.Index,
			})
		} else {
			stack = append(stack, cur)
		}
	}

	return ans
}
