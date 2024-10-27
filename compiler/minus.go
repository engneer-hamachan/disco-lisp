package compiler

import (
	"disco/base"
)

type Minus struct {
	BuiltinCompiler
}

func NewMinus() BuiltinCompilerIF {
	return &Minus{
		BuiltinCompiler: BuiltinCompiler{
			name:              "-",
			returnType:        base.INT,
			firstArgumentType: base.INT,
			minArgumentCount:  2,
			maxArgumentCount:  1000,
		},
	}
}

func init() {
	bc := NewMinus()
	BuiltinCompilers[bc.getName()] = bc
}

func (m *Minus) builtinCompile(
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

	codes = append(codes, base.MINUS)
	codes = append(codes, sLength(s))

	err = m.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (m *Minus) typePropagation(
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

	tstack := TypeEnv.PopMultiStack(sLength(s))

	for _, t := range tstack {
		is_symbol, err := m.isSymbolOrNumberType(t, row, caller)
		if err != nil {
			return err
		}

		if is_symbol {
			err := m.setFunctionArgumentTypes(t, caller, m.firstArgumentType, row)
			if err != nil {
				return err
			}
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], m.getName())

	return nil
}
