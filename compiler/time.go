package compiler

import (
	"disco/base"
)

type Time struct {
	BuiltinCompiler
}

func NewTime() BuiltinCompilerIF {
	return &Time{}
}

func init() {
	bc := NewTime()
	BuiltinCompilers["time"] = bc
}

func (t *Time) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	codes = codeAppend(codes, base.TIME_START, caller, file_name, row)

	codes, err := compiler.Compile(codes, s.GetCar(), caller, file_name, row)
	if err != nil {
		return nil, err
	}

	codes = codeAppend(codes, base.TIME_END, caller, file_name, row)

	err = t.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (t *Time) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], "time")

	return nil
}
