package compiler

import (
	"disco/base"
)

type AddRequestHeader struct {
	BuiltinCompiler
}

func NewAddRequestHeader() BuiltinCompilerIF {
	return &AddRequestHeader{
		BuiltinCompiler{
			name:               "add-request-header",
			returnType:         base.REQUEST,
			firstArgumentType:  base.REQUEST,
			secondArgumentType: base.STRING,
			minArgumentCount:   3,
			maxArgumentCount:   3,
		},
	}
}

func init() {
	bc := NewAddRequestHeader()
	BuiltinCompilers[bc.getName()] = bc
}

func (a *AddRequestHeader) builtinCompile(
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

	codes = codeAppend(codes, base.ADD_REQUEST_HEADER, caller, file_name, row)

	err = a.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (a *AddRequestHeader) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := a.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	a.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopMultiStack(3)

	is_symbol, err :=
		a.isSymbolOrWantType(tstack[0], a.firstArgumentType, row, caller)

	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			a.setFunctionArgumentTypes(tstack[0], caller, a.firstArgumentType, row)

		if err != nil {
			return err
		}
	}

	is_symbol, err =
		a.isSymbolOrWantType(tstack[1], a.secondArgumentType, row, caller)

	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			a.setFunctionArgumentTypes(tstack[1], caller, a.secondArgumentType, row)
		if err != nil {
			return err
		}
	}

	is_symbol, err =
		a.isSymbolOrWantType(tstack[2], a.secondArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			a.setFunctionArgumentTypes(tstack[2], caller, a.secondArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], a.getName())

	return nil
}
