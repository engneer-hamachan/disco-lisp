package compiler

import (
	"disco/base"
)

type DoRequest struct {
	BuiltinCompiler
}

func NewDoRequest() BuiltinCompilerIF {
	return &DoRequest{
		BuiltinCompiler{
			name:              "do-request",
			returnType:        base.STRING,
			firstArgumentType: base.REQUEST,
			minArgumentCount:  1,
			maxArgumentCount:  1,
		},
	}
}

func init() {
	bc := NewDoRequest()
	BuiltinCompilers[bc.getName()] = bc
}

func (d *DoRequest) builtinCompile(
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

	codes = codeAppend(codes, base.DO_REQUEST, caller, file_name, row)

	err = d.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (d *DoRequest) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := d.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	d.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopStack()

	is_symbol, err :=
		d.isSymbolOrWantType(tstack, d.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err := d.setFunctionArgumentTypes(tstack, caller, d.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], d.getName())

	return nil
}
