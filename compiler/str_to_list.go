package compiler

import (
	"disco/base"
)

type StrToList struct {
	BuiltinCompiler
}

func NewStrToList() BuiltinCompilerIF {
	return &StrToList{
		BuiltinCompiler{
			name:              "str-to-list",
			returnType:        base.LIST,
			firstArgumentType: base.STRING,
			minArgumentCount:  1,
			maxArgumentCount:  1,
		},
	}
}

func init() {
	bc := NewStrToList()
	BuiltinCompilers[bc.getName()] = bc
}

func (stl *StrToList) builtinCompile(
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

	codes = codeAppend(codes, base.STR_TO_LIST, caller, file_name, row)

	err = stl.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (stl *StrToList) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := stl.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	stl.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopStack()

	is_symbol, err :=
		stl.isSymbolOrWantType(tstack, stl.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			stl.setFunctionArgumentTypes(tstack, caller, stl.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], stl.getName())

	return nil
}
