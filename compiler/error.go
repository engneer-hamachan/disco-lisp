package compiler

import (
	"disco/base"
)

type Error struct {
	BuiltinCompiler
}

func NewError() BuiltinCompilerIF {
	return &Error{
		BuiltinCompiler{
			name:              "error",
			returnType:        base.ANY,
			firstArgumentType: base.STRING,
			minArgumentCount:  1,
			maxArgumentCount:  1,
		},
	}
}

func init() {
	bc := NewError()
	BuiltinCompilers[bc.getName()] = bc
}

func (e *Error) builtinCompile(
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

	codes = codeAppend(codes, base.ERROR, caller, file_name, row)

	err = e.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (e *Error) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := e.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	e.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopStack()

	is_symbol, err :=
		e.isSymbolOrWantType(tstack, e.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err := e.setFunctionArgumentTypes(tstack, caller, e.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], e.getName())

	return nil
}
