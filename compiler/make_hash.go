package compiler

import (
	"disco/base"
)

type MakeHash struct {
	BuiltinCompiler
}

func NewMakeHash() BuiltinCompilerIF {
	return &MakeHash{
		BuiltinCompiler{
			name:             "make-hash",
			returnType:       base.HASH,
			minArgumentCount: 0,
			maxArgumentCount: 0,
		},
	}
}

func init() {
	bc := NewMakeHash()
	BuiltinCompilers[bc.getName()] = bc
}

func (m *MakeHash) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	codes = append(codes, base.MAKE_HASH)

	err := m.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (m *MakeHash) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := m.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	m.setFunctionReturnTypes(caller)

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], m.getName())

	return nil
}
