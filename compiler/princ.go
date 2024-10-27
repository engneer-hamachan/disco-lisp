package compiler

import (
	"disco/base"
)

type Princ struct {
	BuiltinCompiler
}

func NewPrinc() BuiltinCompilerIF {
	return &Princ{
		BuiltinCompiler{
			name:             "princ",
			returnType:       base.ANY,
			minArgumentCount: 1,
			maxArgumentCount: 1,
		},
	}
}

func init() {
	bc := NewPrinc()
	BuiltinCompilers[bc.getName()] = bc

	alias := NewPrinc()
	alias.setAlias("pc")
	BuiltinCompilers[alias.getName()] = alias
}

func (pr *Princ) builtinCompile(
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

	codes = codeAppend(codes, base.PRINC, caller, file_name, row)

	err = pr.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (pr *Princ) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := pr.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	TypeEnv.PopStack()

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], pr.getName())

	return nil
}
