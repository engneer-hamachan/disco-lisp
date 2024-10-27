package compiler

import (
	"disco/base"
)

type Nilp struct {
	BuiltinCompiler
}

func NewNilp() BuiltinCompilerIF {
	return &Nilp{
		BuiltinCompiler{
			name:             "nil?",
			returnType:       base.ANY,
			minArgumentCount: 1,
			maxArgumentCount: 1,
		},
	}
}

func init() {
	bc := NewNilp()
	BuiltinCompilers[bc.getName()] = bc
}

func (n *Nilp) builtinCompile(
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

	codes = codeAppend(codes, base.NILP, caller, file_name, row)

	err = n.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (n *Nilp) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := n.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	TypeEnv.PopStack()

	n.setFunctionReturnTypes(caller)

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], n.getName())

	return nil
}
