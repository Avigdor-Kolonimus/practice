package problems

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type ParserBooleanExpression struct {
	s   string
	pos int
}

func (p *ParserBooleanExpression) parseExpr() int {
	val := p.parseAnd()

	for p.pos < len(p.s) {
		op := p.s[p.pos]
		if op != '|' && op != '^' {
			break
		}

		p.pos++
		right := p.parseAnd()

		if op == '|' {
			val |= right
		} else {
			val ^= right
		}
	}

	return val
}

func (p *ParserBooleanExpression) parseAnd() int {
	val := p.parseUnary()

	for p.pos < len(p.s) && p.s[p.pos] == '&' {
		p.pos++
		val &= p.parseUnary()
	}

	return val
}

func (p *ParserBooleanExpression) parseUnary() int {
	if p.s[p.pos] == '!' {
		p.pos++
		if p.parseUnary() == 0 {
			return 1
		}

		return 0
	}

	if p.s[p.pos] == '(' {
		p.pos++
		val := p.parseExpr()
		p.pos++ // ')'

		return val
	}

	ch := p.s[p.pos]
	p.pos++

	return int(ch - '0')
}

// https://coderun.yandex.ru/problem/the-value-of-a-boolean-expression
// TheValueOfABooleanExpression - problem 19
func TheValueOfABooleanExpression() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// expression input
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	expr := strings.TrimRight(line, "\r\n")

	p := ParserBooleanExpression{s: expr}

	ans := p.parseExpr()

	if ans == 1 {
		writer.WriteByte('1')
	} else {
		writer.WriteByte('0')
	}
	writer.WriteByte('\n')
}
