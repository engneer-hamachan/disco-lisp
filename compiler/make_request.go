package compiler

import (
	"disco/base"
)

type MakeRequest struct {
	BuiltinCompiler
}

func NewMakeRequest() BuiltinCompilerIF {
	return &MakeRequest{
		BuiltinCompiler{
			name:               "make-request",
			returnType:         base.REQUEST,
			firstArgumentType:  base.QUOTED_SYMBOL,
			secondArgumentType: base.STRING,
			minArgumentCount:   2,
			maxArgumentCount:   3,
		},
	}
}

func init() {
	bc := NewMakeRequest()
	BuiltinCompilers[bc.getName()] = bc
}

func (m *MakeRequest) builtinCompile(
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

	codes = codeAppend(codes, base.MAKE_REQUEST, caller, file_name, row)
	codes = codeAppend(codes, sLength(s), caller, file_name, row)

	err = m.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (m *MakeRequest) typePropagation(
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

	_, err = m.isSymbolOrQuotedSymbol(tstack[0], row, caller)
	if err != nil {
		return err
	}

	if s.GetCar().Type != base.LIST {
		err :=
			m.setFunctionArgumentTypes(tstack[0], caller, m.firstArgumentType, row)

		if err != nil {
			return err
		}
	}

	is_symbol, err :=
		m.isSymbolOrWantType(tstack[1], m.secondArgumentType, row, caller)
	if err != nil {
		*row = base.InformationWhenParsing[s.GetCadr().Val].Row

		return err
	}

	if is_symbol {
		err :=
			m.setFunctionArgumentTypes(tstack[1], caller, m.secondArgumentType, row)

		if err != nil {
			return err
		}
	}

	if sLength(s) == 3 {
		is_symbol, err :=
			m.isSymbolOrWantType(tstack[2], m.secondArgumentType, row, caller)
		if err != nil {
			*row = base.InformationWhenParsing[s.GetCaddr().Val].Row

			return err
		}

		if is_symbol {
			err :=
				m.setFunctionArgumentTypes(tstack[2], caller, m.secondArgumentType, row)
			if err != nil {
				return err
			}
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], m.getName())

	return nil
}
