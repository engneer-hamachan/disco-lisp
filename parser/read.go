package parser

import (
	"disco/base"
	"disco/predicater"
	"errors"
)

func (p *Parser) getToken() {
	if p.ungetFlg {
		p.ungetFlg = false
		return
	}

	if p.lexer.Advance() {
		p.token = p.lexer.Token()
		p.skipLf()
		return
	}

	p.token = base.EOS
}

func (p *Parser) unget() {
	p.ungetFlg = true
}

func (p *Parser) skipLf() {
	for {
		if p.token != '\n' {
			break
		}

		p.Row += 1

		p.getToken()
	}
}

func (p *Parser) Read() (*base.S, error) {
	p.getToken()

	switch p.token {
	case base.INT:
		s := base.MakeInt(p.lexer.Value().(int64))

		base.InformationWhenParsing[s.Val] =
			base.Info{
				FileName: p.FileName,
				Row:      p.Row,
			}

		return s, nil

	case base.FLOAT:
		s := base.MakeFloat(p.lexer.Value().(float64))

		base.InformationWhenParsing[s.Val] =
			base.Info{
				FileName: p.FileName,
				Row:      p.Row,
			}

		return s, nil

	case base.STRING:
		s := base.MakeString(p.lexer.Value().(string))

		base.InformationWhenParsing[s.Val] =
			base.Info{
				FileName: p.FileName,
				Row:      p.Row,
			}

		return s, nil

	case 't':
		return base.TrueAtom, nil

	case base.NIL:
		return base.MakeNil(), nil

	case '+', '-', '/', '*', '>', '<':
		s := base.MakeSym(string(p.token))

		base.InformationWhenParsing[s.Val] =
			base.Info{
				FileName: p.FileName,
				Row:      p.Row,
			}

		return s, nil

	case base.LOAD:
		return base.Load, nil

	case base.SYMBOL:
		sym := p.lexer.Value().(base.Symbol)

		s := base.MakeSym(sym.GetName())

		base.InformationWhenParsing[s.Val] =
			base.Info{
				FileName: p.FileName,
				Row:      p.Row,
			}

		return s, nil

	case '\'':
		car, _ := p.Read()

		base.InformationWhenParsing[car] =
			base.Info{
				FileName: p.FileName,
				Row:      p.Row,
			}

		quoted_object :=
			base.Cons(
				base.MakeSym("quote"),
				base.Cons(
					car,
					base.MakeNil(),
				),
			)

		base.InformationWhenParsing[quoted_object] =
			base.Info{
				FileName: p.FileName,
				Row:      p.Row,
			}

		return quoted_object, nil

	case '#':
		p.getToken()

		if p.token == '\'' {
			car, _ := p.Read()

			base.InformationWhenParsing[car] =
				base.Info{
					FileName: p.FileName,
					Row:      p.Row,
				}

			function := base.MakeSym("function")

			base.InformationWhenParsing[function] =
				base.Info{
					FileName: p.FileName,
					Row:      p.Row,
				}

			function_object :=
				base.Cons(
					function,
					base.Cons(
						car,
						base.MakeNil(),
					),
				)

			base.InformationWhenParsing[function_object] =
				base.Info{
					FileName: p.FileName,
					Row:      p.Row,
				}

			return function_object, nil
		}

		return nil, errors.New("read error!")

	case '`':
		car, _ := p.Read()

		base.InformationWhenParsing[car] =
			base.Info{
				FileName: p.FileName,
				Row:      p.Row,
			}

		quasi_quote := base.MakeSym("quasi-quote")

		base.InformationWhenParsing[quasi_quote] =
			base.Info{
				FileName: p.FileName,
				Row:      p.Row,
			}

		quoted_object :=
			base.Cons(
				quasi_quote,
				base.Cons(
					car,
					base.MakeNil(),
				),
			)

		base.InformationWhenParsing[quoted_object] =
			base.Info{
				FileName: p.FileName,
				Row:      p.Row,
			}

		return quoted_object, nil

	case ',':
		p.getToken()

		if p.token == '@' {
			car, _ := p.Read()

			base.InformationWhenParsing[car] =
				base.Info{
					FileName: p.FileName,
					Row:      p.Row,
				}

			unquote_splicing := base.MakeSym("unquote-splicing")

			base.InformationWhenParsing[unquote_splicing] =
				base.Info{
					FileName: p.FileName,
					Row:      p.Row,
				}

			quoted_object :=
				base.Cons(
					unquote_splicing,
					base.Cons(
						car,
						base.MakeNil(),
					),
				)

			base.InformationWhenParsing[quoted_object] =
				base.Info{
					FileName: p.FileName,
					Row:      p.Row,
				}

			return quoted_object, nil
		}

		p.unget()

		car, _ := p.Read()

		base.InformationWhenParsing[car] =
			base.Info{
				FileName: p.FileName,
				Row:      p.Row,
			}

		unquote := base.MakeSym("unquote")

		base.InformationWhenParsing[unquote] =
			base.Info{
				FileName: p.FileName,
				Row:      p.Row,
			}

		quoted_object :=
			base.Cons(
				unquote,
				base.Cons(
					car,
					base.MakeNil(),
				),
			)

		base.InformationWhenParsing[quoted_object] =
			base.Info{
				FileName: p.FileName,
				Row:      p.Row,
			}

		return quoted_object, nil

	case '(':
		object, err := p.readlist()
		if err != nil {
			return nil, err
		}

		base.InformationWhenParsing[object] =
			base.Info{
				FileName: p.FileName,
				Row:      p.Row,
			}

		return object, nil

	case base.EOS:
		return nil, nil

	default:
		return nil, errors.New("read error!")
	}
}

func (p *Parser) readlist() (*base.S, error) {
	p.getToken()

	switch p.token {
	case ')':
		return base.MakeNil(), nil

	case '.':
		cdr, err := p.Read()
		if err != nil {
			return nil, err
		}

		base.InformationWhenParsing[cdr] =
			base.Info{
				FileName: p.FileName,
				Row:      p.Row,
			}

		if predicater.Symbolp(cdr) ||
			predicater.Intp(cdr) ||
			predicater.Floatp(cdr) ||
			predicater.Stringp(cdr) ||
			predicater.Nilp(cdr) ||
			predicater.Truep(cdr) {

			p.getToken()
		}

		return cdr, nil

	default:
		p.unget()

		car, err := p.Read()
		if err != nil {
			return nil, err
		}

		base.InformationWhenParsing[car] =
			base.Info{
				FileName: p.FileName,
				Row:      p.Row,
			}

		if car == nil {
			return nil, errors.New("read error!")
		}

		cdr, err := p.readlist()
		if err != nil {
			return nil, err
		}

		base.InformationWhenParsing[cdr] =
			base.Info{
				FileName: p.FileName,
				Row:      p.Row,
			}

		base.InformationWhenParsing["list-data"] =
			base.Info{
				FileName: p.FileName,
				Row:      p.Row,
			}

		return base.Cons(car, cdr), nil
	}
}
