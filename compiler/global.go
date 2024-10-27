package compiler

import (
	"disco/base"
)

type Global struct {
	BuiltinCompiler
}

func NewGlobal() BuiltinCompilerIF {
	return &Global{
		BuiltinCompiler{
			name:             "global",
			returnType:       base.ANY,
			minArgumentCount: 2,
			maxArgumentCount: 2,
		},
	}
}

func init() {
	bc := NewGlobal()
	BuiltinCompilers[bc.getName()] = bc
}

func (g *Global) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	var err error

	switch s.GetCar().Type {
	case base.SYMBOL:
		quoted_object :=
			base.Cons(
				base.MakeSym("quote"),
				base.Cons(
					s.GetCar(),
					base.MakeNil(),
				),
			)

		codes, err =
			compiler.Compile(codes, quoted_object, caller, file_name, row)
		if err != nil {
			return nil, err
		}

	default:
		codes, err =
			compiler.Compile(codes, s.GetCar(), caller, file_name, row)
		if err != nil {
			return nil, err
		}
	}

	codes, err =
		compiler.recurisonCompile(codes, s.GetCdr(), caller, file_name, row)
	if err != nil {
		return nil, err
	}

	codes = codeAppend(codes, base.GLOBAL, caller, file_name, row)

	err = g.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (g *Global) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := g.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	tstack := TypeEnv.PopMultiStack(2)

	switch s.GetCadr().Type {
	case base.LIST:
		fname := s.GetCaadr().Val.(string)

		if fname == "lambda" || fname == "fn" {
			g.typePropagationForLambda(s, "", tstack[0])

			break
		}

		if fname == "quote" {
			FunctionArgumentTypes[MapKeys{"", tstack[0].Val.(string)}] =
				s.GetCadr().GetCadr().Type
		}

		f_return_type, ok := FunctionReturnTypes[fname]
		if ok {
			FunctionArgumentTypes[MapKeys{"", tstack[0].Val.(string)}] =
				f_return_type
		}

	default:
		FunctionArgumentTypes[MapKeys{"", tstack[0].Val.(string)}] =
			s.GetCadr().Type
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[""], g.getName())

	return nil
}

func (se *Global) typePropagationForLambda(
	s *base.S,
	caller string,
	tstack *base.S,
) {

	fname := "peek-lambda"
	symbol_name := s.GetCar().Val.(string)
	peek_lambda := base.Globals[fname]

	base.Globals[symbol_name] = peek_lambda

	s_args := peek_lambda.GetCar()

	for {
		if s_args.Type == base.NIL {
			break
		}

		argument_string := s_args.GetCar().Val.(string)

		FunctionArgumentTypes[MapKeys{symbol_name, argument_string}] =
			FunctionArgumentTypes[MapKeys{"peek-lambda", argument_string}]

		s_args = s_args.GetCdr()
	}

	FunctionReturnTypes[symbol_name] = FunctionReturnTypes[fname]

	f_return_type, ok := FunctionReturnTypes[fname]
	if ok {
		FunctionArgumentTypes[MapKeys{"", tstack.Val.(string)}] =
			f_return_type
	}
}
