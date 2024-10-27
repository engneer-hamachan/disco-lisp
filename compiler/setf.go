package compiler

import (
	"disco/base"
)

type Setf struct {
	BuiltinCompiler
}

func NewSetf() BuiltinCompilerIF {
	return &Setf{
		BuiltinCompiler{
			name:             "=",
			returnType:       base.ANY,
			minArgumentCount: 2,
			maxArgumentCount: 2,
		},
	}
}

func init() {
	bc := NewSetf()
	BuiltinCompilers[bc.getName()] = bc
}

func (se *Setf) builtinCompile(
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

	codes = codeAppend(codes, base.SET, caller, file_name, row)

	err = se.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (se *Setf) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := se.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	tstack := TypeEnv.PopMultiStack(2)

	delete(
		FunctionArgumentOptionalTypes,
		MapKeys{caller, tstack[0].Val.(string)},
	)

	switch s.GetCadr().Type {
	case base.NIL:
		FunctionReturnTypes[caller] = base.NIL

	case base.LIST:
		FunctionReturnTypes[caller] = s.GetCadr().Type

		fname := s.GetCaadr().Val.(string)

		if fname == "lambda" || fname == "fn" {
			se.typePropagationForLambda(s, caller, tstack[0])

			break
		}

		if fname == "quote" {
			FunctionArgumentTypes[MapKeys{caller, tstack[0].Val.(string)}] =
				s.GetCadr().GetCadr().Type

			break
		}

		optional_types, ok := FunctionOptionalTypes[fname]
		if ok {
			FunctionArgumentOptionalTypes[MapKeys{caller, tstack[0].Val.(string)}] =
				optional_types

			FunctionArgumentTypes[MapKeys{caller, tstack[0].Val.(string)}] =
				base.OPTIONAL

			break
		}

		if tstack[1].Type != base.SYMBOL {
			FunctionArgumentTypes[MapKeys{caller, tstack[0].Val.(string)}] =
				tstack[1].Type

			break
		}

		f_return_type, ok := FunctionReturnTypes[fname]
		if ok {
			FunctionArgumentTypes[MapKeys{caller, tstack[0].Val.(string)}] =
				f_return_type
		}

	case base.SYMBOL:
		optional_types, ok :=
			FunctionArgumentOptionalTypes[MapKeys{caller, tstack[1].Val.(string)}]

		if ok {
			FunctionReturnTypes[caller] = base.OPTIONAL

			FunctionArgumentTypes[MapKeys{caller, tstack[0].Val.(string)}] =
				base.OPTIONAL

			FunctionArgumentOptionalTypes[MapKeys{caller, tstack[0].Val.(string)}] =
				optional_types

			break
		}

		FunctionReturnTypes[caller] = tstack[1].Type

		FunctionArgumentTypes[MapKeys{caller, tstack[0].Val.(string)}] =
			tstack[1].Type

	default:
		FunctionReturnTypes[caller] = tstack[1].Type

		FunctionArgumentTypes[MapKeys{caller, tstack[0].Val.(string)}] =
			tstack[1].Type

	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], se.getName())

	return nil
}

func (se *Setf) typePropagationForLambda(
	s *base.S,
	caller string,
	tstack *base.S,
) {

	fname := "peek-lambda"
	symbol_name := s.GetCar().Val.(string)
	peek_lambda := base.Globals[fname]

	delete(FunctionOptionalTypes, symbol_name)

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

	f_return_type, ok := FunctionReturnTypes[fname]
	if ok {
		FunctionReturnTypes[symbol_name] = f_return_type

		if f_return_type == base.OPTIONAL {
			FunctionOptionalTypes[symbol_name] = FunctionOptionalTypes[fname]
		}

		FunctionArgumentTypes[MapKeys{caller, tstack.Val.(string)}] = f_return_type
	}
}
