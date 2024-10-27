package compiler

import (
	"disco/base"
)

type Lambda struct {
	BuiltinCompiler
}

func NewLambda() BuiltinCompilerIF {
	return &Lambda{
		BuiltinCompiler{
			name:             "fn",
			returnType:       base.LIST,
			minArgumentCount: 1,
			maxArgumentCount: 1000,
		},
	}
}

func init() {
	bc := NewLambda()
	BuiltinCompilers[bc.getName()] = bc
}

func (l *Lambda) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	body := s.GetCdr()
	s_args := s.GetCar()

	base.Globals["peek-lambda"] = base.MakeFunc("peek-lambda", s_args)

	var args []*base.S

	for {
		if s_args.Type == base.NIL {
			break
		}

		args = append(args, s_args.GetCar())

		s_args = s_args.GetCdr()
	}

	err := l.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	var fbody []any

	fbody, err =
		compiler.recurisonCompile(fbody, body, "peek-lambda", file_name, row)
	if err != nil {
		return nil, err
	}

	f :=
		&base.F{
			Body: fbody,
			Args: args,
			Env:  nil,
		}

	codes = codeAppend(codes, base.LAMBDA, caller, file_name, row)
	codes = codeAppend(codes, f, caller, file_name, row)

	return codes, nil
}

func lambdaCompile(s *base.S, file_name *string, row *int) *base.S {
	base.InformationWhenParsing["funcall"] =
		base.Info{
			FileName: *file_name,
			Row:      *row,
		}

	return base.Cons(
		base.MakeSym("funcall"),
		base.Cons(
			base.Cons(
				base.MakeSym("function"),
				base.Cons(
					s.GetCar(),
					base.NilAtom,
				),
			),
			s.GetCdr(),
		),
	)
}

func (l *Lambda) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := l.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	delete(FunctionReturnTypes, "peek-lambda")
	delete(FunctionOptionalTypes, "peek-lambda")

	l.setFunctionReturnTypes(caller)

	s_args := s.GetCar()

	for {
		if s_args.Type == base.NIL {
			break
		}

		argumment_string := s_args.GetCar().Val.(string)
		FunctionArgumentTypes[MapKeys{"peek-lambda", argumment_string}] = base.ANY

		s_args = s_args.GetCdr()
	}

	return nil
}
