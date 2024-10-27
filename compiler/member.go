package compiler

import (
	"disco/base"
)

type Member struct {
	BuiltinCompiler
}

func NewMember() BuiltinCompilerIF {
	return &Member{
		BuiltinCompiler{
			name:              "member",
			returnType:        base.LIST,
			firstArgumentType: base.LIST,
			minArgumentCount:  2,
			maxArgumentCount:  2,
		},
	}
}

func init() {
	bc := NewMember()
	BuiltinCompilers[bc.getName()] = bc
}

func (m *Member) builtinCompile(
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

	codes = append(codes, base.MEMBER)

	err = m.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (m *Member) typePropagation(
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

	tstack := TypeEnv.PopMultiStack(2)
	list_s := tstack[1]

	is_symbol, err :=
		m.isSymbolOrWantType(list_s, m.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err := m.setFunctionArgumentTypes(list_s, caller, m.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], m.getName())

	return nil
}
