package compiler

import (
	"disco/base"
)

type And struct {
	BuiltinCompiler
}

func NewAnd() BuiltinCompilerIF {
	return &And{
		BuiltinCompiler{
			name:             "and",
			returnType:       base.ANY,
			minArgumentCount: 2,
			maxArgumentCount: 1000,
		},
	}
}

func init() {
	bc := NewAnd()
	BuiltinCompilers[bc.getName()] = bc
}

func (a *And) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	var err error
	args := s

	codes = codeAppend(codes, base.PROGN_START, caller, file_name, row)

	for {
		codes, err = compiler.Compile(codes, args.GetCar(), caller, file_name, row)
		if err != nil {
			return nil, err
		}

		codes = codeAppend(codes, base.AND, caller, file_name, row)
		codes = codeAppend(codes, base.LJMP, caller, file_name, row)
		codes = codeAppend(codes, "AND_END", caller, file_name, row)

		if args.GetCdr().Type != base.NIL {
			args = args.GetCdr()

			continue
		}

		break
	}

	codes = codeAppend(codes, base.LABEL, caller, file_name, row)
	codes = codeAppend(codes, "AND_END", caller, file_name, row)
	codes = codeAppend(codes, base.PROGN_END, caller, file_name, row)

	err = a.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (a *And) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := a.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	a.setFunctionReturnTypes(caller)

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], a.getName())

	return nil
}
