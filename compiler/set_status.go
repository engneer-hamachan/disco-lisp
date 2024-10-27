package compiler

import (
	"disco/base"
)

type SetStatus struct {
	BuiltinCompiler
}

func NewSetStatus() BuiltinCompilerIF {
	return &SetStatus{
		BuiltinCompiler{
			name:              "set-status",
			returnType:        base.ANY,
			firstArgumentType: base.INT,
			minArgumentCount:  1,
			maxArgumentCount:  1,
		},
	}
}

func init() {
	bc := NewSetStatus()
	BuiltinCompilers[bc.getName()] = bc
}

func (se *SetStatus) builtinCompile(
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

	codes = codeAppend(codes, base.SETSTATUS, caller, file_name, row)

	err = se.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (se *SetStatus) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := se.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	se.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopStack()

	is_symbol, err :=
		se.isSymbolOrWantType(tstack, se.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			se.setFunctionArgumentTypes(tstack, caller, se.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], se.getName())

	return nil
}
