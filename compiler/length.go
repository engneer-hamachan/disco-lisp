package compiler

import (
	"disco/base"
)

type Length struct {
	BuiltinCompiler
}

func NewLength() BuiltinCompilerIF {
	return &Length{
		BuiltinCompiler{
			name:              "length",
			returnType:        base.INT,
			firstArgumentType: base.ANY,
			minArgumentCount:  1,
			maxArgumentCount:  1,
		},
	}
}

func init() {
	bc := NewLength()
	BuiltinCompilers[bc.getName()] = bc
}

func (l *Length) builtinCompile(
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

	codes = codeAppend(codes, base.LENGTH, caller, file_name, row)

	err = l.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (l *Length) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := l.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	l.setFunctionReturnTypes(caller)

	TypeEnv.PopStack()

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], l.getName())

	return nil
}
