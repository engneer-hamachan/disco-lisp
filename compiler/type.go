package compiler

import (
	"disco/base"
)

type Type struct {
	BuiltinCompiler
}

func NewType() BuiltinCompilerIF {
	return &Type{
		BuiltinCompiler{
			name:              "type",
			firstArgumentType: base.ANY,
			returnType:        base.STRING,
			minArgumentCount:  1,
			maxArgumentCount:  1,
		},
	}
}

func init() {
	bc := NewType()
	BuiltinCompilers[bc.getName()] = bc
}

func (t *Type) builtinCompile(
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

	codes = codeAppend(codes, base.TYPE, caller, file_name, row)

	err = t.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (t *Type) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := t.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	t.setFunctionReturnTypes(caller)

	TypeEnv.PopStack()

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], t.getName())

	return nil
}
