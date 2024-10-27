package compiler

import (
	"disco/base"
)

type If struct {
	BuiltinCompiler
}

func NewIf() BuiltinCompilerIF {
	return &If{
		BuiltinCompiler{
			name:             "if",
			returnType:       base.ANY,
			minArgumentCount: 1,
			maxArgumentCount: 1000,
		},
	}
}

func init() {
	bc := NewIf()
	BuiltinCompilers["if"] = bc
}

func (i *If) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	FunctionReturnTypes[caller] = i.returnType

	codes, err := compiler.Compile(codes, s.GetCar(), caller, file_name, row)
	if err != nil {
		return nil, err
	}

	codes = codeAppend(codes, base.IF, caller, file_name, row)

	var then_codes []any

	then_codes, err =
		compiler.Compile(then_codes, s.GetCadr(), caller, file_name, row)
	if err != nil {
		return nil, err
	}

	then_return_type := FunctionReturnTypes[caller]

	var else_codes []any

	else_codes, err =
		compiler.Compile(else_codes, s.GetCaddr(), caller, file_name, row)
	if err != nil {
		return nil, err
	}

	else_return_type := FunctionReturnTypes[caller]

	codes = codeAppend(codes, len(then_codes)+2, caller, file_name, row)

	for _, code := range then_codes {
		codes = codeAppend(codes, code, caller, file_name, row)
	}

	codes = codeAppend(codes, base.JMP, caller, file_name, row)
	codes = codeAppend(codes, len(else_codes), caller, file_name, row)

	for _, code := range else_codes {
		codes = codeAppend(codes, code, caller, file_name, row)
	}

	i.typePropagationForOptionalType(then_return_type, else_return_type, caller)

	err = i.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (i *If) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {
	TypeEnv.PopMultiStack(sLength(s))

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], "if")

	return nil
}

func (i *If) typePropagationForOptionalType(
	then_return_type int,
	else_return_type int,
	caller string,
) error {

	i.clearFunctionOptionalTypes(caller)

	switch then_return_type {
	case else_return_type:
		i.setFunctionChoiceReturnTypes(caller, then_return_type)

	default:
		i.setFunctionChoiceReturnTypes(caller, base.OPTIONAL)
		i.appendFunctionOptionalTypes(caller, then_return_type)
		i.appendFunctionOptionalTypes(caller, else_return_type)
	}

	return nil
}
