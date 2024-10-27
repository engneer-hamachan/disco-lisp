package compiler

import (
	"disco/base"
)

type ListFuntction struct {
	BuiltinCompiler
}

func NewListFuntction() BuiltinCompilerIF {
	return &ListFuntction{
		BuiltinCompiler{
			name:             "list",
			returnType:       base.LIST,
			minArgumentCount: 1,
			maxArgumentCount: 1000,
		},
	}
}

func init() {
	bc := NewListFuntction()
	BuiltinCompilers[bc.getName()] = bc
}

func (l *ListFuntction) builtinCompile(
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

	codes = codeAppend(codes, base.LIST_FUNCTION, caller, file_name, row)
	codes = codeAppend(codes, sLength(s), caller, file_name, row)

	err = l.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (l *ListFuntction) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := l.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	TypeEnv.PopMultiStack(sLength(s))

	l.setFunctionReturnTypes(caller)

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], l.getName())

	return nil
}
