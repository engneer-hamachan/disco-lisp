package compiler

import (
	"disco/base"
)

type Vector struct {
	BuiltinCompiler
}

func NewVector() BuiltinCompilerIF {
	return &Vector{
		BuiltinCompiler{
			name:             "vector",
			returnType:       base.VECTOR,
			minArgumentCount: 1,
			maxArgumentCount: 1000,
		},
	}
}

func init() {
	bc := NewVector()
	BuiltinCompilers[bc.getName()] = bc
}

func (v *Vector) builtinCompile(
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

	codes = codeAppend(codes, base.MAKE_VECTOR, caller, file_name, row)
	codes = codeAppend(codes, sLength(s), caller, file_name, row)

	err = v.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (v *Vector) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := v.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	v.setFunctionReturnTypes(caller)

	TypeEnv.PopMultiStack(sLength(s))

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], v.getName())

	return nil
}
