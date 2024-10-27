package compiler

import (
	"disco/base"
)

type Mapcar struct {
	BuiltinCompiler
}

func NewMapcar() BuiltinCompilerIF {
	return &Mapcar{
		BuiltinCompiler{
			name:               "mapcar",
			returnType:         base.LIST,
			firstArgumentType:  base.EXEC_FUNC,
			secondArgumentType: base.LIST,
			minArgumentCount:   2,
			maxArgumentCount:   1000,
		},
	}
}

func init() {
	bc := NewMapcar()
	BuiltinCompilers[bc.getName()] = bc
}

func (m *Mapcar) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	codes, err :=
		compiler.recurisonCompile(codes, s.GetCdr(), caller, file_name, row)
	if err != nil {
		return nil, err
	}

	codes, err = compiler.Compile(codes, s.GetCar(), caller, file_name, row)
	if err != nil {
		return nil, err
	}

	codes = codeAppend(codes, base.MAPCAR, caller, file_name, row)
	codes = codeAppend(codes, sLength(s.GetCdr()), caller, file_name, row)

	err = m.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (m *Mapcar) typePropagation(
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

	tstack := TypeEnv.PopStack()

	is_symbol, err :=
		m.isSymbolOrWantType(tstack, m.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			m.setFunctionArgumentTypes(tstack, caller, m.firstArgumentType, row)

		if err != nil {
			return err
		}
	}

	arguments := TypeEnv.PopMultiStack(sLength(s) - 1)

	for _, argument := range arguments {
		is_symbol, err :=
			m.isSymbolOrWantType(argument, m.secondArgumentType, row, caller)
		if err != nil {
			return err
		}

		if is_symbol {
			err :=
				m.setFunctionArgumentTypes(argument, caller, m.secondArgumentType, row)
			if err != nil {
				return err
			}
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], m.getName())

	return nil
}
