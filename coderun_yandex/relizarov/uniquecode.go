package relizarov

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	MaxChars = 128
	Empty    = byte(0)
)

type Solver struct {
	word1 string
	word2 string

	maxWordSize int

	isPresent [MaxChars]bool
	isFirst   [MaxChars]bool

	letterDigit [MaxChars]byte
	digitLetter [10]byte

	letters []byte

	masks map[string]uint8

	words []Word
}

type Word struct {
	value string
	mask  string
}

func NewSolver(word1, word2 string, n int) *Solver {
	s := &Solver{
		word1:       word1,
		word2:       word2,
		maxWordSize: max(len(word1), len(word2)),
		masks:       make(map[string]uint8, n),
		words:       make([]Word, 0, n),
	}

	for i := 0; i < len(word1); i++ {
		s.isPresent[word1[i]] = true
	}

	for i := 0; i < len(word2); i++ {
		s.isPresent[word2[i]] = true
	}

	s.isFirst[word1[0]] = true
	s.isFirst[word2[0]] = true

	for i := 0; i < MaxChars; i++ {
		if s.isPresent[i] {
			s.letters = append(s.letters, byte(i))
		}
	}

	return s
}

func (s *Solver) MakeMask(word string) (string, int) {
	mapping := [MaxChars]byte{}
	next := byte('0')

	mask := make([]byte, len(word))
	for i := 0; i < len(word); i++ {
		ch := word[i]

		if s.isPresent[ch] {
			mask[i] = ch
			continue
		}

		if mapping[ch] == Empty {
			next++
			mapping[ch] = next
		}

		mask[i] = mapping[ch]
	}

	return string(mask), int(next - '0')
}

func (s *Solver) AddWord(word string) {
	if len(word) < s.maxWordSize ||
		len(word) > s.maxWordSize+1 {
		return
	}

	mask, newLetters := s.MakeMask(word)

	if len(s.letters)+newLetters > 10 {
		return
	}

	s.masks[mask] = 0

	s.words = append(s.words, Word{
		value: word,
		mask:  mask,
	})
}

func (s *Solver) CalcWord(word string) uint64 {
	var result uint64
	for i := 0; i < len(word); i++ {
		result = result*10 +
			uint64(s.letterDigit[word[i]]-'0')
	}

	return result
}

func (s *Solver) CheckSum() {
	sum := s.CalcWord(s.word1) + s.CalcWord(s.word2)

	sumWord := strconv.FormatUint(sum, 10)

	mapping := [MaxChars]byte{}
	next := byte('0')
	mask := make([]byte, len(sumWord))
	for i := 0; i < len(sumWord); i++ {
		digit := sumWord[i]

		letter := s.digitLetter[digit-'0']

		if letter != Empty {
			mask[i] = letter
			continue
		}

		if mapping[digit] == Empty {
			next++
			mapping[digit] = next
		}

		mask[i] = mapping[digit]
	}

	key := string(mask)

	if count, ok := s.masks[key]; ok && count < 2 {
		s.masks[key] = count + 1
	}
}

func (s *Solver) Calculate(pos int) {
	if pos == len(s.letters) {
		s.CheckSum()
		return
	}

	letter := s.letters[pos]

	start := byte('0')

	if s.isFirst[letter] {
		start = '1'
	}

	for digit := start; digit <= '9'; digit++ {
		d := digit - '0'

		if s.digitLetter[d] != Empty {
			continue
		}

		s.digitLetter[d] = letter
		s.letterDigit[letter] = digit

		s.Calculate(pos + 1)

		s.digitLetter[d] = Empty
		s.letterDigit[letter] = Empty
	}
}

func (s *Solver) PrintResult(writer *bufio.Writer) {
	count := 0

	for _, word := range s.words {
		if s.masks[word.mask] == 1 {
			count++
		}
	}

	writer.WriteString(strconv.Itoa(count))
	writer.WriteByte('\n')

	for _, word := range s.words {
		if s.masks[word.mask] == 1 {
			writer.WriteString(word.value)
			writer.WriteByte('\n')
		}
	}
}

// https://coderun.yandex.ru/selections/relizarov/problems/unique-code
// UniqueCode - problem 4
func UniqueCode() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var word1, word2 string
	var n int

	line, _ := reader.ReadString('\n')
	word1 = strings.TrimSpace(line)

	line, _ = reader.ReadString('\n')
	word2 = strings.TrimSpace(line)

	line, _ = reader.ReadString('\n')
	n, _ = strconv.Atoi(strings.TrimSpace(line))

	solver := NewSolver(word1, word2, n)

	for i := 0; i < n; i++ {
		line, _ = reader.ReadString('\n')
		word := strings.TrimSpace(line)

		solver.AddWord(word)
	}

	if len(solver.letters) <= 10 {
		solver.Calculate(0)
	}

	solver.PrintResult(writer)
}
