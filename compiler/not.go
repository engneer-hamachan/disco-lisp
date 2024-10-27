package compiler

import (
	"disco/base"
)

type Not struct {
	BuiltinCompiler
}

func NewNot() BuiltinCompilerIF {
	return &Not{
		BuiltinCompiler{
			name:             "not",
			returnType:       base.ANY,
			minArgumentCount: 1,
			maxArgumentCount: 1,
		},
	}
}

func init() {
	bc := NewNot()
	BuiltinCompilers[bc.getName()] = bc
}

func (n *Not) builtinCompile(
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

	codes = codeAppend(codes, base.NOT, caller, file_name, row)

	err = n.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (n *Not) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := n.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	n.setFunctionReturnTypes(caller)

	TypeEnv.PopStack()

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], n.getName())

	return nil
}
