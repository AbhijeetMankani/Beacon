package main

import (
	"Beacon/lexer"
	"fmt"
)

func main() {
	input := "123 + 456- 789 + 56 ** 8*10 ***12"

	l := lexer.NewLexer(input)

	for {
		tok := l.NextToken()
		fmt.Printf("%+v\n", tok)
		if tok.Type == lexer.EOF {
			break
		}
	}
}
