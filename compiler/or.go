package compiler

import (
	"disco/base"
)

type Or struct {
	BuiltinCompiler
}

func NewOr() BuiltinCompilerIF {
	return &Or{
		BuiltinCompiler{
			name:             "or",
			returnType:       base.ANY,
			minArgumentCount: 1,
			maxArgumentCount: 1,
		},
	}
}

func init() {
	bc := NewOr()
	BuiltinCompilers[bc.getName()] = bc
}

func (o *Or) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	var err error

	codes = codeAppend(codes, base.PROGN_START, caller, file_name, row)

	for {
		codes, err = compiler.Compile(codes, s.GetCar(), caller, file_name, row)
		if err != nil {
			return nil, err
		}

		codes = codeAppend(codes, base.OR, caller, file_name, row)

		codes = codeAppend(codes, base.LJMP, caller, file_name, row)
		codes = codeAppend(codes, "OR_END", caller, file_name, row)

		if s.GetCdr().Type != base.NIL {
			s = s.GetCdr()

			continue
		}

		break
	}

	codes = codeAppend(codes, base.LABEL, caller, file_name, row)
	codes = codeAppend(codes, "OR_END", caller, file_name, row)

	codes = codeAppend(codes, base.PROGN_END, caller, file_name, row)

	err = o.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (o *Or) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := o.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	TypeEnv.PopStack()

	o.setFunctionReturnTypes(caller)

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], o.getName())

	return nil
}
