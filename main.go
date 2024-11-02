package main

import (
	"fmt"
	"os"
)

type TokenType string

type Token struct {
	kind   TokenType
	lexeme string
	line   int
}

type Scanner struct {
	source  string
	list    []Token
	start   int
	current int
	line    int
}

func (s Scanner) scan() string {
	// placeholder
	return s.source
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gobble [file]")
		os.Exit(1)
	}
	s := Scanner{
		source: os.Args[1],
	}
	tokenlist := s.scan()
	for _, v := range tokenlist {
		fmt.Printf("%c\n", v)
	}
	os.Exit(0)
}
