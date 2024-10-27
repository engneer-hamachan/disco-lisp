package compiler

import (
	"disco/base"
)

type Listp struct {
	BuiltinCompiler
}

func NewListp() BuiltinCompilerIF {
	return &Listp{
		BuiltinCompiler{
			name:             "list?",
			returnType:       base.ANY,
			minArgumentCount: 1,
			maxArgumentCount: 1,
		},
	}
}

func init() {
	bc := NewListp()
	BuiltinCompilers[bc.getName()] = bc
}

func (l *Listp) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	codes, err := compiler.Compile(codes, s.GetCar(), caller, file_name, row)
	if err != nil {
		return nil, err
	}

	codes = codeAppend(codes, base.LISTP, caller, file_name, row)

	err = l.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (l *Listp) typePropagation(
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

	l.setFunctionReturnTypes(caller)

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], l.getName())

	return nil
}
