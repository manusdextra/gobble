package main

import (
	"fmt"
	"os"
	"strings"
)

type TokenType int

const (
	EOF TokenType = iota
	SPACE
	WORD
)

type Token struct {
	kind   TokenType
	lexeme string
	line   int
}

type Scanner struct {
	source  []byte
	list    []Token
	start   int
	current int
	line    int
}

func (s Scanner) scan() []Token {
	var sb strings.Builder
	for {
		s.current++
		char := s.source[s.current]
		switch char {
		// delimiters that end a token
		case ' ':
			s.list = append(s.list,
				Token{
					lexeme: sb.String(),
				})
		default:
			sb.WriteByte(char)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gobble [file]")
		os.Exit(1)
	}
	source, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Input file error: %s", err)
		os.Exit(1)
	}
	s := Scanner{
		source:  source,
		start:   0,
		current: 0,
		line:    0,
	}
	tokenlist := s.scan()
	for _, v := range tokenlist {
		fmt.Printf("%c\n", v)
	}
	os.Exit(0)
}
