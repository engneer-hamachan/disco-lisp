package compiler

import (
	"disco/base"
)

type Eq struct {
	BuiltinCompiler
}

func NewEq() BuiltinCompilerIF {
	return &Eq{
		BuiltinCompiler{
			name:             "eq?",
			returnType:       base.ANY,
			minArgumentCount: 2,
			maxArgumentCount: 2,
		},
	}
}

func init() {
	bc := NewEq()
	BuiltinCompilers[bc.getName()] = bc
}

func (e *Eq) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	codes, err := compiler.recurisonCompile(codes, s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	codes = codeAppend(codes, base.EQ, caller, file_name, row)

	err = e.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (e *Eq) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := e.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	e.setFunctionReturnTypes(caller)

	TypeEnv.PopMultiStack(2)

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], e.getName())

	return nil
}
