package compiler

import (
	"disco/base"
)

type Atomp struct {
	BuiltinCompiler
}

func NewAtomp() BuiltinCompilerIF {
	return &Atomp{
		BuiltinCompiler{
			name:              "atom?",
			returnType:        base.ANY,
			firstArgumentType: base.ANY,
			minArgumentCount:  1,
			maxArgumentCount:  1,
		},
	}
}

func init() {
	bc := NewAtomp()
	BuiltinCompilers[bc.getName()] = bc
}

func (a *Atomp) builtinCompile(
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

	codes = codeAppend(codes, base.ATOMP, caller, file_name, row)

	err = a.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (a *Atomp) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := a.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	a.setFunctionReturnTypes(caller)

	TypeEnv.PopStack()

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], a.getName())

	return nil
}
