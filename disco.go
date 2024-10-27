package main

import (
	"bufio"
	"disco/banner"
	"disco/base"
	"disco/compiler"
	"disco/fever"
	"disco/lexer"
	"disco/lexer/reader"
	"disco/parser"
	"disco/printer"
	"fmt"

	//	"github.com/pkg/profile"
	"os"
	"runtime/debug"
)

func getParser(
	br *bufio.Reader,
	file_name string,
	intaractive bool,
) parser.Parser {

	lr := reader.New(*br)
	l := lexer.New(lr)

	return parser.New(l, file_name, intaractive)
}

func load(
	env *base.Environment,
	s *base.S,
	intaractive bool,
	p parser.Parser,
	c *compiler.Compiler,
) error {

	file_name := s.GetCadr().Val.(string)

	fp, err := os.Open(file_name)
	if err != nil {
		return fmt.Errorf("%s is not open. check file path.", file_name)
	}

	br := bufio.NewReader(fp)
	otherp := getParser(br, file_name, intaractive)

	loop(env, otherp, c, intaractive)

	return nil
}

func loop(
	env *base.Environment,
	p parser.Parser,
	c *compiler.Compiler,
	intaractive bool,
) {

	for {
		if intaractive {
			fmt.Print("(disco) => ")
		}

		s, err := p.Read()
		if err != nil {
			p.Fatal(err, intaractive)
		}

		if s != nil && s.GetCar().Type == base.LOAD {
			err := load(env, s, intaractive, p, c)
			if err != nil {
				p.Fatal(err, intaractive)
				continue
			}

			continue
		}

		if s != nil {
			switch intaractive {
			case true:
				fever.VM.Codes, err =
					c.Compile(fever.VM.Codes, s, "", &p.FileName, &p.Row)
				if err != nil {
					p.Fatal(err, intaractive)
					continue
				}

				err := fever.Fever(fever.VM.Codes, env, "")
				if err != nil {
					fever.VM.Codes = make([]any, 0)

					continue
				}

				printer.Print(fever.VM.PopStack(), false)
				fmt.Println("")

				fever.VM.Codes = make([]any, 0)

			default:
				fever.VM.Codes, err =
					c.Compile(fever.VM.Codes, s, "", &p.FileName, &p.Row)
				if err != nil {
					p.Fatal(err, intaractive)
				}
			}

			continue
		}

		break
	}
}

func main() {
	debug.SetGCPercent(700)

	var br *bufio.Reader
	var file_name string

	switch len(os.Args) {
	case 1:
		file_name = "REPL"
		fever.VM.Intaractive = true
		fmt.Printf("%v", banner.ReplTitle)
		br = bufio.NewReader(os.Stdin)

	default:
		file_name = os.Args[1]
		fp, _ := os.Open(file_name)
		br = bufio.NewReader(fp)
	}

	p := getParser(br, file_name, fever.VM.Intaractive)
	c := compiler.NewCompiler()
	env := base.NewEnvironment()

	loop(&env, p, c, fever.VM.Intaractive)

	if len(os.Args) >= 3 && os.Args[2] == "-d" {
		fever.VM.Dump()
		return
	}

	if len(os.Args) >= 3 && os.Args[2] == "-t" {
		return
	}

	fever.FeverPreparaion()

	//	defer profile.Start().Stop()
	fever.Fever(fever.VM.Codes, &env, "")
}
