package tokenizer

type Token struct {
	Typ   TokenType
	Value string
}

type TokenType int

const (
	TokenNumber TokenType = iota
	TokenOperator
	TokenParenOpen
	TokenParenClose
)
