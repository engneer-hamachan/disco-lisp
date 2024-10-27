package compiler

import (
	"disco/base"
	"disco/predicater"
	"fmt"
)

type Compiler struct{}

var compiler = &Compiler{}

func NewCompiler() *Compiler {
	return &Compiler{}
}

func codeAppend(
	codes []any,
	code any,
	caller string,
	file_name *string,
	row *int,
) []any {

	codes = append(codes, code)

	_, ok := base.InformationWhenCompile[caller]
	if !ok {
		base.InformationWhenCompile[caller] = make(map[int]base.Info)
	}

	base.InformationWhenCompile[caller][len(codes)-1] =
		base.NewInfo(*file_name, *row, code)

	return codes
}

func (c *Compiler) compileAtom(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	TypeEnv.PushStack(s)

	if s.Type == base.SYMBOL {
		symbol_type, ok := FunctionArgumentTypes[MapKeys{caller, s.Val.(string)}]
		switch ok {
		case true:
			FunctionReturnTypes[caller] = symbol_type
		default:
			FunctionReturnTypes[caller] = base.ANY
		}
	}

	if s.Type != base.SYMBOL {
		FunctionReturnTypes[caller] = s.Type
	}

	codes = codeAppend(codes, base.LD, caller, file_name, row)
	codes = codeAppend(codes, s, caller, file_name, row)

	return codes, nil
}

func (c *Compiler) compileList(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	if predicater.Intp(s.GetCar()) ||
		predicater.Stringp(s.GetCar()) ||
		predicater.Floatp(s.GetCar()) {

		return nil, fmt.Errorf("operator error %v", s.GetCar().Val)
	}

	if predicater.HasName(s.GetCar(), "quote") {
		s.GetCadr().IsQuoted = true

		TypeEnv.PushStack(s.GetCadr())

		switch s.GetCadr().Type {
		case base.SYMBOL:
			FunctionReturnTypes[caller] = base.QUOTED_SYMBOL
		default:
			FunctionReturnTypes[caller] = s.GetCadr().Type
		}

		codes = codeAppend(codes, base.QUOTE, caller, file_name, row)
		codes = codeAppend(codes, s.GetCadr(), caller, file_name, row)

		return codes, nil
	}

	if predicater.HasName(s.GetCaar(), "lambda") ||
		predicater.HasName(s.GetCaar(), "fn") {

		s = lambdaCompile(s, file_name, row)
	}

	types, ok := FunctionArgumentTypes[MapKeys{caller, s.GetCar().Val.(string)}]
	if ok {
		var err error

		s, err = c.firstArrayOrHashCompile(s, types, row)
		if err != nil {
			return nil, err
		}
	}

	f, ok := BuiltinCompilers[s.GetCar().Val.(string)]
	if ok {
		return f.builtinCompile(codes, s.GetCdr(), caller, file_name, row)
	}

	v, ok := base.Globals[s.GetCar().Val.(string)]
	if ok && v.Type == base.MACRO {
		return compileMacro(codes, v, s.GetCdr(), caller, file_name, row)
	}

	if !ok {
		return c.lastArrayOrHashCompile(codes, s, caller, file_name, row)
	}

	return compileFunction(codes, s.GetCar(), s.GetCdr(), caller, file_name, row)
}

func (c *Compiler) recurisonCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	for {
		car := s.GetCar()

		if car == nil || (car.Type == base.NIL && car.GetCdr() == nil) {
			break
		}

		var err error

		codes, err = c.Compile(codes, car, caller, file_name, row)
		if err != nil {
			return nil, err
		}

		s = s.GetCdr()
	}

	return codes, nil
}

func (c *Compiler) firstArrayOrHashCompile(
	s *base.S,
	types int,
	row *int,
) (*base.S, error) {

	if types == base.VECTOR {
		s = arefCompile(s)
	}

	if types == base.HASH {
		switch sLength(s) {
		case 3:
			s = sethashCompile(s)
		case 2:
			s = gethashCompile(s)
		default:
			*row = base.InformationWhenParsing[s.GetCar().Val].Row
			return nil, fmt.Errorf("undefined method %s.", s.GetCar().Val.(string))
		}
	}

	return s, nil
}

func (c *Compiler) lastArrayOrHashCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	if s.GetCadr().Type == base.INT {
		s = arefCompile(s)
		return compiler.Compile(codes, s, caller, file_name, row)
	}

	switch sLength(s) {
	case 3:
		s = sethashCompile(s)
	case 2:
		s = gethashCompile(s)
	default:
		*row = base.InformationWhenParsing[s.GetCar().Val].Row
		return nil, fmt.Errorf("undefined method %s.", s.GetCar().Val.(string))
	}

	return compiler.Compile(codes, s, caller, file_name, row)
}

func (c *Compiler) Compile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	if predicater.Atomp(s) {
		return c.compileAtom(codes, s, caller, file_name, row)
	}

	return c.compileList(codes, s, caller, file_name, row)
}
