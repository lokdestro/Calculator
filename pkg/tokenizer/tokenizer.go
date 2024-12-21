package tokenizer

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Tokenize(input string) ([]Token, error) {
	var tokens []Token
	runes := []rune(strings.TrimSpace(input))
	i := 0
	for i < len(runes) {
		ch := runes[i]
		if unicode.IsSpace(ch) {
			i++
			continue
		}

		switch {
		case ch == '(':
			tokens = append(tokens, Token{Typ: TokenParenOpen, Value: "("})
			i++
		case ch == ')':
			tokens = append(tokens, Token{Typ: TokenParenClose, Value: ")"})
			i++
		case ch == '+' || ch == '-' || ch == '*' || ch == '/':
			tokens = append(tokens, Token{Typ: TokenOperator, Value: string(ch)})
			i++
		case unicode.IsDigit(ch) || ch == '.':
			start := i
			dotSeen := false
			for i < len(runes) && (unicode.IsDigit(runes[i]) || runes[i] == '.') {
				if runes[i] == '.' {
					if dotSeen {
						return nil, errors.New("invalid number format: multiple decimal points")
					}
					dotSeen = true
				}
				i++
			}
			value := string(runes[start:i])
			if _, err := strconv.ParseFloat(value, 64); err != nil {
				return nil, fmt.Errorf("invalid number format: %v", value)
			}
			tokens = append(tokens, Token{Typ: TokenNumber, Value: value})
		default:
			return nil, fmt.Errorf("invalid character: '%c'", ch)
		}
	}
	return tokens, nil
}
