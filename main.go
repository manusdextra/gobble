package main

type TokenType string

type Token struct {
	kind   TokenType
	lexeme string
	line   int
}

type Scanner struct {
	source string
	list   []Token
}

func (s Scanner) scan() string {
	return s.source
}
