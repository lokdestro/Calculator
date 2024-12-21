package parser

import (
	"calc/pkg/tokenizer"
	"errors"
	"fmt"
	"strconv"
)

type Parser struct {
	Tokens []tokenizer.Token
	Pos    int
}

func (p *Parser) ParseExpression() (float64, error) {
	result, err := p.parseTerm()
	if err != nil {
		return 0, err
	}
	for p.Pos < len(p.Tokens) {
		tok := p.Tokens[p.Pos]
		if tok.Value == "+" || tok.Value == "-" {
			p.Pos++
			nextTerm, err := p.parseTerm()
			if err != nil {
				return 0, err
			}
			if tok.Value == "+" {
				result += nextTerm
			} else {
				result -= nextTerm
			}
		} else {
			break
		}
	}
	return result, nil
}

func (p *Parser) parseTerm() (float64, error) {
	result, err := p.parseFactor()
	if err != nil {
		return 0, err
	}

	for p.Pos < len(p.Tokens) {
		tok := p.Tokens[p.Pos]
		if tok.Value == "*" || tok.Value == "/" {
			p.Pos++
			nextFactor, err := p.parseFactor()
			if err != nil {
				return 0, err
			}
			if tok.Value == "*" {
				result *= nextFactor
			} else {
				if nextFactor == 0 {
					return 0, errors.New("division by zero")
				}
				result /= nextFactor
			}
		} else {
			break
		}
	}
	return result, nil
}

func (p *Parser) parseFactor() (float64, error) {
	if p.Pos >= len(p.Tokens) {
		return 0, errors.New("unexpected end of expression")
	}

	tok := p.Tokens[p.Pos]
	if tok.Typ == tokenizer.TokenNumber {
		p.Pos++
		return strconv.ParseFloat(tok.Value, 64)

	} else if tok.Typ == tokenizer.TokenOperator && tok.Value == "-" {
		p.Pos++
		value, err := p.parseFactor()
		if err != nil {
			return 0, err
		}
		return -value, nil
	} else if tok.Typ == tokenizer.TokenParenOpen {
		p.Pos++
		value, err := p.ParseExpression()
		if err != nil {
			return 0, err
		}
		if p.Pos >= len(p.Tokens) || p.Tokens[p.Pos].Typ != tokenizer.TokenParenClose {
			return 0, errors.New("missing closing parenthesis")
		}
		p.Pos++
		return value, nil
	} else {
		return 0, fmt.Errorf("unexpected token: %v", tok.Value)
	}
}
