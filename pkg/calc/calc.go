package calc

import (
	"errors"
	"fmt"

	"calc/pkg/parser"
	"calc/pkg/tokenizer"
)

type Calculator interface {
	Calc(expression string) (float64, error)
}

type calculator struct {
}

func New() *calculator {
	return &calculator{}
}

func (d *calculator) Calc(expression string) (float64, error) {
	tokens, err := tokenizer.Tokenize(expression)
	if err != nil {
		return 0, errors.New("expression is not valid")
	}
	parser := &parser.Parser{Tokens: tokens}
	result, err := parser.ParseExpression()
	if err != nil {
		return 0, errors.New("expression is not valid")
	}

	if parser.Pos < len(tokens) {
		return 0, fmt.Errorf("expression is not valid")
	}

	return result, nil
}
