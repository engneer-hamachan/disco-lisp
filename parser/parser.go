package parser

import (
	"disco/lexer"
	"fmt"
	"os"
)

type Parser struct {
	lexer       lexer.Lexer
	token       rune
	ungetFlg    bool
	FileName    string
	Row         int
	intaractive bool
}

func New(
	lexer lexer.Lexer,
	file_name string,
	intaractive bool,
) Parser {

	return Parser{
		lexer:       lexer,
		ungetFlg:    false,
		FileName:    file_name,
		Row:         1,
		intaractive: intaractive,
	}
}

func (p *Parser) Fatal(err error, intaractive bool) {
	fmt.Printf("%v::", p.FileName)
	fmt.Printf("%v::", p.Row)
	fmt.Printf("%v", err)

	if !p.intaractive {
		os.Exit(1)
	}

	fmt.Println("")
}
