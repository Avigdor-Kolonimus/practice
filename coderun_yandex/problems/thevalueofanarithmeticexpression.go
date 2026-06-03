package problems

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type ParserArithmeticExpression struct {
	s   string
	pos int
	ok  bool
}

func (p *ParserArithmeticExpression) skipSpaces() {
	for p.pos < len(p.s) && p.s[p.pos] == ' ' {
		p.pos++
	}
}

func (p *ParserArithmeticExpression) parseExpr() int {
	val := p.parseTerm()

	for {
		p.skipSpaces()

		if p.pos >= len(p.s) {
			break
		}

		op := p.s[p.pos]
		if op != '+' && op != '-' {
			break
		}

		p.pos++
		right := p.parseTerm()

		if op == '+' {
			val += right
		} else {
			val -= right
		}
	}

	return val
}

func (p *ParserArithmeticExpression) parseTerm() int {
	val := p.parseFactor()

	for {
		p.skipSpaces()

		if p.pos >= len(p.s) || p.s[p.pos] != '*' {
			break
		}

		p.pos++
		right := p.parseFactor()
		val *= right
	}

	return val
}

func (p *ParserArithmeticExpression) parseFactor() int {
	p.skipSpaces()

	if p.pos >= len(p.s) {
		p.ok = false

		return 0
	}

	if p.s[p.pos] == '+' {
		p.pos++

		return p.parseFactor()
	}

	if p.s[p.pos] == '-' {
		p.pos++

		return -p.parseFactor()
	}

	if p.s[p.pos] == '(' {
		p.pos++

		val := p.parseExpr()

		p.skipSpaces()
		if p.pos >= len(p.s) || p.s[p.pos] != ')' {
			p.ok = false

			return 0
		}

		p.pos++

		return val
	}

	if p.s[p.pos] < '0' || p.s[p.pos] > '9' {
		p.ok = false

		return 0
	}

	val := 0
	for p.pos < len(p.s) && p.s[p.pos] >= '0' && p.s[p.pos] <= '9' {
		val = val*10 + int(p.s[p.pos]-'0')
		p.pos++
	}

	return val
}

// https://coderun.yandex.ru/problem/the-value-of-an-arithmetic-expression
// TheValueOfAnArithmeticExpression - problem 18
func TheValueOfAnArithmeticExpression() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// expression input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	expr := strings.TrimRight(line, "\r\n")

	p := ParserArithmeticExpression{
		s:  expr,
		ok: true,
	}

	ans := p.parseExpr()

	p.skipSpaces()

	if !p.ok || p.pos != len(p.s) {
		writer.WriteString("WRONG")
		writer.WriteByte('\n')

		return
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
