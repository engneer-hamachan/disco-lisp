package compiler

import (
	"disco/base"
)

type Pairp struct {
	BuiltinCompiler
}

func NewPairp() BuiltinCompilerIF {
	return &Pairp{
		BuiltinCompiler{
			name:             "pair?",
			returnType:       base.ANY,
			minArgumentCount: 1,
			maxArgumentCount: 1,
		},
	}
}

func init() {
	bc := NewPairp()
	BuiltinCompilers[bc.getName()] = bc
}

func (pa *Pairp) builtinCompile(
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

	codes = codeAppend(codes, base.PAIRP, caller, file_name, row)

	err = pa.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (pa *Pairp) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := pa.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	pa.setFunctionReturnTypes(caller)

	TypeEnv.PopStack()

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], pa.getName())

	return nil
}
