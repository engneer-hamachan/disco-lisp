package compiler

import (
	"disco/base"
)

type Let struct {
	BuiltinCompiler
}

func NewLet() BuiltinCompilerIF {
	return &Let{
		BuiltinCompiler{
			name:             "with",
			returnType:       base.ANY,
			minArgumentCount: 1,
			maxArgumentCount: 1000,
		},
	}
}

func init() {
	bc := NewLet()
	BuiltinCompilers[bc.getName()] = bc
}

func (l *Let) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	args := s.GetCar()
	body := s.GetCdr()

	codes = codeAppend(codes, base.PUSH_LEXICAL_FRAME, caller, file_name, row)

	codes, err := l.LetBind(codes, args, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	codes, err =
		compiler.recurisonCompile(codes, body, l.getName(), file_name, row)
	if err != nil {
		return nil, err
	}

	codes = codeAppend(codes, base.POP_LEXICAL_FRAME, caller, file_name, row)

	err = l.typePropagation(s, l.getName(), file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (l *Let) LetBind(
	codes []any,
	args *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	for {
		if args.GetCar().Type == base.NIL || args.GetCadr().Type == base.NIL {
			break
		}

		var err error
		codes, err = compiler.Compile(codes, args.GetCadr(), caller, file_name, row)
		if err != nil {
			return nil, err
		}

		codes = codeAppend(codes, base.LET, caller, file_name, row)

		codes =
			codeAppend(codes, args.GetCar().Val.(string), l.getName(), file_name, row)

		switch l.isLambda(args) {
		case true:
			l.typePropagationWhenBindForLambda(args)
		default:
			l.typePropagationWhenBind(args)
		}

		args = args.GetCddr()
	}

	return codes, nil
}

func (l *Let) isLambda(args *base.S) bool {
	if args.GetCadr().Type == base.LIST &&
		(args.GetCadr().GetCar().Val.(string) == "lambda" ||
			args.GetCadr().GetCar().Val.(string) == "fn") {

		return true
	}

	return false
}

func (l *Let) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := l.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	TypeEnv.PopStack()

	args := s.GetCar()

	err = l.checkArgumentCountIsEven(args, row)
	if err != nil {
		return err
	}

	for {
		if args.GetCar().Type == base.NIL || args.GetCadr().Type == base.NIL {
			break
		}

		delete(
			FunctionArgumentOptionalTypes,
			MapKeys{caller, args.GetCar().Val.(string)},
		)

		FunctionArgumentTypes[MapKeys{caller, args.GetCar().Val.(string)}] =
			base.ANY

		args = args.GetCddr()
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], l.getName())

	return nil
}

func (l *Let) typePropagationWhenBind(args *base.S) {
	tstack := TypeEnv.PopStack()

	delete(
		FunctionArgumentOptionalTypes,
		MapKeys{l.getName(), args.GetCar().Val.(string)},
	)

	FunctionArgumentTypes[MapKeys{l.getName(), args.GetCar().Val.(string)}] =
		tstack.Type

	if tstack.Type == base.INT {
		return
	}

	optional_types, ok := FunctionOptionalTypes[tstack.Val.(string)]
	if ok {
		FunctionArgumentOptionalTypes[MapKeys{
			l.getName(),
			args.GetCar().Val.(string),
		}] = optional_types
	}
}

func (l *Let) typePropagationWhenBindForLambda(args *base.S) {
	fname := "peek-lambda"
	symbol_name := args.GetCar().Val.(string)
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

	f_return_type, ok := FunctionReturnTypes[fname]
	if ok {
		FunctionReturnTypes[symbol_name] = f_return_type
		FunctionArgumentTypes[MapKeys{l.getName(), symbol_name}] = f_return_type
	}

	if f_return_type == base.OPTIONAL {
		FunctionOptionalTypes[symbol_name] = FunctionOptionalTypes[fname]
	}
}
